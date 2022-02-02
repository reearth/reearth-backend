package memory

import (
	"context"
	"sync"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/reearth/reearth-backend/pkg/rerror"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
)

type Property struct {
	lock   sync.Mutex
	data   property.Map
	filter *id.SceneIDSet
}

func NewProperty() repo.Property {
	return &Property{
		data: property.Map{},
	}
}
func (r *Property) Filtered(scenes []id.SceneID) repo.Property {
	return &Property{
		data:   r.data,
		filter: id.NewSceneIDSet(scenes...),
	}
}

func (r *Property) FindByID(ctx context.Context, id id.PropertyID, f []id.SceneID) (*property.Property, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	p, ok := r.data[id]
	if ok && isSceneIncludes(p.Scene(), f) {
		return p, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Property) FindByIDs(ctx context.Context, ids []id.PropertyID, f []id.SceneID) (property.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := property.List{}
	for _, id := range ids {
		if p, ok := r.data[id]; ok {
			if isSceneIncludes(p.Scene(), f) {
				result = append(result, p)
				continue
			}
		}
		result = append(result, nil)
	}
	return result, nil
}

func (r *Property) FindByDataset(ctx context.Context, sid id.DatasetSchemaID, did id.DatasetID) (property.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := property.List{}
	for _, p := range r.data {
		if p.IsDatasetLinked(sid, did) {
			result = append(result, p)
		}
	}
	return result, nil
}

func (r *Property) FindLinkedAll(ctx context.Context, s id.SceneID) (property.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := property.List{}
	for _, p := range r.data {
		if p.Scene() != s {
			continue
		}
		if p.HasLinkedField() {
			result = append(result, p)
		}
	}
	return result, nil
}

func (r *Property) Save(ctx context.Context, d *property.Property) error {
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

func (r *Property) SaveAll(ctx context.Context, list property.List) error {
	for _, d := range list {
		if !r.ok(d) {
			return repo.ErrOperationDenied
		}
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	for _, d := range list {
		if d == nil {
			continue
		}
		r.data[d.ID()] = d
	}
	return nil
}

func (r *Property) Remove(ctx context.Context, id id.PropertyID) error {
	if d, ok := r.data[id]; !ok || !r.ok(d) {
		return repo.ErrNotFound
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.data, id)
	return nil
}

func (r *Property) RemoveAll(ctx context.Context, ids []id.PropertyID) error {
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

func (r *Property) RemoveByScene(ctx context.Context, sceneID id.SceneID) error {
	if r.filter != nil && !r.filter.Has(sceneID) {
		return repo.ErrOperationDenied
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	for pid, p := range r.data {
		if p.Scene() == sceneID {
			delete(r.data, pid)
		}
	}
	return nil
}

func (r *Property) ok(d *property.Property) bool {
	return r.filter == nil || d != nil && r.filter.Has(d.Scene())
}

func (r *Property) applyFilter(list property.List) property.List {
	if len(list) == 0 {
		return nil
	}

	res := make(property.List, 0, len(list))
	for _, e := range list {
		if r.ok(e) {
			res = append(res, e)
		}
	}

	return res
}
