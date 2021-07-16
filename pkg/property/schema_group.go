package property

import (
	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/id"
)

// SchemaGroup represents a group of property that has some fields
type SchemaGroup struct {
	id                  id.PropertySchemaFieldID
	sid                 id.PropertySchemaID
	fields              []*SchemaField
	list                bool
	isAvailableIf       *Condition
	title               i18n.String
	representativeField *id.PropertySchemaFieldID
}

// ID returns id
func (s *SchemaGroup) ID() id.PropertySchemaFieldID {
	if s == nil {
		return id.PropertySchemaFieldID("")
	}
	return s.id
}

func (s *SchemaGroup) IDRef() *id.PropertySchemaFieldID {
	if s == nil {
		return nil
	}
	return s.id.Ref()
}

func (s *SchemaGroup) Schema() id.PropertySchemaID {
	if s == nil {
		return id.PropertySchemaID{}
	}
	return s.sid
}

func (s *SchemaGroup) SchemaRef() *id.PropertySchemaID {
	if s == nil {
		return nil
	}
	return &s.sid
}

// Fields returns a slice of fields
func (s *SchemaGroup) Fields() []*SchemaField {
	if s == nil {
		return nil
	}
	return append([]*SchemaField{}, s.fields...)
}

// Field returns a field whose id is specified
func (s *SchemaGroup) Field(fid id.PropertySchemaFieldID) *SchemaField {
	if s == nil {
		return nil
	}
	for _, f := range s.fields {
		if f.ID() == fid {
			return f
		}
	}
	return nil
}

// FieldByPointer returns a field whose id is specified
func (s *SchemaGroup) FieldByPointer(ptr *Pointer) *SchemaField {
	if s == nil {
		return nil
	}
	fid, ok := ptr.Field()
	if !ok {
		return nil
	}
	return s.Field(fid)
}

func (s *SchemaGroup) HasField(i id.PropertySchemaFieldID) bool {
	return s.Field(i) != nil
}

// IsList returns true if this group is list
func (s *SchemaGroup) IsList() bool {
	if s == nil {
		return false
	}
	return s.list
}

// IsAvailableIf returns condition of availability
func (s *SchemaGroup) IsAvailableIf() *Condition {
	if s == nil {
		return nil
	}
	return s.isAvailableIf.Clone()
}

// Title returns a title of the group
func (s *SchemaGroup) Title() i18n.String {
	if s == nil {
		return nil
	}
	return s.title.Copy()
}

// RepresentativeFieldID returns the representative field ID of the group
func (s *SchemaGroup) RepresentativeFieldID() *id.PropertySchemaFieldID {
	if s == nil {
		return nil
	}
	return s.representativeField
}

// RepresentativeField returns the representative field of the group
func (s *SchemaGroup) RepresentativeField() *SchemaField {
	if s == nil || s.representativeField == nil {
		return nil
	}
	return s.Field(*s.representativeField)
}

func (s *SchemaGroup) SetTitle(t i18n.String) {
	s.title = t.Copy()
}