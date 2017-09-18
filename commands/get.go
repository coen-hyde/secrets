package commands

import (
	"fmt"

	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/libsecrets"
)

// Get gets a value by key from a scope
func Get(c *cli.Context) {
	context := c.GlobalString("context")
	scope, err := libsecrets.GetScope(context)

	if err != nil {
		g.LogError(err)
	}

	if len(c.Args()) != 1 {
		err = fmt.Errorf("The get command requires exactly one argument")
		g.LogError(err)
	}

	key := c.Args().First()
	fmt.Println(scope.Get(key))
}
