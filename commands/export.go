package commands

import (
	"fmt"

	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/libsecrets"
)

// Export gets all values from the secrets
func Export(c *cli.Context) {
	scope, err := libsecrets.GetScope("default")
	if err != nil {
		g.LogError(err)
	}

	export, err := scope.Export(c.String("format"))
	if err != nil {
		g.LogError(err)
	}

	fmt.Print(export)
}
