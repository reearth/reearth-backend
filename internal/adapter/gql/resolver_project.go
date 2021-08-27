package gql

import (
	"context"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

func (r *Resolver) Project() ProjectResolver {
	return &projectResolver{r}
}

type projectResolver struct{ *Resolver }

func (r *projectResolver) Team(ctx context.Context, obj *Project) (*Team, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Team.Load(id.TeamID(obj.TeamID))
}

func (r *projectResolver) Scene(ctx context.Context, obj *Project) (*Scene, error) {
	exit := trace(ctx)
	defer exit()

	s, err := r.controllers.Scene.FindByProject(ctx, id.ProjectID(obj.ID), getOperator(ctx))
	if err != nil && err != rerror.ErrNotFound {
		return nil, err
	}
	return s, nil
}
