package commands

import (
	"fmt"

	cli "github.com/urfave/cli"

	"github.com/coen-hyde/secrets/libsecrets"
)

func printAll(s *libsecrets.Scope) {
	for key, value := range s.Data {
		fmt.Println(key + ": " + value)
	}
}

// Export gets all values from the secrets
func Export(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		g.LogError(err)
	}

	export, err := scope.Export(c.String("format"))
	if err != nil {
		g.LogError(err)
	}

	fmt.Print(export)
}
