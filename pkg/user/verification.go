package user

import (
	"math/rand"
	"time"
)

type Verification struct {
	verified   bool
	code       string
	expiration time.Time
}

func (v *Verification) IsVerified() bool {
	if v == nil {
		return false
	}
	return v.verified
}

func (v *Verification) Code() string {
	if v == nil {
		return ""
	}
	return v.code
}

func (v *Verification) Expiration() time.Time {
	if v == nil {
		return time.Time{}
	}
	return v.expiration
}

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 5)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func (v *Verification) IsExpired() bool {
	if v == nil {
		return true
	}
	now := time.Now()
	return now.After(v.expiration)
}

func (v *Verification) SetVerified(b bool) {
	if v == nil {
		return
	}
	v.verified = b
}

func NewVerification() *Verification {
	v := &Verification{
		verified:   false,
		code:       generateCode(),
		expiration: time.Now().Add(time.Hour * 24),
	}
	return v
}

func NewVerificationFrom(c string, e time.Time, b bool) *Verification {
	v := &Verification{
		verified:   b,
		code:       c,
		expiration: e,
	}
	return v
}
