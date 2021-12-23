package memory

import (
	"context"
	"sort"
	"sync"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/reearth/reearth-backend/pkg/rerror"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
)

type Property struct {
	lock sync.Mutex
	data map[id.PropertyID]property.Property
}

func NewProperty() repo.Property {
	return &Property{
		data: map[id.PropertyID]property.Property{},
	}
}

func (r *Property) FindByID(ctx context.Context, id id.PropertyID, f []id.SceneID) (*property.Property, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	p, ok := r.data[id]
	if ok && isSceneIncludes(p.Scene(), f) {
		return &p, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Property) FindByIDs(ctx context.Context, ids []id.PropertyID, f []id.SceneID) (property.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := property.List{}
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

func (r *Property) FindByDataset(ctx context.Context, sid id.DatasetSchemaID, did id.DatasetID) (property.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := property.List{}
	for _, p := range r.data {
		if p.IsDatasetLinked(sid, did) {
			result = append(result, &p)
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
			p2 := p
			result = append(result, &p2)
		}
	}
	return result, nil
}

func (r *Property) FindBySchema(_ context.Context, schemas []id.PropertySchemaID, s id.SceneID) (property.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := property.List{}
	for _, p := range r.data {
		if p.Scene() != s {
			continue
		}
		for _, s := range schemas {
			if p.Schema().Equal(s) {
				p2 := p
				result = append(result, &p2)
				break
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID().ID().Compare(result[j].ID().ID()) < 0
	})
	return result, nil
}

func (r *Property) Save(ctx context.Context, p *property.Property) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.data[p.ID()] = *p
	return nil
}

func (r *Property) SaveAll(ctx context.Context, pl property.List) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, p := range pl {
		r.data[p.ID()] = *p
	}
	return nil
}

func (r *Property) Remove(ctx context.Context, id id.PropertyID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.data, id)
	return nil
}

func (r *Property) RemoveAll(ctx context.Context, ids []id.PropertyID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, id := range ids {
		delete(r.data, id)
	}
	return nil
}

func (r *Property) RemoveByScene(ctx context.Context, sceneID id.SceneID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for pid, p := range r.data {
		if p.Scene() == sceneID {
			delete(r.data, pid)
		}
	}
	return nil
}
