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

type DatasetSchema struct {
	lock   sync.Mutex
	data   dataset.SchemaMap
	filter *id.SceneIDSet
}

func NewDatasetSchema() repo.DatasetSchema {
	return &DatasetSchema{
		data: dataset.SchemaMap{},
	}
}

func (r *DatasetSchema) Filtered(scenes []id.SceneID) repo.DatasetSchema {
	return &DatasetSchema{
		data:   r.data,
		filter: id.NewSceneIDSet(scenes...),
	}
}

func (r *DatasetSchema) FindByID(ctx context.Context, id id.DatasetSchemaID, f []id.SceneID) (*dataset.Schema, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	p, ok := r.data[id]
	if ok {
		return p, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *DatasetSchema) FindByIDs(ctx context.Context, ids []id.DatasetSchemaID, f []id.SceneID) (dataset.SchemaList, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := dataset.SchemaList{}
	for _, id := range ids {
		if d, ok := r.data[id]; ok {
			d2 := d
			result = append(result, d2)
		} else {
			result = append(result, nil)
		}
	}
	return result, nil
}

func (r *DatasetSchema) FindByScene(ctx context.Context, s id.SceneID, p *usecase.Pagination) (dataset.SchemaList, *usecase.PageInfo, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := dataset.SchemaList{}
	for _, d := range r.data {
		if d.Scene() == s {
			d2 := d
			result = append(result, d2)
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

func (r *DatasetSchema) FindBySceneAll(ctx context.Context, s id.SceneID) (dataset.SchemaList, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := dataset.SchemaList{}
	for _, d := range r.data {
		if d.Scene() == s {
			d2 := d
			result = append(result, d2)
		}
	}
	return result, nil
}

func (r *DatasetSchema) FindAllDynamicByScene(ctx context.Context, s id.SceneID) (dataset.SchemaList, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := dataset.SchemaList{}
	for _, d := range r.data {
		if d.Scene() == s && d.Dynamic() {
			d2 := d
			result = append(result, d2)
		}
	}
	return result, nil
}

func (r *DatasetSchema) FindDynamicByID(ctx context.Context, id id.DatasetSchemaID) (*dataset.Schema, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	p, ok := r.data[id]
	if ok && p.Dynamic() {
		return p, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *DatasetSchema) FindBySceneAndSource(ctx context.Context, s id.SceneID, src string) (dataset.SchemaList, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := dataset.SchemaList{}
	for _, d := range r.data {
		if d.Scene() == s && d.Source() == src {
			d2 := d
			result = append(result, d2)
		}
	}
	return result, nil
}

func (r *DatasetSchema) Save(ctx context.Context, d *dataset.Schema) error {
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

func (r *DatasetSchema) SaveAll(ctx context.Context, dl dataset.SchemaList) error {
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

func (r *DatasetSchema) Remove(ctx context.Context, id id.DatasetSchemaID) error {
	if d, ok := r.data[id]; !ok || !r.ok(d) {
		return repo.ErrNotFound
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.data, id)
	return nil
}

func (r *DatasetSchema) RemoveAll(ctx context.Context, ids []id.DatasetSchemaID) error {
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

func (r *DatasetSchema) RemoveByScene(ctx context.Context, sceneID id.SceneID) error {
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

func (r *DatasetSchema) ok(d *dataset.Schema) bool {
	return r.filter == nil || d != nil && r.filter.Has(d.Scene())
}

func (r *DatasetSchema) applyFilter(list dataset.SchemaList) dataset.SchemaList {
	if len(list) == 0 {
		return nil
	}

	res := make(dataset.SchemaList, 0, len(list))
	for _, a := range list {
		if r.ok(a) {
			res = append(res, a)
		}
	}

	return res
}
