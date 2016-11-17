package libsecrets

import (
	"time"

	"github.com/keybase/client/go/protocol/keybase1"
)

// Member is a secret member
type Member struct {
	Username  string
	Uid       keybase1.UID
	AddedBy   keybase1.UID
	DateAdded time.Time
}

// NewMember instantiate a scope struct
func NewMember(username string, uid keybase1.UID) *Member {
	return &Member{
		Username:  username,
		Uid:       uid,
		DateAdded: time.Now(),
	}
}

// NewMemberFromKeybaseUser Creates a member from a Keybase User type
func NewMemberFromKeybaseUser(user *keybase1.User) *Member {
	return NewMember(user.Username, user.Uid)
}
