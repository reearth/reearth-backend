package interactor

import (
	"context"
	"net/url"
	"path"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/executor"
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/asset"
	"github.com/reearth/reearth-backend/pkg/id"
)

type Asset struct {
	common
	e        *executor.Exectutor
	repos    *repo.Container
	gateways *gateway.Container
}

func NewAsset(r *repo.Container, g *gateway.Container) interfaces.Asset {
	return &Asset{
		e:        executor.New(r.Transaction),
		repos:    r,
		gateways: g,
	}
}

func (i *Asset) Create(ctx context.Context, inp interfaces.CreateAssetParam, operator *usecase.Operator) (result *asset.Asset, err error) {
	err = i.e.Run(ctx, func(ctx context.Context) error {
		if inp.File == nil {
			return interfaces.ErrFileNotIncluded
		}

		url, err := i.gateways.File.UploadAsset(ctx, inp.File)
		if err != nil {
			return err
		}

		result, err = asset.New().
			NewID().
			Team(inp.TeamID).
			Name(path.Base(inp.File.Path)).
			Size(inp.File.Size).
			URL(url.String()).
			Build()
		if err != nil {
			return err
		}

		if err := i.repos.Asset.Save(ctx, result); err != nil {
			return err
		}

		return nil
	}, i.e.CanWriteTeam(inp.TeamID, operator), i.e.Transaction())
	return
}

func (i *Asset) Remove(ctx context.Context, aid id.AssetID, operator *usecase.Operator) (result id.AssetID, err error) {
	err = i.e.Run(ctx, func(ctx context.Context) error {
		asset, err := i.repos.Asset.FindByID(ctx, aid)
		if err != nil {
			return err
		}
		if err := operator.CanWriteTeam(asset.Team()); err != nil {
			return err
		}

		if url, _ := url.Parse(asset.URL()); url != nil {
			if err = i.gateways.File.RemoveAsset(ctx, url); err != nil {
				return err
			}
		}

		if err = i.repos.Asset.Remove(ctx, aid); err != nil {
			return err
		}

		return nil
	}, i.e.Transaction())
	return
}
