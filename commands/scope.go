package commands

import (
	"fmt"
	"path/filepath"
	"strings"

	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/libsecrets"
)

// ScopeAdd creates a new scope
func ScopeAdd(c *cli.Context) {
	args := c.Args()

	if len(args) != 1 {
		err := fmt.Errorf("The add command requires exactly one argument")
		g.LogError(err)
	}

	scope, err := libsecrets.CreateScope(args[0])
	if err != nil {
		g.LogError(err)
	}

	g.Log.Notice("Created the \"%s\" scope", scope.Name)
}

// ScopeRemove removes a scope
func ScopeRemove(c *cli.Context) {
	args := c.Args()

	if len(args) != 1 {
		err := fmt.Errorf("The remove command requires exactly one argument")
		g.LogError(err)
	}
	scope := args[0]
	err := libsecrets.RemoveScope(scope)

	if err != nil {
		err := fmt.Errorf("Error removing the \"%s\" scope: %s", scope, err)
		g.LogError(err)
	}

	g.Log.Notice("Removed the \"%s\" scope", scope)
}

// ScopeList lists all scopes
func ScopeList(c *cli.Context) {
	files, _ := filepath.Glob(g.Dir() + "/*.keybase")

	for i := 0; i < len(files); i++ {
		file := files[i]
		basename := filepath.Base(file)
		filename := strings.TrimSuffix(basename, filepath.Ext(basename))
		fmt.Println(filename)
	}
}
