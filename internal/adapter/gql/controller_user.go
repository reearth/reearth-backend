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

// data loader

type UserDataLoader interface {
	Load(id.UserID) (*User, error)
	LoadAll([]id.UserID) ([]*User, []error)
}

func (c *UserController) DataLoader(ctx context.Context) UserDataLoader {
	return NewUserLoader(UserLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.UserID) ([]*User, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *UserController) OrdinaryDataLoader(ctx context.Context) UserDataLoader {
	return &ordinaryUserLoader{
		fetch: func(keys []id.UserID) ([]*User, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryUserLoader struct {
	fetch func(keys []id.UserID) ([]*User, []error)
}

func (l *ordinaryUserLoader) Load(key id.UserID) (*User, error) {
	res, errs := l.fetch([]id.UserID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryUserLoader) LoadAll(keys []id.UserID) ([]*User, []error) {
	return l.fetch(keys)
}
