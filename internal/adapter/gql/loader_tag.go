package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
)

type TagLoader struct {
	r repo.Tag
}

func NewTagLoader(r repo.Tag) *TagLoader {
	return &TagLoader{r: r}
}

func (c *TagLoader) Fetch(ctx context.Context, ids []id.TagID) ([]*gqlmodel.Tag, []error) {
	res, err := c.r.FindByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	tags := make([]*gqlmodel.Tag, 0, len(res))
	for _, t := range res {
		if t != nil {
			tag := gqlmodel.ToTag(*t)
			tags = append(tags, &tag)
		}
	}

	return tags, nil
}

func (c *TagLoader) FetchGroup(ctx context.Context, ids []id.TagID) ([]*gqlmodel.TagGroup, []error) {
	res, err := c.r.FindGroupByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	tagGroups := make([]*gqlmodel.TagGroup, 0, len(res))
	for _, t := range res {
		tg := gqlmodel.ToTagGroup(t)
		if tg != nil {
			tagGroups = append(tagGroups, tg)
		}
	}

	return tagGroups, nil
}

func (c *TagLoader) FetchItem(ctx context.Context, ids []id.TagID) ([]*gqlmodel.TagItem, []error) {
	res, err := c.r.FindItemByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	tagItems := make([]*gqlmodel.TagItem, 0, len(res))
	for _, t := range res {
		ti := gqlmodel.ToTagItem(t)
		if ti != nil {
			tagItems = append(tagItems, ti)
		}
	}

	return tagItems, nil
}

func (c *TagLoader) FetchByScene(ctx context.Context, s id.SceneID) ([]gqlmodel.Tag, error) {
	res, err := c.r.FindByScene(ctx, s)
	if err != nil {
		return nil, err
	}

	tags := make([]gqlmodel.Tag, 0, len(res))
	for _, t := range res {
		if t != nil {
			tag := gqlmodel.ToTag(*t)
			tags = append(tags, tag)
		}
	}

	return tags, nil
}

// data loaders

type TagDataLoader interface {
	Load(id.TagID) (*gqlmodel.Tag, error)
	LoadAll([]id.TagID) ([]*gqlmodel.Tag, []error)
}

func (c *TagLoader) DataLoader(ctx context.Context) TagDataLoader {
	return gqldataloader.NewTagLoader(gqldataloader.TagLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.TagID) ([]*gqlmodel.Tag, []error) {
			return c.Fetch(ctx, keys)
		},
	})
}

func (c *TagLoader) OrdinaryDataLoader(ctx context.Context) TagDataLoader {
	return &ordinaryTagLoader{
		fetch: func(keys []id.TagID) ([]*gqlmodel.Tag, []error) {
			return c.Fetch(ctx, keys)
		},
	}
}

type ordinaryTagLoader struct {
	fetch func(keys []id.TagID) ([]*gqlmodel.Tag, []error)
}

func (t *ordinaryTagLoader) Load(key id.TagID) (*gqlmodel.Tag, error) {
	res, errs := t.fetch([]id.TagID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (t *ordinaryTagLoader) LoadAll(keys []id.TagID) ([]*gqlmodel.Tag, []error) {
	return t.fetch(keys)
}

type TagItemDataLoader interface {
	Load(id.TagID) (*gqlmodel.TagItem, error)
	LoadAll([]id.TagID) ([]*gqlmodel.TagItem, []error)
}

func (c *TagLoader) ItemDataLoader(ctx context.Context) TagItemDataLoader {
	return gqldataloader.NewTagItemLoader(gqldataloader.TagItemLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.TagID) ([]*gqlmodel.TagItem, []error) {
			return c.FetchItem(ctx, keys)
		},
	})
}

func (c *TagLoader) ItemOrdinaryDataLoader(ctx context.Context) TagItemDataLoader {
	return &ordinaryTagItemLoader{
		fetch: func(keys []id.TagID) ([]*gqlmodel.TagItem, []error) {
			return c.FetchItem(ctx, keys)
		},
	}
}

type ordinaryTagItemLoader struct {
	fetch func(keys []id.TagID) ([]*gqlmodel.TagItem, []error)
}

func (t *ordinaryTagItemLoader) Load(key id.TagID) (*gqlmodel.TagItem, error) {
	res, errs := t.fetch([]id.TagID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (t *ordinaryTagItemLoader) LoadAll(keys []id.TagID) ([]*gqlmodel.TagItem, []error) {
	return t.fetch(keys)
}

type TagGroupDataLoader interface {
	Load(id.TagID) (*gqlmodel.TagGroup, error)
	LoadAll([]id.TagID) ([]*gqlmodel.TagGroup, []error)
}

func (c *TagLoader) GroupDataLoader(ctx context.Context) TagGroupDataLoader {
	return gqldataloader.NewTagGroupLoader(gqldataloader.TagGroupLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.TagID) ([]*gqlmodel.TagGroup, []error) {
			return c.FetchGroup(ctx, keys)
		},
	})
}

func (c *TagLoader) GroupOrdinaryDataLoader(ctx context.Context) TagGroupDataLoader {
	return &ordinaryTagGroupLoader{
		fetch: func(keys []id.TagID) ([]*gqlmodel.TagGroup, []error) {
			return c.FetchGroup(ctx, keys)
		},
	}
}

type ordinaryTagGroupLoader struct {
	fetch func(keys []id.TagID) ([]*gqlmodel.TagGroup, []error)
}

func (t *ordinaryTagGroupLoader) Load(key id.TagID) (*gqlmodel.TagGroup, error) {
	res, errs := t.fetch([]id.TagID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (t *ordinaryTagGroupLoader) LoadAll(keys []id.TagID) ([]*gqlmodel.TagGroup, []error) {
	return t.fetch(keys)
}
