package property

import (
	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/id"
)

type SchemaGroupBuilder struct {
	p *SchemaGroup
}

func NewSchemaGroup() *SchemaGroupBuilder {
	return &SchemaGroupBuilder{
		p: &SchemaGroup{},
	}
}

func (b *SchemaGroupBuilder) Build() *SchemaGroup {
	if b.p.id == "" {
		return nil
	}
	return b.p
}

func (b *SchemaGroupBuilder) MustBuild() *SchemaGroup {
	p := b.Build()
	if p == nil {
		panic("invalid property schema group")
	}
	return p
}

func (b *SchemaGroupBuilder) ID(id id.PropertySchemaGroupID) *SchemaGroupBuilder {
	b.p.id = id
	return b
}

func (b *SchemaGroupBuilder) Fields(fields []*SchemaField) *SchemaGroupBuilder {
	newFields := []*SchemaField{}
	ids := map[id.PropertySchemaFieldID]struct{}{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		if _, ok := ids[f.ID()]; ok {
			continue
		}
		ids[f.ID()] = struct{}{}
		newFields = append(newFields, f)
	}
	b.p.fields = newFields
	return b
}

func (b *SchemaGroupBuilder) IsList(list bool) *SchemaGroupBuilder {
	b.p.list = list
	return b
}

func (b *SchemaGroupBuilder) IsAvailableIf(cond *Condition) *SchemaGroupBuilder {
	b.p.isAvailableIf = cond.Clone()
	return b
}

func (b *SchemaGroupBuilder) Title(title i18n.String) *SchemaGroupBuilder {
	b.p.title = title.Copy()
	return b
}

func (b *SchemaGroupBuilder) RepresentativeField(representativeField *id.PropertySchemaFieldID) *SchemaGroupBuilder {
	b.p.representativeField = representativeField.CopyRef()
	return b
}
