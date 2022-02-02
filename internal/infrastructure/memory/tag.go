package memory

import (
	"context"
	"sync"

	"github.com/reearth/reearth-backend/pkg/rerror"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
)

type Tag struct {
	lock   sync.Mutex
	data   tag.Map
	filter *id.SceneIDSet
}

func NewTag() repo.Tag {
	return &Tag{
		data: tag.Map{},
	}
}

func (r *Tag) Filtered(scenes []id.SceneID) repo.Tag {
	return &Tag{
		data:   r.data,
		filter: id.NewSceneIDSet(scenes...),
	}
}

func (r *Tag) FindByID(ctx context.Context, tagID id.TagID, ids []id.SceneID) (tag.Tag, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	res, ok := r.data[tagID]
	if ok && isSceneIncludes(res.Scene(), ids) {
		return res, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Tag) FindByIDs(ctx context.Context, tids []id.TagID, ids []id.SceneID) ([]*tag.Tag, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	var res []*tag.Tag
	for _, id := range tids {
		if d, ok := r.data[id]; ok {
			if isSceneIncludes(d.Scene(), ids) {
				res = append(res, &d)
				continue
			}
		}
		res = append(res, nil)
	}
	return res, nil
}

func (r *Tag) FindByScene(ctx context.Context, sceneID id.SceneID) ([]*tag.Tag, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	return r.data.All().FilterByScene(sceneID).Refs(), nil
}

func (r *Tag) FindItemByID(ctx context.Context, tagID id.TagID, ids []id.SceneID) (*tag.Item, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if d, ok := r.data[tagID]; ok {
		if res := tag.ItemFrom(d); res != nil {
			if isSceneIncludes(res.Scene(), ids) {
				return res, nil
			}
		}
	}
	return nil, rerror.ErrNotFound
}

func (r *Tag) FindItemByIDs(ctx context.Context, tagIDs []id.TagID, ids []id.SceneID) ([]*tag.Item, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	var res []*tag.Item
	for _, id := range tagIDs {
		if d, ok := r.data[id]; ok {
			if ti := tag.ItemFrom(d); ti != nil {
				if isSceneIncludes(ti.Scene(), ids) {
					res = append(res, ti)
				}
			}
		}
	}
	return res, nil
}

func (r *Tag) FindGroupByID(ctx context.Context, tagID id.TagID, ids []id.SceneID) (*tag.Group, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if d, ok := r.data[tagID]; ok {
		if res := tag.GroupFrom(d); res != nil {
			if isSceneIncludes(res.Scene(), ids) {
				return res, nil
			}
		}
	}
	return nil, rerror.ErrNotFound
}

func (r *Tag) FindGroupByIDs(ctx context.Context, tagIDs []id.TagID, ids []id.SceneID) ([]*tag.Group, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	var res []*tag.Group
	for _, id := range tagIDs {
		if d, ok := r.data[id]; ok {
			if tg := tag.GroupFrom(d); tg != nil {
				if isSceneIncludes(tg.Scene(), ids) {
					res = append(res, tg)
				}
			}
		}
	}
	return res, nil
}

func (r *Tag) FindRootsByScene(ctx context.Context, sceneID id.SceneID) ([]*tag.Tag, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	return r.data.All().FilterByScene(sceneID).Roots().Refs(), nil
}

func (r *Tag) FindGroupByItem(ctx context.Context, tagID id.TagID, s []id.SceneID) (*tag.Group, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, tg := range r.data {
		if res := tag.GroupFrom(tg); res != nil {
			tags := res.Tags()
			for _, item := range tags.Tags() {
				if item == tagID {
					return res, nil
				}
			}
		}
	}

	return nil, rerror.ErrNotFound
}

func (r *Tag) Save(ctx context.Context, d tag.Tag) error {
	if d == nil {
		return nil
	}
	if !r.ok(d) {
		return repo.ErrOperationDenied
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	r.data[d.ID()] = d
	return nil
}

func (r *Tag) SaveAll(ctx context.Context, list []*tag.Tag) error {
	for _, d := range list {
		if d != nil && !r.ok(*d) {
			return repo.ErrOperationDenied
		}
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	for _, d := range list {
		if d == nil {
			continue
		}
		d := *d
		if d == nil {
			continue
		}
		r.data[d.ID()] = d
	}
	return nil
}

func (r *Tag) Remove(ctx context.Context, id id.TagID) error {
	if t, ok := r.data[id]; !ok || !r.ok(t) {
		return repo.ErrNotFound
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.data, id)
	return nil
}

func (r *Tag) RemoveAll(ctx context.Context, ids []id.TagID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, id := range ids {
		if d, ok := r.data[id]; !ok || !r.ok(d) {
			continue
		}
		delete(r.data, id)
	}
	return nil
}

func (r *Tag) RemoveByScene(ctx context.Context, sceneID id.SceneID) error {
	if r.filter != nil && !r.filter.Has(sceneID) {
		return repo.ErrOperationDenied
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	for tid, v := range r.data {
		if v.Scene() == sceneID {
			delete(r.data, tid)
		}
	}
	return nil
}

func (r *Tag) ok(d tag.Tag) bool {
	return r.filter == nil || d != nil && r.filter.Has(d.Scene())
}

func (r *Tag) applyFilter(list tag.List) tag.List {
	if len(list) == 0 {
		return nil
	}

	res := make(tag.List, 0, len(list))
	for _, e := range list {
		if r.ok(e) {
			res = append(res, e)
		}
	}

	return res
}
