package repo

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/asset"
	"github.com/reearth/reearth-backend/pkg/id"
)

type Asset interface {
	Save(context.Context, *asset.Asset) error
	Remove(context.Context, id.AssetID) error
	FindByTeam(context.Context, id.TeamID, *interfaces.AssetFilterType, *usecase.Pagination) ([]*asset.Asset, *usecase.PageInfo, error)
	FindByID(context.Context, id.AssetID, []id.TeamID) (*asset.Asset, error)
	FindByIDs(context.Context, []id.AssetID, []id.TeamID) ([]*asset.Asset, error)
}
