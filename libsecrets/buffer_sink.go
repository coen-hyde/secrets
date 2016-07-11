package libsecrets

import (
	"bytes"
)

// BufferSink is used to capture Keybase decrypted data
type BufferSink struct {
	buf  *bytes.Buffer
	open bool
}

// Open is a stub to fit the Sink interface
func (s *BufferSink) Open() error {
	s.open = true
	return nil
}

// Close is a stub to fit the Sink interface
func (s *BufferSink) Close() error {
	s.open = false
	return nil
}

// Write writes bytes to an internal buffer
func (s *BufferSink) Write(b []byte) (n int, err error) {
	return s.buf.Write(b)
}

// HitError is a stub to fit the Sink interface
func (s *BufferSink) HitError(e error) error { return nil }

// String returns the buffer as a string
func (s *BufferSink) String() string {
	return s.buf.String()
}
