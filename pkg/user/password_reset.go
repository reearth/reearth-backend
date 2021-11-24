package user

import (
	"time"

	"github.com/dgryski/trifles/uuid"
)

type PasswordReset struct {
	Token     string
	CreatedAt time.Time
}

func NewPasswordReset() *PasswordReset {
	return &PasswordReset{
		Token:     generateToken(),
		CreatedAt: time.Now(),
	}
}

func generateToken() string {
	return uuid.UUIDv4()
}

func (pr *PasswordReset) IsValidRequest(token string) bool {
	return pr != nil && pr.Token == token && pr.CreatedAt.Add(24*time.Hour).After(time.Now())
}
