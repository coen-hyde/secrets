package commands

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli"

	"github.com/coen-hyde/secrets/libsecrets"
)

// Get gets a value by key from a scope
func Get(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(c.Args()) != 1 {
		libsecrets.G.Log.Error("The get command requires exactly one argument")
		os.Exit(1)
	}

	key := c.Args().First()
	fmt.Println(scope.Get(key))
}
