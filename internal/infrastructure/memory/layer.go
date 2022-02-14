package memory

import (
	"context"
	"sync"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

type Layer struct {
	lock sync.Mutex
	data map[id.LayerID]layer.Layer
}

func NewLayer() repo.Layer {
	return &Layer{
		data: map[id.LayerID]layer.Layer{},
	}
}

func (r *Layer) FindByID(ctx context.Context, id id.LayerID, f []id.SceneID) (layer.Layer, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	res, ok := r.data[id]
	if ok && isSceneIncludes(res.Scene(), f) {
		return res, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Layer) FindByIDs(ctx context.Context, ids []id.LayerID, f []id.SceneID) (layer.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := layer.List{}
	for _, id := range ids {
		if d, ok := r.data[id]; ok {
			if isSceneIncludes(d.Scene(), f) {
				result = append(result, &d)
				continue
			}
		}
		result = append(result, nil)
	}
	return result, nil
}

func (r *Layer) FindGroupByIDs(ctx context.Context, ids []id.LayerID, f []id.SceneID) (layer.GroupList, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := layer.GroupList{}
	for _, id := range ids {
		if d, ok := r.data[id]; ok {
			if lg := layer.GroupFromLayer(d); lg != nil {
				if isSceneIncludes(lg.Scene(), f) {
					result = append(result, lg)
					continue
				}
			}
			result = append(result, nil)
		}
	}
	return result, nil
}

func (r *Layer) FindItemByIDs(ctx context.Context, ids []id.LayerID, f []id.SceneID) (layer.ItemList, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := layer.ItemList{}
	for _, id := range ids {
		if d, ok := r.data[id]; ok {
			if li := layer.ItemFromLayer(d); li != nil {
				if isSceneIncludes(li.Scene(), f) {
					result = append(result, li)
					continue
				}
			}
			result = append(result, nil)
		}
	}
	return result, nil
}

func (r *Layer) FindItemByID(ctx context.Context, id id.LayerID, f []id.SceneID) (*layer.Item, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	d, ok := r.data[id]
	if !ok {
		return &layer.Item{}, nil
	}
	if li := layer.ItemFromLayer(d); li != nil {
		if isSceneIncludes(li.Scene(), f) {
			return li, nil
		}
	}
	return nil, rerror.ErrNotFound
}

func (r *Layer) FindGroupByID(ctx context.Context, id id.LayerID, f []id.SceneID) (*layer.Group, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	d, ok := r.data[id]
	if !ok {
		return &layer.Group{}, nil
	}
	if lg := layer.GroupFromLayer(d); lg != nil {
		if isSceneIncludes(lg.Scene(), f) {
			return lg, nil
		}
	}
	return nil, rerror.ErrNotFound
}

func (r *Layer) FindGroupBySceneAndLinkedDatasetSchema(ctx context.Context, s id.SceneID, ds id.DatasetSchemaID) (layer.GroupList, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := layer.GroupList{}
	for _, l := range r.data {
		if l.Scene() != s {
			continue
		}
		if lg, ok := l.(*layer.Group); ok {
			if dsid := lg.LinkedDatasetSchema(); dsid != nil && *dsid == ds {
				result = append(result, lg)
			}
		}
	}
	return result, nil
}

func (r *Layer) FindParentsByIDs(_ context.Context, ids []id.LayerID, scenes []id.SceneID) (layer.GroupList, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	res := layer.GroupList{}
	for _, l := range r.data {
		if !isSceneIncludes(l.Scene(), scenes) {
			continue
		}
		gl, ok := l.(*layer.Group)
		if !ok {
			continue
		}
		for _, cl := range gl.Layers().Layers() {
			if cl.Contains(ids) {
				res = append(res, gl)
			}
		}
	}

	return res, nil
}

func (r *Layer) FindByPluginAndExtension(_ context.Context, pid id.PluginID, eid *id.PluginExtensionID, scenes []id.SceneID) (layer.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	res := layer.List{}
	for _, l := range r.data {
		l := l
		if !isSceneIncludes(l.Scene(), scenes) {
			continue
		}

		if p := l.Plugin(); p != nil && p.Equal(pid) {
			e := l.Extension()
			if eid == nil || e != nil && *e == *eid {
				res = append(res, &l)
			}
		}
	}

	return res, nil
}

func (r *Layer) FindByProperty(ctx context.Context, id id.PropertyID, f []id.SceneID) (layer.Layer, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, l := range r.data {
		if !isSceneIncludes(l.Scene(), f) {
			continue
		}
		if pid := l.Property(); pid != nil && *pid == id {
			return l, nil
		}
		if pid := l.Infobox().PropertyRef(); pid != nil && *pid == id {
			return l, nil
		}
		for _, f := range l.Infobox().Fields() {
			if f.Property() == id {
				return l, nil
			}
		}
	}
	return nil, rerror.ErrNotFound
}

func (r *Layer) FindParentByID(ctx context.Context, id id.LayerID, f []id.SceneID) (*layer.Group, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, l := range r.data {
		if !isSceneIncludes(l.Scene(), f) {
			continue
		}
		gl, ok := l.(*layer.Group)
		if !ok {
			continue
		}
		for _, cl := range gl.Layers().Layers() {
			if cl == id {
				return gl, nil
			}
		}
	}
	return nil, rerror.ErrNotFound
}

func (r *Layer) FindByScene(ctx context.Context, sceneID id.SceneID) (layer.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	res := layer.List{}
	for _, l := range r.data {
		l := l
		if l.Scene() == sceneID {
			res = append(res, &l)
		}
	}
	return res, nil
}

func (r *Layer) FindAllByDatasetSchema(ctx context.Context, datasetSchemaID id.DatasetSchemaID) (layer.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	res := layer.List{}
	for _, l := range r.data {
		if d := layer.ToLayerGroup(l).LinkedDatasetSchema(); d != nil && *d == datasetSchemaID {
			res = append(res, &l)
		}
	}
	return res, nil
}

func (r *Layer) Save(ctx context.Context, l layer.Layer) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.data[l.ID()] = l
	return nil
}

func (r *Layer) SaveAll(ctx context.Context, ll layer.List) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, l := range ll {
		layer := *l
		r.data[layer.ID()] = layer
	}
	return nil
}

func (r *Layer) UpdatePlugin(ctx context.Context, old id.PluginID, new id.PluginID, scenes []id.SceneID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, l := range r.data {
		p := l.Plugin()
		if p != nil && p.Equal(old) && isSceneIncludes(l.Scene(), scenes) {
			l.SetPlugin(&new)
			r.data[l.ID()] = l
		}
	}
	return nil
}

func (r *Layer) Remove(ctx context.Context, id id.LayerID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.data, id)
	return nil
}

func (r *Layer) RemoveAll(ctx context.Context, ids []id.LayerID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, id := range ids {
		delete(r.data, id)
	}
	return nil
}

func (r *Layer) RemoveByScene(ctx context.Context, sceneID id.SceneID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for lid, p := range r.data {
		if p.Scene() == sceneID {
			delete(r.data, lid)
		}
	}
	return nil
}

func (r *Layer) FindByTag(ctx context.Context, tagID id.TagID, s []id.SceneID) (layer.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	var res layer.List
	for _, l := range r.data {
		l := l
		if l.Tags().Has(tagID) {
			res = append(res, &l)
		}
	}

	return res, nil
}
