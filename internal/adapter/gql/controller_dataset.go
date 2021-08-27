package gql

import (
	"context"

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

func (c *DatasetController) Fetch(ctx context.Context, ids []id.DatasetID, operator *usecase.Operator) ([]*Dataset, []error) {
	res, err := c.usecase.Fetch(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	datasets := make([]*Dataset, 0, len(res))
	for _, d := range res {
		datasets = append(datasets, toDataset(d))
	}

	return datasets, nil
}

func (c *DatasetController) FetchSchema(ctx context.Context, ids []id.DatasetSchemaID, operator *usecase.Operator) ([]*DatasetSchema, []error) {
	res, err := c.usecase.FetchSchema(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	schemas := make([]*DatasetSchema, 0, len(res))
	for _, d := range res {
		schemas = append(schemas, toDatasetSchema(d))
	}

	return schemas, nil
}

func (c *DatasetController) GraphFetch(ctx context.Context, i id.DatasetID, depth int, operator *usecase.Operator) ([]*Dataset, []error) {
	res, err := c.usecase.GraphFetch(ctx, i, depth, operator)
	if err != nil {
		return nil, []error{err}
	}

	datasets := make([]*Dataset, 0, len(res))
	for _, d := range res {
		datasets = append(datasets, toDataset(d))
	}

	return datasets, nil
}

func (c *DatasetController) GraphFetchSchema(ctx context.Context, i id.ID, depth int, operator *usecase.Operator) ([]*DatasetSchema, []error) {
	res, err := c.usecase.GraphFetchSchema(ctx, id.DatasetSchemaID(i), depth, operator)
	if err != nil {
		return nil, []error{err}
	}

	schemas := make([]*DatasetSchema, 0, len(res))
	for _, d := range res {
		schemas = append(schemas, toDatasetSchema(d))
	}

	return schemas, nil
}

func (c *DatasetController) FindSchemaByScene(ctx context.Context, i id.ID, first *int, last *int, before *usecase.Cursor, after *usecase.Cursor, operator *usecase.Operator) (*DatasetSchemaConnection, error) {
	res, pi, err := c.usecase.FindSchemaByScene(ctx, id.SceneID(i), usecase.NewPagination(first, last, before, after), operator)
	if err != nil {
		return nil, err
	}

	edges := make([]*DatasetSchemaEdge, 0, len(res))
	nodes := make([]*DatasetSchema, 0, len(res))
	for _, dataset := range res {
		ds := toDatasetSchema(dataset)
		edges = append(edges, &DatasetSchemaEdge{
			Node:   ds,
			Cursor: usecase.Cursor(ds.ID.String()),
		})
		nodes = append(nodes, ds)
	}

	return &DatasetSchemaConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   toPageInfo(pi),
		TotalCount: pi.TotalCount(),
	}, nil
}

func (c *DatasetController) FindDynamicSchemasByScene(ctx context.Context, sid id.ID) ([]*DatasetSchema, error) {
	res, err := c.usecase.FindDynamicSchemaByScene(ctx, id.SceneID(sid))
	if err != nil {
		return nil, err
	}

	dss := []*DatasetSchema{}
	for _, dataset := range res {
		dss = append(dss, toDatasetSchema(dataset))
	}

	return dss, nil
}

func (c *DatasetController) FindBySchema(ctx context.Context, dsid id.ID, first *int, last *int, before *usecase.Cursor, after *usecase.Cursor, operator *usecase.Operator) (*DatasetConnection, error) {
	p := usecase.NewPagination(first, last, before, after)
	res, pi, err2 := c.usecase.FindBySchema(ctx, id.DatasetSchemaID(dsid), p, operator)
	if err2 != nil {
		return nil, err2
	}

	edges := make([]*DatasetEdge, 0, len(res))
	nodes := make([]*Dataset, 0, len(res))
	for _, dataset := range res {
		ds := toDataset(dataset)
		edges = append(edges, &DatasetEdge{
			Node:   ds,
			Cursor: usecase.Cursor(ds.ID.String()),
		})
		nodes = append(nodes, ds)
	}

	conn := &DatasetConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   toPageInfo(pi),
		TotalCount: pi.TotalCount(),
	}

	return conn, nil
}

// data loader

type DatasetDataLoader interface {
	Load(id.DatasetID) (*Dataset, error)
	LoadAll([]id.DatasetID) ([]*Dataset, []error)
}

func (c *DatasetController) DataLoader(ctx context.Context) DatasetDataLoader {
	return NewDatasetLoader(DatasetLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.DatasetID) ([]*Dataset, []error) {
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

func (l *ordinaryDatasetLoader) Load(key id.DatasetID) (*Dataset, error) {
	res, errs := l.c.Fetch(l.ctx, []id.DatasetID{key}, getOperator(l.ctx))
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryDatasetLoader) LoadAll(keys []id.DatasetID) ([]*Dataset, []error) {
	return l.c.Fetch(l.ctx, keys, getOperator(l.ctx))
}

type DatasetSchemaDataLoader interface {
	Load(id.DatasetSchemaID) (*DatasetSchema, error)
	LoadAll([]id.DatasetSchemaID) ([]*DatasetSchema, []error)
}

func (c *DatasetController) SchemaDataLoader(ctx context.Context) DatasetSchemaDataLoader {
	return NewDatasetSchemaLoader(DatasetSchemaLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.DatasetSchemaID) ([]*DatasetSchema, []error) {
			return c.FetchSchema(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *DatasetController) SchemaOrdinaryDataLoader(ctx context.Context) DatasetSchemaDataLoader {
	return &ordinaryDatasetSchemaLoader{
		fetch: func(keys []id.DatasetSchemaID) ([]*DatasetSchema, []error) {
			return c.FetchSchema(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryDatasetSchemaLoader struct {
	fetch func(keys []id.DatasetSchemaID) ([]*DatasetSchema, []error)
}

func (l *ordinaryDatasetSchemaLoader) Load(key id.DatasetSchemaID) (*DatasetSchema, error) {
	res, errs := l.fetch([]id.DatasetSchemaID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryDatasetSchemaLoader) LoadAll(keys []id.DatasetSchemaID) ([]*DatasetSchema, []error) {
	return l.fetch(keys)
}
