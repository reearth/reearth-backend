package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *mutationResolver) Signup(ctx context.Context, input SignupInput) (*SignupPayload, error) {
	exit := trace(ctx)
	defer exit()

	secret := ""
	if input.Secret != nil {
		secret = *input.Secret
	}

	u, team, err := r.usecases.User.Signup(ctx, interfaces.SignupParam{
		Sub:    getSub(ctx),
		Lang:   input.Lang,
		Theme:  toTheme(input.Theme),
		UserID: id.UserIDFromRefID(input.UserID),
		TeamID: id.TeamIDFromRefID(input.TeamID),
		Secret: secret,
	})
	if err != nil {
		return nil, err
	}

	return &SignupPayload{User: ToUser(u), Team: toTeam(team)}, nil
}

func (r *mutationResolver) UpdateMe(ctx context.Context, input UpdateMeInput) (*UpdateMePayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.User.UpdateMe(ctx, interfaces.UpdateMeParam{
		Name:                 input.Name,
		Email:                input.Email,
		Lang:                 input.Lang,
		Theme:                toTheme(input.Theme),
		Password:             input.Password,
		PasswordConfirmation: input.PasswordConfirmation,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &UpdateMePayload{User: ToUser(res)}, nil
}

func (r *mutationResolver) RemoveMyAuth(ctx context.Context, input RemoveMyAuthInput) (*UpdateMePayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.User.RemoveMyAuth(ctx, input.Auth, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &UpdateMePayload{User: ToUser(res)}, nil
}

func (r *mutationResolver) DeleteMe(ctx context.Context, input DeleteMeInput) (*DeleteMePayload, error) {
	exit := trace(ctx)
	defer exit()

	if err := r.usecases.User.DeleteMe(ctx, id.UserID(input.UserID), getOperator(ctx)); err != nil {
		return nil, err
	}

	return &DeleteMePayload{UserID: input.UserID}, nil
}
