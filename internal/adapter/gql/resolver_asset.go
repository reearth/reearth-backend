package gql

import (
	"context"

	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *Resolver) Asset() AssetResolver {
	return &assetResolver{r}
}

type assetResolver struct{ *Resolver }

func (r *assetResolver) Team(ctx context.Context, obj *Asset) (*Team, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Team.Load(id.TeamID(obj.TeamID))
}
