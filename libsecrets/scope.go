package libsecrets

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// Scope encapsulates a logical set of secrets
type Scope struct {
	Name    string
	Members []Member
	Data    map[string]interface{}
}

// Member is a keybase user
type Member struct {
	username string
}

// NewScope instanciates a scope struct
func NewScope(name string) (scope *Scope, err error) {
	scope = &Scope{
		Name:    name,
		Members: make([]Member, 0),
		Data:    make(map[string]interface{}),
	}

	// Load existing data if it exists
	if scope.exists() {
		err = scope.Load()
	}

	return
}

// CreateScope creates a new scope
func CreateScope(name string) (scope *Scope, err error) {
	if fileExists(makeScopePath(name)) {
		return nil, errors.New("Can not create scope " + name + ". Scope already exists")
	}

	scope, err = NewScope(name)
	if err != nil {
		return
	}

	// TODO: Add current keybase user as Member
	err = scope.Save()

	return
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

// Path returns the file path of the secret file
func (s *Scope) Path() string {
	return makeScopePath(s.Name)
}

func (s *Scope) exists() bool {
	return fileExists(s.Path())
}

// Load reads the secret scope from disk
func (s *Scope) Load() error {
	if !s.exists() {
		return errors.New("Can not load scope " + s.Name + " from location " + s.Path() + ". No such file")
	}

	// Read secrets from disk
	data, err := ioutil.ReadFile(s.Path())
	if err != nil {
		return err
	}

	// Import the data
	return json.Unmarshal(data, &s)
}

// Save writes the secret scope to disk
func (s *Scope) Save() error {
	data, err := s.ToJSON()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.Path(), data, 0644)
}

// ToJSON converts this scope to json
func (s *Scope) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}
