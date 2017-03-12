package libsecrets

import (
	"time"

	"github.com/keybase/client/go/protocol/keybase1"
)

// Member is a secret member
type Member struct {
	Identifier string
	Type       string
	KeybaseUid keybase1.UID
	AddedBy    string
	DateAdded  time.Time
}

// NewMember instantiate a scope struct
func NewKeybaseMember(username string, uid keybase1.UID) *Member {
	return &Member{
		Identifier: username,
		Type:       "keybase",
		KeybaseUid: uid,
		DateAdded:  time.Now(),
	}
}

// NewMemberFromKeybaseUser Creates a member from a Keybase User type
func NewMemberFromKeybaseUser(user *keybase1.User) *Member {
	return NewKeybaseMember(user.Username, user.Uid)
}

func GetMemberListIdentifiers(members []*Member) (identifiers []string) {
	for _, memberPointer := range members {
		member := *memberPointer
		if member.Type != "keybase" {
			continue
		}
		identifiers = append(identifiers, member.Identifier)
	}

	return identifiers
}
