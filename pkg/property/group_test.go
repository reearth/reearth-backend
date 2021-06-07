package property

import (
	"errors"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestGroup_IDRef(t *testing.T) {
	gid := id.NewPropertyItemID()
	var g *Group
	assert.Nil(t, g.IDRef())
	g = NewGroup().ID(gid).MustBuild()
	assert.Equal(t, gid.Ref(), g.IDRef())
}

func TestGroup_SchemaGroup(t *testing.T) {
	var g *Group
	assert.Nil(t, g.SchemaGroupRef())
	assert.Equal(t, id.PropertySchemaFieldID(""), g.SchemaGroup())
	pfid := id.PropertySchemaFieldID("aa")
	g = NewGroup().NewID().Schema(id.MustPropertySchemaID("xx/aa"), pfid).MustBuild()
	assert.Equal(t, pfid, g.SchemaGroup())
	assert.Equal(t, pfid.Ref(), g.SchemaGroupRef())
}

func TestGroup_HasLinkedField(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFromUnsafe("vvv")
	l := NewLink(id.NewDatasetID(), id.NewDatasetSchemaID(), id.NewDatasetSchemaFieldID())
	ls := NewLinks([]*Link{l})
	f := NewField(sf).Value(v).Link(ls).MustBuild()
	f2 := NewField(sf).Value(v).MustBuild()

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
	v := ValueTypeString.ValueFromUnsafe("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	l := NewLink(dsid, dssid, id.NewDatasetSchemaFieldID())
	ls := NewLinks([]*Link{l})
	f := NewField(sf).Value(v).Link(ls).MustBuild()
	f2 := NewField(sf).Value(v).MustBuild()

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

func TestGroup_CollectDatasets(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFromUnsafe("vvv")
	dsid := id.NewDatasetID()
	l := NewLink(dsid, id.NewDatasetSchemaID(), id.NewDatasetSchemaFieldID())
	ls := NewLinks([]*Link{l})
	f := NewField(sf).Value(v).Link(ls).MustBuild()

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
			res := tc.Group.CollectDatasets()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestGroup_FieldsByLinkedDataset(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFromUnsafe("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	l := NewLink(dsid, dssid, id.NewDatasetSchemaFieldID())
	ls := NewLinks([]*Link{l})
	f := NewField(sf).Value(v).Link(ls).MustBuild()

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
	v := ValueTypeString.ValueFromUnsafe("vvv")
	f := NewField(sf).Value(v).MustBuild()
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
	v := ValueTypeString.ValueFromUnsafe("vvv")
	f := NewField(sf).Value(v).MustBuild()
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
			assert.Equal(tt, tc.Expected, tc.Group.Fields())
		})
	}
}

func TestGroup_GetOrCreateField(t *testing.T) {
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	f := NewField(sf).MustBuild()
	sg := NewSchemaGroup().ID("aa").Schema(id.MustPropertySchemaID("xx/aa")).Fields([]*SchemaField{sf}).MustBuild()
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
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx/aa"), "aa").MustBuild(),
		},
		{
			Name:  "group schema doesn't equal to ps",
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xxx/aaa"), "aa").MustBuild(),
			PS:    NewSchema().ID(id.MustPropertySchemaID("xx/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
		},
		{
			Name:  "create field",
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx/aa"), "aa").MustBuild(),
			PS:    NewSchema().ID(id.MustPropertySchemaID("xx/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
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
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx/aa"), "aa").Fields([]*Field{f}).MustBuild(),
			PS:    NewSchema().ID(id.MustPropertySchemaID("xx/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
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
	v := ValueTypeString.ValueFromUnsafe("vvv")
	f := NewField(sf).Value(v).MustBuild()
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
			assert.Equal(tt, tc.Expected, tc.Group.Fields())
		})
	}
}

func TestGroup_FieldIDs(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	sf2 := NewSchemaField().ID("b").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFromUnsafe("vvv")
	f := NewField(sf).Value(v).MustBuild()
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
	v := ValueTypeString.ValueFromUnsafe("vvv")
	f := NewField(sf).Value(v).MustBuild()
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
	sg := NewSchemaGroup().ID("aa").Schema(id.MustPropertySchemaID("xx/aa")).Fields([]*SchemaField{sf}).MustBuild()
	sg2 := NewSchemaGroup().ID("bb").Schema(id.MustPropertySchemaID("xx/bb")).Fields([]*SchemaField{sf}).MustBuild()
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
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xx/aa"), "aa").MustBuild(),
		},
		{
			Name:  "group schema doesn't equal to ps",
			Group: NewGroup().NewID().Schema(id.MustPropertySchemaID("xxx/aaa"), "aa").MustBuild(),
			PS:    NewSchema().ID(id.MustPropertySchemaID("xx/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
		},
		{
			Name:     "update value",
			Group:    NewGroup().NewID().Schema(id.MustPropertySchemaID("xx/aa"), "aa").MustBuild(),
			PS:       NewSchema().ID(id.MustPropertySchemaID("xx/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Value:    ValueTypeString.ValueFromUnsafe("abc"),
			FID:      "aa",
			Expected: NewField(sf).Value(ValueTypeString.ValueFromUnsafe("abc")).MustBuild(),
		},
		{
			Name:     "invalid property field",
			Group:    NewGroup().NewID().Schema(id.MustPropertySchemaID("xx/aa"), "aa").MustBuild(),
			PS:       NewSchema().ID(id.MustPropertySchemaID("xx/bb")).Groups([]*SchemaGroup{sg2}).MustBuild(),
			Value:    ValueTypeString.ValueFromUnsafe("abc"),
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

func TestGroup_MigrateGroup(t *testing.T) {
	sid := id.MustPropertySchemaID("xxx/aaa")
	sid2 := id.MustPropertySchemaID("xxx/bbb")
	sid3 := id.MustPropertySchemaID("xxx/ccc")
	sfid1 := id.PropertySchemaFieldID("a")
	sfid2 := id.PropertySchemaFieldID("b")
	sfid3 := id.PropertySchemaFieldID("c")
	nsg := NewSchemaGroup().ID("b").Schema(sid2).Fields([]*SchemaField{
		NewSchemaField().ID(sfid1).Type(ValueTypeString).MustBuild(),
	}).MustBuild()
	nsg2 := NewSchemaGroup().ID("b").Schema(sid3).Fields([]*SchemaField{
		NewSchemaField().ID(sfid1).Type(ValueTypeString).MustBuild(),
		NewSchemaField().ID(sfid3).Type(ValueTypeString).MustBuild(),
	}).MustBuild()
	osg := NewSchemaGroup().ID("c").Schema(sid).Fields([]*SchemaField{
		NewSchemaField().ID(sfid1).Type(ValueTypeString).MustBuild(),
		NewSchemaField().ID(sfid2).Type(ValueTypeString).MustBuild(),
	}).MustBuild()
	fields := []*Field{
		// should remain
		NewFieldUnsafe().FieldUnsafe(sfid1).
			ValueUnsafe(ValueTypeString.ValueFromUnsafe("foobar")).
			Build(),
		// should remove
		NewFieldUnsafe().FieldUnsafe(sfid2).
			ValueUnsafe(ValueTypeString.ValueFromUnsafe("xxx")).
			Build(),
	}
	fields2 := []*Field{
		// should remain
		NewFieldUnsafe().FieldUnsafe(sfid1).
			ValueUnsafe(ValueTypeString.ValueFromUnsafe("foobar")).
			Build(),
		// should be renamed
		NewFieldUnsafe().FieldUnsafe(sfid2).
			ValueUnsafe(ValueTypeString.ValueFromUnsafe("xxx")).
			Build(),
	}
	oldGroup := NewGroup().NewID().Schema(sid, osg.ID()).Fields(nil).MustBuild()
	oldGroup1 := NewGroup().NewID().Schema(sid, osg.ID()).Fields(fields).MustBuild()
	oldGroup2 := NewGroup().NewID().Schema(sid, osg.ID()).Fields(fields2).MustBuild()
	testCases := []struct {
		Name           string
		NewSchemaGroup *SchemaGroup
		Plan           MigrationPlan
		Schema         id.PropertySchemaID
		OldGroup       *Group
		Expected       *Group
	}{
		{
			Name:           "no migration -> nil old group",
			NewSchemaGroup: nsg,
			OldGroup:       nil,
		},
		{
			Name:           "no migration -> nil schemaGroup",
			NewSchemaGroup: nil,
			OldGroup:       oldGroup,
			Expected:       oldGroup,
		},
		{
			Name:           "Rename field",
			NewSchemaGroup: nsg2,
			OldGroup:       oldGroup2,
			Schema:         sid3,
			Plan: MigrationPlan{
				From: "b",
				To:   "c",
			},
			Expected: NewGroup().ID(oldGroup2.ID()).Schema(sid3, nsg2.ID()).Fields([]*Field{
				NewFieldUnsafe().FieldUnsafe(sfid1).
					ValueUnsafe(ValueTypeString.ValueFromUnsafe("foobar")).
					Build(),
				NewFieldUnsafe().FieldUnsafe(sfid3).
					ValueUnsafe(ValueTypeString.ValueFromUnsafe("xxx")).
					Build()}).MustBuild(),
		},
		{
			Name:           "Remove: no migration plan",
			NewSchemaGroup: nsg,
			OldGroup:       oldGroup1,
			Schema:         sid2,
			Expected: NewGroup().ID(oldGroup1.ID()).Schema(sid2, nsg.ID()).Fields([]*Field{
				NewFieldUnsafe().FieldUnsafe(sfid1).
					ValueUnsafe(ValueTypeString.ValueFromUnsafe("foobar")).
					Build()}).MustBuild(),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.OldGroup.MigrateGroup(tc.Schema, tc.NewSchemaGroup, tc.Plan)
			assert.Equal(tt, tc.Expected, tc.OldGroup)
		})
	}
}
