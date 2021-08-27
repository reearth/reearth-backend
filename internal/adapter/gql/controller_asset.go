package gql

import (
	"context"

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

func (c *AssetController) FindByTeam(ctx context.Context, teamID id.ID, first *int, last *int, before *usecase.Cursor, after *usecase.Cursor, operator *usecase.Operator) (*AssetConnection, error) {
	p := usecase.NewPagination(first, last, before, after)
	assets, pi, err := c.usecase.FindByTeam(ctx, id.TeamID(teamID), p, operator)
	if err != nil {
		return nil, err
	}

	edges := make([]*AssetEdge, 0, len(assets))
	nodes := make([]*Asset, 0, len(assets))
	for _, a := range assets {
		asset := toAsset(a)
		edges = append(edges, &AssetEdge{
			Node:   asset,
			Cursor: usecase.Cursor(asset.ID.String()),
		})
		nodes = append(nodes, asset)
	}

	return &AssetConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   toPageInfo(pi),
		TotalCount: pi.TotalCount(),
	}, nil
}
