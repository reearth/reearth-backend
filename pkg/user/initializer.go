package user

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"golang.org/x/text/language"
)

type InitParams struct {
	Email    string
	Name     string
	Sub      *Auth
	Password string
	Lang     *language.Tag
	Theme    *Theme
	UserID   *id.UserID
	TeamID   *id.TeamID
}

func Init(p InitParams) (*User, *Team, error) {
	if p.UserID == nil {
		p.UserID = id.NewUserID().Ref()
	}
	if p.TeamID == nil {
		p.TeamID = id.NewTeamID().Ref()
	}
	if p.Lang == nil {
		p.Lang = &language.Tag{}
	}
	if p.Theme == nil {
		t := ThemeDefault
		p.Theme = &t
	}
	if p.Sub == nil {
		p.Sub = GenReearthSub(p.UserID.String())
	}

	u, err := New().
		ID(*p.UserID).
		Name(p.Name).
		Email(p.Email).
		Auths([]Auth{*p.Sub}).
		Lang(*p.Lang).
		Password(p.Password).
		Theme(*p.Theme).
		Build()
	if err != nil {
		return nil, nil, err
	}

	// create a user's own team
	t, err := NewTeam().
		ID(*p.TeamID).
		Name(p.Name).
		Members(map[id.UserID]Role{u.ID(): RoleOwner}).
		Personal(true).
		Build()
	if err != nil {
		return nil, nil, err
	}
	u.UpdateTeam(t.ID())

	return u, t, err
}
