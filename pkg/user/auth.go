package user

import "strings"

type Auth struct {
	Provider string
	Sub      string
}

func AuthFromAuth0Sub(sub string) Auth {
	s := strings.SplitN(sub, "|", 2)
	if len(s) != 2 {
		return Auth{}
	}
	return Auth{Provider: s[0], Sub: sub}
}

func (a Auth) IsAuth0() bool {
	return a.Provider == "auth0"
}
