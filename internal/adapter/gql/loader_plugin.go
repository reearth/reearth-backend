package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
)

type PluginLoader struct {
	r  repo.Plugin
	rr repo.PluginRegistry
}

func NewPluginLoader(r repo.Plugin, rr repo.PluginRegistry) *PluginLoader {
	return &PluginLoader{r: r, rr: rr}
}

func (c *PluginLoader) Fetch(ctx context.Context, ids []id.PluginID) ([]*gqlmodel.Plugin, []error) {
	res, err := c.r.FindByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	plugins := make([]*gqlmodel.Plugin, 0, len(res))
	for _, pl := range res {
		plugins = append(plugins, gqlmodel.ToPlugin(pl))
	}

	return plugins, nil
}

func (c *PluginLoader) FetchPluginMetadata(ctx context.Context) ([]*gqlmodel.PluginMetadata, error) {
	res, err := c.rr.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	pluginMetaList := make([]*gqlmodel.PluginMetadata, 0, len(res))
	for _, md := range res {
		pm, err := gqlmodel.ToPluginMetadata(md)
		if err != nil {
			return nil, err
		}
		pluginMetaList = append(pluginMetaList, pm)
	}

	return pluginMetaList, nil
}

// data loader

type PluginDataLoader interface {
	Load(id.PluginID) (*gqlmodel.Plugin, error)
	LoadAll([]id.PluginID) ([]*gqlmodel.Plugin, []error)
}

func (c *PluginLoader) DataLoader(ctx context.Context) PluginDataLoader {
	return gqldataloader.NewPluginLoader(gqldataloader.PluginLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.PluginID) ([]*gqlmodel.Plugin, []error) {
			return c.Fetch(ctx, keys)
		},
	})
}

func (c *PluginLoader) OrdinaryDataLoader(ctx context.Context) PluginDataLoader {
	return &ordinaryPluginLoader{
		fetch: func(keys []id.PluginID) ([]*gqlmodel.Plugin, []error) {
			return c.Fetch(ctx, keys)
		},
	}
}

type ordinaryPluginLoader struct {
	fetch func(keys []id.PluginID) ([]*gqlmodel.Plugin, []error)
}

func (l *ordinaryPluginLoader) Load(key id.PluginID) (*gqlmodel.Plugin, error) {
	res, errs := l.fetch([]id.PluginID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryPluginLoader) LoadAll(keys []id.PluginID) ([]*gqlmodel.Plugin, []error) {
	return l.fetch(keys)
}
