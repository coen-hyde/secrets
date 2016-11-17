package libsecrets

import (
	"golang.org/x/net/context"

	"github.com/keybase/client/go/client"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

// Encrypt encrypts data from the source stream to the sink stream for the
// list of provided members
func Encrypt(source client.Source, sink client.Sink, members []string) error {
	cli, err := client.GetSaltpackClient(libkb.G)
	if err != nil {
		return err
	}

	protocols := []rpc.Protocol{
		client.NewStreamUIProtocol(libkb.G),
		client.NewSecretUIProtocol(libkb.G),
		client.NewIdentifyUIProtocol(libkb.G),
	}
	if err = client.RegisterProtocolsWithContext(protocols, libkb.G); err != nil {
		return err
	}

	filter := NewStreamFilter(source, sink)

	snk, src, err := filter.ClientOpen()
	if err != nil {
		return err
	}

	opts := keybase1.SaltpackEncryptOptions{
		Recipients:     members,
		NoSelfEncrypt:  false,
		Binary:         false,
		HideRecipients: false,
	}

	arg := keybase1.SaltpackEncryptArg{Source: src, Sink: snk, Opts: opts}
	err = cli.SaltpackEncrypt(context.TODO(), arg)
	ferr := filter.Close(err)
	return libkb.PickFirstError(err, ferr)
}
