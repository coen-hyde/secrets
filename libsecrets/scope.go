package libsecrets

import (
	"errors"
	"fmt"
	"os"
)

// Scope encapsulates a logical set of secrets
type Scope struct {
	Name     string
	Location string
	Members  []Member
	Data     map[string]interface{}
}

// Member is a keybase user
type Member struct {
	username string
}

// NewScope instanciates a scope struct
func NewScope(name string) (scope *Scope, err error) {
	location := makeScopePath(name)
	scope = &Scope{
		Name:     name,
		Location: location,
	}

	// Load existing data if it exists
	if fileExists(makeScopePath(name)) {
		err = scope.load()
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
	err = scope.save()

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

	fmt.Println(err)
	os.Exit(1)
	return false
}

func (s *Scope) load() error {
	if !fileExists(s.Location) {
		return errors.New("Can not load scope " + s.Name + " from location " + s.Location + ". No such file")
	}

	return nil
}

func (s *Scope) save() error {
	// errors.New("Can not save scope " + scope.Name + " at location " + scope.Location)
	return nil
}
