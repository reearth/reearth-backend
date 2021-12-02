package property

import (
	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/value"
)

type ValueAndDatasetValue struct {
	t ValueType
	d *dataset.Value
	p *Value
}

func NewValueAndDatasetValue(ty ValueType, d *dataset.Value, p *Value) *ValueAndDatasetValue {
	if !ty.Valid() {
		return nil
	}

	if d != nil && ValueType(d.Type()) != ty {
		d = nil
	}

	if p != nil && p.Type() != ty {
		p = nil
	}

	return &ValueAndDatasetValue{
		t: ty,
		d: d,
		p: p,
	}
}

func (v *ValueAndDatasetValue) Type() ValueType {
	if v == nil {
		return ValueTypeUnknown
	}
	return v.t
}

func (v *ValueAndDatasetValue) DatasetValue() *dataset.Value {
	if v == nil || v.t == ValueTypeUnknown {
		return nil
	}
	return v.d
}

func (v *ValueAndDatasetValue) PropertyValue() *Value {
	if v == nil || v.t == ValueTypeUnknown {
		return nil
	}
	return v.p
}

func (v *ValueAndDatasetValue) Value() *Value {
	if v == nil || v.t == ValueTypeUnknown {
		return nil
	}
	if v.d != nil {
		dv := valueFromDataset(v.d)
		return dv
	}
	return v.p
}

func valueFromDataset(v *dataset.Value) *Value {
	return ValueType(value.Type(v.Type())).ValueFrom(v.Value())
}
