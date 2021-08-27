package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type ProjectController struct {
	usecase interfaces.Project
}

func NewProjectController(usecase interfaces.Project) *ProjectController {
	return &ProjectController{usecase: usecase}
}

func (c *ProjectController) Fetch(ctx context.Context, ids []id.ProjectID, operator *usecase.Operator) ([]*Project, []error) {
	res, err := c.usecase.Fetch(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	projects := make([]*Project, 0, len(res))
	for _, project := range res {
		projects = append(projects, toProject(project))
	}

	return projects, nil
}

func (c *ProjectController) FindByTeam(ctx context.Context, teamID id.TeamID, first *int, last *int, before *usecase.Cursor, after *usecase.Cursor, operator *usecase.Operator) (*ProjectConnection, error) {
	res, pi, err := c.usecase.FindByTeam(ctx, teamID, usecase.NewPagination(first, last, before, after), operator)
	if err != nil {
		return nil, err
	}

	edges := make([]*ProjectEdge, 0, len(res))
	nodes := make([]*Project, 0, len(res))
	for _, p := range res {
		prj := toProject(p)
		edges = append(edges, &ProjectEdge{
			Node:   prj,
			Cursor: usecase.Cursor(prj.ID.String()),
		})
		nodes = append(nodes, prj)
	}

	return &ProjectConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   toPageInfo(pi),
		TotalCount: pi.TotalCount(),
	}, nil
}

func (c *ProjectController) CheckAlias(ctx context.Context, alias string) (*CheckProjectAliasPayload, error) {
	ok, err := c.usecase.CheckAlias(ctx, alias)
	if err != nil {
		return nil, err
	}

	return &CheckProjectAliasPayload{Alias: alias, Available: ok}, nil
}

// data loaders

type ProjectDataLoader interface {
	Load(id.ProjectID) (*Project, error)
	LoadAll([]id.ProjectID) ([]*Project, []error)
}

func (c *ProjectController) DataLoader(ctx context.Context) ProjectDataLoader {
	return NewProjectLoader(ProjectLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.ProjectID) ([]*Project, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *ProjectController) OrdinaryDataLoader(ctx context.Context) ProjectDataLoader {
	return &ordinaryProjectLoader{
		fetch: func(keys []id.ProjectID) ([]*Project, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryProjectLoader struct {
	fetch func(keys []id.ProjectID) ([]*Project, []error)
}

func (l *ordinaryProjectLoader) Load(key id.ProjectID) (*Project, error) {
	res, errs := l.fetch([]id.ProjectID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryProjectLoader) LoadAll(keys []id.ProjectID) ([]*Project, []error) {
	return l.fetch(keys)
}
