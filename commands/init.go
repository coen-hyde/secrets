package commands

import (
	"os"

	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/libsecrets"
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
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		g.LogError(err)
	}

	err = scope.Save()
	if err != nil {
		g.LogError(err)
	}

	g.Log.Info("Initialized empty secrets repository at %s", g.Dir())
}
