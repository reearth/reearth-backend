package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type PropertyController struct {
	usecase interfaces.Property
}

func NewPropertyController(usecase interfaces.Property) *PropertyController {
	return &PropertyController{usecase: usecase}
}

func (c *PropertyController) Fetch(ctx context.Context, ids []id.PropertyID, operator *usecase.Operator) ([]*Property, []error) {
	res, err := c.usecase.Fetch(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	properties := make([]*Property, 0, len(res))
	for _, property := range res {
		properties = append(properties, toProperty(property))
	}

	return properties, nil
}

func (c *PropertyController) FetchSchema(ctx context.Context, ids []id.PropertySchemaID, operator *usecase.Operator) ([]*PropertySchema, []error) {
	res, err := c.usecase.FetchSchema(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	schemas := make([]*PropertySchema, 0, len(res))
	for _, propertySchema := range res {
		schemas = append(schemas, toPropertySchema(propertySchema))
	}

	return schemas, nil
}

func (c *PropertyController) FetchMerged(ctx context.Context, org, parent, linked *id.ID, operator *usecase.Operator) (*MergedProperty, error) {
	res, err := c.usecase.FetchMerged(ctx, id.PropertyIDFromRefID(org), id.PropertyIDFromRefID(parent), id.DatasetIDFromRefID(linked), operator)

	if err != nil {
		return nil, err
	}

	return toMergedProperty(res), nil
}

// data loader

type PropertyDataLoader interface {
	Load(id.PropertyID) (*Property, error)
	LoadAll([]id.PropertyID) ([]*Property, []error)
}

func (c *PropertyController) DataLoader(ctx context.Context) *PropertyLoader {
	return NewPropertyLoader(PropertyLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.PropertyID) ([]*Property, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *PropertyController) OrdinaryDataLoader(ctx context.Context) PropertyDataLoader {
	return &ordinaryPropertyLoader{
		fetch: func(keys []id.PropertyID) ([]*Property, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryPropertyLoader struct {
	fetch func(keys []id.PropertyID) ([]*Property, []error)
}

func (l *ordinaryPropertyLoader) Load(key id.PropertyID) (*Property, error) {
	res, errs := l.fetch([]id.PropertyID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryPropertyLoader) LoadAll(keys []id.PropertyID) ([]*Property, []error) {
	return l.fetch(keys)
}

type PropertySchemaDataLoader interface {
	Load(id.PropertySchemaID) (*PropertySchema, error)
	LoadAll([]id.PropertySchemaID) ([]*PropertySchema, []error)
}

func (c *PropertyController) SchemaDataLoader(ctx context.Context) PropertySchemaDataLoader {
	return NewPropertySchemaLoader(PropertySchemaLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.PropertySchemaID) ([]*PropertySchema, []error) {
			return c.FetchSchema(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *PropertyController) SchemaOrdinaryDataLoader(ctx context.Context) PropertySchemaDataLoader {
	return &ordinaryPropertySchemaLoader{
		fetch: func(keys []id.PropertySchemaID) ([]*PropertySchema, []error) {
			return c.FetchSchema(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinaryPropertySchemaLoader struct {
	fetch func(keys []id.PropertySchemaID) ([]*PropertySchema, []error)
}

func (l *ordinaryPropertySchemaLoader) Load(key id.PropertySchemaID) (*PropertySchema, error) {
	res, errs := l.fetch([]id.PropertySchemaID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryPropertySchemaLoader) LoadAll(keys []id.PropertySchemaID) ([]*PropertySchema, []error) {
	return l.fetch(keys)
}
