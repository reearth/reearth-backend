package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

var (
	testGroupList1 = NewGroupList().NewID().Schema(testSchema1.ID(), testSchemaGroup2.ID()).Groups([]*Group{testGroup2}).MustBuild()
)

func TestGroupList_HasLinkedField(t *testing.T) {
	pid := id.NewPropertyItemID()
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	f := NewField(sf).Value(OptionalValueFrom(v)).Link(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())})).MustBuild()
	groups := []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()}
	groups2 := []*Group{NewGroup().ID(pid).MustBuild()}
	testCases := []struct {
		Name     string
		GL       *GroupList
		Expected bool
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "has linked field",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups(groups).MustBuild(),
			Expected: true,
		},
		{
			Name:     "no linked field",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups(groups2).MustBuild(),
			Expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.GL.HasLinkedField())
			assert.Equal(tt, tc.Expected, tc.GL.IsDatasetLinked(dssid, dsid))
		})
	}
}

func TestGroupList_Datasets(t *testing.T) {
	pid := id.NewPropertyItemID()
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	f := NewField(sf).Value(OptionalValueFrom(v)).Link(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())})).MustBuild()
	groups := []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()}
	groups2 := []*Group{NewGroup().ID(pid).MustBuild()}
	testCases := []struct {
		Name     string
		GL       *GroupList
		Expected []id.DatasetID
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "one dataset",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups(groups).MustBuild(),
			Expected: []id.DatasetID{dsid},
		},
		{
			Name:     "empty list",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups(groups2).MustBuild(),
			Expected: []id.DatasetID{},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.GL.Datasets())
		})
	}
}

func TestGroupList_FieldsByLinkedDataset(t *testing.T) {
	pid := id.NewPropertyItemID()
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	f := NewField(sf).Value(OptionalValueFrom(v)).Link(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())})).MustBuild()
	groups := []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()}
	groups2 := []*Group{NewGroup().ID(pid).MustBuild()}
	testCases := []struct {
		Name     string
		GL       *GroupList
		Expected []*Field
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "one field list",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups(groups).MustBuild(),
			Expected: []*Field{f},
		},
		{
			Name:     "empty list",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups(groups2).MustBuild(),
			Expected: []*Field{},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.GL.FieldsByLinkedDataset(dssid, dsid))
		})
	}
}

func TestGroupList_IsEmpty(t *testing.T) {
	pid := id.NewPropertyItemID()
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	f := NewField(sf).Value(OptionalValueFrom(v)).Link(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())})).MustBuild()
	groups := []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()}
	testCases := []struct {
		Name     string
		GL       *GroupList
		Expected bool
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "is empty",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").MustBuild(),
			Expected: true,
		},
		{
			Name:     "is not empty",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups(groups).MustBuild(),
			Expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.GL.IsEmpty())
		})
	}
}

func TestGroupList_Prune(t *testing.T) {
	sf := NewSchemaField().ID("a").Type(ValueTypeString).MustBuild()
	v := ValueTypeString.ValueFrom("vvv")
	f := NewField(sf).Value(OptionalValueFrom(v)).MustBuild()
	f2 := NewField(sf).MustBuild()
	pid := id.NewPropertyItemID()
	groups := []*Group{NewGroup().ID(pid).Fields([]*Field{f, f2}).MustBuild()}
	pruned := []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()}
	testCases := []struct {
		Name     string
		GL       *GroupList
		Expected []*Group
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "pruned list",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups(groups).MustBuild(),
			Expected: pruned,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.GL.Prune()
			assert.Equal(tt, tc.Expected, tc.GL.Groups())
		})
	}
}

func TestGroupList_GetGroup(t *testing.T) {
	pid := id.NewPropertyItemID()
	g := NewGroup().ID(pid).MustBuild()
	testCases := []struct {
		Name     string
		Input    id.PropertyItemID
		GL       *GroupList
		Expected *Group
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "found",
			Input:    pid,
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g}).MustBuild(),
			Expected: g,
		},
		{
			Name:     "not found",
			Input:    id.NewPropertyItemID(),
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g}).MustBuild(),
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.GL.Group(tc.Input))
		})
	}
}

func TestGroupList_GroupAt(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	testCases := []struct {
		Name     string
		Index    int
		GL       *GroupList
		Expected *Group
	}{
		{
			Name: "nil group list",
		},
		{
			Name:  "index < 0",
			Index: -1,
		},
		{
			Name:  "index > len(g)-1",
			Index: 4,
		},
		{
			Name:     "found",
			Index:    2,
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: g3,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.GL.GroupAt(tc.Index))
		})
	}
}

func TestGroupList_Has(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	testCases := []struct {
		Name     string
		Input    id.PropertyItemID
		GL       *GroupList
		Expected bool
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "found",
			Input:    g2.ID(),
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: true,
		},
		{
			Name:     "not found",
			Input:    g3.ID(),
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g4}).MustBuild(),
			Expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.GL.Has(tc.Input))
		})
	}
}

func TestGroupList_Count(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	testCases := []struct {
		Name     string
		GL       *GroupList
		Expected int
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "not found",
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: 4,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.GL.Count())
		})
	}
}

