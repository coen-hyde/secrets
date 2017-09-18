package commands

import (
	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/libsecrets"
)

// Set sets a value in secrets
func Del(c *cli.Context) {
	scope, err := libsecrets.GetScope("default")
	if err != nil {
		g.LogError(err)
	}

	scope.Del(c.Args().First())
	scope.Save()
}
