package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type LayerController struct {
	usecase interfaces.Layer
}

func NewLayerController(usecase interfaces.Layer) *LayerController {
	return &LayerController{usecase: usecase}
}

func (c *LayerController) Fetch(ctx context.Context, ids []id.LayerID, operator *usecase.Operator) ([]*Layer, []error) {
	res, err := c.usecase.Fetch(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	layers := make([]*Layer, 0, len(res))
	for _, l := range res {
		if l == nil {
			layers = append(layers, nil)
		} else {
			layer := toLayer(*l, nil)
			layers = append(layers, &layer)
		}
	}

	return layers, nil
}

func (c *LayerController) FetchGroup(ctx context.Context, ids []id.LayerID, operator *usecase.Operator) ([]*LayerGroup, []error) {
	res, err := c.usecase.FetchGroup(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	layerGroups := make([]*LayerGroup, 0, len(res))
	for _, l := range res {
		layerGroups = append(layerGroups, toLayerGroup(l, nil))
	}

	return layerGroups, nil
}

func (c *LayerController) FetchItem(ctx context.Context, ids []id.LayerID, operator *usecase.Operator) ([]*LayerItem, []error) {
	res, err := c.usecase.FetchItem(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	layerItems := make([]*LayerItem, 0, len(res))
	for _, l := range res {
		layerItems = append(layerItems, toLayerItem(l, nil))
	}

	return layerItems, nil
}

func (c *LayerController) FetchParent(ctx context.Context, lid id.LayerID, operator *usecase.Operator) (*LayerGroup, error) {
	res, err := c.usecase.FetchParent(ctx, id.LayerID(lid), operator)
	if err != nil {
		return nil, err
	}

	return toLayerGroup(res, nil), nil
}

func (c *LayerController) FetchByProperty(ctx context.Context, pid id.PropertyID, operator *usecase.Operator) (Layer, error) {
	res, err := c.usecase.FetchByProperty(ctx, pid, operator)
	if err != nil {
		return nil, err
	}

	return toLayer(res, nil), nil
}

func (c *LayerController) FetchMerged(ctx context.Context, org id.LayerID, parent *id.LayerID, operator *usecase.Operator) (*MergedLayer, error) {
	res, err2 := c.usecase.FetchMerged(ctx, org, parent, operator)
	if err2 != nil {
		return nil, err2
	}

	return toMergedLayer(res), nil
}

func (c *LayerController) FetchParentAndMerged(ctx context.Context, org id.LayerID, operator *usecase.Operator) (*MergedLayer, error) {
	res, err2 := c.usecase.FetchParentAndMerged(ctx, org, operator)
	if err2 != nil {
		return nil, err2
	}

	return toMergedLayer(res), nil
}

// data loader

type LayerDataLoader interface {
	Load(id.LayerID) (*Layer, error)
	LoadAll([]id.LayerID) ([]*Layer, []error)
}

func (c *LayerController) DataLoader(ctx context.Context) LayerDataLoader {
	return NewLayerLoader(LayerLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.LayerID) ([]*Layer, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *LayerController) OrdinaryDataLoader(ctx context.Context) LayerDataLoader {
	return &ordinaryLayerLoader{
		fetch: func(keys []id.LayerID) ([]*Layer, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryLayerLoader struct {
	fetch func(keys []id.LayerID) ([]*Layer, []error)
}

func (l *ordinaryLayerLoader) Load(key id.LayerID) (*Layer, error) {
	res, errs := l.fetch([]id.LayerID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryLayerLoader) LoadAll(keys []id.LayerID) ([]*Layer, []error) {
	return l.fetch(keys)
}

type LayerItemDataLoader interface {
	Load(id.LayerID) (*LayerItem, error)
	LoadAll([]id.LayerID) ([]*LayerItem, []error)
}

func (c *LayerController) ItemDataLoader(ctx context.Context) LayerItemDataLoader {
	return NewLayerItemLoader(LayerItemLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.LayerID) ([]*LayerItem, []error) {
			return c.FetchItem(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *LayerController) ItemOrdinaryDataLoader(ctx context.Context) LayerItemDataLoader {
	return &ordinaryLayerItemLoader{
		fetch: func(keys []id.LayerID) ([]*LayerItem, []error) {
			return c.FetchItem(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryLayerItemLoader struct {
	fetch func(keys []id.LayerID) ([]*LayerItem, []error)
}

func (l *ordinaryLayerItemLoader) Load(key id.LayerID) (*LayerItem, error) {
	res, errs := l.fetch([]id.LayerID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryLayerItemLoader) LoadAll(keys []id.LayerID) ([]*LayerItem, []error) {
	return l.fetch(keys)
}

type LayerGroupDataLoader interface {
	Load(id.LayerID) (*LayerGroup, error)
	LoadAll([]id.LayerID) ([]*LayerGroup, []error)
}

func (c *LayerController) GroupDataLoader(ctx context.Context) LayerGroupDataLoader {
	return NewLayerGroupLoader(LayerGroupLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.LayerID) ([]*LayerGroup, []error) {
			return c.FetchGroup(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *LayerController) GroupOrdinaryDataLoader(ctx context.Context) LayerGroupDataLoader {
	return &ordinaryLayerGroupLoader{
		fetch: func(keys []id.LayerID) ([]*LayerGroup, []error) {
			return c.FetchGroup(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryLayerGroupLoader struct {
	fetch func(keys []id.LayerID) ([]*LayerGroup, []error)
}

func (l *ordinaryLayerGroupLoader) Load(key id.LayerID) (*LayerGroup, error) {
	res, errs := l.fetch([]id.LayerID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryLayerGroupLoader) LoadAll(keys []id.LayerID) ([]*LayerGroup, []error) {
	return l.fetch(keys)
}
