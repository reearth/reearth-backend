package dataset

type GraphPointer struct {
	pointers []*Pointer
}

func NewGraphPointer(p []*Pointer) *GraphPointer {
	if len(p) == 0 {
		return nil
	}
	pointers := make([]*Pointer, 0, len(p))
	for _, l := range p {
		if l == nil {
			continue
		}
		pointers = append(pointers, l.Clone())
	}
	return &GraphPointer{
		pointers: pointers,
	}
}

func (l *GraphPointer) Clone() *GraphPointer {
	if l.IsEmpty() {
		return nil
	}
	return &GraphPointer{
		pointers: append([]*Pointer{}, l.pointers...),
	}
}

func (l *GraphPointer) WithDataset(ds ID) *GraphPointer {
	if l.IsEmpty() {
		return nil
	}

	links := l.Clone()
	if links.First().Dataset() == nil {
		links.pointers[0] = links.pointers[0].PointAt(&ds)
	}
	return links
}

func (l *GraphPointer) IsEmpty() bool {
	return l == nil || len(l.pointers) == 0
}

func (l *GraphPointer) IsLinkedFully() bool {
	return !l.IsEmpty() && len(l.Datasets()) == len(l.pointers)
}

func (l *GraphPointer) Len() int {
	if l.IsEmpty() {
		return 0
	}
	return len(l.pointers)
}

func (l *GraphPointer) First() *Pointer {
	if l.IsEmpty() {
		return nil
	}
	return l.pointers[0]
}

func (l *GraphPointer) Last() *Pointer {
	if l.IsEmpty() {
		return nil
	}
	return l.pointers[len(l.pointers)-1]
}

func (l *GraphPointer) Pointers() []*Pointer {
	if l == nil || len(l.pointers) == 0 {
		return nil
	}
	return append([]*Pointer{}, l.pointers...)
}

func (l *GraphPointer) Datasets() []ID {
	if l.IsEmpty() {
		return nil
	}
	datasets := make([]ID, 0, len(l.pointers))
	for _, i := range l.pointers {
		if d := i.Dataset(); d != nil {
			datasets = append(datasets, *d)
		} else {
			return datasets
		}
	}
	return datasets
}

func (l *GraphPointer) Schemas() []SchemaID {
	if l.IsEmpty() {
		return nil
	}
	schemas := make([]SchemaID, 0, len(l.pointers))
	for _, i := range l.pointers {
		schemas = append(schemas, i.Schema())
	}
	return schemas
}

func (l *GraphPointer) Fields() []FieldID {
	if l.IsEmpty() {
		return nil
	}
	fields := make([]FieldID, 0, len(l.pointers))
	for _, i := range l.pointers {
		fields = append(fields, i.Field())
	}
	return fields
}

func (l *GraphPointer) HasDataset(did ID) bool {
	if l.IsEmpty() {
		return false
	}
	for _, l2 := range l.pointers {
		if d := l2.Dataset(); d != nil && *d == did {
			return true
		}
	}
	return false
}

func (l *GraphPointer) HasSchema(dsid SchemaID) bool {
	if l.IsEmpty() {
		return false
	}
	for _, l2 := range l.pointers {
		if l2.Schema() == dsid {
			return true
		}
	}
	return false
}

func (l *GraphPointer) HasSchemaAndDataset(dsid SchemaID, did ID) bool {
	if l.IsEmpty() {
		return false
	}

	for _, l2 := range l.pointers {
		if l2 == nil || l2.Schema() != dsid {
			continue
		}
		if d := l2.Dataset(); d != nil && *d == did {
			return true
		}
	}

	return false
}
