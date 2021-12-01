package dataset

type Pointer struct {
	dataset *ID
	schema  SchemaID
	field   FieldID
}

func PointAt(d ID, s SchemaID, f FieldID) *Pointer {
	if d.IsNil() || s.IsNil() || f.IsNil() {
		return nil
	}
	return &Pointer{
		dataset: d.CopyRef(),
		schema:  s,
		field:   f,
	}
}

func PointAtField(s SchemaID, f FieldID) *Pointer {
	if s.IsNil() || f.IsNil() {
		return nil
	}
	return &Pointer{
		schema: s,
		field:  f,
	}
}

func (p *Pointer) Dataset() *ID {
	if p == nil {
		return nil
	}
	return p.dataset.CopyRef()
}

func (p *Pointer) Schema() SchemaID {
	if p == nil {
		return SchemaID{}
	}
	return p.schema
}

func (p *Pointer) Field() FieldID {
	if p == nil {
		return FieldID{}
	}
	return p.field
}

func (p *Pointer) IsEmpty() bool {
	return p == nil || p.field.IsNil() || p.schema.IsNil()
}

func (l *Pointer) Clone() *Pointer {
	if l == nil || l.IsEmpty() {
		return nil
	}
	return &Pointer{
		dataset: l.Dataset(),
		schema:  l.Schema(),
		field:   l.Field(),
	}
}

func (l *Pointer) PointAt(d *ID) *Pointer {
	if l == nil || l.IsEmpty() {
		return nil
	}

	return &Pointer{
		dataset: d.CopyRef(),
		schema:  l.Schema(),
		field:   l.Field(),
	}
}
