package interfaces

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/asset"
	"github.com/reearth/reearth-backend/pkg/file"
	"github.com/reearth/reearth-backend/pkg/id"
)

type CreateAssetParam struct {
	TeamID id.TeamID
	File   *file.File
}

var (
	ErrCreateAssetFailed error = errors.New("failed to create asset")
)

type Asset interface {
	Create(context.Context, CreateAssetParam, *usecase.Operator) (*asset.Asset, error)
	Remove(context.Context, id.AssetID, *usecase.Operator) (id.AssetID, error)
	FindByTeam(context.Context, id.TeamID, *usecase.Pagination, *usecase.Operator) ([]*asset.Asset, *usecase.PageInfo, error)
}