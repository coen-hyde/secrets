package libsecrets

import (
	"time"

	"github.com/keybase/client/go/protocol/keybase1"
	"github.com/satori/go.uuid"
)

// Member is a secret member
type Member struct {
	Uuid        uuid.UUID
	DisplayName string
	Type        string
	KeybaseUid  keybase1.UID
	AddedBy     uuid.UUID
	DateAdded   time.Time
}

// NewMember instantiate a scope struct
func NewKeybaseMember(username string, uid keybase1.UID) *Member {
	return &Member{
		Uuid:        uuid.NewV4(),
		DisplayName: username,
		Type:        "keybase",
		KeybaseUid:  uid,
		DateAdded:   time.Now(),
	}
}

// NewMemberFromKeybaseUser Creates a member from a Keybase User type
func NewMemberFromKeybaseUser(user *keybase1.User) *Member {
	return NewKeybaseMember(user.Username, user.Uid)
}
