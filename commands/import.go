package commands

import (
	"bufio"
	"os"

	cli "github.com/urfave/cli"

	"github.com/coen-hyde/secrets/libsecrets"
)

// Export gets all values from the secrets
func Import(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
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
