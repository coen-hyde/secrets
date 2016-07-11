package libsecrets

import (
	"bytes"
)

func NewBufferSource(b *[]byte) *BufferSource {
	return &BufferSource{buf: bytes.NewBuffer(*b)}
}

type BufferSource struct {
	buf *bytes.Buffer
}

// Open is a stub to fit the Source interface
func (b *BufferSource) Open() error {
	return nil
}

// Read lenth p bytes
func (b *BufferSource) Read(p []byte) (n int, err error) {
	return b.buf.Read(p)
}

// Close is a stub to fit the Source interface
func (b *BufferSource) Close() error { return nil }

// CloseWithError is a stub to fit the Source interface
func (b *BufferSource) CloseWithError(error) error { return nil }
