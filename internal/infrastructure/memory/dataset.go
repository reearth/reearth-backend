package memory

import (
	"context"
	"sync"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

type Dataset struct {
	lock   sync.Mutex
	data   dataset.Map
	filter *id.SceneIDSet
}

func NewDataset() repo.Dataset {
	return &Dataset{
		data: dataset.Map{},
	}
}

func (r *Dataset) Filtered(scenes []id.SceneID) repo.Dataset {
	return &Dataset{
		data:   r.data,
		filter: id.NewSceneIDSet(scenes...),
	}
}

func (r *Dataset) FindByID(ctx context.Context, id id.DatasetID, f []id.SceneID) (*dataset.Dataset, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	p, ok := r.data[id]
	if ok && isSceneIncludes(p.Scene(), f) {
		return p, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Dataset) FindByIDs(ctx context.Context, ids []id.DatasetID, f []id.SceneID) (dataset.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := dataset.List{}
	for _, id := range ids {
		if d, ok := r.data[id]; ok {
			if isSceneIncludes(d.Scene(), f) {
				result = append(result, d)
				continue
			}
		}
		result = append(result, nil)
	}
	return result, nil
}

func (r *Dataset) FindBySchema(ctx context.Context, id id.DatasetSchemaID, f []id.SceneID, p *usecase.Pagination) (dataset.List, *usecase.PageInfo, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := dataset.List{}
	for _, d := range r.data {
		if d.Schema() == id && isSceneIncludes(d.Scene(), f) {
			dd := d
			result = append(result, dd)
		}
	}

	var startCursor, endCursor *usecase.Cursor
	if len(result) > 0 {
		_startCursor := usecase.Cursor(result[0].ID().String())
		_endCursor := usecase.Cursor(result[len(result)-1].ID().String())
		startCursor = &_startCursor
		endCursor = &_endCursor
	}

	return result, usecase.NewPageInfo(
		len(r.data),
		startCursor,
		endCursor,
		true,
		true,
	), nil
}

func (r *Dataset) FindBySchemaAll(ctx context.Context, id id.DatasetSchemaID) (dataset.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := dataset.List{}
	for _, d := range r.data {
		if d.Schema() == id {
			dd := d
			result = append(result, dd)
		}
	}
	return result, nil
}

func (r *Dataset) FindGraph(ctx context.Context, i id.DatasetID, f []id.SceneID, fields []id.DatasetSchemaFieldID) (dataset.List, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := make(dataset.List, 0, len(fields))
	next := i
	for _, nextField := range fields {
		d, _ := r.FindByID(ctx, next, f)
		if d != nil {
			result = append(result, d)
			if f := d.Field(nextField); f != nil {
				if f.Type() == dataset.ValueTypeRef {
					if l := f.Value().ValueRef(); l != nil {
						if did, err := id.DatasetIDFrom(*l); err == nil {
							next = did
							continue
						}
					}
				}
			}
		}
	}
	return result, nil
}

func (r *Dataset) Save(ctx context.Context, d *dataset.Dataset) error {
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

func (r *Dataset) SaveAll(ctx context.Context, dl dataset.List) error {
	for _, d := range dl {
		if !r.ok(d) {
			return repo.ErrOperationDenied
		}
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	for _, d := range dl {
		r.data[d.ID()] = d
	}
	return nil
}

func (r *Dataset) Remove(ctx context.Context, id id.DatasetID) error {
	if d, ok := r.data[id]; !ok || !r.ok(d) {
		return repo.ErrNotFound
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.data, id)
	return nil
}

func (r *Dataset) RemoveAll(ctx context.Context, ids []id.DatasetID) error {
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

func (r *Dataset) RemoveByScene(ctx context.Context, sceneID id.SceneID) error {
	if r.filter != nil && !r.filter.Has(sceneID) {
		return repo.ErrOperationDenied
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	for did, d := range r.data {
		if d.Scene() == sceneID {
			delete(r.data, did)
		}
	}
	return nil
}

func (r *Dataset) ok(d *dataset.Dataset) bool {
	return r.filter == nil || d != nil && r.filter.Has(d.Scene())
}

func (r *Dataset) applyFilter(list dataset.List) dataset.List {
	if len(list) == 0 {
		return nil
	}

	res := make(dataset.List, 0, len(list))
	for _, a := range list {
		if r.ok(a) {
			res = append(res, a)
		}
	}

	return res
}
