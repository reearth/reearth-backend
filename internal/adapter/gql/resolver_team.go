package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *Resolver) Team() TeamResolver {
	return &teamResolver{r}
}

func (r *Resolver) TeamMember() TeamMemberResolver {
	return &teamMemberResolver{r}
}

type teamResolver struct{ *Resolver }

func (r *teamResolver) Assets(ctx context.Context, obj *Team, first *int, last *int, after *usecase.Cursor, before *usecase.Cursor) (*AssetConnection, error) {
	exit := trace(ctx)
	defer exit()

	return r.controllers.Asset.FindByTeam(ctx, obj.ID, first, last, before, after, getOperator(ctx))
}

func (r *teamResolver) Projects(ctx context.Context, obj *Team, includeArchived *bool, first *int, last *int, after *usecase.Cursor, before *usecase.Cursor) (*ProjectConnection, error) {
	exit := trace(ctx)
	defer exit()

	return r.controllers.Project.FindByTeam(ctx, id.TeamID(obj.ID), first, last, before, after, getOperator(ctx))
}

type teamMemberResolver struct{ *Resolver }

func (r *teamMemberResolver) User(ctx context.Context, obj *TeamMember) (*User, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).User.Load(id.UserID(obj.UserID))
}
