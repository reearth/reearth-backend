package memory

import (
	"context"
	"sync"
	"time"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/project"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

type Project struct {
	lock   sync.Mutex
	data   map[id.ProjectID]*project.Project
	filter *id.TeamIDSet
}

func NewProject() repo.Project {
	return &Project{
		data: map[id.ProjectID]*project.Project{},
	}
}

func (r *Project) Filtered(teams []id.TeamID) repo.Project {
	return &Project{
		data:   r.data,
		filter: id.NewTeamIDSet(teams...),
	}
}

func (r *Project) FindByTeam(ctx context.Context, id id.TeamID, p *usecase.Pagination) ([]*project.Project, *usecase.PageInfo, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := []*project.Project{}
	for _, d := range r.data {
		if d.Team() == id {
			result = append(result, d)
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

func (r *Project) FindByIDs(ctx context.Context, ids []id.ProjectID, filter []id.TeamID) ([]*project.Project, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := []*project.Project{}
	for _, id := range ids {
		if d, ok := r.data[id]; ok {
			if isTeamIncludes(d.Team(), filter) {
				result = append(result, d)
				continue
			}
		}
		result = append(result, nil)
	}
	return result, nil
}

func (r *Project) FindByID(ctx context.Context, id id.ProjectID, filter []id.TeamID) (*project.Project, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	p, ok := r.data[id]
	if ok && isTeamIncludes(p.Team(), filter) {
		return p, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Project) FindByPublicName(ctx context.Context, name string) (*project.Project, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if name == "" {
		return nil, nil
	}
	for _, p := range r.data {
		if p.MatchWithPublicName(name) {
			return p, nil
		}
	}
	return nil, rerror.ErrNotFound
}

func (r *Project) CountByTeam(ctx context.Context, team id.TeamID) (c int, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, p := range r.data {
		if p.Team() == team {
			c++
		}
	}
	return
}

func (r *Project) Save(ctx context.Context, d *project.Project) error {
	if d == nil {
		return nil
	}
	if !r.ok(d) {
		return repo.ErrOperationDenied
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	d.SetUpdatedAt(time.Now())
	r.data[d.ID()] = d
	return nil
}

func (r *Project) Remove(ctx context.Context, id id.ProjectID) error {
	if d, ok := r.data[id]; !ok || !r.ok(d) {
		return repo.ErrNotFound
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.data, id)
	return nil
}

func (r *Project) ok(d *project.Project) bool {
	return r.filter == nil || d != nil && r.filter.Has(d.Team())
}

func (r *Project) applyFilter(list []*project.Project) []*project.Project {
	if len(list) == 0 {
		return nil
	}

	res := make([]*project.Project, 0, len(list))
	for _, e := range list {
		if r.ok(e) {
			res = append(res, e)
		}
	}

	return res
}
