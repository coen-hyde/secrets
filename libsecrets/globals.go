package libsecrets

import (
	"fmt"
	"os"
	"sync"

	"github.com/keybase/client/go/client"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/logger"
)

var libkbGlobals = libkb.G
var libkbOnce sync.Once

// Init Initializes the secrets app
func Init() {
	// Force Production Mode for the moment
	os.Setenv("KEYBASE_RUN_MODE", "prod")

	initLibkb()
	log := logger.NewWithCallDepth("", 1, os.Stderr)

	client.InitUI()
	if err := client.GlobUI.Configure(); err != nil {
		log.Warning("problem configuring UI: %s", err)
		log.Warning("ignoring for now...")
	}
}

func initLibkb() {
	libkbOnce.Do(func() {
		libkb.G.Init()
		libkb.G.ConfigureConfig()
		libkb.G.ConfigureLogging()
		libkb.G.ConfigureCaches()
		libkb.G.ConfigureMerkleClient()
		libkb.G.ConfigureSocketInfo()
	})
}

// CurrentUser Get the current Keybase User
// func CurrentUser() (keybase1.User, error) {
// 	configCli, err := client.GetConfigClient(libkb.G)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	currentStatus, err := configCli.GetCurrentStatus(context.TODO(), 0)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !currentStatus.LoggedIn {
// 		return nil, fmt.Errorf("Not logged in.")
// 	}
// 	myUID := currentStatus.User.Uid
//
// 	userCli, err := client.GetUserClient()
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	me, err := userCli.LoadUser(context.TODO(), keybase1.LoadUserArg{Uid: myUID})
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return me, nil
// }

// Dir returns the directory where the secrets should be stored
func Dir() string {
	pwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return pwd + "/.secrets"
}

// DirExists does the secret directory exist?
func DirExists() bool {
	_, err := os.Stat(Dir())
	return (err == nil)
}
