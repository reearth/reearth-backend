package fs

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	err1 "github.com/reearth/reearth-backend/pkg/error"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
)

type propertySchema struct {
	basePath string
}

func NewPropertySchema(basePath string) repo.PropertySchema {
	return &propertySchema{
		basePath: basePath,
	}
}

func (r *propertySchema) FindByID(ctx context.Context, i id.PropertySchemaID) (*property.Schema, error) {
	pid, err := id.PluginIDFrom(i.Plugin())
	if err != nil {
		return nil, err1.ErrNotFound
	}

	m, err := readManifest(r.basePath, pid)
	if err != nil {
		return nil, err1.ErrInternalBy(err)
	}

	if m.Schema != nil && m.Schema.ID() == i {
		return m.Schema, nil
	}

	for _, ps := range m.ExtensionSchema {
		if ps == nil {
			continue
		}
		if ps.ID() == i {
			return ps, nil
		}
	}

	return nil, err1.ErrNotFound
}

func (r *propertySchema) FindByIDs(ctx context.Context, ids []id.PropertySchemaID) (property.SchemaList, error) {
	results := make(property.SchemaList, 0, len(ids))
	for _, id := range ids {
		res, err := r.FindByID(ctx, id)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (r *propertySchema) Save(ctx context.Context, p *property.Schema) error {
	return err1.ErrInternalBy(errors.New("read only"))
}

func (r *propertySchema) SaveAll(ctx context.Context, p property.SchemaList) error {
	return err1.ErrInternalBy(errors.New("read only"))
}

func (r *propertySchema) Remove(ctx context.Context, pid id.PropertySchemaID) error {
	return err1.ErrInternalBy(errors.New("read only"))
}

func (r *propertySchema) RemoveAll(ctx context.Context, pid []id.PropertySchemaID) error {
	return err1.ErrInternalBy(errors.New("read only"))
}
