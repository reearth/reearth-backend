package user

import (
	"math/rand"

	"github.com/reearth/reearth-backend/pkg/id"
	"golang.org/x/text/language"
)

type User struct {
	id       id.UserID
	name     string
	email    string
	password []byte
	team     id.TeamID
	auths    []Auth
	lang     language.Tag
	theme    Theme
}

func (u *User) ID() id.UserID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Team() id.TeamID {
	return u.team
}

func (u *User) Lang() language.Tag {
	return u.lang
}

func (u *User) Theme() Theme {
	return u.theme
}

func (u *User) UpdateName(name string) {
	u.name = name
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

func (u *User) UpdateTeam(team id.TeamID) {
	u.team = team
}

func (u *User) UpdateLang(lang language.Tag) {
	u.lang = lang
}

func (u *User) UpdateTheme(t Theme) {
	u.theme = t
}

func (u *User) Auths() []Auth {
	if u == nil {
		return nil
	}
	return append([]Auth{}, u.auths...)
}

func (u *User) ContainAuth(a Auth) bool {
	if u == nil {
		return false
	}
	for _, b := range u.auths {
		if a == b || a.Provider == b.Provider {
			return true
		}
	}
	return false
}

func (u *User) AddAuth(a Auth) bool {
	if u == nil {
		return false
	}
	if !u.ContainAuth(a) {
		u.auths = append(u.auths, a)
		return true
	}
	return false
}

func (u *User) RemoveAuth(a Auth) bool {
	if u == nil || a.IsAuth0() {
		return false
	}
	for i, b := range u.auths {
		if a == b {
			u.auths = append(u.auths[:i], u.auths[i+1:]...)
			return true
		}
	}
	return false
}

func (u *User) RemoveAuthByProvider(provider string) bool {
	if u == nil || provider == "auth0" {
		return false
	}
	for i, b := range u.auths {
		if provider == b.Provider {
			u.auths = append(u.auths[:i], u.auths[i+1:]...)
			return true
		}
	}
	return false
}

func (u *User) ClearAuths() {
	u.auths = []Auth{}
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
