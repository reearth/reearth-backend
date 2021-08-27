package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type DatasetController struct {
	usecase interfaces.Dataset
}

func NewDatasetController(usecase interfaces.Dataset) *DatasetController {
	return &DatasetController{usecase: usecase}
}

func (c *DatasetController) Fetch(ctx context.Context, ids []id.DatasetID, operator *usecase.Operator) ([]*gqlmodel.Dataset, []error) {
	res, err := c.usecase.Fetch(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	datasets := make([]*gqlmodel.Dataset, 0, len(res))
	for _, d := range res {
		datasets = append(datasets, gqlmodel.ToDataset(d))
	}

	return datasets, nil
}

func (c *DatasetController) FetchSchema(ctx context.Context, ids []id.DatasetSchemaID, operator *usecase.Operator) ([]*gqlmodel.DatasetSchema, []error) {
	res, err := c.usecase.FetchSchema(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	schemas := make([]*gqlmodel.DatasetSchema, 0, len(res))
	for _, d := range res {
		schemas = append(schemas, gqlmodel.ToDatasetSchema(d))
	}

	return schemas, nil
}

func (c *DatasetController) GraphFetch(ctx context.Context, i id.DatasetID, depth int, operator *usecase.Operator) ([]*gqlmodel.Dataset, []error) {
	res, err := c.usecase.GraphFetch(ctx, i, depth, operator)
	if err != nil {
		return nil, []error{err}
	}

	datasets := make([]*gqlmodel.Dataset, 0, len(res))
	for _, d := range res {
		datasets = append(datasets, gqlmodel.ToDataset(d))
	}

	return datasets, nil
}

func (c *DatasetController) GraphFetchSchema(ctx context.Context, i id.ID, depth int, operator *usecase.Operator) ([]*gqlmodel.DatasetSchema, []error) {
	res, err := c.usecase.GraphFetchSchema(ctx, id.DatasetSchemaID(i), depth, operator)
	if err != nil {
		return nil, []error{err}
	}

	schemas := make([]*gqlmodel.DatasetSchema, 0, len(res))
	for _, d := range res {
		schemas = append(schemas, gqlmodel.ToDatasetSchema(d))
	}

	return schemas, nil
}

func (c *DatasetController) FindSchemaByScene(ctx context.Context, i id.ID, first *int, last *int, before *usecase.Cursor, after *usecase.Cursor, operator *usecase.Operator) (*gqlmodel.DatasetSchemaConnection, error) {
	res, pi, err := c.usecase.FindSchemaByScene(ctx, id.SceneID(i), usecase.NewPagination(first, last, before, after), operator)
	if err != nil {
		return nil, err
	}

	edges := make([]*gqlmodel.DatasetSchemaEdge, 0, len(res))
	nodes := make([]*gqlmodel.DatasetSchema, 0, len(res))
	for _, dataset := range res {
		ds := gqlmodel.ToDatasetSchema(dataset)
		edges = append(edges, &gqlmodel.DatasetSchemaEdge{
			Node:   ds,
			Cursor: usecase.Cursor(ds.ID.String()),
		})
		nodes = append(nodes, ds)
	}

	return &gqlmodel.DatasetSchemaConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   gqlmodel.ToPageInfo(pi),
		TotalCount: pi.TotalCount(),
	}, nil
}

func (c *DatasetController) FindDynamicSchemasByScene(ctx context.Context, sid id.ID) ([]*gqlmodel.DatasetSchema, error) {
	res, err := c.usecase.FindDynamicSchemaByScene(ctx, id.SceneID(sid))
	if err != nil {
		return nil, err
	}

	dss := []*gqlmodel.DatasetSchema{}
	for _, dataset := range res {
		dss = append(dss, gqlmodel.ToDatasetSchema(dataset))
	}

	return dss, nil
}

func (c *DatasetController) FindBySchema(ctx context.Context, dsid id.ID, first *int, last *int, before *usecase.Cursor, after *usecase.Cursor, operator *usecase.Operator) (*gqlmodel.DatasetConnection, error) {
	p := usecase.NewPagination(first, last, before, after)
	res, pi, err2 := c.usecase.FindBySchema(ctx, id.DatasetSchemaID(dsid), p, operator)
	if err2 != nil {
		return nil, err2
	}

	edges := make([]*gqlmodel.DatasetEdge, 0, len(res))
	nodes := make([]*gqlmodel.Dataset, 0, len(res))
	for _, dataset := range res {
		ds := gqlmodel.ToDataset(dataset)
		edges = append(edges, &gqlmodel.DatasetEdge{
			Node:   ds,
			Cursor: usecase.Cursor(ds.ID.String()),
		})
		nodes = append(nodes, ds)
	}

	conn := &gqlmodel.DatasetConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   gqlmodel.ToPageInfo(pi),
		TotalCount: pi.TotalCount(),
	}

	return conn, nil
}

// data loader

type DatasetDataLoader interface {
	Load(id.DatasetID) (*gqlmodel.Dataset, error)
	LoadAll([]id.DatasetID) ([]*gqlmodel.Dataset, []error)
}

func (c *DatasetController) DataLoader(ctx context.Context) DatasetDataLoader {
	return gqldataloader.NewDatasetLoader(gqldataloader.DatasetLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.DatasetID) ([]*gqlmodel.Dataset, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *DatasetController) OrdinaryDataLoader(ctx context.Context) DatasetDataLoader {
	return &ordinaryDatasetLoader{ctx: ctx, c: c}
}

type ordinaryDatasetLoader struct {
	ctx context.Context
	c   *DatasetController
}

func (l *ordinaryDatasetLoader) Load(key id.DatasetID) (*gqlmodel.Dataset, error) {
	res, errs := l.c.Fetch(l.ctx, []id.DatasetID{key}, getOperator(l.ctx))
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryDatasetLoader) LoadAll(keys []id.DatasetID) ([]*gqlmodel.Dataset, []error) {
	return l.c.Fetch(l.ctx, keys, getOperator(l.ctx))
}

type DatasetSchemaDataLoader interface {
	Load(id.DatasetSchemaID) (*gqlmodel.DatasetSchema, error)
	LoadAll([]id.DatasetSchemaID) ([]*gqlmodel.DatasetSchema, []error)
}

func (c *DatasetController) SchemaDataLoader(ctx context.Context) DatasetSchemaDataLoader {
	return gqldataloader.NewDatasetSchemaLoader(gqldataloader.DatasetSchemaLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.DatasetSchemaID) ([]*gqlmodel.DatasetSchema, []error) {
			return c.FetchSchema(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *DatasetController) SchemaOrdinaryDataLoader(ctx context.Context) DatasetSchemaDataLoader {
	return &ordinaryDatasetSchemaLoader{
		fetch: func(keys []id.DatasetSchemaID) ([]*gqlmodel.DatasetSchema, []error) {
			return c.FetchSchema(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryDatasetSchemaLoader struct {
	fetch func(keys []id.DatasetSchemaID) ([]*gqlmodel.DatasetSchema, []error)
}

func (l *ordinaryDatasetSchemaLoader) Load(key id.DatasetSchemaID) (*gqlmodel.DatasetSchema, error) {
	res, errs := l.fetch([]id.DatasetSchemaID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryDatasetSchemaLoader) LoadAll(keys []id.DatasetSchemaID) ([]*gqlmodel.DatasetSchema, []error) {
	return l.fetch(keys)
}
