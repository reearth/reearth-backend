package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *mutationResolver) CreateTagItem(ctx context.Context, input gqlmodel.CreateTagItemInput) (*gqlmodel.CreateTagItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	tag, err := r.usecases.Tag.CreateItem(ctx, interfaces.CreateTagItemParam{
		Label:                 input.Label,
		SceneID:               id.SceneID(input.SceneID),
		LinkedDatasetSchemaID: id.DatasetSchemaIDFromRefID(input.LinkedDatasetSchemaID),
		LinkedDatasetID:       id.DatasetIDFromRefID(input.LinkedDatasetID),
		LinkedDatasetField:    id.DatasetSchemaFieldIDFromRefID(input.LinkedDatasetField),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}
	return &gqlmodel.CreateTagItemPayload{
		Tag: gqlmodel.ToTagItem(tag),
	}, nil
}

func (r *mutationResolver) CreateTagGroup(ctx context.Context, input gqlmodel.CreateTagGroupInput) (*gqlmodel.CreateTagGroupPayload, error) {
	exit := trace(ctx)
	defer exit()

	tag, err := r.usecases.Tag.CreateGroup(ctx, interfaces.CreateTagGroupParam{
		Label:   input.Label,
		SceneID: id.SceneID(input.SceneID),
		Tags:    id.TagIDsFromIDRef(input.Tags),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}
	return &gqlmodel.CreateTagGroupPayload{
		Tag: gqlmodel.ToTagGroup(tag),
	}, nil
}
