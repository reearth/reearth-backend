package http

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type UserController struct {
	usecase interfaces.User
}

func NewUserController(usecase interfaces.User) *UserController {
	return &UserController{
		usecase: usecase,
	}
}

type PasswordResetRequestInput struct {
	Email *string `json:"email"`
}

type PasswordResetConfirmInput struct {
	Token    *string `json:"token"`
	Password *string `json:"password"`
}

type SignupInput struct {
	Sub      *string    `json:"sub"`
	Secret   *string    `json:"secret"`
	UserID   *id.UserID `json:"userId"`
	TeamID   *id.TeamID `json:"teamId"`
	Name     *string    `json:"username"`
	Email    *string    `json:"email"`
	Password *string    `json:"password"`
}

type SignupOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *UserController) Signup(ctx context.Context, input SignupInput) (interface{}, error) {
	u, _, err := c.usecase.Signup(ctx, interfaces.SignupParam{
		Sub:      input.Sub,
		Secret:   input.Secret,
		UserID:   input.UserID,
		TeamID:   input.TeamID,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return SignupOutput{
		ID:    u.ID().String(),
		Name:  u.Name(),
		Email: u.Email(),
	}, nil
}

func (c *UserController) PasswordResetRequest(ctx context.Context, input PasswordResetRequestInput) (interface{}, error) {
	ok, err := c.usecase.PasswordResetRequest(ctx, interfaces.PasswordResetRequestParam{
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}

	if ok {
		//	send password reset link via email
	}

	return true, nil
}

func (c *UserController) PasswordResetConfirm(ctx context.Context, input PasswordResetConfirmInput) (interface{}, error) {
	return c.usecase.PasswordResetConfirm(ctx, interfaces.PasswordResetConfirmParam{
		Token:    input.Token,
		Password: input.Password,
	})
}
