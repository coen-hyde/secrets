package libsecrets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/net/context"

	"github.com/keybase/client/go/client"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol"
	rpc "github.com/keybase/go-framed-msgpack-rpc"
)

// Scope encapsulates a logical set of secrets
type Scope struct {
	Name    string
	Members []Member
	Data    map[string]string
}

// NewScope instanciates a scope struct
func NewScope(name string) (scope *Scope, err error) {
	scope = &Scope{
		Name:    name,
		Members: make([]Member, 0),
		Data:    make(map[string]string),
	}

	// Load existing data if it exists
	if scope.exists() {
		err = scope.Load()
	}

	return
}

// CreateScope creates a new scope
func CreateScope(name string) (*Scope, error) {
	if fileExists(makeScopePath(name)) {
		return nil, fmt.Errorf("Can not create scope %s. Scope already exists", name)
	}

	scope, err := NewScope(name)
	if err != nil {
		return nil, err
	}

	// Add the creator of this scope as a member
	member := NewMemberFromKeybaseUser(G.KeybaseUser)
	scope.AddMember(member, member)

	err = scope.Save()

	return scope, err
}

// MakeScopeLocation constructs the path of a scope
func makeScopePath(name string) string {
	return Dir() + "/" + name
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

func (s *Scope) exists() bool {
	return fileExists(s.Path())
}

// Get returns a secret from
func (s *Scope) Get(key string) string {
	return s.Data[key]
}

// Set returns a secret from
func (s *Scope) Set(key string, value string) {
	s.Data[key] = value
}

// Path returns the file path of the secret file
func (s *Scope) Path() string {
	return makeScopePath(s.Name)
}

// Load reads the secret scope from disk
func (s *Scope) Load() error {
	if !s.exists() {
		return fmt.Errorf("Can not load scope %s from location %s. No such file", s.Name, s.Path())
	}

	// Read secrets from disk
	data, err := ioutil.ReadFile(s.Path())
	if err != nil {
		return err
	}

	// Import the data
	return json.Unmarshal(data, &s)
}

// AddMember adds a member to this scope
func (s *Scope) AddMember(m *Member, adder *Member) {
	m.AddedBy = adder.Uid
	s.Members = append(s.Members, *m)
}

// GetMemberUsernames returns an array of member usernames
func (s *Scope) GetMemberUsernames() (usernames []string) {
	for _, member := range s.Members {
		usernames = append(usernames, member.Username)
	}

	return usernames
}

// Save writes the secret scope to disk
func (s *Scope) Save() error {
	data, err := s.ToJSON()
	if err != nil {
		return err
	}

	cli, err := client.GetSaltpackClient(libkb.G)
	if err != nil {
		return err
	}

	protocols := []rpc.Protocol{
		client.NewStreamUIProtocol(libkb.G),
		client.NewSecretUIProtocol(libkb.G),
		client.NewIdentifyUIProtocol(libkb.G),
	}
	if err = client.RegisterProtocolsWithContext(protocols, libkb.G); err != nil {
		return err
	}

	filter := &client.UnixFilter{}

	if err = filter.FilterInit(string(data), "", s.Path()+".enc"); err != nil {
		return err
	}

	snk, src, err := filter.ClientFilterOpen()
	if err != nil {
		return err
	}

	opts := keybase1.SaltpackEncryptOptions{
		Recipients:     s.GetMemberUsernames(),
		NoSelfEncrypt:  false,
		Binary:         false,
		HideRecipients: false,
	}

	arg := keybase1.SaltpackEncryptArg{Source: src, Sink: snk, Opts: opts}
	err = cli.SaltpackEncrypt(context.TODO(), arg)
	ferr := filter.Close(err)
	return libkb.PickFirstError(err, ferr)
}

// ToJSON converts this scope to json
func (s *Scope) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}
