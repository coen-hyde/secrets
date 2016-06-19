package main

import (
	"os"

	cli "github.com/codegangsta/cli"
	"github.com/coen-hyde/secrets/commands"
	"github.com/coen-hyde/secrets/libsecrets"
)

func main() {
	libsecrets.G.Init()

	app := cli.NewApp()
	app.Name = "Secrets"
	app.Usage = "Managing your application secrets"
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "Initialize a Secrets respository in the current directory",
			Action: func(c *cli.Context) error {
				commands.Init(c)
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Get a value",
			Action: func(c *cli.Context) error {
				commands.Get(c)
				return nil
			},
		},
		{
			Name:  "set",
			Usage: "Set a value",
			Action: func(c *cli.Context) error {
				commands.Set(c)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
