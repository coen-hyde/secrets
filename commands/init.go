package commands

import (
	"fmt"
	"os"

	cli "github.com/codegangsta/cli"
	"github.com/coen-hyde/secrets/libsecrets"
)

var g = libsecrets.G

// Init initializes a secrets repository in the current directory
func Init(c *cli.Context) {
	if libsecrets.DirExists() {
		g.Log.Warning("Secrets repository has already been initialized")
		os.Exit(0)
	}

	// Create secrets directory
	if err := os.Mkdir(libsecrets.Dir(), 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: Create initial scopes
	_, err := libsecrets.CreateScope("default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	g.Log.Info("Initialized empty secrets repository at %s", libsecrets.Dir())
}
