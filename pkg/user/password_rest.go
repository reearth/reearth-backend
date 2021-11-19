package user

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

type PasswordReset struct {
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
	IsUsed    bool
}

func NewPasswordReset() *PasswordReset {
	return &PasswordReset{
		Token:     generateToken(),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
		IsUsed:    false,
	}
}

func generateToken() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func (pr *PasswordReset) IsValidRequest(token string) bool {
	return !pr.IsUsed && pr.Token == token && time.Now().Before(pr.ExpiresAt)
}

func (pr *PasswordReset) MarkAsUsed() {
	pr.IsUsed = true
}
