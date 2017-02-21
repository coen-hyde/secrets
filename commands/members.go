package commands

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	cli "github.com/codegangsta/cli"

	"github.com/coen-hyde/secrets/libsecrets"
	"github.com/keybase/client/go/client"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

// MembersList gets all or a specific value from the secrets
func MembersList(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, member := range scope.Members {
		fmt.Println(member.DisplayName)
	}
}

// MembersAdd add a new member to the scope
func MembersAdd(c *cli.Context) {
	scope, err := libsecrets.NewScope("default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	userCli, err := client.GetUserClient(libkb.G)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	loadUserArgs := keybase1.LoadUserByNameArg{
		Username: c.Args().First(),
	}

	user, err := userCli.LoadUserByName(context.TODO(), loadUserArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	member := libsecrets.NewMemberFromKeybaseUser(&user)
	adder := libsecrets.NewMemberFromKeybaseUser(libsecrets.G.KeybaseUser)
	scope.AddMember(member, adder)

	err = scope.Save()
	if err != nil {
		g.Log.Error(err.Error())
		os.Exit(1)
	}
}

// MembersRemove add a new member to the scope
func MembersRemove(c *cli.Context) {

}
