package commands

import (
	"fmt"

	cli "github.com/urfave/cli"
)

// "github.com/bugcrowd/secrets/libsecrets"

// ScopeAdd creates a new scope
func ScopeAdd(c *cli.Context) {
	scope := c.String("scope")
	fmt.Printf("scope add %s", scope)
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
