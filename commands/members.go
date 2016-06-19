package commands

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	cli "github.com/codegangsta/cli"

	"github.com/keybase/client/go/client"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol"

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
		fmt.Println(member.Username)
	}
}

// MembersAdd add a new member to the scope
func MembersAdd(c *cli.Context) {
	_, err := libsecrets.NewScope("default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	userCli, err := client.GetUserClient(libkb.G)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	user, err := userCli.LoadUserByName(context.TODO(), keybase1.LoadUserByNameArg{Username: "codesoda"})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(user)
	os.Exit(1)
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
