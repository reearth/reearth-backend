package memory

import (
	"context"
	"sync"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/asset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

type Asset struct {
	lock   sync.Mutex
	data   map[id.AssetID]*asset.Asset
	filter *id.TeamIDSet
}

func NewAsset() repo.Asset {
	return &Asset{
		data: map[id.AssetID]*asset.Asset{},
	}
}

func (r *Asset) Filtered(filter []id.TeamID) repo.Asset {
	return &Asset{
		data:   r.data,
		filter: id.NewTeamIDSet(filter...),
	}
}

func (r *Asset) FindByID(ctx context.Context, id id.AssetID, teams []id.TeamID) (*asset.Asset, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	d, ok := r.data[id]
	if ok && r.ok(d) {
		return d, nil
	}

	return nil, rerror.ErrNotFound
}

func (r *Asset) FindByIDs(ctx context.Context, ids []id.AssetID, teams []id.TeamID) ([]*asset.Asset, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := []*asset.Asset{}
	for _, id := range ids {
		if d, ok := r.data[id]; ok {
			if isTeamIncludes(d.Team(), teams) {
				result = append(result, d)
				continue
			}
		}
		result = append(result, nil)
	}
	return r.applyFilter(result), nil
}

func (r *Asset) FindByTeam(ctx context.Context, id id.TeamID, pagination *usecase.Pagination) ([]*asset.Asset, *usecase.PageInfo, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := []*asset.Asset{}
	for _, d := range r.data {
		if d.Team() == id {
			result = append(result, d)
		}
	}
	result = r.applyFilter(result)

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

func (r *Asset) Save(ctx context.Context, d *asset.Asset) error {
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

func (r *Asset) Remove(ctx context.Context, id id.AssetID) error {
	if d, ok := r.data[id]; !ok || !r.ok(d) {
		return repo.ErrNotFound
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.data, id)
	return nil
}

func (r *Asset) ok(d *asset.Asset) bool {
	return r.filter == nil || d != nil && r.filter.Has(d.Team())
}

func (r *Asset) applyFilter(list []*asset.Asset) []*asset.Asset {
	if len(list) == 0 {
		return nil
	}

	res := make([]*asset.Asset, 0, len(list))
	for _, a := range list {
		if r.ok(a) {
			res = append(res, a)
		}
	}

	return res
}
