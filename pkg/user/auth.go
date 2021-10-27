package user

import (
	"math/rand"
	"strings"
)

type Auth struct {
	Provider string
	Sub      string
}

func AuthFromAuth0Sub(sub string) Auth {
	s := strings.Split(sub, "|")
	if len(s) != 2 {
		return Auth{}
	}
	return Auth{Provider: s[0], Sub: sub}
}

func (a Auth) IsAuth0() bool {
	return a.Provider == "auth0"
}

func GenReearthSub(length int) Auth {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return Auth{
		Provider: "reearth",
		Sub:      string(b),
	}
}
