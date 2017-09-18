package libsecrets

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/keybase/client/go/client"
)

// Scope encapsulates a logical set of secrets
type Scope struct {
	Name    string
	Members []Member
	Data    map[string]string
}

// ImportOptions is a set of options for importing secrets
type ImportOptions struct {
	Format string
	regex  *regexp.Regexp
}

// NewScope instantiates a Scope struct.
// If the Scope already exists then creation will fail.
// In that case use `GetScope`.
func NewScope(name string) (scope *Scope) {
	scope = &Scope{
		Name:    name,
		Members: make([]Member, 0),
		Data:    make(map[string]string),
	}

	return
}

// GetScope returns an existing Scope
func GetScope(name string) (scope *Scope, err error) {
	scope = NewScope(name)
	err = scope.Load()

	return scope, err
}

// CreateScope creates a new Scope and saves it to disk.
func CreateScope(name string) (scope *Scope, err error) {
	scope = NewScope(name)
	if scope.Exists() {
		return nil, fmt.Errorf("The \"%s\" scope already exists", scope.Name)
	}

	err = scope.Save()
	if err != nil {
		return nil, err
	}

	return scope, nil
}

// RemoveScope deletes an existing scope
func RemoveScope(name string) (err error) {
	scope := NewScope(name)
	path := scope.KeybaseSinkPath()

	if !scope.Exists() {
		return fmt.Errorf("The \"%s\" scope does not exist", name)
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return
}

// MakeScopeLocation constructs the path of a scope
func makeScopePath(name string) string {
	return G.Dir() + "/" + name
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

// Exists returns whether a Scope exists on disk
func (s *Scope) Exists() bool {
	return fileExists(s.KeybaseSinkPath())
}

// Get returns a secret from a Scope
func (s *Scope) Get(key string) string {
	return s.Data[key]
}

// Set returns a secret from a Scope
func (s *Scope) Set(key string, value string) {
	s.Data[key] = value
}

// Del deletes a secret from a Scope
func (s *Scope) Del(key string) {
	delete(s.Data, key)
}

// Path returns the file path of the secret file
func (s *Scope) Path() string {
	return makeScopePath(s.Name)
}

// KeybaseSinkPath is the location of the keybase encryption
func (s *Scope) KeybaseSinkPath() string {
	return s.Path() + ".keybase"
}

// Load reads the secret scope from disk
func (s *Scope) Load() error {
	if !s.Exists() {
		return fmt.Errorf("The \"%s\" scope does not exist", s.Name)
	}

	src := client.NewFileSource(s.KeybaseSinkPath())
	sink := NewBufferSink()

	err := Decrypt(src, sink, true, false)

	if err != nil {
		return err
	}

	return json.Unmarshal(sink.Bytes(), &s)
}

// AddMembers adds a list of members to the Scope
func (s *Scope) AddMembers(members []*Member, adder *Member) []*Member {
	membersAdded := []*Member{}

	for i := 0; i < len(members); i++ {
		member := members[i]

		// Member already exists
		if s.MemberExists(member.Identifier) {
			G.Log.Warning("%s is already a member of this scope", member.Identifier)
			continue
		}

		member.AddedBy = adder.Identifier
		s.Members = append(s.Members, *member)
		membersAdded = append(membersAdded, member)
	}

	return membersAdded
}

// RemoveMembersByIdentifiers removes members from a Scope
func (s *Scope) RemoveMembersByIdentifiers(members []string) []*Member {
	membersKept := []Member{}
	membersRemoved := []*Member{}

	for i := 0; i < len(s.Members); i++ {
		member := s.Members[i]
		keep := true
		for j := 0; j < len(members); j++ {
			memberIdentifier := members[j]
			if memberIdentifier == member.Identifier {
				keep = false
			}
		}

		if keep {
			membersKept = append(membersKept, member)
		} else {
			membersRemoved = append(membersRemoved, &member)
		}
	}

	s.Members = membersKept

	return membersRemoved
}

// MemberPointers returns a list with pointers to the members in this scope
func (s *Scope) MemberPointers() (members []*Member) {
	for _, member := range s.Members {
		members = append(members, &member)
	}

	return members
}

// MemberExists tests if a member already exists in this scope
func (s *Scope) MemberExists(identifier string) bool {
	result := false

	for _, member := range s.Members {
		if member.Identifier == identifier {
			result = true
		}
	}

	return result
}

// Save writes the secret scope to disk
func (s *Scope) Save() error {
	// If this is a new Scope (not on disk yet), then we need to add the
	// current keybase user as a member
	if !s.Exists() {
		member := NewMemberFromKeybaseUser(G.KeybaseUser)
		s.AddMembers([]*Member{member}, member)
	}

	data, err := s.ToJSON()
	if err != nil {
		return err
	}

	src := NewBufferSource(&data)
	sink := client.NewFileSink(s.KeybaseSinkPath())

	return Encrypt(src, sink, GetMemberListIdentifiers(s.MemberPointers()))
}

// Export returns this Scope's data in the requested format
func (s *Scope) Export(format string) (string, error) {
	var formatter Formatter

	switch format {
	case "json":
		formatter = NewExporterJSON(&s.Data)
	case "yaml":
		formatter = NewExporterYAML(&s.Data)
	case "human":
		formatter = NewExporterHuman(&s.Data)
	case "env":
		formatter = NewExporterEnv(&s.Data)
	default:
		return "", fmt.Errorf("Unknown export format %s", format)
	}

	return formatter.String(), nil
}

// Import adds `contents` to the Scope with the given options
func (s *Scope) Import(contents string, options ImportOptions) error {
	var parser Importer

	switch options.Format {
	case "json":
		parser = NewImporterJSON(options)
	case "yaml":
		parser = NewImporterYAML(options)
	default:
		return fmt.Errorf("Unknown import format %s", options.Format)
	}

	structuredData, err := parser.Parse(contents)

	if err != nil {
		return err
	}

	for k, v := range structuredData {
		switch valueType := v.(type) {
		case string:
			s.Data[k] = v.(string)
		case float64:
			s.Data[k] = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		case int64:
			s.Data[k] = strconv.FormatInt(v.(int64), 10)
		case int:
			s.Data[k] = strconv.Itoa(v.(int))
		default:
			return fmt.Errorf("Can not use %s as string", valueType)
		}
	}

	return nil
}

// ToJSON converts this scope to json
func (s *Scope) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}
