package property

import "github.com/reearth/reearth-backend/pkg/id"

type Schema struct {
	id       id.PropertySchemaID
	version  int
	groups   *SchemaGroupList
	linkable LinkableFields
}

type SchemaFieldPointer struct {
	SchemaGroup SchemaGroupID
	Field       FieldID
}

type LinkableFields struct {
	LatLng *SchemaFieldPointer
	URL    *SchemaFieldPointer
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

func (p *Schema) Groups() *SchemaGroupList {
	if p == nil {
		return nil
	}
	return p.groups
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
		if f := s.groups.Field(l.LatLng.Field); f == nil {
			return false
		}
	}
	if l.URL != nil {
		if f := s.groups.Field(l.URL.Field); f == nil {
			return false
		}
	}
	return true
}

func (l LinkableFields) PointerByType(ty ValueType) *SchemaFieldPointer {
	switch ty {
	case ValueTypeLatLng:
		return l.LatLng
	case ValueTypeURL:
		return l.URL
	}
	return nil
}

func (l LinkableFields) FieldByType(ty ValueType) *FieldID {
	p := l.PointerByType(ty)
	if p == nil {
		return nil
	}
	return p.Field.Ref()
}

func (p *SchemaFieldPointer) Clone() *SchemaFieldPointer {
	if p == nil {
		return p
	}
	p2 := *p
	return &p2
}
