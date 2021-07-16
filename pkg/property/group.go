package property

import (
	"context"
	"errors"
	"fmt"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
)

// Group represents a group of property
type Group struct {
	itemBase
	fields []*Field
}

// Group implements Item interface
var _ Item = &Group{}

// ID returns id
func (g *Group) ID() id.PropertyItemID {
	if g == nil {
		return id.PropertyItemID{}
	}
	return g.itemBase.ID
}

// IDRef returns a reference of id
func (g *Group) IDRef() *id.PropertyItemID {
	if g == nil {
		return nil
	}
	return g.itemBase.ID.Ref()
}

// SchemaGroup returns id of schema group
func (g *Group) SchemaGroup() id.PropertySchemaFieldID {
	if g == nil {
		return id.PropertySchemaFieldID("")
	}
	return g.itemBase.SchemaGroup
}

// SchemaGroupRef _
func (g *Group) SchemaGroupRef() *id.PropertySchemaFieldID {
	if g == nil {
		return nil
	}
	return g.itemBase.SchemaGroup.Ref()
}

// Schema _
func (g *Group) Schema() id.PropertySchemaID {
	if g == nil {
		return id.PropertySchemaID{}
	}
	return g.itemBase.Schema
}

// SchemaRef _
func (g *Group) SchemaRef() *id.PropertySchemaID {
	if g == nil {
		return nil
	}
	return g.itemBase.Schema.Ref()
}

// HasLinkedField _
func (g *Group) HasLinkedField() bool {
	if g == nil {
		return false
	}
	for _, f := range g.fields {
		if f.HasLinkedField() {
			return true
		}
	}
	return false
}

// CollectDatasets _
func (g *Group) CollectDatasets() []id.DatasetID {
	if g == nil {
		return nil
	}
	res := []id.DatasetID{}

	for _, f := range g.fields {
		res = append(res, f.CollectDatasets()...)
	}

	return res
}

// FieldsByLinkedDataset _
func (g *Group) FieldsByLinkedDataset(s id.DatasetSchemaID, i id.DatasetID) []*Field {
	if g == nil {
		return nil
	}
	res := []*Field{}
	for _, f := range g.fields {
		if f.Links().IsDatasetLinked(s, i) {
			res = append(res, f)
		}
	}
	return res
}

// IsDatasetLinked _
func (g *Group) IsDatasetLinked(s id.DatasetSchemaID, i id.DatasetID) bool {
	if g == nil {
		return false
	}
	for _, f := range g.fields {
		if f.IsDatasetLinked(s, i) {
			return true
		}
	}
	return false
}

// IsEmpty _
func (g *Group) IsEmpty() bool {
	if g != nil {
		for _, f := range g.fields {
			if !f.IsEmpty() {
				return false
			}
		}
	}
	return true
}

// Prune _
func (g *Group) Prune() {
	if g == nil {
		return
	}
	for _, f := range g.fields {
		if f.IsEmpty() {
			g.RemoveField(f.Field())
		}
	}
}

// TODO: group migration
func (g *Group) MigrateSchema(ctx context.Context, newSchema *Schema, dl dataset.Loader) {
	if g == nil || dl == nil {
		return
	}

	g.itemBase.Schema = newSchema.ID()

	for _, f := range g.fields {
		if !f.MigrateSchema(ctx, newSchema, dl) {
			g.RemoveField(f.Field())
		}
	}

	g.Prune()
}

// GetOrCreateField _
func (g *Group) GetOrCreateField(ps *Schema, fid id.PropertySchemaFieldID) (*Field, bool) {
	if g == nil || ps == nil || g.Schema() != ps.ID() {
		return nil, false
	}
	psg := ps.Group(g.SchemaGroup())
	if psg == nil {
		return nil, false
	}

	psf := psg.Field(fid)
	if psf == nil {
		return nil, false
	}

	psfid := psf.ID()
	field := g.Field(psfid)
	if field != nil {
		return field, false
	}

	// if the field does not exist, create it here
	field, _ = NewField(psf).Build()
	if field == nil {
		return nil, false
	}

	g.fields = append(g.fields, field)
	return field, true
}

// RemoveField _
func (g *Group) RemoveField(fid id.PropertySchemaFieldID) {
	if g == nil {
		return
	}
	for i, f := range g.fields {
		if f.Field() == fid {
			g.fields = append(g.fields[:i], g.fields[i+1:]...)
			return
		}
	}
}

// FieldIDs _
func (g *Group) FieldIDs() []id.PropertySchemaFieldID {
	if g == nil {
		return nil
	}
	fields := make([]id.PropertySchemaFieldID, 0, len(g.fields))
	for _, f := range g.fields {
		fields = append(fields, f.Field())
	}
	return fields
}

// Fields returns a slice of fields
func (g *Group) Fields() []*Field {
	if g == nil {
		return nil
	}
	return append([]*Field{}, g.fields...)
}

// Field returns a field whose id is specified
func (g *Group) Field(fid id.PropertySchemaFieldID) *Field {
	if g == nil {
		return nil
	}
	for _, f := range g.fields {
		if f.Field() == fid {
			return f
		}
	}
	return nil
}

// MigrateDataset _
func (g *Group) MigrateDataset(q DatasetMigrationParam) {
	if g == nil {
		return
	}
	for _, f := range g.fields {
		f.MigrateDataset(q)
	}
}

func (g *Group) UpdateNameFieldValue(ps *Schema, value *Value) error {
	if g == nil || ps == nil || g.Schema() != ps.ID() {
		return nil
	}
	if psg := ps.GroupByPointer(NewPointer(&g.itemBase.SchemaGroup, nil, nil)); psg != nil {
		if representativeField := psg.RepresentativeFieldID(); representativeField != nil {
			if f, _ := g.GetOrCreateField(ps, *representativeField); f != nil {
				return f.Update(value, psg.Field(*representativeField))
			}
		}
	}
	return ErrInvalidPropertyField
}

func (p *Group) ValidateSchema(ps *SchemaGroup) error {
	if p == nil {
		return nil
	}
	if ps == nil {
		return errors.New("invalid schema")
	}
	if p.Schema() != ps.Schema() {
		return errors.New("invalid schema id")
	}
	if p.SchemaGroup() != ps.ID() {
		return errors.New("invalid schema group id")
	}

	for _, i := range p.fields {
		if err := i.ValidateSchema(ps.Field(i.Field())); err != nil {
			return fmt.Errorf("%s: %w", i.Field(), err)
		}
	}

	return nil
}