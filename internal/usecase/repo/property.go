package repo

import (
	"context"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
)

type Property interface {
	FindByID(context.Context, id.PropertyID, []id.SceneID) (*property.Property, error)
	FindByIDs(context.Context, []id.PropertyID, []id.SceneID) (property.List, error)
	FindLinkedAll(context.Context, id.SceneID) (property.List, error)
	FindByDataset(context.Context, id.DatasetSchemaID, id.DatasetID) (property.List, error)
	FindBySchema(context.Context, []id.PropertySchemaID, id.SceneID) (property.List, error)
	FindByPlugin(context.Context, id.PluginID, id.SceneID) (property.List, error)
	Save(context.Context, *property.Property) error
	SaveAll(context.Context, property.List) error
	UpdateSchemaPlugin(context.Context, id.PluginID, id.PluginID, id.SceneID) error
	Remove(context.Context, id.PropertyID) error
	RemoveAll(context.Context, []id.PropertyID) error
	RemoveByScene(context.Context, id.SceneID) error
}

func PropertyLoaderFrom(r Property, scenes []id.SceneID) property.Loader {
	return func(ctx context.Context, ids ...id.PropertyID) (property.List, error) {
		return r.FindByIDs(ctx, ids, scenes)
	}
}