func TestGroupList_Add(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	testCases := []struct {
		Name     string
		GL       *GroupList
		Gr       *Group
		Index    int
		Expected struct {
			Gr    *Group
			Index int
		}
	}{
		{
			Name: "nil group list",
		},
		{
			Name:  "index < 0",
			Index: -1,
			Gr:    g2,
			GL:    NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
			Expected: struct {
				Gr    *Group
				Index int
			}{
				Gr:    g2,
				Index: 3,
			},
		},
		{
			Name:  "len(g) > index > 0 ",
			Index: 2,
			Gr:    g2,
			GL:    NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
			Expected: struct {
				Gr    *Group
				Index int
			}{
				Gr:    g2,
				Index: 2,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.GL.Add(tc.Gr, tc.Index)
			assert.Equal(tt, tc.Expected.Gr, tc.GL.GroupAt(tc.Expected.Index))
		})
	}
}

func TestGroupList_AddOrMove(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	testCases := []struct {
		Name     string
		GL       *GroupList
		Gr       *Group
		Index    int
		Expected struct {
			Gr    *Group
			Index int
		}
	}{
		{
			Name: "nil group list",
		},
		{
			Name:  "index < 0",
			Index: -1,
			Gr:    g2,
			GL:    NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
			Expected: struct {
				Gr    *Group
				Index int
			}{
				Gr:    g2,
				Index: 3,
			},
		},
		{
			Name:  "len(g) > index > 0 ",
			Index: 2,
			Gr:    g2,
			GL:    NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
			Expected: struct {
				Gr    *Group
				Index int
			}{
				Gr:    g2,
				Index: 2,
			},
		},
		{
			Name:  "move group",
			Index: 2,
			Gr:    g1,
			GL:    NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
			Expected: struct {
				Gr    *Group
				Index int
			}{
				Gr:    g1,
				Index: 2,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.GL.AddOrMove(tc.Gr, tc.Index)
			assert.Equal(tt, tc.Expected.Gr, tc.GL.GroupAt(tc.Expected.Index))
		})
	}
}

func TestGroupList_Move(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	testCases := []struct {
		Name     string
		GL       *GroupList
		Id       id.PropertyItemID
		ToIndex  int
		Expected struct {
			Id    id.PropertyItemID
			Index int
		}
	}{
		{
			Name: "nil group list",
		},
		{
			Name:    "success",
			Id:      g1.ID(),
			ToIndex: 2,
			GL:      NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: struct {
				Id    id.PropertyItemID
				Index int
			}{Id: g1.ID(), Index: 2},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.GL.Move(tc.Id, tc.ToIndex)
			assert.Equal(tt, tc.Expected.Id, tc.GL.GroupAt(tc.Expected.Index).ID())
		})
	}
}

func TestGroupList_MoveAt(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	testCases := []struct {
		Name               string
		GL                 *GroupList
		FromIndex, ToIndex int
		Expected           []*Group
	}{
		{
			Name: "nil group list",
		},
		{
			Name:      "from = to",
			FromIndex: 2,
			ToIndex:   2,
			GL:        NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected:  []*Group{g1, g2, g3, g4},
		},
		{
			Name:      "from < 0",
			FromIndex: -1,
			ToIndex:   2,
			GL:        NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected:  []*Group{g1, g2, g3, g4},
		},
		{
			Name:      "success move",
			FromIndex: 0,
			ToIndex:   2,
			GL:        NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected:  []*Group{g2, g3, g1, g4},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.GL.MoveAt(tc.FromIndex, tc.ToIndex)
			assert.Equal(tt, tc.Expected, tc.GL.Groups())
		})
	}
}

