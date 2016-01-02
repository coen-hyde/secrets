package libsecrets

// Scope manages read and write to secrets
type Scope struct {
	file string
}

// NewScope initializse a new secrets scope instance
func NewScope(location string) *Scope {
	return &Scope{location}
}
