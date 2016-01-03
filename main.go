package main

import (
	"os"

	cli "github.com/coen-hyde/secrets/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/coen-hyde/secrets/commands"
)

func main() {
	// g := libkb.G
	// g.Init()

	app := cli.NewApp()
	app.Name = "Secrets"
	app.Usage = "Managing your application secrets"
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "Initialize a Secrets respository in the current directory",
			Action: func(c *cli.Context) {
				commands.Init(c)
			},
		},
	}

	app.Run(os.Args)
}
