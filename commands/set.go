package commands

import (
	"fmt"
	"strings"

	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/libsecrets"
)

// Set sets a value in secrets
func Set(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		g.LogError(err)
	}

	kv := strings.SplitN(c.Args().First(), "=", 2)
	if len(kv) != 2 {
		err := fmt.Errorf("Please provide key and value for set command in the format key=value")
		g.LogError(err)
	}

	scope.Set(kv[0], kv[1])
	scope.Save()
}
