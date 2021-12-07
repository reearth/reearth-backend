package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestField_ActualValue(t *testing.T) {
	p := NewSchemaField().ID("A").Type(ValueTypeString).MustBuild()
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	dssfid := id.NewDatasetSchemaFieldID()
	l := dataset.PointAt(dsid, dssid, dssfid)
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})

	testCases := []struct {
		Name     string
		Field    *Field
		DS       *dataset.Dataset
		Expected *ValueAndDatasetValue
	}{
		{
			Name:     "nil links",
			Field:    NewField(p).Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).MustBuild(),
			Expected: NewValueAndDatasetValue(ValueTypeString, nil, ValueTypeString.ValueFrom("vvv")),
		},
		{
			Name:     "empty links",
			Field:    NewField(p).Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(&dataset.GraphPointer{}).MustBuild(),
			Expected: NewValueAndDatasetValue(ValueTypeString, nil, ValueTypeString.ValueFrom("vvv")),
		},
		{
			Name:  "dataset value",
			Field: NewField(p).Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(ls).MustBuild(),
			DS: dataset.New().
				ID(dsid).Schema(dssid).
				Fields([]*dataset.Field{
					dataset.NewField(dssfid, dataset.ValueTypeString.ValueFrom("xxx"), "")},
				).
				MustBuild(),
			Expected: NewValueAndDatasetValue(ValueTypeString, dataset.ValueTypeString.ValueFrom("xxx"), ValueTypeString.ValueFrom("vvv")),
		},
		{
			Name:     "dataset value missing",
			Field:    NewField(p).Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(ls).MustBuild(),
			DS:       dataset.New().ID(dsid).Schema(dssid).MustBuild(),
			Expected: NewValueAndDatasetValue(ValueTypeString, nil, ValueTypeString.ValueFrom("vvv")),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Field.ActualValue(tc.DS)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestField_Datasets(t *testing.T) {
	p := NewSchemaField().ID("A").Type(ValueTypeString).MustBuild()
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	dssfid := id.NewDatasetSchemaFieldID()
	l := dataset.PointAt(dsid, dssid, dssfid)
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})

	testCases := []struct {
		Name     string
		Field    *Field
		Expected []id.DatasetID
	}{
		{
			Name:     "list of one datasets",
			Field:    NewField(p).Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(ls).MustBuild(),
			Expected: []id.DatasetID{dsid},
		},
		{
			Name:     "nil field",
			Expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Field.Datasets()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestField_Clone(t *testing.T) {
	p := NewSchemaField().ID("A").Type(ValueTypeString).MustBuild()
	l := dataset.PointAt(id.NewDatasetID(), id.NewDatasetSchemaID(), id.NewDatasetSchemaFieldID())
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})
	b := NewField(p).Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(ls).MustBuild()
	r := b.Clone()
	assert.Equal(t, b, r)
}

func TestField(t *testing.T) {
	did := id.NewDatasetID()
	dsid := id.NewDatasetSchemaID()
	p := NewSchemaField().ID("A").Type(ValueTypeString).MustBuild()
	b := NewField(p).MustBuild()
	assert.True(t, b.IsEmpty())
	l := dataset.PointAt(did, dsid, id.NewDatasetSchemaFieldID())
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})
	b.Link(ls)
	assert.True(t, b.IsDatasetLinked(dsid, did))
	b.Unlink()
	assert.Nil(t, b.Links())
}

func TestField_Update(t *testing.T) {
	p := NewSchemaField().ID("A").Type(ValueTypeString).MustBuild()
	b := NewField(p).Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).MustBuild()
	v := ValueTypeString.ValueFrom("xxx")
	b.UpdateUnsafe(v)
	assert.Equal(t, v, b.Value())
}

func TestField_Cast(t *testing.T) {
	dgp := dataset.NewGraphPointer([]*dataset.Pointer{
		dataset.PointAt(id.NewDatasetID(), id.NewDatasetSchemaID(), id.NewDatasetSchemaFieldID()),
	})

	type args struct {
		t ValueType
	}
	tests := []struct {
		name   string
		target *Field
		args   args
		want   *Field
	}{
		{
			name: "ok",
			target: &Field{
				field: FieldID("foobar"),
				v:     OptionalValueFrom(ValueTypeString.ValueFrom("-123")),
				links: dgp.Clone(),
			},
			args: args{t: ValueTypeNumber},
			want: &Field{
				field: FieldID("foobar"),
				v:     OptionalValueFrom(ValueTypeNumber.ValueFrom(-123)),
			},
		},
		{
			name: "failed",
			target: &Field{
				field: FieldID("foobar"),
				v:     OptionalValueFrom(ValueTypeString.ValueFrom("foo")),
				links: dgp.Clone(),
			},
			args: args{t: ValueTypeLatLng},
			want: &Field{
				field: FieldID("foobar"),
				v:     NewOptionalValue(ValueTypeLatLng, nil),
			},
		},
		{
			name:   "empty",
			target: &Field{},
			args:   args{t: ValueTypeNumber},
			want:   &Field{},
		},
		{
			name:   "nil",
			target: nil,
			args:   args{t: ValueTypeNumber},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.target.Cast(tt.args.t)
			assert.Equal(t, tt.want, tt.target)
		})
	}
}
