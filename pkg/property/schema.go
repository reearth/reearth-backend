package property

import "github.com/reearth/reearth-backend/pkg/id"

type Schema struct {
	id       id.PropertySchemaID
	version  int
	groups   []*SchemaGroup
	linkable LinkableFields
}

type SchemaGroupAndField struct {
	Group *SchemaGroup
	Field *SchemaField
}

type LinkableFields struct {
	LatLng *Pointer
	URL    *Pointer
}

func (p *Schema) ID() id.PropertySchemaID {
	return p.id
}

func (p *Schema) IDRef() *id.PropertySchemaID {
	if p == nil {
		return nil
	}
	return p.id.Ref()
}

func (p *Schema) Version() int {
	return p.version
}

func (p *Schema) Fields() []*SchemaField {
	if p == nil {
		return nil
	}
	fields := []*SchemaField{}
	for _, g := range p.groups {
		fields = append(fields, g.Fields()...)
	}
	return fields
}

func (p *Schema) Field(id id.PropertySchemaFieldID) *SchemaField {
	if p == nil {
		return nil
	}
	for _, g := range p.groups {
		if f := g.Field(id); f != nil {
			return f
		}
	}
	return nil
}

func (p *Schema) FieldByPointer(ptr *Pointer) *SchemaField {
	if p == nil {
		return nil
	}
	g := p.GroupByPointer(ptr)
	if g == nil {
		return nil
	}
	return g.FieldByPointer(ptr)
}

func (p *Schema) Groups() []*SchemaGroup {
	if p == nil {
		return nil
	}
	return append([]*SchemaGroup{}, p.groups...)
}

func (p *Schema) Group(id id.PropertySchemaGroupID) *SchemaGroup {
	if p == nil {
		return nil
	}
	for _, f := range p.groups {
		if f.ID() == id {
			return f
		}
	}
	return nil
}

func (p *Schema) GroupByField(id id.PropertySchemaFieldID) *SchemaGroup {
	if p == nil {
		return nil
	}
	for _, f := range p.groups {
		if f.HasField(id) {
			return f
		}
	}
	return nil
}

func (p *Schema) GroupByPointer(ptr *Pointer) *SchemaGroup {
	if p == nil {
		return nil
	}

	if gid, ok := ptr.ItemBySchemaGroup(); ok {
		return p.Group(gid)
	}
	if fid, ok := ptr.Field(); ok {
		for _, g := range p.groups {
			if g.HasField(fid) {
				return g
			}
		}
	}

	return nil
}

func (gf SchemaGroupAndField) IsEmpty() bool {
	return gf.Group == nil && gf.Field == nil
}

func (p *Schema) GroupAndFields() []SchemaGroupAndField {
	if p == nil {
		return nil
	}
	fields := []SchemaGroupAndField{}
	for _, g := range p.groups {
		for _, f := range g.Fields() {
			fields = append(fields, SchemaGroupAndField{Group: g, Field: f})
		}
	}
	return fields
}

func (p *Schema) GroupAndField(f FieldID) *SchemaGroupAndField {
	if p == nil {
		return nil
	}
	for _, g := range p.groups {
		if gf := g.Field(f); gf != nil {
			return &SchemaGroupAndField{Group: g, Field: gf}
		}
	}
	return nil
}

func (s *Schema) DetectDuplicatedFields() []id.PropertySchemaFieldID {
	duplicated := []id.PropertySchemaFieldID{}
	ids := map[id.PropertySchemaFieldID]struct{}{}
	for _, f := range s.Fields() {
		i := f.ID()
		if _, ok := ids[i]; ok {
			duplicated = append(duplicated, i)
			return duplicated
		}
		ids[i] = struct{}{}
	}
	return nil
}

func (p *Schema) LinkableFields() LinkableFields {
	if p == nil {
		return LinkableFields{}
	}
	return p.linkable.Clone()
}

func (l LinkableFields) Clone() LinkableFields {
	return LinkableFields{
		LatLng: l.LatLng.Clone(),
		URL:    l.URL.Clone(),
	}
}

func (l LinkableFields) Validate(s *Schema) bool {
	if s == nil {
		return false
	}
	if l.LatLng != nil {
		if f := s.FieldByPointer(l.LatLng); f == nil {
			return false
		}
	}
	if l.URL != nil {
		if f := s.FieldByPointer(l.URL); f == nil {
			return false
		}
	}
	return true
}
