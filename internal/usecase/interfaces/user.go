package interfaces

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/pkg/user"
	"golang.org/x/text/language"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
)

var (
	ErrUserInvalidPasswordConfirmation = errors.New("invalid password confirmation")
	ErrUserInvalidLang                 = errors.New("invalid lang")
	ErrSignupInvalidSecret             = errors.New("invalid secret")
	ErrInvalidUserCredentials          = errors.New("invalid credentials")
)

type SignupParam struct {
	Sub    string
	Lang   *language.Tag
	Theme  *user.Theme
	UserID *id.UserID
	TeamID *id.TeamID
	Secret string
}

type GetUserByCredentials struct {
	Email    string
	Password string
}

type UpdateMeParam struct {
	Name                 *string
	Email                *string
	Lang                 *language.Tag
	Theme                *user.Theme
	Password             *string
	PasswordConfirmation *string
}

type User interface {
	Fetch(context.Context, []id.UserID, *usecase.Operator) ([]*user.User, error)
	Signup(context.Context, SignupParam) (*user.User, *user.Team, error)
	GetUserByCredentials(context.Context, GetUserByCredentials) (*user.User, error)
	GetUserBySubject(context.Context, string, *usecase.Operator) (*user.User, error)
	UpdateMe(context.Context, UpdateMeParam, *usecase.Operator) (*user.User, error)
	RemoveMyAuth(context.Context, string, *usecase.Operator) (*user.User, error)
	SearchUser(context.Context, string, *usecase.Operator) (*user.User, error)
	DeleteMe(context.Context, id.UserID, *usecase.Operator) error
}
