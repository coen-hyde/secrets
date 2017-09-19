package commands

import (
	"bufio"
	"os"

	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/libsecrets"
)

// Import bulk adds data to a secrets scope
func Import(c *cli.Context) {
	scopeName := c.String("scope")
	scope, err := libsecrets.GetScope(scopeName)
	if err != nil {
		g.LogError(err)
	}

	data := c.String("data")

	// Attempt to fetch data from stdin if no data was passed via arguments
	if len(data) == 0 {
		reader := bufio.NewReader(os.Stdin)
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			data += "\n" + scanner.Text()
		}
	}

	options := libsecrets.ImportOptions{
		Format: c.String("format"),
	}

	err = scope.Import(data, options)
	if err != nil {
		g.LogError(err)
	}

	err = scope.Save()
	if err != nil {
		g.LogError(err)
	}
}
