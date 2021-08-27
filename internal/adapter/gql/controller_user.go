package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type UserController struct {
	usecase interfaces.User
}

func NewUserController(usecase interfaces.User) *UserController {
	return &UserController{usecase: usecase}
}

func (c *UserController) Fetch(ctx context.Context, ids []id.UserID, operator *usecase.Operator) ([]*User, []error) {
	res, err := c.usecase.Fetch(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	users := make([]*User, 0, len(res))
	for _, u := range res {
		users = append(users, ToUser(u))
	}

	return users, nil
}

func (c *UserController) SearchUser(ctx context.Context, nameOrEmail string, operator *usecase.Operator) (*SearchedUser, error) {
	res, err := c.usecase.SearchUser(ctx, nameOrEmail, operator)
	if err != nil {
		return nil, err
	}

	return toSearchedUser(res), nil
}
