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
	Log            logger.Logger
	KeybaseContext *libkb.GlobalContext
	KeybaseUser    *keybase1.User
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

	me, err := g.CurrentUser()
	if err != nil {
		g.LogError(err)
	}

	g.KeybaseUser = me

	client.InitUI()
	if err := client.GlobUI.Configure(); err != nil {
		g.Log.Warning("problem configuring UI: %s", err)
		g.Log.Warning("ignoring for now...")
	}
}

func (g *GlobalContext) initLibkb() {
	g.KeybaseContext = libkb.NewGlobalContext()
	g.KeybaseContext.Init()

	libkbOnce.Do(func() {
		g.KeybaseContext.Init()
		g.KeybaseContext.ConfigureConfig()
		g.KeybaseContext.ConfigureLogging()
		g.KeybaseContext.ConfigureCaches()
		g.KeybaseContext.ConfigureMerkleClient()
		g.KeybaseContext.ConfigureKeyring()
		g.KeybaseContext.ConfigureExportedStreams()
		g.KeybaseContext.ConfigureSocketInfo()
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
func (g *GlobalContext) CurrentUser() (*keybase1.User, error) {
	configCli, err := client.GetConfigClient(g.KeybaseContext)
	if err != nil {
		return nil, err
	}

	currentStatus, err := configCli.GetCurrentStatus(context.TODO(), 0)
	if err != nil {
		return nil, err
	}

	if !currentStatus.LoggedIn {
		return nil, fmt.Errorf("Please login to Keybase before using Secrets. You can do this by issuing the command `keybase login`")
	}

	myUID := currentStatus.User.Uid

	userCli, err := client.GetUserClient(g.KeybaseContext)
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
func (g *GlobalContext) Dir() string {
	pwd, err := os.Getwd()

	if err != nil {
		g.LogError(err)
	}

	return pwd + "/.secrets"
}

// DirExists does the secret directory exist?
func (g *GlobalContext) DirExists() bool {
	_, err := os.Stat(g.Dir())
	return (err == nil)
}
