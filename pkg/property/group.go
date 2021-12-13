package property

import (
	"context"
	"errors"

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

func (g *Group) ID() id.PropertyItemID {
	if g == nil {
		return id.PropertyItemID{}
	}
	return g.itemBase.ID
}

func (g *Group) SchemaGroup() id.PropertySchemaGroupID {
	if g == nil {
		return id.PropertySchemaGroupID("")
	}
	return g.itemBase.SchemaGroup
}

func (g *Group) HasLinkedField() bool {
	if g == nil {
		return false
	}
	for _, f := range g.fields {
		if f.Links() != nil {
			return true
		}
	}
	return false
}

func (g *Group) Datasets() []id.DatasetID {
	if g == nil {
		return nil
	}
	res := []id.DatasetID{}

	for _, f := range g.fields {
		res = append(res, f.Datasets()...)
	}

	return res
}

func (g *Group) FieldsByLinkedDataset(s id.DatasetSchemaID, i id.DatasetID) []*Field {
	if g == nil {
		return nil
	}
	res := []*Field{}
	for _, f := range g.fields {
		if f.Links().HasSchemaAndDataset(s, i) {
			res = append(res, f)
		}
	}
	return res
}

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

func (g *Group) Prune() (res bool) {
	if g == nil {
		return
	}
	for _, f := range g.fields {
		if f.IsEmpty() {
			res = g.RemoveField(f.Field())
		}
	}
	return
}

// TODO: group migration
func (g *Group) MigrateSchema(ctx context.Context, newSchema *Schema, dl dataset.Loader) {
	if g == nil || dl == nil {
		return
	}

	for _, f := range g.fields {
		if !f.MigrateSchema(ctx, newSchema, dl) {
			g.RemoveField(f.Field())
		}
	}

	g.Prune()
}

func (g *Group) GetOrCreateField(ps *Schema, fid id.PropertySchemaFieldID) (*Field, bool) {
	if g == nil || ps == nil {
		return nil, false
	}
	psg := ps.Groups().Group(g.SchemaGroup())
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
	field = NewField().Field(fid).Value(NewOptionalValue(psf.Type(), nil)).Build()
	if field == nil {
		return nil, false
	}

	g.AddFields(field)
	return field, true
}

func (g *Group) AddFields(fields ...*Field) {
	if g == nil {
		return
	}
	for _, f := range fields {
		_ = g.RemoveField(f.Field())
		g.fields = append(g.fields, f)
	}
}

func (g *Group) RemoveField(fid id.PropertySchemaFieldID) (res bool) {
	if g == nil {
		return
	}
	for i, f := range g.fields {
		if f.Field() == fid {
			g.fields = append(g.fields[:i], g.fields[i+1:]...)
			return true
		}
	}
	return
}

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

func (g *Group) UpdateNameFieldValue(ps *Schema, value *Value) error {
	if g == nil || ps == nil {
		return nil
	}
	if psg := ps.Groups().Group(g.itemBase.SchemaGroup); psg != nil {
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
	if p.SchemaGroup() != ps.ID() {
		return errors.New("invalid schema group id")
	}

	for _, i := range p.fields {
		f := ps.Field(i.Field())
		if f.Type() != i.Type() {
			return errors.New("invalid field type")
		}
	}

	return nil
}

func (p *Group) Clone() *Group {
	if p == nil {
		return nil
	}
	fields := make([]*Field, 0, len(p.fields))
	for _, f := range p.fields {
		fields = append(fields, f.Clone())
	}
	return &Group{
		fields:   fields,
		itemBase: p.itemBase,
	}
}

func (p *Group) CloneItem() Item {
	return p.Clone()
}

func (g *Group) Fields(p *Pointer) []*Field {
	if g == nil || len(g.fields) == 0 || (p != nil && !p.TestItem(g.SchemaGroup(), g.ID())) {
		return nil
	}

	if fid, ok := p.Field(); ok {
		if f := g.Field(fid); f != nil {
			return []*Field{f}
		}
		return nil
	}

	return append(g.fields[:0:0], g.fields...)
}

func (g *Group) RemoveFields(ptr *Pointer) (res bool) {
	if g == nil || ptr == nil {
		return false
	}
	if f, ok := ptr.FieldIfItemIs(g.SchemaGroup(), g.ID()); ok {
		if g.RemoveField(f) {
			res = true
		}
	}
	return
}

func (p *Group) GroupAndFields() []GroupAndField {
	if p == nil || len(p.fields) == 0 {
		return nil
	}
	res := []GroupAndField{}
	for _, f := range p.fields {
		res = append(res, GroupAndField{
			Group: p,
			Field: f,
		})
	}
	return res
}
