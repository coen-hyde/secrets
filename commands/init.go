package commands

import (
	"fmt"
	"os"

	cli "github.com/coen-hyde/secrets/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/coen-hyde/secrets/libsecrets"
)

// Init initializes a secrets repository in the current directory
func Init(c *cli.Context) {
	if libsecrets.DirExists() {
		fmt.Println("Secrets repository has already been initialized")
		os.Exit(0)
	}

	// Create secrets directory
	if err := os.Mkdir(libsecrets.Dir(), 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: Create initial scopes

	fmt.Println("Initialized empty secrets repository at", libsecrets.Dir())
}
