package commands

import (
	"fmt"
	"os"

	cli "github.com/codegangsta/cli"
	"github.com/coen-hyde/secrets/libsecrets"
)

// MembersList gets all or a specific value from the secrets
func MembersList(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, member := range scope.Members {
		fmt.Println(member)
	}
}

// MembersAdd add a new member to the scope
func MembersAdd(c *cli.Context) {
	_, err := libsecrets.NewScope("default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// err = scope.AddMember(c.Args().First())
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	//
	// fmt.Println(err)
}

// MembersRemove add a new member to the scope
func MembersRemove(c *cli.Context) {

}
