package commands

import (
	"fmt"
	"os"

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
		fmt.Println(err)
		os.Exit(1)
	}

	export, err := scope.Export(c.String("format"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(export)
}
