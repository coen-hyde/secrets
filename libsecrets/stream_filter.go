package libsecrets

import (
	"github.com/keybase/client/go/client"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

// NewStreamFilter creates a new StreamFilter
func NewStreamFilter(source client.Source, sink client.Sink) *StreamFilter {
	return &StreamFilter{
		source: source,
		sink:   sink,
	}
}

// StreamFilter manages the input and output streams.
type StreamFilter struct {
	source client.Source
	sink   client.Sink
}

// Open the streams for read and write
func (s *StreamFilter) Open() error {
	err := s.sink.Open()
	if err == nil {
		err = s.source.Open()
	}
	return err
}

// Close the streams
func (s *StreamFilter) Close(inerr error) error {
	e1 := s.source.CloseWithError(inerr)
	e2 := s.sink.Close()
	e3 := s.sink.HitError(inerr)
	return libkb.PickFirstError(e1, e2, e3)
}

// ClientOpen Connect the input and output streams with the Keybase client
func (s *StreamFilter) ClientOpen() (snk, src keybase1.Stream, err error) {
	if err = s.Open(); err != nil {
		return
	}

	snk = libkb.G.XStreams.ExportWriter(s.sink)
	src = libkb.G.XStreams.ExportReader(s.source)
	return
}
