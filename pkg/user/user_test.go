package user

import (
	"testing"
	"time"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestUser(t *testing.T) {
	uid := id.NewUserID()
	tid := id.NewTeamID()
	testCases := []struct {
		Name     string
		User     *User
		Expected struct {
			Id    id.UserID
			Name  string
			Email string
			Team  id.TeamID
			Auths []Auth
			Lang  language.Tag
		}
	}{
		{
			Name: "create user",
			User: New().ID(uid).
				Team(tid).
				Name("xxx").
				LangFrom("en").
				Email("ff@xx.zz").
				Auths([]Auth{{
					Provider: "aaa",
					Sub:      "sss",
				}}).MustBuild(),
			Expected: struct {
				Id    id.UserID
				Name  string
				Email string
				Team  id.TeamID
				Auths []Auth
				Lang  language.Tag
			}{
				Id:    uid,
				Name:  "xxx",
				Email: "ff@xx.zz",
				Team:  tid,
				Auths: []Auth{{
					Provider: "aaa",
					Sub:      "sss",
				}},
				Lang: language.Make("en"),
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected.Id, tc.User.ID())
			assert.Equal(tt, tc.Expected.Name, tc.User.Name())
			assert.Equal(tt, tc.Expected.Team, tc.User.Team())
			assert.Equal(tt, tc.Expected.Auths, tc.User.Auths())
			assert.Equal(tt, tc.Expected.Email, tc.User.Email())
			assert.Equal(tt, tc.Expected.Lang, tc.User.Lang())
		})
	}
}

