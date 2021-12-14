package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

var (
	testField1 = NewField().Field(testSchemaField1.ID()).Value(OptionalValueFrom(ValueTypeString.ValueFrom("aaa"))).MustBuild()
	testField2 = NewField().Field(testSchemaField3.ID()).Value(NewOptionalValue(ValueTypeLatLng, nil)).MustBuild()
)

func TestField_ActualValue(t *testing.T) {
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	dssfid := id.NewDatasetSchemaFieldID()
	l := dataset.PointAt(dsid, dssid, dssfid)
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})

	tests := []struct {
		Name     string
		Field    *Field
		DS       *dataset.Dataset
		Expected *ValueAndDatasetValue
	}{
		{
			Name:     "nil links",
			Field:    NewField().Field("a").Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Build(),
			Expected: NewValueAndDatasetValue(ValueTypeString, nil, ValueTypeString.ValueFrom("vvv")),
		},
		{
			Name:     "empty links",
			Field:    NewField().Field("a").Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(&dataset.GraphPointer{}).Build(),
			Expected: NewValueAndDatasetValue(ValueTypeString, nil, ValueTypeString.ValueFrom("vvv")),
		},
		{
			Name:  "dataset value",
			Field: NewField().Field("a").Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(ls).Build(),
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
			Field:    NewField().Field("a").Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(ls).Build(),
			DS:       dataset.New().ID(dsid).Schema(dssid).MustBuild(),
			Expected: NewValueAndDatasetValue(ValueTypeString, nil, ValueTypeString.ValueFrom("vvv")),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res := tt.Field.ActualValue(tt.DS)
			assert.Equal(t, tt.Expected, res)
		})
	}
}

func TestField_Datasets(t *testing.T) {
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	dssfid := id.NewDatasetSchemaFieldID()
	l := dataset.PointAt(dsid, dssid, dssfid)
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})

	tests := []struct {
		Name     string
		Field    *Field
		Expected []id.DatasetID
	}{
		{
			Name:     "list of one datasets",
			Field:    NewField().Field("a").Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(ls).Build(),
			Expected: []id.DatasetID{dsid},
		},
		{
			Name:     "nil field",
			Expected: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res := tt.Field.Datasets()
			assert.Equal(t, tt.Expected, res)
		})
	}
}

func TestField_Clone(t *testing.T) {
	l := dataset.PointAt(id.NewDatasetID(), id.NewDatasetSchemaID(), id.NewDatasetSchemaFieldID())
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})
	b := NewField().Field("a").Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Link(ls).Build()

	tests := []struct {
		name   string
		target *Field
		want   *Field
	}{
		{
			name:   "ok",
			target: b,
			want:   b,
		},
		{
			name:   "nil",
			target: nil,
			want:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := b.Clone()
			assert.Equal(t, b, r)
			if tt.want != nil {
				assert.NotSame(t, b, r)
			}
		})
	}
}

func TestField_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		target *Field
		want   bool
	}{
		{
			name:   "empty",
			target: &Field{},
			want:   true,
		},
		{
			name:   "empty value",
			target: NewField().Field("a").Value(NewOptionalValue(ValueTypeString, nil)).Build(),
			want:   true,
		},
		{
			name:   "not empty",
			target: NewField().Field("a").Value(OptionalValueFrom(ValueTypeString.ValueFrom("x"))).Build(),
			want:   false,
		},
		{
			name:   "nil",
			target: nil,
			want:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.target.IsEmpty())
		})
	}
}

func TestField_Link(t *testing.T) {
	did := id.NewDatasetID()
	dsid := id.NewDatasetSchemaID()
	dfid := id.NewDatasetSchemaFieldID()
	l := dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(did, dsid, dfid)})

	tests := []struct {
		name   string
		target *Field
		args   *dataset.GraphPointer
	}{
		{
			name:   "link",
			target: testField1.Clone(),
			args:   l,
		},
		{
			name:   "unlink",
			target: NewField().Field("a").Value(NewOptionalValue(ValueTypeString, nil)).Link(l).Build(),
			args:   nil,
		},
		{
			name:   "empty",
			target: &Field{},
			args:   nil,
		},
		{
			name:   "nil",
			target: nil,
			args:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.target.Link(tt.args)
			if tt.target != nil {
				assert.Equal(t, tt.args, tt.target.links)
			}
		})
	}
}

func TestField_Update(t *testing.T) {
	b := NewField().Field("a").Value(OptionalValueFrom(ValueTypeString.ValueFrom("vvv"))).Build()
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
		name      string
		target    *Field
		args      args
		want      bool
		wantField *Field
	}{
		{
			name: "ok",
			target: &Field{
				field: FieldID("foobar"),
				v:     OptionalValueFrom(ValueTypeString.ValueFrom("-123")),
				links: dgp.Clone(),
			},
			args: args{t: ValueTypeNumber},
			want: true,
			wantField: &Field{
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
			want: true,
			wantField: &Field{
				field: FieldID("foobar"),
				v:     NewOptionalValue(ValueTypeLatLng, nil),
			},
		},
		{
			name:      "empty",
			target:    &Field{},
			args:      args{t: ValueTypeNumber},
			want:      false,
			wantField: &Field{},
		},
		{
			name:      "nil",
			target:    nil,
			args:      args{t: ValueTypeNumber},
			want:      false,
			wantField: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Cast(tt.args.t))
			assert.Equal(t, tt.wantField, tt.target)
		})
	}
}

func TestField_GuessSchema(t *testing.T) {
	tests := []struct {
		name   string
		target *Field
		want   *SchemaField
	}{
		{
			name:   "ok",
			target: &Field{field: "a", v: NewOptionalValue(ValueTypeLatLng, nil)},
			want:   &SchemaField{id: "a", propertyType: ValueTypeLatLng},
		},
		{
			name:   "empty",
			target: &Field{},
			want:   nil,
		},
		{
			name:   "nil",
			target: nil,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.GuessSchema())
		})
	}
}
