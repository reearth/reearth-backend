package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *mutationResolver) CreateAsset(ctx context.Context, input CreateAssetInput) (*CreateAssetPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Asset.Create(ctx, interfaces.CreateAssetParam{
		TeamID: id.TeamID(input.TeamID),
		File:   fromFile(&input.File),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &CreateAssetPayload{Asset: toAsset(res)}, nil
}

func (r *mutationResolver) RemoveAsset(ctx context.Context, input RemoveAssetInput) (*RemoveAssetPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err2 := r.usecases.Asset.Remove(ctx, id.AssetID(input.AssetID), getOperator(ctx))
	if err2 != nil {
		return nil, err2
	}

	return &RemoveAssetPayload{AssetID: res.ID()}, nil
}
