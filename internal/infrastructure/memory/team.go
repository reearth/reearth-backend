package memory

import (
	"context"
	"sync"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/reearth/reearth-backend/pkg/user"
)

type Team struct {
	lock sync.Mutex
	data map[id.TeamID]*user.Team
}

func NewTeam() repo.Team {
	return &Team{
		data: map[id.TeamID]*user.Team{},
	}
}

func (r *Team) FindByUser(ctx context.Context, i id.UserID) ([]*user.Team, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := []*user.Team{}
	for _, d := range r.data {
		if d.Members().ContainsUser(i) {
			result = append(result, d)
		}
	}
	return result, nil
}

func (r *Team) FindByIDs(ctx context.Context, ids []id.TeamID) ([]*user.Team, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := []*user.Team{}
	for _, id := range ids {
		if d, ok := r.data[id]; ok {
			result = append(result, d)
		} else {
			result = append(result, nil)
		}
	}
	return result, nil
}

func (r *Team) FindByID(ctx context.Context, id id.TeamID) (*user.Team, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	d, ok := r.data[id]
	if ok {
		return d, nil
	}
	return &user.Team{}, rerror.ErrNotFound
}

func (r *Team) Save(ctx context.Context, d *user.Team) error {
	if d == nil {
		return nil
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	r.data[d.ID()] = d
	return nil
}

func (r *Team) SaveAll(ctx context.Context, list []*user.Team) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, t := range list {
		if t == nil {
			continue
		}
		r.data[t.ID()] = t
	}
	return nil
}

func (r *Team) Remove(ctx context.Context, id id.TeamID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.data, id)
	return nil
}

func (r *Team) RemoveAll(ctx context.Context, ids []id.TeamID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, id := range ids {
		delete(r.data, id)
	}
	return nil
}