func TestGroupList_RemoveAt(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	testCases := []struct {
		Name     string
		GL       *GroupList
		Index    int
		Expected []*Group
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "success",
			Index:    1,
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: []*Group{g1, g3, g4},
		},
		{
			Name:     "index < 0",
			Index:    -1,
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: []*Group{g1, g2, g3, g4},
		},
		{
			Name:     "index > length",
			Index:    5,
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: []*Group{g1, g2, g3, g4},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.GL.RemoveAt(tc.Index)
			assert.Equal(tt, tc.Expected, tc.GL.Groups())
		})
	}
}
func TestGroupList_Remove(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	testCases := []struct {
		Name     string
		GL       *GroupList
		Input    id.PropertyItemID
		Expected bool
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "success",
			Input:    g1.ID(),
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: true,
		},
		{
			Name:     "not found",
			Input:    g4.ID(),
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "xx").Groups([]*Group{g1, g2, g3}).MustBuild(),
			Expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.GL.Remove(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestGroupList_GetOrCreateField(t *testing.T) {
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	sg := NewSchemaGroup().ID("aa").Schema(id.MustPropertySchemaID("xx~1.0.0/aa")).Fields([]*SchemaField{sf}).MustBuild()
	g := NewGroup().ID(id.NewPropertyItemID()).Schema(sg.Schema(), sg.ID()).MustBuild()
	testCases := []struct {
		Name     string
		GL       *GroupList
		Schema   *Schema
		Ptr      *Pointer
		Expected struct {
			Ok    bool
			Field *Field
		}
	}{
		{
			Name:   "success",
			GL:     NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").Groups([]*Group{g}).MustBuild(),
			Schema: NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Ptr:    NewPointer(nil, g.ID().Ref(), sf.ID().Ref()),
			Expected: struct {
				Ok    bool
				Field *Field
			}{
				Ok:    true,
				Field: NewField(sf).MustBuild(),
			},
		},
		{
			Name:   "can't get a group",
			GL:     NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").MustBuild(),
			Schema: NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Ptr:    NewPointer(nil, g.ID().Ref(), sf.ID().Ref()),
		},
		{
			Name:   "FieldByItem not ok: sg!=nil",
			GL:     NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").Groups([]*Group{g}).MustBuild(),
			Schema: NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Ptr:    NewPointer(sg.IDRef(), g.ID().Ref(), sf.ID().Ref()),
		},
		{
			Name:   "psg == nil",
			GL:     NewGroupList().NewID().Groups([]*Group{g}).MustBuild(),
			Schema: NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Ptr:    NewPointer(nil, g.ID().Ref(), sf.ID().Ref()),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res, ok := tc.GL.GetOrCreateField(tc.Schema, tc.Ptr)
			assert.Equal(tt, tc.Expected.Field, res)
			assert.Equal(tt, tc.Expected.Ok, ok)
		})
	}
}

func TestGroupList_CreateAndAddListItem(t *testing.T) {
	getIntRef := func(i int) *int { return &i }
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	sg := NewSchemaGroup().ID("aa").Schema(id.MustPropertySchemaID("xx~1.0.0/aa")).Fields([]*SchemaField{sf}).MustBuild()
	g := NewGroup().ID(id.NewPropertyItemID()).Schema(sg.Schema(), sg.ID()).MustBuild()
	testCases := []struct {
		Name     string
		GL       *GroupList
		Schema   *Schema
		Index    *int
		Expected *Group
	}{
		{
			Name:     "success",
			Index:    getIntRef(0),
			GL:       NewGroupList().NewID().Schema(id.MustPropertySchemaID("xx~1.0.0/aa"), "aa").MustBuild(),
			Schema:   NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Expected: g,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.GL.CreateAndAddListItem(tc.Schema, tc.Index)
			assert.Equal(tt, tc.Expected.Schema(), res.Schema())
			assert.Equal(tt, tc.Expected.Fields(nil), res.Fields(nil))
			assert.Equal(tt, tc.Expected.SchemaGroup(), res.SchemaGroup())
		})
	}
}

func TestGroupList_Clone(t *testing.T) {
	tests := []struct {
		name   string
		target *GroupList
		n      bool
	}{
		{
			name:   "ok",
			target: testGroupList1.Clone(),
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

func TestGroupList_CloneItem(t *testing.T) {
	tests := []struct {
		name   string
		target *GroupList
		n      bool
	}{
		{
			name:   "ok",
			target: testGroupList1.Clone(),
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

func TestGroupList_Fields(t *testing.T) {
	type args struct {
		p *Pointer
	}
	tests := []struct {
		name   string
		target *GroupList
		args   args
		want   []*Field
	}{
		{
			name:   "all",
			target: testGroupList1,
			args:   args{p: nil},
			want:   []*Field{testField2},
		},
		{
			name:   "specified",
			target: testGroupList1,
			args:   args{p: PointFieldOnly(testField2.Field())},
			want:   []*Field{testField2},
		},
		{
			name:   "not found",
			target: testGroupList1,
			args:   args{p: PointFieldOnly("xxxxxx")},
			want:   []*Field{},
		},
		{
			name:   "empty",
			target: &GroupList{},
			args:   args{p: PointFieldOnly(testField2.Field())},
			want:   nil,
		},
		{
			name:   "nil",
			target: nil,
			args:   args{p: PointFieldOnly(testField2.Field())},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Fields(tt.args.p))
		})
	}
}

func TestGroupList_RemoveFields(t *testing.T) {
	type args struct {
		p *Pointer
	}
	tests := []struct {
		name   string
		target *GroupList
		args   args
		want   []*Field
	}{
		{
			name:   "nil pointer",
			target: testGroupList1.Clone(),
			args:   args{p: nil},
			want:   []*Field{testField2},
		},
		{
			name:   "specified",
			target: testGroupList1.Clone(),
			args:   args{p: PointFieldOnly(testField2.Field())},
			want:   nil,
		},
		{
			name:   "specified schema group",
			target: testGroupList1.Clone(),
			args:   args{p: PointItemBySchema(testGroupList1.SchemaGroup())},
			want:   nil,
		},
		{
			name:   "specified item",
			target: testGroupList1.Clone(),
			args:   args{p: PointItem(testGroupList1.ID())},
			want:   nil,
		},
		{
			name:   "not found",
			target: testGroupList1.Clone(),
			args:   args{p: PointFieldOnly("xxxxxx")},
			want:   []*Field{testField2},
		},
		{
			name:   "empty",
			target: &GroupList{},
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
				assert.Equal(t, tt.want, tt.target.Fields(nil))
			}
		})
	}
}
