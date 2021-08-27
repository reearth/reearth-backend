package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type AssetController struct {
	usecase interfaces.Asset
}

func NewAssetController(usecase interfaces.Asset) *AssetController {
	return &AssetController{usecase: usecase}
}

func (c *AssetController) Fetch(ctx context.Context, ids []id.AssetID) ([]*gqlmodel.Asset, []error) {
	res, err := c.usecase.Fetch(ctx, ids, getOperator(ctx))
	if err != nil {
		return nil, []error{err}
	}

	assets := make([]*gqlmodel.Asset, 0, len(res))
	for _, a := range res {
		assets = append(assets, gqlmodel.ToAsset(a))
	}

	return assets, nil
}

func (c *AssetController) FindByTeam(ctx context.Context, teamID id.ID, first *int, last *int, before *usecase.Cursor, after *usecase.Cursor) (*gqlmodel.AssetConnection, error) {
	p := usecase.NewPagination(first, last, before, after)
	assets, pi, err := c.usecase.FindByTeam(ctx, id.TeamID(teamID), p, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	edges := make([]*gqlmodel.AssetEdge, 0, len(assets))
	nodes := make([]*gqlmodel.Asset, 0, len(assets))
	for _, a := range assets {
		asset := gqlmodel.ToAsset(a)
		edges = append(edges, &gqlmodel.AssetEdge{
			Node:   asset,
			Cursor: usecase.Cursor(asset.ID.String()),
		})
		nodes = append(nodes, asset)
	}

	return &gqlmodel.AssetConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   gqlmodel.ToPageInfo(pi),
		TotalCount: pi.TotalCount(),
	}, nil
}

// data loader

type AssetDataLoader interface {
	Load(id.AssetID) (*gqlmodel.Asset, error)
	LoadAll([]id.AssetID) ([]*gqlmodel.Asset, []error)
}

func (c *AssetController) DataLoader(ctx context.Context) AssetDataLoader {
	return gqldataloader.NewAssetLoader(gqldataloader.AssetLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.AssetID) ([]*gqlmodel.Asset, []error) {
			return c.Fetch(ctx, keys)
		},
	})
}

func (c *AssetController) OrdinaryDataLoader(ctx context.Context) AssetDataLoader {
	return &ordinaryAssetLoader{ctx: ctx, c: c}
}

type ordinaryAssetLoader struct {
	ctx context.Context
	c   *AssetController
}

func (l *ordinaryAssetLoader) Load(key id.AssetID) (*gqlmodel.Asset, error) {
	res, errs := l.c.Fetch(l.ctx, []id.AssetID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryAssetLoader) LoadAll(keys []id.AssetID) ([]*gqlmodel.Asset, []error) {
	return l.c.Fetch(l.ctx, keys)
}
