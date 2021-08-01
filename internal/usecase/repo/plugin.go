package repo

import (
	"context"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
)

type Plugin interface {
	FindByID(context.Context, id.PluginID, []id.SceneID) (*plugin.Plugin, error)
	FindByIDs(context.Context, []id.PluginID, []id.SceneID) ([]*plugin.Plugin, error)
	Save(context.Context, *plugin.Plugin) error
	Remove(context.Context, id.PluginID) error
}

func PluginLoaderFrom(r Plugin) plugin.Loader {
	return func(ctx context.Context, ids []id.PluginID, sids []id.SceneID) ([]*plugin.Plugin, error) {
		return r.FindByIDs(ctx, ids, sids)
	}
}
