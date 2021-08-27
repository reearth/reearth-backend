package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *Resolver) DatasetSchema() DatasetSchemaResolver {
	return &datasetSchemaResolver{r}
}

func (r *Resolver) DatasetSchemaField() DatasetSchemaFieldResolver {
	return &datasetSchemaFieldResolver{r}
}

type datasetSchemaResolver struct{ *Resolver }

func (r *datasetSchemaResolver) Scene(ctx context.Context, obj *DatasetSchema) (*Scene, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Scene.Load(id.SceneID(obj.SceneID))
}

func (r *datasetSchemaResolver) RepresentativeField(ctx context.Context, obj *DatasetSchema) (*DatasetSchemaField, error) {
	exit := trace(ctx)
	defer exit()

	if obj.RepresentativeFieldID == nil {
		return nil, nil
	}
	nf := *obj.RepresentativeFieldID
	for _, f := range obj.Fields {
		if f.ID == nf {
			return f, nil
		}
	}
	return nil, nil
}

func (r *datasetSchemaResolver) Datasets(ctx context.Context, obj *DatasetSchema, first *int, last *int, after *usecase.Cursor, before *usecase.Cursor) (*DatasetConnection, error) {
	exit := trace(ctx)
	defer exit()

	return r.controllers.Dataset.FindBySchema(ctx, obj.ID, first, last, before, after, getOperator(ctx))
}

type datasetSchemaFieldResolver struct{ *Resolver }

func (r *datasetSchemaFieldResolver) Schema(ctx context.Context, obj *DatasetSchemaField) (*DatasetSchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).DatasetSchema.Load(id.DatasetSchemaID(obj.SchemaID))
}

func (r *datasetSchemaFieldResolver) Ref(ctx context.Context, obj *DatasetSchemaField) (*DatasetSchema, error) {
	exit := trace(ctx)
	defer exit()

	if obj.RefID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).DatasetSchema.Load(id.DatasetSchemaID(*obj.RefID))
}
