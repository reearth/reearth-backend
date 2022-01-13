package gqlmodel

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/user"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestToSearchedUser(t *testing.T) {
	uid := id.NewUserID()
	tests := []struct {
		name  string
		input *user.User
		want  *SearchedUser
	}{
		{
			name:  "should return nil",
			input: nil,
			want:  nil,
		},
		{
			name:  "shuld return SearchedUser",
			input: user.New().ID(uid).Name("foo").Email("user@email.com").MustBuild(),
			want: &SearchedUser{
				UserID:    uid.ID(),
				UserName:  "foo",
				UserEmail: "user@email.com",
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			got := ToSearchedUser(tc.input)
			assert.Equal(tt, tc.want, got)
		})
	}
}

func TestToTheme(t *testing.T) {
	th := Theme("")
	dark := ThemeDark
	light := ThemeLight
	userDefault := user.ThemeDefault
	userDark := user.ThemeDark
	userLight := user.ThemeLight
	tests := []struct {
		name  string
		input *Theme
		want  *user.Theme
	}{
		{
			name:  "should return nil",
			input: nil,
			want:  nil,
		},
		{
			name:  "should return ThemeDefault",
			input: &th,
			want:  &userDefault,
		},
		{
			name:  "should return ThemeDark",
			input: &dark,
			want:  &userDark,
		},
		{
			name:  "should return ThemeLight",
			input: &light,
			want:  &userLight,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			got := ToTheme(tc.input)
			assert.Equal(tt, tc.want, got)
		})
	}
}

func TestToUser(t *testing.T) {
	uid := id.NewUserID()
	tid := id.NewTeamID()
	tests := []struct {
		name  string
		input *user.User
		want  *User
	}{
		{
			name:  "should return nil",
			input: nil,
			want:  nil,
		},
		{
			name: "should return user",
			input: user.New().ID(uid).Name("foo").Team(tid).Theme(user.ThemeDark).Email("user@email.com").LangFrom("en-US").Auths([]user.Auth{
				{
					Provider: "test",
					Sub:      "xxx",
				},
			}).MustBuild(),
			want: &User{
				ID:       uid.ID(),
				Name:     "foo",
				Email:    "user@email.com",
				Lang:     language.MustParse("en-US"),
				Theme:    Theme(user.ThemeDark),
				MyTeamID: tid.ID(),
				Auths:    []string{"test"},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			got := ToUser(tc.input)
			assert.Equal(tt, tc.want, got)
		})
	}
}
