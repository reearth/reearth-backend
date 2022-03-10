package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
)

type LayerLoader struct {
	r repo.Layer
	u interfaces.Layer
}

func NewLayerLoader(r repo.Layer, u interfaces.Layer) *LayerLoader {
	return &LayerLoader{r: r, u: u}
}

func (c *LayerLoader) Fetch(ctx context.Context, ids []id.LayerID) ([]*gqlmodel.Layer, []error) {
	res, err := c.r.FindByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	layers := make([]*gqlmodel.Layer, 0, len(res))
	for _, l := range res {
		if l == nil {
			layers = append(layers, nil)
		} else {
			layer := gqlmodel.ToLayer(*l, nil)
			layers = append(layers, &layer)
		}
	}

	return layers, nil
}

func (c *LayerLoader) FetchGroup(ctx context.Context, ids []id.LayerID) ([]*gqlmodel.LayerGroup, []error) {
	res, err := c.r.FindGroupByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	layerGroups := make([]*gqlmodel.LayerGroup, 0, len(res))
	for _, l := range res {
		layerGroups = append(layerGroups, gqlmodel.ToLayerGroup(l, nil))
	}

	return layerGroups, nil
}

func (c *LayerLoader) FetchItem(ctx context.Context, ids []id.LayerID) ([]*gqlmodel.LayerItem, []error) {
	res, err := c.r.FindItemByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	layerItems := make([]*gqlmodel.LayerItem, 0, len(res))
	for _, l := range res {
		layerItems = append(layerItems, gqlmodel.ToLayerItem(l, nil))
	}

	return layerItems, nil
}

func (c *LayerLoader) FetchParent(ctx context.Context, lid id.LayerID) (*gqlmodel.LayerGroup, error) {
	res, err := c.r.FindParentByID(ctx, id.LayerID(lid))
	if err != nil {
		return nil, err
	}

	return gqlmodel.ToLayerGroup(res, nil), nil
}

func (c *LayerLoader) FetchByProperty(ctx context.Context, pid id.PropertyID) (gqlmodel.Layer, error) {
	res, err := c.r.FindByProperty(ctx, pid)
	if err != nil {
		return nil, err
	}

	return gqlmodel.ToLayer(res, nil), nil
}

func (c *LayerLoader) FetchMerged(ctx context.Context, org id.LayerID, parent *id.LayerID) (*gqlmodel.MergedLayer, error) {
	res, err2 := c.u.FetchMerged(ctx, org, parent, getOperator(ctx))
	if err2 != nil {
		return nil, err2
	}

	return gqlmodel.ToMergedLayer(res), nil
}

func (c *LayerLoader) FetchParentAndMerged(ctx context.Context, org id.LayerID) (*gqlmodel.MergedLayer, error) {
	res, err2 := c.u.FetchParentAndMerged(ctx, org, getOperator(ctx))
	if err2 != nil {
		return nil, err2
	}

	return gqlmodel.ToMergedLayer(res), nil
}

func (c *LayerLoader) FetchByTag(ctx context.Context, tag id.TagID) ([]gqlmodel.Layer, error) {
	res, err2 := c.r.FindByTag(ctx, tag)
	if err2 != nil {
		return nil, err2
	}

	layers := make([]gqlmodel.Layer, 0, len(res))
	for _, l := range res {
		if l == nil {
			layers = append(layers, nil)
		} else {
			layers = append(layers, gqlmodel.ToLayer(*l, nil))
		}
	}

	return layers, nil
}

// data loader

type LayerDataLoader interface {
	Load(id.LayerID) (*gqlmodel.Layer, error)
	LoadAll([]id.LayerID) ([]*gqlmodel.Layer, []error)
}

func (c *LayerLoader) DataLoader(ctx context.Context) LayerDataLoader {
	return gqldataloader.NewLayerLoader(gqldataloader.LayerLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.LayerID) ([]*gqlmodel.Layer, []error) {
			return c.Fetch(ctx, keys)
		},
	})
}

func (c *LayerLoader) OrdinaryDataLoader(ctx context.Context) LayerDataLoader {
	return &ordinaryLayerLoader{
		fetch: func(keys []id.LayerID) ([]*gqlmodel.Layer, []error) {
			return c.Fetch(ctx, keys)
		},
	}
}

type ordinaryLayerLoader struct {
	fetch func(keys []id.LayerID) ([]*gqlmodel.Layer, []error)
}

func (l *ordinaryLayerLoader) Load(key id.LayerID) (*gqlmodel.Layer, error) {
	res, errs := l.fetch([]id.LayerID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryLayerLoader) LoadAll(keys []id.LayerID) ([]*gqlmodel.Layer, []error) {
	return l.fetch(keys)
}

type LayerItemDataLoader interface {
	Load(id.LayerID) (*gqlmodel.LayerItem, error)
	LoadAll([]id.LayerID) ([]*gqlmodel.LayerItem, []error)
}

func (c *LayerLoader) ItemDataLoader(ctx context.Context) LayerItemDataLoader {
	return gqldataloader.NewLayerItemLoader(gqldataloader.LayerItemLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.LayerID) ([]*gqlmodel.LayerItem, []error) {
			return c.FetchItem(ctx, keys)
		},
	})
}

func (c *LayerLoader) ItemOrdinaryDataLoader(ctx context.Context) LayerItemDataLoader {
	return &ordinaryLayerItemLoader{
		fetch: func(keys []id.LayerID) ([]*gqlmodel.LayerItem, []error) {
			return c.FetchItem(ctx, keys)
		},
	}
}

type ordinaryLayerItemLoader struct {
	fetch func(keys []id.LayerID) ([]*gqlmodel.LayerItem, []error)
}

func (l *ordinaryLayerItemLoader) Load(key id.LayerID) (*gqlmodel.LayerItem, error) {
	res, errs := l.fetch([]id.LayerID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryLayerItemLoader) LoadAll(keys []id.LayerID) ([]*gqlmodel.LayerItem, []error) {
	return l.fetch(keys)
}

type LayerGroupDataLoader interface {
	Load(id.LayerID) (*gqlmodel.LayerGroup, error)
	LoadAll([]id.LayerID) ([]*gqlmodel.LayerGroup, []error)
}

func (c *LayerLoader) GroupDataLoader(ctx context.Context) LayerGroupDataLoader {
	return gqldataloader.NewLayerGroupLoader(gqldataloader.LayerGroupLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.LayerID) ([]*gqlmodel.LayerGroup, []error) {
			return c.FetchGroup(ctx, keys)
		},
	})
}

func (c *LayerLoader) GroupOrdinaryDataLoader(ctx context.Context) LayerGroupDataLoader {
	return &ordinaryLayerGroupLoader{
		fetch: func(keys []id.LayerID) ([]*gqlmodel.LayerGroup, []error) {
			return c.FetchGroup(ctx, keys)
		},
	}
}

type ordinaryLayerGroupLoader struct {
	fetch func(keys []id.LayerID) ([]*gqlmodel.LayerGroup, []error)
}

func (l *ordinaryLayerGroupLoader) Load(key id.LayerID) (*gqlmodel.LayerGroup, error) {
	res, errs := l.fetch([]id.LayerID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryLayerGroupLoader) LoadAll(keys []id.LayerID) ([]*gqlmodel.LayerGroup, []error) {
	return l.fetch(keys)
}
