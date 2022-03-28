package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter"
	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/user"
)

func (r *mutationResolver) Signup(ctx context.Context, input gqlmodel.SignupInput) (*gqlmodel.SignupPayload, error) {
	sub := adapter.Sub(ctx)
	accessToken := adapter.AccessToken(ctx)
	issuer := adapter.Issuer(ctx)

	var u *user.User
	var t *user.Team
	var err error

	if sub != "" && accessToken != "" && issuer != "" {
		u, t, err = usecases(ctx).User.SignupOIDC(ctx, interfaces.SignupOIDCParam{
			Sub:         sub,
			AccessToken: accessToken,
			Issuer:      issuer,
			// Email:       email,
			// Name:        name,
			Secret: input.Secret,
			User: interfaces.SignupUserParam{
				Lang:   input.Lang,
				Theme:  gqlmodel.ToTheme(input.Theme),
				UserID: id.UserIDFromRefID(input.UserID),
				TeamID: id.TeamIDFromRefID(input.TeamID),
			},
		})
	} else {
		return nil, interfaces.ErrOperationDenied
	}

	if err != nil {
		return nil, err
	}

	return &gqlmodel.SignupPayload{User: gqlmodel.ToUser(u), Team: gqlmodel.ToTeam(t)}, nil
}

func (r *mutationResolver) UpdateMe(ctx context.Context, input gqlmodel.UpdateMeInput) (*gqlmodel.UpdateMePayload, error) {
	res, err := usecases(ctx).User.UpdateMe(ctx, interfaces.UpdateMeParam{
		Name:                 input.Name,
		Email:                input.Email,
		Lang:                 input.Lang,
		Theme:                gqlmodel.ToTheme(input.Theme),
		Password:             input.Password,
		PasswordConfirmation: input.PasswordConfirmation,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.UpdateMePayload{User: gqlmodel.ToUser(res)}, nil
}

func (r *mutationResolver) RemoveMyAuth(ctx context.Context, input gqlmodel.RemoveMyAuthInput) (*gqlmodel.UpdateMePayload, error) {
	res, err := usecases(ctx).User.RemoveMyAuth(ctx, input.Auth, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.UpdateMePayload{User: gqlmodel.ToUser(res)}, nil
}

func (r *mutationResolver) DeleteMe(ctx context.Context, input gqlmodel.DeleteMeInput) (*gqlmodel.DeleteMePayload, error) {
	if err := usecases(ctx).User.DeleteMe(ctx, id.UserID(input.UserID), getOperator(ctx)); err != nil {
		return nil, err
	}

	return &gqlmodel.DeleteMePayload{UserID: input.UserID}, nil
}
