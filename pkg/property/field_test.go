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

func TestField_CollectDatasets(t *testing.T) {
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
			res := tc.Field.CollectDatasets()
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
