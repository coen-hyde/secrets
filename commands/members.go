package commands

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"

	cli "github.com/urfave/cli"

	"github.com/bugcrowd/secrets/libsecrets"
	"github.com/keybase/client/go/client"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

func loadMembersFromArgs(c *cli.Context) ([]*libsecrets.Member, error) {
	userCli, err := client.GetUserClient(libkb.G)
	if err != nil {
		return nil, err
	}

	args := c.Args()
	members := []*libsecrets.Member{}

	for i := 0; i < len(args); i++ {
		loadUserArgs := keybase1.LoadUserByNameArg{
			Username: args[i],
		}

		user, err := userCli.LoadUserByName(context.TODO(), loadUserArgs)
		if err != nil {
			err := fmt.Errorf("Could not load Keybase user \"%s\"", loadUserArgs.Username)
			return nil, err
		}

		member := libsecrets.NewMemberFromKeybaseUser(&user)
		members = append(members, member)
	}

	return members, nil
}

// MembersList gets all or a specific value from the secrets
func MembersList(c *cli.Context) {
	scope, err := libsecrets.GetScope("default")
	if err != nil {
		g.LogError(err)
	}

	for _, member := range scope.Members {
		fmt.Println(member.Identifier)
	}
}

// MembersAdd add a new member to the scope
func MembersAdd(c *cli.Context) {
	scope, err := libsecrets.GetScope("default")
	if err != nil {
		g.LogError(err)
	}

	members, err := loadMembersFromArgs(c)
	if err != nil {
		g.LogError(err)
	}

	adder := libsecrets.NewMemberFromKeybaseUser(libsecrets.G.KeybaseUser)
	membersAdded := scope.AddMembers(members, adder)

	err = scope.Save()
	if err != nil {
		g.LogError(err)
	}

	if len(membersAdded) == 0 {
		g.Log.Warning("No members were added")
		return
	}

	g.Log.Notice("Added members %s", strings.Join(libsecrets.GetMemberListIdentifiers(membersAdded), ", "))
}

// MembersRemove add a new member to the scope
func MembersRemove(c *cli.Context) {
	scope, err := libsecrets.GetScope("default")
	if err != nil {
		g.LogError(err)
	}

	membersRemoved := scope.RemoveMembersByIdentifiers(c.Args())

	err = scope.Save()
	if err != nil {
		g.LogError(err)
	}

	g.Log.Notice("Removed members %s", strings.Join(libsecrets.GetMemberListIdentifiers(membersRemoved), ", "))
}
