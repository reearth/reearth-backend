package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPasswordReset(t *testing.T) {
	pr := NewPasswordReset()
	assert.NotNil(t, pr)
	assert.NotEmpty(t, pr.Token)
	assert.True(t, pr.CreatedAt.Sub(time.Now()) < 1*time.Second)
}

func TestPasswordReset_IsValidRequest(t *testing.T) {
	tests := []struct {
		name  string
		pr    *PasswordReset
		token string
		want  bool
	}{
		{
			name: "valid",
			pr: &PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Now(),
			},
			token: "xyz",
			want:  true,
		},
		{
			name: "wrong token",
			pr: &PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Now(),
			},
			token: "xxx",
			want:  false,
		},
		{
			name: "old request",
			pr: &PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Now().Add(-24 * time.Hour),
			},
			token: "xyz",
			want:  false,
		},
		{
			name:  "nil request",
			pr:    nil,
			token: "xyz",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.pr.IsValidRequest(tt.token))
		})
	}
}

func Test_generateToken(t *testing.T) {
	t1 := generateToken()
	t2 := generateToken()

	assert.NotNil(t, t1)
	assert.NotNil(t, t2)
	assert.NotEmpty(t, t1)
	assert.NotEmpty(t, t2)
	assert.NotEqual(t, t1, t2)

}
