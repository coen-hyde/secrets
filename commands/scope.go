package commands

import (
	"fmt"

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
	scope := c.String("scope")
	fmt.Printf("scope remove %s", scope)
}

// ScopeList lists all scopes
func ScopeList(c *cli.Context) {
	fmt.Println("scope list")
}
