package libsecrets

import (
	"time"

	keybase1 "github.com/keybase/client/go/protocol"
)

// Member is a secret member
type Member struct {
	Username  string
	Uid       keybase1.UID
	AddedBy   keybase1.UID
	DateAdded time.Time
}

// NewMember instanciates a scope struct
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
