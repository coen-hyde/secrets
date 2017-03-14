package commands

import (
	cli "github.com/urfave/cli"

	"github.com/coen-hyde/secrets/libsecrets"
)

// Export gets all values from the secrets
func Import(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		g.LogError(err)
	}

	rawData := c.Args().First()
	err = scope.Import(rawData, c.String("format"))
	if err != nil {
		g.LogError(err)
	}

	err = scope.Save()
	if err != nil {
		g.LogError(err)
	}
}
