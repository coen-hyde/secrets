package commands

import (
	"fmt"
	"os"
	"strings"

	cli "github.com/codegangsta/cli"
	"github.com/coen-hyde/secrets/libsecrets"
)

// Set sets a value in secrets
func Set(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	kv := strings.SplitN(c.Args().First(), "=", 2)
	if len(kv) != 2 {
		fmt.Println("Please provide key and value for set command in the format key=value")
		os.Exit(1)
	}

	scope.Set(kv[0], kv[1])
	scope.Save()
}
