package libsecrets

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
	"golang.org/x/net/context"

	"github.com/keybase/client/go/client"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

// Decrypt decrypts data from the source stream to the sink stream
func Decrypt(source client.Source, sink client.Sink, interactive bool, force bool) error {
	cli, err := client.GetSaltpackClient(libkb.G)
	if err != nil {
		return err
	}

	// Setup SaltpackUI for user interaction
	spui := &SaltpackUI{
		Contextified: libkb.NewContextified(libkb.G),
		terminal:     libkb.G.UI.GetTerminalUI(),
		interactive:  interactive,
		force:        force,
	}

	protocols := []rpc.Protocol{
		client.NewStreamUIProtocol(libkb.G),
		client.NewSecretUIProtocol(libkb.G),
		client.NewIdentifyUIProtocol(libkb.G),
		keybase1.SaltpackUiProtocol(spui),
	}
	if err = client.RegisterProtocolsWithContext(protocols, libkb.G); err != nil {
		return err
	}

	// Hookup the source and the sink streams to Keybase client
	filter := NewStreamFilter(source, sink)

	snk, src, err := filter.ClientOpen()
	if err != nil {
		return err
	}

	arg := keybase1.SaltpackDecryptArg{
		Source: src,
		Sink:   snk,
		Opts: keybase1.SaltpackDecryptOptions{
			Interactive: interactive,
			UsePaperKey: false,
		},
	}

	var info keybase1.SaltpackEncryptedMessageInfo
	info, err = cli.SaltpackDecrypt(context.TODO(), arg)

	if _, ok := err.(libkb.NoDecryptionKeyError); ok {
		explainDecryptionFailure(&info)
	}

	cerr := filter.Close(err)
	return libkb.PickFirstError(err, cerr)
}

// explainDecryptionFailure is a port from Keybase client.CmdDecrypt
// Had to copy it as it was not exported
func explainDecryptionFailure(info *keybase1.SaltpackEncryptedMessageInfo) {
	if info == nil {
		return
	}
	out := libkb.G.UI.GetTerminalUI().ErrorWriter()
	prnt := func(s string, args ...interface{}) {
		fmt.Fprintf(out, s, args...)
	}
	if len(info.Devices) > 0 {
		prnt("Decryption failed; try one of these devices instead:\n")
		for _, d := range info.Devices {
			t := keybase1.FromTime(d.CTime)
			prnt("  * %s (%s); provisioned %s (%s)\n", client.ColorString("bold", d.Name), d.Type,
				humanize.Time(t), t.Format("2006-01-02 15:04:05 MST"))
		}
		if info.NumAnonReceivers > 0 {
			prnt("Additionally, there were %d hidden receivers for this message\n", info.NumAnonReceivers)
		}
	} else if info.NumAnonReceivers > 0 {
		prnt("Decryption failed; it was encrypted for %d hidden receivers, which may or may not be you\n", info.NumAnonReceivers)
	} else {
		prnt("Decryption failed; message wasn't encrypted for any of your known keys\n")
	}
}
