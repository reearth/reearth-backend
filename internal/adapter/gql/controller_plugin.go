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

// data loader

type PluginDataLoader interface {
	Load(id.PluginID) (*Plugin, error)
	LoadAll([]id.PluginID) ([]*Plugin, []error)
}

func (c *PluginController) DataLoader(ctx context.Context) *PluginLoader {
	return NewPluginLoader(PluginLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.PluginID) ([]*Plugin, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *PluginController) OrdinaryDataLoader(ctx context.Context) PluginDataLoader {
	return &ordinaryPluginLoader{
		fetch: func(keys []id.PluginID) ([]*Plugin, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryPluginLoader struct {
	fetch func(keys []id.PluginID) ([]*Plugin, []error)
}

func (l *ordinaryPluginLoader) Load(key id.PluginID) (*Plugin, error) {
	res, errs := l.fetch([]id.PluginID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryPluginLoader) LoadAll(keys []id.PluginID) ([]*Plugin, []error) {
	return l.fetch(keys)
}
