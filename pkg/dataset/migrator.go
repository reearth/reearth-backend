package dataset

import "context"

type Migrator struct {
	d func(ID) *ID
	s func(SchemaID) *SchemaID
	f func(FieldID) *FieldID
}

func NewMigrator(d func(ID) *ID, s func(SchemaID) *SchemaID, f func(FieldID) *FieldID) *Migrator {
	if d == nil || s == nil || f == nil {
		return nil
	}
	return &Migrator{d: d, s: s, f: f}
}

func MigratorFrom(d map[ID]ID, s map[SchemaID]SchemaID, f map[FieldID]FieldID) *Migrator {
	df := func(i ID) *ID {
		f, ok := d[i]
		if !ok {
			return nil
		}
		return f.Ref()
	}
	sf := func(i SchemaID) *SchemaID {
		f, ok := s[i]
		if !ok {
			return nil
		}
		return f.CopyRef()
	}
	ff := func(i FieldID) *FieldID {
		f, ok := f[i]
		if !ok {
			return nil
		}
		return f.CopyRef()
	}
	return &Migrator{d: df, s: sf, f: ff}
}

func (m *Migrator) MigrateAndValidateGraphPointer(ctx context.Context, p *GraphPointer, l GraphLoader) (*GraphPointer, error) {
	if m == nil || p == nil {
		return nil, nil
	}

	np := m.MigrateGraphPointer(p)
	d, f, err := l.ByGraphPointer(ctx, np)
	if err != nil || d == nil || f == nil {
		return nil, err
	}

	return np, nil
}

func (m *Migrator) MigrateGraphPointer(p *GraphPointer) *GraphPointer {
	if m == nil || p == nil {
		return nil
	}

	pointers := p.Pointers()
	pointers2 := make([]*Pointer, 0, len(pointers))
	for _, p := range pointers {
		pointers2 = append(pointers2, m.MigratePointer(p))
	}
	return NewGraphPointer(pointers2)
}

func (m *Migrator) MigratePointer(p *Pointer) *Pointer {
	if m == nil || p == nil {
		return nil
	}

	ns := m.getS(p.Schema())
	nf := m.getF(p.Field())

	if d := p.Dataset(); d != nil {
		return PointAt(m.getD(*d), ns, nf)
	}

	return PointAtField(ns, nf)
}

func (m *Migrator) getD(d ID) ID {
	if m == nil || m.d == nil {
		return d
	}
	if nd := m.d(d); nd != nil {
		return *nd
	}
	return d
}

func (m *Migrator) getS(s SchemaID) SchemaID {
	if m == nil || m.s == nil {
		return s
	}
	if ns := m.s(s); ns != nil {
		return *ns
	}
	return s
}

func (m *Migrator) getF(f FieldID) FieldID {
	if m == nil || m.f == nil {
		return f
	}
	if nf := m.f(f); nf != nil {
		return *nf
	}
	return f
}
