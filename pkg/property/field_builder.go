package property

import (
	"github.com/reearth/reearth-backend/pkg/dataset"
)

type FieldBuilder struct {
	p   *Field
	psf *SchemaField
}

type FieldUnsafeBuilder struct {
	p *Field
}

func NewField() *FieldBuilder {
	return &FieldBuilder{
		p: &Field{},
	}
}

func (b *FieldBuilder) Build() *Field {
	if b.p.field == FieldID("") || b.p.v == nil {
		return nil
	}
	return b.p
}

func (b *FieldBuilder) MustBuild() *Field {
	p := b.Build()
	if p == nil {
		panic("invalid field")
	}
	return p
}

func (b *FieldBuilder) Field(f FieldID) *FieldBuilder {
	b.p.field = f
	return b
}

func (b *FieldBuilder) Value(v *OptionalValue) *FieldBuilder {
	b.p.v = v.Clone()
	return b
}

func (b *FieldBuilder) Link(l *dataset.GraphPointer) *FieldBuilder {
	b.p.links = l.Clone()
	return b
}
