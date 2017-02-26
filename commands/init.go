package commands

import (
	"os"

	cli "github.com/urfave/cli"

	"github.com/coen-hyde/secrets/libsecrets"
)

var g = libsecrets.G

// Init initializes a secrets repository in the current directory
func Init(c *cli.Context) {
	if g.DirExists() {
		g.Log.Warning("Secrets repository has already been initialized")
		os.Exit(0)
	}

	// Create secrets directory
	if err := os.Mkdir(g.Dir(), 0755); err != nil {
		g.LogError(err)
	}

	// TODO: Create initial scopes
	_, err := libsecrets.CreateScope("default")
	if err != nil {
		g.LogError(err)
	}

	g.Log.Info("Initialized empty secrets repository at %s", g.Dir())
}
