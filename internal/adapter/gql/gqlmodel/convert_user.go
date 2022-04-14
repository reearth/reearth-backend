package gqlmodel

import (
	"github.com/reearth/reearth-backend/pkg/user"
	"github.com/reearth/reearth-backend/pkg/util"
)

func ToUser(u *user.User) *User {
	if u == nil {
		return nil
	}

	return &User{
		ID:       IDFrom(u.ID()),
		Name:     u.Name(),
		Email:    u.Email(),
		Lang:     u.Lang(),
		Theme:    Theme(u.Theme()),
		MyTeamID: IDFrom(u.Team()),
		Auths: util.Map(u.Auths(), func(a user.Auth) string {
			return a.Provider
		}),
	}
}

func ToSearchedUser(u *user.User) *SearchedUser {
	if u == nil {
		return nil
	}

	return &SearchedUser{
		UserID:    IDFrom(u.ID()),
		UserName:  u.Name(),
		UserEmail: u.Email(),
	}
}

func ToTheme(t *Theme) *user.Theme {
	if t == nil {
		return nil
	}

	th := user.ThemeDefault
	switch *t {
	case ThemeDark:
		th = user.ThemeDark
	case ThemeLight:
		th = user.ThemeLight
	}
	return &th
}

func ToTeam(t *user.Team) *Team {
	if t == nil {
		return nil
	}

	memberMap := t.Members().Members()
	members := make([]*TeamMember, 0, len(memberMap))
	for u, r := range memberMap {
		members = append(members, &TeamMember{
			UserID: IDFrom(u),
			Role:   ToRole(r),
		})
	}

	return &Team{
		ID:       IDFrom(t.ID()),
		Name:     t.Name(),
		Personal: t.IsPersonal(),
		Members:  members,
	}
}

func FromRole(r Role) user.Role {
	switch r {
	case RoleReader:
		return user.RoleReader
	case RoleWriter:
		return user.RoleWriter
	case RoleOwner:
		return user.RoleOwner
	}
	return user.Role("")
}

func ToRole(r user.Role) Role {
	switch r {
	case user.RoleReader:
		return RoleReader
	case user.RoleWriter:
		return RoleWriter
	case user.RoleOwner:
		return RoleOwner
	}
	return Role("")
}
