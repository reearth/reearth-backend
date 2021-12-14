package property

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
)

var (
	ErrInvalidPropertyValue = errors.New("invalid property value")
	ErrCannotLinkDataset    = errors.New("cannot link dataset")
	ErrInvalidPropertyType  = errors.New("invalid property type")
	ErrInvalidPropertyField = errors.New("invalid property field")
)

type Field struct {
	field id.PropertySchemaFieldID
	links *dataset.GraphPointer
	v     *OptionalValue
}

func (p *Field) Clone() *Field {
	if p == nil {
		return nil
	}
	return &Field{
		field: p.field,
		links: p.links.Clone(),
		v:     p.v.Clone(),
	}
}

func (p *Field) Field() id.PropertySchemaFieldID {
	if p == nil {
		return FieldID("")
	}
	return p.field
}

func (p *Field) Links() *dataset.GraphPointer {
	if p == nil {
		return nil
	}
	return p.links
}

func (p *Field) Type() ValueType {
	if p == nil {
		return ValueTypeUnknown
	}
	return p.v.Type()
}

func (p *Field) Value() *Value {
	if p == nil {
		return nil
	}
	return p.v.Value()
}

func (p *Field) ActualValue(ds *dataset.Dataset) *ValueAndDatasetValue {
	if p == nil {
		return nil
	}

	var dv *dataset.Value
	if p.links != nil {
		if l := p.links.Last(); !l.IsEmpty() {
			d := l.Dataset()
			if d != nil && ds.ID() == *d {
				dv = ds.Field(l.Field()).Value()
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
	return NewValueAndDatasetValue(p.Type(), dv, p.Value())
}

func (p *Field) Datasets() []id.DatasetID {
	if p == nil {
		return nil
	}

	res := []id.DatasetID{}
	if p.Links().IsLinkedFully() {
		dsid := p.Links().Last().Dataset()
		if dsid != nil {
			res = append(res, *dsid)
		}
	}

	return res
}

func (p *Field) IsDatasetLinked(s id.DatasetSchemaID, i id.DatasetID) bool {
	return p.Links().HasSchemaAndDataset(s, i)
}

func (p *Field) Update(value *Value, field *SchemaField) error {
	if p == nil {
		return nil
	}
	if field == nil || p.field != field.ID() || !field.Validate(p.v) {
		return ErrInvalidPropertyValue
	}
	p.v.SetValue(value)
	return nil
}

func (p *Field) Cast(t ValueType) bool {
	if p.IsEmpty() || p.Type() == t {
		return false
	}
	p.v = p.v.Cast(t)
	p.Unlink()
	return true
}

func (p *Field) UpdateUnsafe(value *Value) {
	if p == nil {
		return
	}
	p.v.SetValue(value)
}

func (p *Field) Link(links *dataset.GraphPointer) {
	if p == nil {
		return
	}
	p.links = links.Clone()
}

func (p *Field) Unlink() {
	p.Link(nil)
}

func (p *Field) UpdateField(field id.PropertySchemaFieldID) {
	if p == nil {
		return
	}
	p.field = field
}

func (p *Field) IsEmpty() bool {
	return p == nil || (p.v.Value() == nil && p.Links().IsEmpty())
}

func (p *Field) MigrateSchema(ctx context.Context, newSchema *Schema, dl dataset.Loader) bool {
	if p == nil || dl == nil || newSchema == nil {
		return false
	}

	fid := p.Field()
	schemaField := newSchema.Groups().Field(fid)

	// If field is not found in new schema, this field should be removed
	invalid := schemaField == nil

	// if value is not compatible for type, value will be cleared
	if !schemaField.Validate(p.v) {
		p.UpdateUnsafe(nil)
	}

	// If linked dataset is not compatible for type, it will be unlinked
	l := p.Links()
	if dl != nil && l.IsLinkedFully() {
		if dsid, dsfid := l.Last().Dataset(), l.Last().Field(); dsid != nil {
			dss, _ := dl(ctx, *dsid)
			if dsf := dss[0].Field(dsfid); dsf != nil {
				if schemaField.Type() != ValueType(dsf.Type()) {
					p.Unlink()
				}
			}
		}
	}

	return !invalid
}

func (f *Field) GuessSchema() *SchemaField {
	if f == nil {
		return nil
	}
	if f, err := NewSchemaField().ID(f.Field()).Type(f.Type()).Build(); err == nil {
		return f
	}
	return nil
}