func TestUser_AddAuth(t *testing.T) {
	testCases := []struct {
		Name     string
		User     *User
		A        Auth
		Expected bool
	}{
		{
			Name:     "nil user",
			User:     nil,
			Expected: false,
		},
		{
			Name: "add new auth",
			User: New().NewID().MustBuild(),
			A: Auth{
				Provider: "xxx",
				Sub:      "zzz",
			},
			Expected: true,
		},
		{
			Name: "existing auth",
			User: New().NewID().Auths([]Auth{{
				Provider: "xxx",
				Sub:      "zzz",
			}}).MustBuild(),
			A: Auth{
				Provider: "xxx",
				Sub:      "zzz",
			},
			Expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.User.AddAuth(tc.A)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestUser_RemoveAuth(t *testing.T) {
	testCases := []struct {
		Name     string
		User     *User
		A        Auth
		Expected bool
	}{
		{
			Name:     "nil user",
			User:     nil,
			Expected: false,
		},
		{
			Name: "remove auth0",
			User: New().NewID().MustBuild(),
			A: Auth{
				Provider: "auth0",
				Sub:      "zzz",
			},
			Expected: false,
		},
		{
			Name: "existing auth",
			User: New().NewID().Auths([]Auth{{
				Provider: "xxx",
				Sub:      "zzz",
			}}).MustBuild(),
			A: Auth{
				Provider: "xxx",
				Sub:      "zzz",
			},
			Expected: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.User.RemoveAuth(tc.A)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestUser_ContainAuth(t *testing.T) {
	testCases := []struct {
		Name     string
		User     *User
		A        Auth
		Expected bool
	}{
		{
			Name:     "nil user",
			User:     nil,
			Expected: false,
		},
		{
			Name: "not existing auth",
			User: New().NewID().MustBuild(),
			A: Auth{
				Provider: "auth0",
				Sub:      "zzz",
			},
			Expected: false,
		},
		{
			Name: "existing auth",
			User: New().NewID().Auths([]Auth{{
				Provider: "xxx",
				Sub:      "zzz",
			}}).MustBuild(),
			A: Auth{
				Provider: "xxx",
				Sub:      "zzz",
			},
			Expected: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.User.ContainAuth(tc.A)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestUser_RemoveAuthByProvider(t *testing.T) {
	testCases := []struct {
		Name     string
		User     *User
		Provider string
		Expected bool
	}{
		{
			Name:     "nil user",
			User:     nil,
			Expected: false,
		},
		{
			Name:     "remove auth0",
			User:     New().NewID().MustBuild(),
			Provider: "auth0",
			Expected: false,
		},
		{
			Name: "existing auth",
			User: New().NewID().Auths([]Auth{{
				Provider: "xxx",
				Sub:      "zzz",
			}}).MustBuild(),
			Provider: "xxx",
			Expected: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.User.RemoveAuthByProvider(tc.Provider)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestUser_ClearAuths(t *testing.T) {
	u := New().NewID().Auths([]Auth{{
		Provider: "xxx",
		Sub:      "zzz",
	}}).MustBuild()
	u.ClearAuths()
	assert.Equal(t, 0, len(u.Auths()))
}

func TestUser_Auths(t *testing.T) {
	var u *User
	assert.Equal(t, []Auth(nil), u.Auths())
}

func TestUser_UpdateEmail(t *testing.T) {
	u := New().NewID().MustBuild()
	u.UpdateEmail("ff@xx.zz")
	assert.Equal(t, "ff@xx.zz", u.Email())
}

func TestUser_UpdateLang(t *testing.T) {
	u := New().NewID().MustBuild()
	u.UpdateLang(language.Make("en"))
	assert.Equal(t, language.Make("en"), u.Lang())
}

func TestUser_UpdateTeam(t *testing.T) {
	tid := id.NewTeamID()
	u := New().NewID().MustBuild()
	u.UpdateTeam(tid)
	assert.Equal(t, tid, u.Team())
}

func TestUser_UpdateName(t *testing.T) {
	u := New().NewID().MustBuild()
	u.UpdateName("xxx")
	assert.Equal(t, "xxx", u.Name())
}

func TestUser_GetAuthByProvider(t *testing.T) {
	testCases := []struct {
		Name     string
		User     *User
		Provider string
		Expected *Auth
	}{
		{
			Name: "existing auth",
			User: New().NewID().Auths([]Auth{{
				Provider: "xxx",
				Sub:      "zzz",
			}}).MustBuild(),
			Provider: "xxx",
			Expected: &Auth{
				Provider: "xxx",
				Sub:      "zzz",
			},
		},
		{
			Name: "not existing auth",
			User: New().NewID().Auths([]Auth{{
				Provider: "xxx",
				Sub:      "zzz",
			}}).MustBuild(),
			Provider: "yyy",
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.User.GetAuthByProvider(tc.Provider)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestUser_MatchPassword(t *testing.T) {
	encodedPass, _ := encodePassword("test")
	type args struct {
		pass string
	}
	tests := []struct {
		name     string
		password []byte
		args     args
		want     bool
		wantErr  bool
	}{
		{
			name:     "passwords should match",
			password: encodedPass,
			args: args{
				pass: "test",
			},
			want:    true,
			wantErr: false,
		},
		{
			name:     "passwords shouldn't match",
			password: encodedPass,
			args: args{
				pass: "xxx",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			u := &User{
				password: tc.password,
			}
			got, err := u.MatchPassword(tc.args.pass)
			assert.Equal(tt, tc.want, got)
			if tc.wantErr {
				assert.Error(tt, err)
			} else {
				assert.NoError(tt, err)
			}
		})
	}
}

func TestUser_SetPassword(t *testing.T) {
	type args struct {
		pass string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should set the password",
			args: args{
				pass: "test",
			},
			want: "test",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			u := &User{}
			_ = u.SetPassword(tc.args.pass)
			got, err := verifyPassword(tc.want, u.password)
			assert.NoError(tt, err)
			assert.True(tt, got)
		})
	}
}

func TestUser_PasswordReset(t *testing.T) {
	testCases := []struct {
		Name     string
		User     *User
		Expected *PasswordReset
	}{
		{
			Name:     "not password request",
			User:     New().NewID().MustBuild(),
			Expected: nil,
		},
		{
			Name: "create new password request over existing one",
			User: New().NewID().PasswordReset(&PasswordReset{"xzy", time.Unix(0, 0)}).MustBuild(),
			Expected: &PasswordReset{
				Token:     "xzy",
				CreatedAt: time.Unix(0, 0),
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.User.PasswordReset())
		})
	}
}

func TestUser_SetPasswordReset(t *testing.T) {
	tests := []struct {
		Name     string
		User     *User
		Pr       *PasswordReset
		Expected *PasswordReset
	}{
		{
			Name:     "nil",
			User:     New().NewID().MustBuild(),
			Pr:       nil,
			Expected: nil,
		},
		{
			Name: "nil",
			User: New().NewID().MustBuild(),
			Pr: &PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Unix(1, 1),
			},
			Expected: &PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Unix(1, 1),
			},
		},
		{
			Name: "create new password request",
			User: New().NewID().MustBuild(),
			Pr: &PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Unix(1, 1),
			},
			Expected: &PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Unix(1, 1),
			},
		},
		{
			Name: "create new password request over existing one",
			User: New().NewID().PasswordReset(&PasswordReset{"xzy", time.Now()}).MustBuild(),
			Pr: &PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Unix(1, 1),
			},
			Expected: &PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Unix(1, 1),
			},
		},
		{
			Name:     "remove none existing password request",
			User:     New().NewID().MustBuild(),
			Pr:       nil,
			Expected: nil,
		},
		{
			Name:     "remove existing password request",
			User:     New().NewID().PasswordReset(&PasswordReset{"xzy", time.Now()}).MustBuild(),
			Pr:       nil,
			Expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.User.SetPasswordReset(tt.Pr)
			assert.Equal(t, tt.Expected, tt.User.PasswordReset())
		})
	}
}

func TestUser_SetVerification(t *testing.T) {
	input := &User{}
	v := &Verification{
		verified:   false,
		code:       "xxx",
		expiration: time.Time{},
	}
	input.SetVerification(v)
	assert.Equal(t, v, input.verification)
}

func TestUser_Verification(t *testing.T) {
	v := NewVerification()
	tests := []struct {
		name         string
		verification *Verification
		want         *Verification
	}{
		{
			name:         "should return the same verification",
			verification: v,
			want:         v,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				verification: tt.verification,
			}
			assert.Equal(t, tt.want, u.Verification())
		})
	}
}
