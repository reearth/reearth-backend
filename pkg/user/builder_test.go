package user

import (
	"errors"
	"testing"
	"time"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestBuilder_ID(t *testing.T) {
	uid := id.NewUserID()
	b := New().ID(uid).MustBuild()
	assert.Equal(t, uid, b.ID())
	assert.Nil(t, b.passwordReset)
}

func TestBuilder_Name(t *testing.T) {
	b := New().NewID().Name("xxx").MustBuild()
	assert.Equal(t, "xxx", b.Name())
}

func TestBuilder_NewID(t *testing.T) {
	b := New().NewID().MustBuild()
	assert.NotNil(t, b.ID())
}

func TestBuilder_Team(t *testing.T) {
	tid := id.NewTeamID()
	b := New().NewID().Team(tid).MustBuild()
	assert.Equal(t, tid, b.Team())
}

func TestBuilder_Auths(t *testing.T) {
	b := New().NewID().Auths([]Auth{
		{
			Provider: "xxx",
			Sub:      "aaa",
		},
	}).MustBuild()
	assert.Equal(t, []Auth{
		{
			Provider: "xxx",
			Sub:      "aaa",
		},
	}, b.Auths())
}

func TestBuilder_Email(t *testing.T) {
	b := New().NewID().Email("xx@yy.zz").MustBuild()
	assert.Equal(t, "xx@yy.zz", b.Email())
}

func TestBuilder_Lang(t *testing.T) {
	l := language.Make("en")
	b := New().NewID().Lang(l).MustBuild()
	assert.Equal(t, l, b.Lang())
}

func TestBuilder_LangFrom(t *testing.T) {
	testCases := []struct {
		Name, Lang string
		Expected   language.Tag
	}{
		{
			Name:     "success creating language",
			Lang:     "en",
			Expected: language.Make("en"),
		},
		{
			Name:     "empty language and empty tag",
			Lang:     "",
			Expected: language.Tag{},
		},
		{
			Name:     "empty tag of parse err",
			Lang:     "xxxxxxxxxxx",
			Expected: language.Tag{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			b := New().NewID().LangFrom(tc.Lang).MustBuild()
			assert.Equal(t, tc.Expected, b.Lang())
		})
	}
}

func TestBuilder_PasswordReset(t *testing.T) {
	testCases := []struct {
		Name, Token string
		CreatedAt   time.Time
		Expected    PasswordReset
	}{
		{
			Name:      "Test1",
			Token:     "xyz",
			CreatedAt: time.Unix(0, 0),
			Expected: PasswordReset{
				Token:     "xyz",
				CreatedAt: time.Unix(0, 0),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			// u := New().NewID().PasswordReset(tc.Token, tc.CreatedAt).MustBuild()
			// assert.Equal(t, tc.Expected, *u.passwordReset)
		})
	}
}

func TestNew(t *testing.T) {
	b := New()
	assert.NotNil(t, b)
	assert.IsType(t, &Builder{}, b)
}

func TestBuilder_Build(t *testing.T) {
	uid := id.NewUserID()
	tid := id.NewTeamID()
	pass, _ := encodePassword("pass")

	testCases := []struct {
		Name, UserName, Lang, Email string
		UID                         id.UserID
		TID                         id.TeamID
		Auths                       []Auth
		PasswordBin                 []byte
		Expected                    *User
		err                         error
	}{
		{
			Name:        "Success build user",
			UserName:    "xxx",
			Email:       "xx@yy.zz",
			Lang:        "en",
			UID:         uid,
			PasswordBin: pass,
			TID:         tid,
			Auths: []Auth{
				{
					Provider: "ppp",
					Sub:      "sss",
				},
			},
			Expected: &User{
				id:       uid,
				name:     "xxx",
				email:    "xx@yy.zz",
				password: pass,
				team:     tid,
				auths: []Auth{
					{
						Provider: "ppp",
						Sub:      "sss",
					},
				},
				lang: language.MustParse("en"),
			},
			err: nil,
		},
		{
			Name:     "failed invalid id",
			Expected: nil,
			err:      id.ErrInvalidID,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res, err := New().ID(tc.UID).Name(tc.UserName).Auths(tc.Auths).LangFrom(tc.Lang).Email(tc.Email).Password(tc.PasswordBin).Team(tc.TID).Build()
			if err == nil {
				assert.Equal(tt, tc.Expected, res)
			} else {
				assert.True(tt, errors.As(tc.err, &err))
			}
		})
	}
}

func TestBuilder_MustBuild(t *testing.T) {
	uid := id.NewUserID()
	tid := id.NewTeamID()
	pass, _ := encodePassword("pass")
	testCases := []struct {
		Name, UserName, Lang, Email string
		UID                         id.UserID
		TID                         id.TeamID
		PasswordBin                 []byte
		Auths                       []Auth
		Expected                    *User
		err                         error
	}{

		{
			Name:        "Success build user",
			UserName:    "xxx",
			Email:       "xx@yy.zz",
			Lang:        "en",
			UID:         uid,
			PasswordBin: pass,
			TID:         tid,
			Auths: []Auth{
				{
					Provider: "ppp",
					Sub:      "sss",
				},
			},
			Expected: &User{
				id:       uid,
				name:     "xxx",
				email:    "xx@yy.zz",
				password: pass,
				team:     tid,
				auths: []Auth{
					{
						Provider: "ppp",
						Sub:      "sss",
					},
				},
				lang: language.MustParse("en"),
			},
			err: nil,
		},
		{
			Name:     "failed invalid id",
			Expected: nil,
			err:      id.ErrInvalidID,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			var res *User
			defer func() {
				if r := recover(); r == nil {
					assert.Equal(tt, tc.Expected, res)
				}
			}()

			res = New().
				ID(tc.UID).
				Name(tc.UserName).
				Auths(tc.Auths).
				Password(tc.PasswordBin).
				LangFrom(tc.Lang).
				Email(tc.Email).
				Team(tc.TID).
				MustBuild()
		})
	}
}

func TestBuilder_Verification(t *testing.T) {
	tests := []struct {
		name  string
		input *Verification
		want  *Builder
	}{
		{
			name: "should return verification",
			input: &Verification{
				verified:   true,
				code:       "xxx",
				expiration: time.Time{},
			},

			want: &Builder{
				u: &User{
					verification: &Verification{
						verified:   true,
						code:       "xxx",
						expiration: time.Time{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New()
			b.Verification(tt.input)
			assert.Equal(t, tt.want, b)
		})
	}
}
