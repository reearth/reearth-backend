package property

import (
	"errors"
	"testing"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

var (
	testGroup1 = NewGroup().NewID().Schema(testSchema1.ID(), testSchemaGroup1.ID()).Fields([]*Field{testField1}).MustBuild()
	testGroup2 = NewGroup().NewID().Schema(testSchema1.ID(), testSchemaGroup2.ID()).Fields([]*Field{testField2}).MustBuild()
)

func TestGroup_SchemaGroup(t *testing.T) {
	var g *Group
	assert.Equal(t, id.PropertySchemaGroupID(""), g.SchemaGroup())
	pfid := id.PropertySchemaGroupID("aa")
	g = NewGroup().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), pfid).MustBuild()
	assert.Equal(t, pfid, g.SchemaGroup())
}

func TestGroup_HasLinkedField(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	l := dataset.PointAt(id.NewDatasetID(), id.NewDatasetSchemaID(), id.NewDatasetSchemaFieldID())
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})
	f := NewField(sf).Value(OptionalValueFrom(v)).Link(ls).MustBuild()
	f2 := NewField(sf).Value(OptionalValueFrom(v)).MustBuild()

	testCases := []struct {
		Name     string
		Group    *Group
		Expected bool
	}{
		{
			Name:     "nil group",
			Group:    nil,
			Expected: false,
		},
		{
			Name:     "true",
			Group:    NewGroup().NewID().Fields([]*Field{f}).MustBuild(),
			Expected: true,
		},
		{
			Name:     "false",
			Group:    NewGroup().NewID().Fields([]*Field{f2}).MustBuild(),
			Expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Group.HasLinkedField()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
func TestGroup_IsDatasetLinked(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	l := dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})
	f := NewField(sf).Value(OptionalValueFrom(v)).Link(ls).MustBuild()
	f2 := NewField(sf).Value(OptionalValueFrom(v)).MustBuild()

	testCases := []struct {
		Name          string
		Group         *Group
		DatasetSchema id.DatasetSchemaID
		Dataset       id.DatasetID
		Expected      bool
	}{
		{
			Name: "nil group",
		},
		{
			Name:          "true",
			Group:         NewGroup().NewID().Fields([]*Field{f}).MustBuild(),
			Dataset:       dsid,
			DatasetSchema: dssid,
			Expected:      true,
		},
		{
			Name:     "false",
			Group:    NewGroup().NewID().Fields([]*Field{f2}).MustBuild(),
			Expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Group.IsDatasetLinked(tc.DatasetSchema, tc.Dataset)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestGroup_Datasets(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	l := dataset.PointAt(dsid, id.NewDatasetSchemaID(), id.NewDatasetSchemaFieldID())
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})
	f := NewField(sf).Value(OptionalValueFrom(v)).Link(ls).MustBuild()

	testCases := []struct {
		Name     string
		Group    *Group
		Expected []id.DatasetID
	}{
		{
			Name:     "nil group",
			Group:    nil,
			Expected: nil,
		},
		{
			Name:     "normal case",
			Group:    NewGroup().NewID().Fields([]*Field{f}).MustBuild(),
			Expected: []id.DatasetID{dsid},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Group.Datasets()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestGroup_FieldsByLinkedDataset(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	l := dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())
	ls := dataset.NewGraphPointer([]*dataset.Pointer{l})
	f := NewField(sf).Value(OptionalValueFrom(v)).Link(ls).MustBuild()

	testCases := []struct {
		Name          string
		Group         *Group
		DatasetSchema id.DatasetSchemaID
		DataSet       id.DatasetID
		Expected      []*Field
	}{
		{
			Name: "nil group",
		},
		{
			Name:          "normal case",
			DataSet:       dsid,
			DatasetSchema: dssid,
			Group:         NewGroup().NewID().Fields([]*Field{f}).MustBuild(),
			Expected:      []*Field{f},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Group.FieldsByLinkedDataset(tc.DatasetSchema, tc.DataSet)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestGroup_IsEmpty(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	f := NewField(sf).Value(OptionalValueFrom(v)).MustBuild()
	f2 := NewField(sf).MustBuild()

	testCases := []struct {
		Name     string
		Group    *Group
		Expected bool
	}{

		{
			Name:     "true case",
			Group:    NewGroup().NewID().Fields([]*Field{f2}).MustBuild(),
			Expected: true,
		},
		{
			Name:     "false case",
			Group:    NewGroup().NewID().Fields([]*Field{f}).MustBuild(),
			Expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Group.IsEmpty()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestGroup_Prune(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	f := NewField(sf).Value(OptionalValueFrom(v)).MustBuild()
	f2 := NewField(sf).MustBuild()

	testCases := []struct {
		Name     string
		Group    *Group
		Expected []*Field
	}{

		{
			Name: "nil group",
		},
		{
			Name:     "normal case",
			Group:    NewGroup().NewID().Fields([]*Field{f, f2}).MustBuild(),
			Expected: []*Field{f},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.Group.Prune()
			assert.Equal(tt, tc.Expected, tc.Group.Fields(nil))
		})
	}
}

func TestGroup_GetOrCreateField(t *testing.T) {
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	f := NewField(sf).MustBuild()
	sg := NewSchemaGroup().ID("aa").Schema(id.MustPropertySchemaID("xx~1.0.0/aa")).Fields([]*SchemaField{sf}).MustBuild()
	testCases := []struct {
		Name     string
		Group    *Group
		PS       *Schema
		FID      id.PropertySchemaFieldID
		Expected struct {
			Field *Field
			Bool  bool
		}
	}{
		{
			Name: "nil group",
		},
		{
			Name:  "nil ps",
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").MustBuild(),
		},
		{
			Name:  "group schema doesn't equal to ps",
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aaa"), "aa").MustBuild(),
			PS:    NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
		},
		{
			Name:  "create field",
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").MustBuild(),
			PS:    NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			FID:   "aa",
			Expected: struct {
				Field *Field
				Bool  bool
			}{
				Field: NewField(sf).MustBuild(),
				Bool:  true,
			},
		},
		{
			Name:  "get field",
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").Fields([]*Field{f}).MustBuild(),
			PS:    NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			FID:   "aa",
			Expected: struct {
				Field *Field
				Bool  bool
			}{
				Field: NewField(sf).MustBuild(),
				Bool:  false,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res, b := tc.Group.GetOrCreateField(tc.PS, tc.FID)
			assert.Equal(tt, tc.Expected.Field, res)
			assert.Equal(tt, tc.Expected.Bool, b)
		})
	}
}

func TestGroup_RemoveField(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	sf2 := NewSchemaField().ID("b").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	f := NewField(sf).Value(OptionalValueFrom(v)).MustBuild()
	f2 := NewField(sf2).MustBuild()

	testCases := []struct {
		Name     string
		Group    *Group
		Input    id.PropertySchemaFieldID
		Expected []*Field
	}{

		{
			Name: "nil group",
		},
		{
			Name:     "normal case",
			Input:    "b",
			Group:    NewGroup().NewID().Fields([]*Field{f, f2}).MustBuild(),
			Expected: []*Field{f},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.Group.RemoveField(tc.Input)
			assert.Equal(tt, tc.Expected, tc.Group.Fields(nil))
		})
	}
}

func TestGroup_FieldIDs(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	sf2 := NewSchemaField().ID("b").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	f := NewField(sf).Value(OptionalValueFrom(v)).MustBuild()
	f2 := NewField(sf2).MustBuild()

	testCases := []struct {
		Name     string
		Group    *Group
		Expected []id.PropertySchemaFieldID
	}{

		{
			Name: "nil group",
		},
		{
			Name:     "normal case",
			Group:    NewGroup().NewID().Fields([]*Field{f, f2}).MustBuild(),
			Expected: []id.PropertySchemaFieldID{"a", "b"},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Group.FieldIDs()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestGroup_Field(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	sf2 := NewSchemaField().ID("b").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	f := NewField(sf).Value(OptionalValueFrom(v)).MustBuild()
	f2 := NewField(sf2).MustBuild()

	testCases := []struct {
		Name     string
		Group    *Group
		Input    id.PropertySchemaFieldID
		Expected *Field
	}{

		{
			Name: "nil group",
		},
		{
			Name:     "normal case",
			Group:    NewGroup().NewID().Fields([]*Field{f, f2}).MustBuild(),
			Input:    "a",
			Expected: f,
		},
		{
			Name:     "normal case",
			Group:    NewGroup().NewID().Fields([]*Field{f, f2}).MustBuild(),
			Input:    "x",
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Group.Field(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestGroup_UpdateNameFieldValue(t *testing.T) {
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	//f := NewField(sf).MustBuild()
	sg := NewSchemaGroup().ID("aa").Schema(id.MustPropertySchemaID("xx~1.0.0/aa")).Fields([]*SchemaField{sf}).MustBuild()
	sg2 := NewSchemaGroup().ID("bb").Schema(id.MustPropertySchemaID("xx~1.0.0/bb")).Fields([]*SchemaField{sf}).MustBuild()
	testCases := []struct {
		Name     string
		Group    *Group
		PS       *Schema
		Value    *Value
		FID      id.PropertySchemaFieldID
		Expected *Field
		Err      error
	}{
		{
			Name: "nil group",
		},
		{
			Name:  "nil ps",
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").MustBuild(),
		},
		{
			Name:  "group schema doesn't equal to ps",
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aaa"), "aa").MustBuild(),
			PS:    NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
		},
		{
			Name:     "update value",
			Group:    NewGroup().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").MustBuild(),
			PS:       NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Value:    ValueTypeString.ValueFrom("abc"),
			FID:      "aa",
			Expected: NewField(sf).Value(OptionalValueFrom(ValueTypeString.ValueFrom("abc"))).MustBuild(),
		},
		{
			Name:     "invalid property field",
			Group:    NewGroup().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").MustBuild(),
			PS:       NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/bb")).Groups([]*SchemaGroup{sg2}).MustBuild(),
			Value:    ValueTypeString.ValueFrom("abc"),
			FID:      "aa",
			Expected: nil,
			Err:      ErrInvalidPropertyField,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.Group.UpdateNameFieldValue(tc.PS, tc.Value)
			if res == nil {
				assert.Equal(tt, tc.Expected, tc.Group.Field(tc.FID))
			} else {
				assert.True(tt, errors.As(res, &tc.Err))
			}
		})
	}
}

func TestGroup_Clone(t *testing.T) {
	tests := []struct {
		name   string
		target *Group
		n      bool
	}{
		{
			name:   "ok",
			target: testGroup1.Clone(),
		},
		{
			name: "nil",
			n:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.target.Clone()
			if tt.n {
				assert.Nil(t, res)
			} else {
				assert.Equal(t, tt.target, res)
				assert.NotSame(t, tt.target, res)
			}
		})
	}
}

func TestGroup_CloneItem(t *testing.T) {
	tests := []struct {
		name   string
		target *Group
		n      bool
	}{
		{
			name:   "ok",
			target: testGroup1.Clone(),
		},
		{
			name: "nil",
			n:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.target.CloneItem()
			if tt.n {
				assert.Nil(t, res)
			} else {
				assert.Equal(t, tt.target, res)
				assert.NotSame(t, tt.target, res)
			}
		})
	}
}

func TestGroup_Fields(t *testing.T) {
	type args struct {
		p *Pointer
	}
	tests := []struct {
		name   string
		target *Group
		args   args
		want   []*Field
	}{
		{
			name:   "all",
			target: testGroup1,
			args:   args{p: nil},
			want:   []*Field{testField1},
		},
		{
			name:   "specified",
			target: testGroup1,
			args:   args{p: PointFieldOnly(testField1.Field())},
			want:   []*Field{testField1},
		},
		{
			name:   "specified schema group",
			target: testGroup1,
			args:   args{p: PointItemBySchema(testGroup1.SchemaGroup())},
			want:   []*Field{testField1},
		},
		{
			name:   "specified item",
			target: testGroup1,
			args:   args{p: PointItem(testGroup1.ID())},
			want:   []*Field{testField1},
		},
		{
			name:   "not found",
			target: testGroup1,
			args:   args{p: PointFieldOnly("xxxxxx")},
			want:   nil,
		},
		{
			name:   "empty",
			target: &Group{},
			args:   args{p: PointFieldOnly(testField1.Field())},
			want:   nil,
		},
		{
			name:   "nil",
			target: nil,
			args:   args{p: PointFieldOnly(testField1.Field())},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Fields(tt.args.p))
		})
	}
}

func TestGroup_RemoveFields(t *testing.T) {
	type args struct {
		p *Pointer
	}
	tests := []struct {
		name   string
		target *Group
		args   args
		want   []*Field
	}{
		{
			name:   "nil pointer",
			target: testGroup1.Clone(),
			args:   args{p: nil},
			want:   []*Field{testField1},
		},
		{
			name:   "specified",
			target: testGroup1.Clone(),
			args:   args{p: PointFieldOnly(testField1.Field())},
			want:   []*Field{},
		},
		{
			name:   "not found",
			target: testGroup1.Clone(),
			args:   args{p: PointFieldOnly("xxxxxx")},
			want:   []*Field{testField1},
		},
		{
			name:   "empty",
			target: &Group{},
			args:   args{p: PointFieldOnly(testField1.Field())},
			want:   nil,
		},
		{
			name:   "nil",
			target: nil,
			args:   args{p: PointFieldOnly(testField1.Field())},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.target.RemoveFields(tt.args.p)
			if tt.target != nil {
				assert.Equal(t, tt.want, tt.target.fields)
			}
		})
	}
}
