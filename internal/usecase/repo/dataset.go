package repo

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
)

type Dataset interface {
	Filtered(SceneFilter) Dataset
	FindByID(context.Context, id.DatasetID, []id.SceneID) (*dataset.Dataset, error)
	FindByIDs(context.Context, []id.DatasetID, []id.SceneID) (dataset.List, error)
	FindBySchema(context.Context, id.DatasetSchemaID, []id.SceneID, *usecase.Pagination) (dataset.List, *usecase.PageInfo, error)
	FindBySchemaAll(context.Context, id.DatasetSchemaID) (dataset.List, error)
	FindGraph(context.Context, id.DatasetID, []id.SceneID, []id.DatasetSchemaFieldID) (dataset.List, error)
	Save(context.Context, *dataset.Dataset) error
	SaveAll(context.Context, dataset.List) error
	Remove(context.Context, id.DatasetID) error
	RemoveAll(context.Context, []id.DatasetID) error
	RemoveByScene(context.Context, id.SceneID) error
}

func DatasetLoaderFrom(r Dataset, scenes []id.SceneID) dataset.Loader {
	return func(ctx context.Context, ids ...id.DatasetID) (dataset.List, error) {
		return r.FindByIDs(ctx, ids, scenes)
	}
}

func DatasetGraphLoaderFrom(r Dataset, scenes []id.SceneID) dataset.GraphLoader {
	return func(ctx context.Context, root id.DatasetID, fields ...id.DatasetSchemaFieldID) (dataset.List, *dataset.Field, error) {
		if len(fields) <= 1 {
			d, err := r.FindByID(ctx, root, scenes)
			if err != nil {
				return nil, nil, err
			}
			var field *dataset.Field
			if len(fields) == 1 {
				field = d.Field(fields[0])
			}
			return dataset.List{d}, field, nil
		}

		list2, err := r.FindGraph(ctx, root, scenes, fields)
		if err != nil {
			return nil, nil, err
		}
		return list2, list2.Last().Field(fields[len(fields)-1]), nil
	}
}
