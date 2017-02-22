package libsecrets

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/net/context"

	"github.com/keybase/client/go/client"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/logger"
	"github.com/keybase/client/go/protocol/keybase1"
)

var libkbOnce sync.Once

// GlobalContext stores the application global context
type GlobalContext struct {
	Log         logger.Logger
	KeybaseUser *keybase1.User
}

// NewGlobalContext initializes a new global context
func NewGlobalContext() *GlobalContext {
	return &GlobalContext{}
}

// G is the current global context
var G *GlobalContext

func init() {
	G = NewGlobalContext()
}

// Init Initializes the secrets app
func (g *GlobalContext) Init() {
	// Force Production Mode for the moment
	// os.Setenv("KEYBASE_DEBUG", "1")
	os.Setenv("KEYBASE_RUN_MODE", "prod")

	g.initLibkb()
	g.Log = g.intLogger()

	me, err := CurrentUser()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	g.KeybaseUser = me

	client.InitUI()
	if err := client.GlobUI.Configure(); err != nil {
		g.Log.Warning("problem configuring UI: %s", err)
		g.Log.Warning("ignoring for now...")
	}
}

func (g *GlobalContext) initLibkb() {
	libkbOnce.Do(func() {
		libkb.G.Init()
		libkb.G.ConfigureConfig()
		libkb.G.ConfigureLogging()
		libkb.G.ConfigureCaches()
		libkb.G.ConfigureMerkleClient()
		libkb.G.ConfigureKeyring()
		libkb.G.ConfigureExportedStreams()
		libkb.G.ConfigureSocketInfo()
	})
}

func (g *GlobalContext) intLogger() logger.Logger {
	log := logger.NewWithCallDepth("secrets", 1)

	return log
}

func (g *GlobalContext) LogError(err error) {
	g.Log.Error(err.Error())
	os.Exit(1)
}

// CurrentUser Get the current Keybase User
func CurrentUser() (*keybase1.User, error) {
	configCli, err := client.GetConfigClient(libkb.G)
	if err != nil {
		return nil, err
	}

	currentStatus, err := configCli.GetCurrentStatus(context.TODO(), 0)
	if err != nil {
		return nil, err
	}

	// If the user isnt logged in then there is nothing we can do, exit
	if !currentStatus.LoggedIn {
		G.Log.Error("Please login to Keybase before using Secrets. You can do this by issuing the command `keybase login`")
		os.Exit(1)
	}
	myUID := currentStatus.User.Uid

	userCli, err := client.GetUserClient(libkb.G)
	if err != nil {
		return nil, err
	}

	me, err := userCli.LoadUser(context.TODO(), keybase1.LoadUserArg{Uid: myUID})
	if err != nil {
		return nil, err
	}

	return &me, nil
}

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
