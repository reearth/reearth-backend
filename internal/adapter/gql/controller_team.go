package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type TeamController struct {
	usecase interfaces.Team
}

func NewTeamController(usecase interfaces.Team) *TeamController {
	return &TeamController{usecase: usecase}
}

func (c *TeamController) Fetch(ctx context.Context, ids []id.TeamID, operator *usecase.Operator) ([]*Team, []error) {
	res, err := c.usecase.Fetch(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	teams := make([]*Team, 0, len(res))
	for _, t := range res {
		teams = append(teams, toTeam(t))
	}
	return teams, nil
}

func (c *TeamController) FindByUser(ctx context.Context, uid id.UserID, operator *usecase.Operator) ([]*Team, error) {
	res, err := c.usecase.FindByUser(ctx, uid, operator)
	if err != nil {
		return nil, err
	}
	teams := make([]*Team, 0, len(res))
	for _, t := range res {
		teams = append(teams, toTeam(t))
	}
	return teams, nil
}

// data loader

type TeamDataLoader interface {
	Load(id.TeamID) (*Team, error)
	LoadAll([]id.TeamID) ([]*Team, []error)
}

func (c *TeamController) DataLoader(ctx context.Context) TeamDataLoader {
	return NewTeamLoader(TeamLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.TeamID) ([]*Team, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *TeamController) OrdinaryDataLoader(ctx context.Context) TeamDataLoader {
	return &ordinaryTeamLoader{
		fetch: func(keys []id.TeamID) ([]*Team, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryTeamLoader struct {
	fetch func(keys []id.TeamID) ([]*Team, []error)
}

func (l *ordinaryTeamLoader) Load(key id.TeamID) (*Team, error) {
	res, errs := l.fetch([]id.TeamID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryTeamLoader) LoadAll(keys []id.TeamID) ([]*Team, []error) {
	return l.fetch(keys)
}
