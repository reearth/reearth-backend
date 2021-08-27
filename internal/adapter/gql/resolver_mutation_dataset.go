package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *mutationResolver) UpdateDatasetSchema(ctx context.Context, input UpdateDatasetSchemaInput) (*UpdateDatasetSchemaPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Dataset.UpdateDatasetSchema(ctx, interfaces.UpdateDatasetSchemaParam{
		SchemaId: id.DatasetSchemaID(input.SchemaID),
		Name:     input.Name,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &UpdateDatasetSchemaPayload{DatasetSchema: toDatasetSchema(res)}, nil
}

func (r *mutationResolver) AddDynamicDatasetSchema(ctx context.Context, input AddDynamicDatasetSchemaInput) (*AddDynamicDatasetSchemaPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Dataset.AddDynamicDatasetSchema(ctx, interfaces.AddDynamicDatasetSchemaParam{
		SceneId: id.SceneID(input.SceneID),
	})
	if err != nil {
		return nil, err
	}

	return &AddDynamicDatasetSchemaPayload{DatasetSchema: toDatasetSchema(res)}, nil
}

func (r *mutationResolver) AddDynamicDataset(ctx context.Context, input AddDynamicDatasetInput) (*AddDynamicDatasetPayload, error) {
	exit := trace(ctx)
	defer exit()

	dss, ds, err := r.usecases.Dataset.AddDynamicDataset(ctx, interfaces.AddDynamicDatasetParam{
		SchemaId: id.DatasetSchemaID(input.DatasetSchemaID),
		Author:   input.Author,
		Content:  input.Content,
		Lat:      input.Lat,
		Lng:      input.Lng,
		Target:   input.Target,
	})
	if err != nil {
		return nil, err
	}

	return &AddDynamicDatasetPayload{DatasetSchema: toDatasetSchema(dss), Dataset: toDataset(ds)}, nil
}

func (r *mutationResolver) SyncDataset(ctx context.Context, input SyncDatasetInput) (*SyncDatasetPayload, error) {
	exit := trace(ctx)
	defer exit()

	dss, ds, err := r.usecases.Dataset.Sync(ctx, id.SceneID(input.SceneID), input.URL, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	schemas := make([]*DatasetSchema, 0, len(dss))
	datasets := make([]*Dataset, 0, len(ds))
	for _, d := range dss {
		schemas = append(schemas, toDatasetSchema(d))
	}
	for _, d := range ds {
		datasets = append(datasets, toDataset(d))
	}

	return &SyncDatasetPayload{
		SceneID:       input.SceneID,
		URL:           input.URL,
		DatasetSchema: schemas,
		Dataset:       datasets,
	}, nil
}

func (r *mutationResolver) RemoveDatasetSchema(ctx context.Context, input RemoveDatasetSchemaInput) (*RemoveDatasetSchemaPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Dataset.RemoveDatasetSchema(ctx, interfaces.RemoveDatasetSchemaParam{
		SchemaId: id.DatasetSchemaID(input.SchemaID),
		Force:    input.Force,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &RemoveDatasetSchemaPayload{SchemaID: res.ID()}, nil
}

func (r *mutationResolver) AddDatasetSchema(ctx context.Context, input AddDatasetSchemaInput) (*AddDatasetSchemaPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err2 := r.usecases.Dataset.AddDatasetSchema(ctx, interfaces.AddDatasetSchemaParam{
		SceneId:             id.SceneID(input.SceneID),
		Name:                input.Name,
		RepresentativeField: id.DatasetSchemaFieldIDFromRefID(input.Representativefield),
	}, getOperator(ctx))
	if err2 != nil {
		return nil, err2
	}

	return &AddDatasetSchemaPayload{DatasetSchema: toDatasetSchema(res)}, nil
}

func (r *mutationResolver) ImportDataset(ctx context.Context, input ImportDatasetInput) (*ImportDatasetPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Dataset.ImportDataset(ctx, interfaces.ImportDatasetParam{
		SceneId:  id.SceneID(input.SceneID),
		SchemaId: id.DatasetSchemaIDFromRefID(input.DatasetSchemaID),
		File:     fromFile(&input.File),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &ImportDatasetPayload{DatasetSchema: toDatasetSchema(res)}, nil
}

func (r *mutationResolver) ImportDatasetFromGoogleSheet(ctx context.Context, input ImportDatasetFromGoogleSheetInput) (*ImportDatasetPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Dataset.ImportDatasetFromGoogleSheet(ctx, interfaces.ImportDatasetFromGoogleSheetParam{
		Token:     input.AccessToken,
		FileID:    input.FileID,
		SheetName: input.SheetName,
		SceneId:   id.SceneID(input.SceneID),
		SchemaId:  id.DatasetSchemaIDFromRefID(input.DatasetSchemaID),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &ImportDatasetPayload{DatasetSchema: toDatasetSchema(res)}, nil
}
