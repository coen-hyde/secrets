package main

import (
	"os"

	cli "github.com/urfave/cli"

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
		{
			Name:  "export",
			Usage: "Export all data in a scope",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format, f",
					Value: "human",
					Usage: "Format to output the data in",
				},
			},
			Action: func(c *cli.Context) error {
				commands.Export(c)
				return nil
			},
		},
		{
			Name:  "members",
			Usage: "Member management",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "List members in a scope",
					Action: func(c *cli.Context) error {
						commands.MembersList(c)
						return nil
					},
				},
				{
					Name:  "add",
					Usage: "Add a new member",
					Action: func(c *cli.Context) error {
						commands.MembersAdd(c)
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
