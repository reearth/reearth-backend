package gql

import (
	"context"

	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *mutationResolver) CreateTeam(ctx context.Context, input CreateTeamInput) (*CreateTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Team.Create(ctx, input.Name, getUser(ctx).ID())
	if err != nil {
		return nil, err
	}

	return &CreateTeamPayload{Team: toTeam(res)}, nil
}

func (r *mutationResolver) DeleteTeam(ctx context.Context, input DeleteTeamInput) (*DeleteTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	if err := r.usecases.Team.Remove(ctx, id.TeamID(input.TeamID), getOperator(ctx)); err != nil {
		return nil, err
	}

	return &DeleteTeamPayload{TeamID: input.TeamID}, nil
}

func (r *mutationResolver) UpdateTeam(ctx context.Context, input UpdateTeamInput) (*UpdateTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Team.Update(ctx, id.TeamID(input.TeamID), input.Name, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &UpdateTeamPayload{Team: toTeam(res)}, nil
}

func (r *mutationResolver) AddMemberToTeam(ctx context.Context, input AddMemberToTeamInput) (*AddMemberToTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Team.AddMember(ctx, id.TeamID(input.TeamID), id.UserID(input.UserID), fromRole(input.Role), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &AddMemberToTeamPayload{Team: toTeam(res)}, nil
}

func (r *mutationResolver) RemoveMemberFromTeam(ctx context.Context, input RemoveMemberFromTeamInput) (*RemoveMemberFromTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Team.RemoveMember(ctx, id.TeamID(input.TeamID), id.UserID(input.UserID), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &RemoveMemberFromTeamPayload{Team: toTeam(res)}, nil
}

func (r *mutationResolver) UpdateMemberOfTeam(ctx context.Context, input UpdateMemberOfTeamInput) (*UpdateMemberOfTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Team.UpdateMember(ctx, id.TeamID(input.TeamID), id.UserID(input.UserID), fromRole(input.Role), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &UpdateMemberOfTeamPayload{Team: toTeam(res)}, nil
}
