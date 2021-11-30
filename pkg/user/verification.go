package user

import (
	"crypto/rand"
	"errors"
	"math/big"
	"time"
)

var codeAlphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

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

func generateCode() (string, error) {
	code := ""
	for i := 0; i < 5; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(codeAlphabet))))
		if err != nil {
			return "", err
		}
		code += string(codeAlphabet[n.Int64()])
	}
	return code, nil
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

func NewVerification() (*Verification, error) {
	c, err := generateCode()
	if err != nil {
		return nil, errors.New("error generating verification code")
	}
	v := &Verification{
		verified:   false,
		code:       c,
		expiration: time.Now().Add(time.Hour * 24),
	}
	return v, nil
}

func VerificationFrom(c string, e time.Time, b bool) *Verification {
	v := &Verification{
		verified:   b,
		code:       c,
		expiration: e,
	}
	return v
}
