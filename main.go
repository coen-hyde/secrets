package main

import (
	"os"

	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/commands"
	"github.com/bugcrowd/secrets/libsecrets"
)

var (
	Version string = "dev"
)

func main() {
	libsecrets.G.Init()

	app := cli.NewApp()
	app.Name = "Secrets"
	app.Usage = "Managing your application secrets"
	app.Version = Version
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "scope, s",
					Value: "default",
					Usage: "Scope to use",
				},
			},
		},
		{
			Name:  "set",
			Usage: "Set a value",
			Action: func(c *cli.Context) error {
				commands.Set(c)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "scope, s",
					Value: "default",
					Usage: "Scope to use",
				},
			},
		},
		{
			Name:    "del",
			Aliases: []string{"remove", "delete"},
			Usage:   "Delete a value",
			Action: func(c *cli.Context) error {
				commands.Del(c)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "scope, s",
					Value: "default",
					Usage: "Scope to use",
				},
			},
		},
		{
			Name:  "export",
			Usage: "Export all data in a scope",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format, f",
					Value: "human",
					Usage: "Valid formats are 'human', 'json', 'yaml' and 'env'.",
				},
			},
			Action: func(c *cli.Context) error {
				commands.Export(c)
				return nil
			},
		},
		{
			Name:  "import",
			Usage: "Import data into a scope",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format, f",
					Value: "env",
					Usage: "Valid formats are 'json' and 'yaml'.",
				},
				cli.StringFlag{
					Name:  "data, d",
					Value: "",
					Usage: "Data to import",
				},
			},
			Action: func(c *cli.Context) error {
				commands.Import(c)
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
					Usage: "Add members",
					Action: func(c *cli.Context) error {
						commands.MembersAdd(c)
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "Remove members",
					Action: func(c *cli.Context) error {
						commands.MembersRemove(c)
						return nil
					},
				},
			},
		},
		{
			Name:  "scope",
			Usage: "Scope management",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "List existing scopes",
					Action: func(c *cli.Context) error {
						commands.ScopeList(c)
						return nil
					},
				},
				{
					Name:  "add",
					Usage: "Add a new scope",
					Action: func(c *cli.Context) error {
						commands.ScopeAdd(c)
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "Remove a scope",
					Action: func(c *cli.Context) error {
						commands.ScopeRemove(c)
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
