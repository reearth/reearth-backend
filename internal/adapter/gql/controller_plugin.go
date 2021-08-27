package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type PluginController struct {
	usecase interfaces.Plugin
}

func NewPluginController(usecase interfaces.Plugin) *PluginController {
	return &PluginController{usecase: usecase}
}

func (c *PluginController) Fetch(ctx context.Context, ids []id.PluginID, operator *usecase.Operator) ([]*Plugin, []error) {
	res, err := c.usecase.Fetch(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	plugins := make([]*Plugin, 0, len(res))
	for _, pl := range res {
		plugins = append(plugins, toPlugin(pl))
	}

	return plugins, nil
}

func (c *PluginController) FetchPluginMetadata(ctx context.Context, operator *usecase.Operator) ([]*PluginMetadata, error) {
	res, err := c.usecase.FetchPluginMetadata(ctx, operator)
	if err != nil {
		return nil, err
	}

	pluginMetaList := make([]*PluginMetadata, 0, len(res))
	for _, md := range res {
		pm, err := toPluginMetadata(md)
		if err != nil {
			return nil, err
		}
		pluginMetaList = append(pluginMetaList, pm)
	}

	return pluginMetaList, nil
}
