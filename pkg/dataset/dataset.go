package dataset

import "github.com/reearth/reearth-backend/pkg/id"

type Dataset struct {
	id     id.DatasetID
	source string
	schema id.DatasetSchemaID
	fields map[id.DatasetSchemaFieldID]*Field
	order  []id.DatasetSchemaFieldID
	scene  id.SceneID
}

func (d *Dataset) ID() (i id.DatasetID) {
	if d == nil {
		return
	}
	return d.id
}

func (d *Dataset) Scene() (i id.SceneID) {
	if d == nil {
		return
	}
	return d.scene
}

func (d *Dataset) Source() string {
	if d == nil {
		return ""
	}
	return d.source
}

func (d *Dataset) Schema() (i id.DatasetSchemaID) {
	if d == nil {
		return
	}
	return d.schema
}

func (d *Dataset) Fields() []*Field {
	if d == nil || d.order == nil {
		return nil
	}
	fields := make([]*Field, 0, len(d.fields))
	for _, id := range d.order {
		fields = append(fields, d.fields[id])
	}
	return fields
}

func (d *Dataset) Field(id id.DatasetSchemaFieldID) *Field {
	if d == nil || d.fields == nil {
		return nil
	}
	return d.fields[id]
}

func (d *Dataset) FieldByPointer(p *Pointer) *Field {
	if d == nil || p.IsEmpty() || d.Schema() != p.Schema() {
		return nil
	} else if pd := p.Dataset(); pd != nil && d.ID() != *pd {
		return nil
	}
	return d.Field(p.Field())
}

func (d *Dataset) FieldRef(id *id.DatasetSchemaFieldID) *Field {
	if d == nil || id == nil {
		return nil
	}
	return d.fields[*id]
}

func (d *Dataset) NameField(ds *Schema) *Field {
	if d == nil {
		return nil
	}
	if d.Schema() != ds.ID() {
		return nil
	}
	f := ds.RepresentativeField()
	if f == nil {
		return nil
	}
	return d.fields[f.ID()]
}

func (d *Dataset) FieldBySource(source string) *Field {
	if d == nil {
		return nil
	}
	for _, f := range d.fields {
		if f.source == source {
			return f
		}
	}
	return nil
}

func (d *Dataset) FieldByType(t ValueType) *Field {
	if d == nil {
		return nil
	}
	for _, f := range d.fields {
		if f.Type() == t {
			return f
		}
	}
	return nil
}

// Interface returns a simple and human-readable representation of the dataset
func (d *Dataset) Interface(s *Schema) map[string]interface{} {
	if d == nil || s == nil || d.Schema() != s.ID() {
		return nil
	}
	m := map[string]interface{}{}
	for _, f := range d.fields {
		key := s.Field(f.Field()).Name()
		m[key] = f.Value().Interface()
	}
	return m
}

// Interface is almost same as Interface, but keys of the map are IDs of fields.
func (d *Dataset) InterfaceWithFieldIDs() map[string]interface{} {
	if d == nil {
		return nil
	}
	m := map[string]interface{}{}
	for _, f := range d.fields {
		key := f.Field().String()
		m[key] = f.Value().Interface()
	}
	return m
}
