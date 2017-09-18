package commands

import (
	"fmt"

	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/libsecrets"
)

// Del deletes a value from secrets
func Del(c *cli.Context) {
	scopeName := c.String("scope")
	scope, err := libsecrets.GetScope(scopeName)

	if err != nil {
		g.LogError(err)
	}

	if len(c.Args()) != 1 {
		err = fmt.Errorf("The del command requires exactly one argument")
		g.LogError(err)
	}

	scope.Del(c.Args().First())
	scope.Save()
}
