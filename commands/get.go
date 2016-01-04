package commands

import (
	"fmt"
	"os"

	cli "github.com/coen-hyde/secrets/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/coen-hyde/secrets/libsecrets"
)

func printAll(s *libsecrets.Scope) {
	for key, value := range s.Data {
		fmt.Println(key + ": " + value)
	}
}

// Get gets all or a specific value from the secrets
func Get(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch len(c.Args()) {
	case 0:
		printAll(scope)
	case 1:
		key := c.Args().First()
		fmt.Println(scope.Get(key))
	}
}
