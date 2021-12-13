package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *mutationResolver) Signup(ctx context.Context, input gqlmodel.SignupInput) (*gqlmodel.SignupPayload, error) {
	exit := trace(ctx)
	defer exit()

	secret := ""
	if input.Secret != nil {
		secret = *input.Secret
	}
	sub := getSub(ctx)
	u, team, err := r.usecases.User.Signup(ctx, interfaces.SignupParam{
		Sub:    &sub,
		Lang:   input.Lang,
		Theme:  gqlmodel.ToTheme(input.Theme),
		UserID: id.UserIDFromRefID(input.UserID),
		TeamID: id.TeamIDFromRefID(input.TeamID),
		Secret: &secret,
	})
	if err != nil {
		return nil, err
	}

	return &gqlmodel.SignupPayload{User: gqlmodel.ToUser(u), Team: gqlmodel.ToTeam(team)}, nil
}

func (r *mutationResolver) UpdateMe(ctx context.Context, input gqlmodel.UpdateMeInput) (*gqlmodel.UpdateMePayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.User.UpdateMe(ctx, interfaces.UpdateMeParam{
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
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.User.RemoveMyAuth(ctx, input.Auth, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.UpdateMePayload{User: gqlmodel.ToUser(res)}, nil
}

func (r *mutationResolver) DeleteMe(ctx context.Context, input gqlmodel.DeleteMeInput) (*gqlmodel.DeleteMePayload, error) {
	exit := trace(ctx)
	defer exit()

	if err := r.usecases.User.DeleteMe(ctx, id.UserID(input.UserID), getOperator(ctx)); err != nil {
		return nil, err
	}

	return &gqlmodel.DeleteMePayload{UserID: input.UserID}, nil
}
