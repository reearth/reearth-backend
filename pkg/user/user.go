package user

import (
	"errors"

	"github.com/matthewhartstonge/argon2"
	"github.com/reearth/reearth-backend/pkg/id"
	"golang.org/x/text/language"
)

var (
	ErrEncodingPassword = errors.New("error encoding password")
)

type User struct {
	id            id.UserID
	name          string
	email         string
	password      []byte
	team          id.TeamID
	auths         []Auth
	lang          language.Tag
	theme         Theme
	passwordReset *PasswordReset
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

func (u *User) Password() []byte {
	return u.password
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

func (u *User) GetAuthByProvider(provider string) *Auth {
	if u == nil || u.auths == nil {
		return nil
	}
	for _, b := range u.auths {
		if provider == b.Provider {
			return &b
		}
	}
	return nil
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

func (u *User) SetPassword(pass string) error {
	p, err := encodePassword(pass)
	if err != nil {
		return err
	}
	u.password = p
	return nil
}

func (u *User) MatchPassword(pass string) (bool, error) {
	return verifyPassword(pass, u.password)
}

func encodePassword(pass string) ([]byte, error) {
	argon := argon2.DefaultConfig()
	encodedPass, err := argon.HashEncoded([]byte(pass))
	if err != nil {
		return nil, err
	}
	return encodedPass, nil
}

func verifyPassword(toVerify string, encoded []byte) (bool, error) {
	raw, err := argon2.Decode(encoded)
	if err != nil {
		return false, err
	}

	ok, err := raw.Verify([]byte(toVerify))
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (u *User) PasswordReset() *PasswordReset {
	return u.passwordReset
}

func (u *User) CreatePasswordReset() {
	u.passwordReset = NewPasswordReset()
}
