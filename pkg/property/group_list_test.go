package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

var (
	testGroupList1 = NewGroupList().NewID().Schema(testSchemaGroup2.ID()).Groups([]*Group{testGroup2}).MustBuild()
)

func TestGroupList_HasLinkedField(t *testing.T) {
	pid := id.NewPropertyItemID()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	f := NewField().Field("a").Value(OptionalValueFrom(v)).Link(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())})).Build()
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
			GL:       NewGroupList().NewID().Schema("xx").Groups(groups).MustBuild(),
			Expected: true,
		},
		{
			Name:     "no linked field",
			GL:       NewGroupList().NewID().Schema("xx").Groups(groups2).MustBuild(),
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
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	f := NewField().Field("a").Value(OptionalValueFrom(v)).Link(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())})).Build()
	groups := []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()}
	groups2 := []*Group{NewGroup().ID(pid).MustBuild()}

	tests := []struct {
		Name     string
		Target   *GroupList
		Expected []id.DatasetID
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "one dataset",
			Target:   NewGroupList().NewID().Schema("xx").Groups(groups).MustBuild(),
			Expected: []id.DatasetID{dsid},
		},
		{
			Name:     "empty list",
			Target:   NewGroupList().NewID().Schema("xx").Groups(groups2).MustBuild(),
			Expected: []id.DatasetID{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.Expected, tt.Target.Datasets())
		})
	}
}

func TestGroupList_FieldsByLinkedDataset(t *testing.T) {
	pid := id.NewPropertyItemID()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	f := NewField().Field("a").Value(OptionalValueFrom(v)).Link(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())})).Build()
	groups := []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()}
	groups2 := []*Group{NewGroup().ID(pid).MustBuild()}

	tests := []struct {
		Name     string
		Target   *GroupList
		Expected []*Field
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "one field list",
			Target:   NewGroupList().NewID().Schema("xx").Groups(groups).MustBuild(),
			Expected: []*Field{f},
		},
		{
			Name:     "empty list",
			Target:   NewGroupList().NewID().Schema("xx").Groups(groups2).MustBuild(),
			Expected: []*Field{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.Expected, tt.Target.FieldsByLinkedDataset(dssid, dsid))
		})
	}
}

func TestGroupList_IsEmpty(t *testing.T) {
	pid := id.NewPropertyItemID()
	v := ValueTypeString.ValueFrom("vvv")
	dsid := id.NewDatasetID()
	dssid := id.NewDatasetSchemaID()
	f := NewField().Field("a").Value(OptionalValueFrom(v)).Link(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(dsid, dssid, id.NewDatasetSchemaFieldID())})).Build()
	groups := []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()}

	tests := []struct {
		Name     string
		Target   *GroupList
		Expected bool
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "is empty",
			Target:   NewGroupList().NewID().Schema("xx").MustBuild(),
			Expected: true,
		},
		{
			Name:     "is not empty",
			Target:   NewGroupList().NewID().Schema("xx").Groups(groups).MustBuild(),
			Expected: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.Expected, tt.Target.IsEmpty())
		})
	}
}

func TestGroupList_Prune(t *testing.T) {
	v := ValueTypeString.ValueFrom("vvv")
	f := NewField().Field("a").Value(OptionalValueFrom(v)).Build()
	f2 := NewField().Field("b").Value(NewOptionalValue(ValueTypeString, nil)).Build()
	pid := id.NewPropertyItemID()

	tests := []struct {
		name       string
		target     *GroupList
		want       bool
		wantGroups []*Group
	}{
		{
			name: "ok",
			target: NewGroupList().NewID().Schema("xx").Groups(
				[]*Group{NewGroup().ID(pid).Fields([]*Field{f, f2}).MustBuild()},
			).MustBuild(),
			want:       true,
			wantGroups: []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()},
		},
		{
			name: "no empty fields",
			target: NewGroupList().NewID().Schema("xx").Groups(
				[]*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()},
			).MustBuild(),
			want:       false,
			wantGroups: []*Group{NewGroup().ID(pid).Fields([]*Field{f}).MustBuild()},
		},
		{
			name:       "nil",
			target:     nil,
			want:       false,
			wantGroups: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.target.Prune())
			assert.Equal(t, tt.wantGroups, tt.target.Groups())
		})
	}
}

func TestGroupList_Group(t *testing.T) {
	pid := id.NewPropertyItemID()
	g := NewGroup().ID(pid).MustBuild()

	tests := []struct {
		Name     string
		Input    id.PropertyItemID
		Target   *GroupList
		Expected *Group
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "found",
			Input:    pid,
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g}).MustBuild(),
			Expected: g,
		},
		{
			Name:     "not found",
			Input:    id.NewPropertyItemID(),
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g}).MustBuild(),
			Expected: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.Expected, tt.Target.Group(tt.Input))
		})
	}
}

func TestGroupList_GroupByPointer(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	gl := NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild()

	tests := []struct {
		Name     string
		Target   *GroupList
		Args     *Pointer
		Expected *Group
	}{
		{
			Name:     "nil",
			Target:   nil,
			Args:     PointItem(g3.ID()),
			Expected: nil,
		},
		{
			Name:     "nil pointer",
			Target:   gl,
			Args:     nil,
			Expected: nil,
		},
		{
			Name:     "not found",
			Target:   gl,
			Args:     PointItem(NewItemID()),
			Expected: nil,
		},
		{
			Name:     "found",
			Target:   gl,
			Args:     PointItem(g3.ID()),
			Expected: g3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.Expected, tt.Target.GroupByPointer(tt.Args))
		})
	}
}

func TestGroupList_GroupAt(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	gl := NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild()

	tests := []struct {
		Name     string
		Index    int
		Target   *GroupList
		Expected *Group
	}{
		{
			Name: "nil group list",
		},
		{
			Name:   "index < 0",
			Index:  -1,
			Target: gl,
		},
		{
			Name:   "index > len(g)-1",
			Index:  4,
			Target: gl,
		},
		{
			Name:     "found",
			Index:    2,
			Target:   gl,
			Expected: g3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.Expected, tt.Target.GroupAt(tt.Index))
		})
	}
}

func TestGroupList_Has(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()

	tests := []struct {
		Name     string
		Input    id.PropertyItemID
		Target   *GroupList
		Expected bool
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "found",
			Input:    g2.ID(),
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: true,
		},
		{
			Name:     "not found",
			Input:    g3.ID(),
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g4}).MustBuild(),
			Expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.Expected, tt.Target.Has(tt.Input))
		})
	}
}

func TestGroupList_Count(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()

	tests := []struct {
		Name     string
		Target   *GroupList
		Expected int
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "not found",
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: 4,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.Target.Count())
		})
	}
}

func TestGroupList_Add(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()

	tests := []struct {
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
			GL:    NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
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
			GL:    NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
			Expected: struct {
				Gr    *Group
				Index int
			}{
				Gr:    g2,
				Index: 2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			tt.GL.Add(tt.Gr, tt.Index)
			assert.Equal(t, tt.Expected.Gr, tt.GL.GroupAt(tt.Expected.Index))
		})
	}
}

func TestGroupList_AddOrMove(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()

	tests := []struct {
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
			GL:    NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
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
			GL:    NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
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
			GL:    NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g3, g4}).MustBuild(),
			Expected: struct {
				Gr    *Group
				Index int
			}{
				Gr:    g1,
				Index: 2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			tt.GL.AddOrMove(tt.Gr, tt.Index)
			assert.Equal(t, tt.Expected.Gr, tt.GL.GroupAt(tt.Expected.Index))
		})
	}
}

func TestGroupList_Move(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()

	tests := []struct {
		Name     string
		Target   *GroupList
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
			Target:  NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: struct {
				Id    id.PropertyItemID
				Index int
			}{Id: g1.ID(), Index: 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			tt.Target.Move(tt.Id, tt.ToIndex)
			assert.Equal(t, tt.Expected.Id, tt.Target.GroupAt(tt.Expected.Index).ID())
		})
	}
}

func TestGroupList_MoveAt(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()

	tests := []struct {
		Name               string
		Target             *GroupList
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
			Target:    NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected:  []*Group{g1, g2, g3, g4},
		},
		{
			Name:      "from < 0",
			FromIndex: -1,
			ToIndex:   2,
			Target:    NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected:  []*Group{g1, g2, g3, g4},
		},
		{
			Name:      "success move",
			FromIndex: 0,
			ToIndex:   2,
			Target:    NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected:  []*Group{g2, g3, g1, g4},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			tt.Target.MoveAt(tt.FromIndex, tt.ToIndex)
			assert.Equal(t, tt.Expected, tt.Target.Groups())
		})
	}
}

func TestGroupList_RemoveAt(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()

	tests := []struct {
		Name     string
		Target   *GroupList
		Index    int
		Expected []*Group
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "success",
			Index:    1,
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: []*Group{g1, g3, g4},
		},
		{
			Name:     "index < 0",
			Index:    -1,
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: []*Group{g1, g2, g3, g4},
		},
		{
			Name:     "index > length",
			Index:    5,
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: []*Group{g1, g2, g3, g4},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			tt.Target.RemoveAt(tt.Index)
			assert.Equal(t, tt.Expected, tt.Target.Groups())
		})
	}
}

func TestGroupList_Remove(t *testing.T) {
	g1 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g2 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g3 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()
	g4 := NewGroup().ID(id.NewPropertyItemID()).MustBuild()

	tests := []struct {
		Name     string
		Target   *GroupList
		Input    id.PropertyItemID
		Expected bool
	}{
		{
			Name: "nil group list",
		},
		{
			Name:     "success",
			Input:    g1.ID(),
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3, g4}).MustBuild(),
			Expected: true,
		},
		{
			Name:     "not found",
			Input:    g4.ID(),
			Target:   NewGroupList().NewID().Schema("xx").Groups([]*Group{g1, g2, g3}).MustBuild(),
			Expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res := tt.Target.Remove(tt.Input)
			assert.Equal(t, tt.Expected, res)
		})
	}
}

func TestGroupList_GetOrCreateField(t *testing.T) {
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	sg := NewSchemaGroup().ID("aa").Fields([]*SchemaField{sf}).MustBuild()
	g := NewGroup().ID(id.NewPropertyItemID()).Schema(sg.ID()).MustBuild()

	tests := []struct {
		Name     string
		Target   *GroupList
		Schema   *Schema
		Ptr      *Pointer
		Expected struct {
			Ok    bool
			Field *Field
		}
	}{
		{
			Name:   "success",
			Target: NewGroupList().NewID().Schema("aa").Groups([]*Group{g}).MustBuild(),
			Schema: NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Ptr:    NewPointer(nil, g.ID().Ref(), sf.ID().Ref()),
			Expected: struct {
				Ok    bool
				Field *Field
			}{
				Ok:    true,
				Field: NewField().Field("aa").Value(NewOptionalValue(ValueTypeString, nil)).Build(),
			},
		},
		{
			Name:   "can't get a group",
			Target: NewGroupList().NewID().Schema("aa").MustBuild(),
			Schema: NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Ptr:    NewPointer(nil, g.ID().Ref(), sf.ID().Ref()),
		},
		{
			Name:   "FieldByItem not ok: sg!=nil",
			Target: NewGroupList().NewID().Schema("aa").Groups([]*Group{g}).MustBuild(),
			Schema: NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Ptr:    NewPointer(sg.ID().Ref(), g.ID().Ref(), sf.ID().Ref()),
		},
		{
			Name:   "psg == nil",
			Target: NewGroupList().NewID().Groups([]*Group{g}).MustBuild(),
			Schema: NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Ptr:    NewPointer(nil, g.ID().Ref(), sf.ID().Ref()),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res, ok := tt.Target.GetOrCreateField(tt.Schema, tt.Ptr)
			assert.Equal(t, tt.Expected.Field, res)
			assert.Equal(t, tt.Expected.Ok, ok)
		})
	}
}

func TestGroupList_CreateAndAddListItem(t *testing.T) {
	getIntRef := func(i int) *int { return &i }
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	sg := NewSchemaGroup().ID("aa").Fields([]*SchemaField{sf}).MustBuild()
	g := NewGroup().ID(id.NewPropertyItemID()).Schema(sg.ID()).MustBuild()

	tests := []struct {
		Name     string
		GL       *GroupList
		Schema   *Schema
		Index    *int
		Expected *Group
	}{
		{
			Name:     "success",
			Index:    getIntRef(0),
			GL:       NewGroupList().NewID().Schema("aa").MustBuild(),
			Schema:   NewSchema().ID(id.MustPropertySchemaID("xx~1.0.0/aa")).Groups([]*SchemaGroup{sg}).MustBuild(),
			Expected: g,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res := tt.GL.CreateAndAddListItem(tt.Schema, tt.Index)
			assert.Equal(t, tt.Expected.Fields(nil), res.Fields(nil))
			assert.Equal(t, tt.Expected.SchemaGroup(), res.SchemaGroup())
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
			want:   nil,
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
		name       string
		target     *GroupList
		args       args
		want       bool
		wantFields []*Field
	}{
		{
			name:       "nil pointer",
			target:     testGroupList1.Clone(),
			args:       args{p: nil},
			want:       false,
			wantFields: []*Field{testField2},
		},
		{
			name:       "specified",
			target:     testGroupList1.Clone(),
			args:       args{p: PointFieldOnly(testField2.Field())},
			want:       true,
			wantFields: nil,
		},
		{
			name:       "specified schema group",
			target:     testGroupList1.Clone(),
			args:       args{p: PointItemBySchema(testGroupList1.SchemaGroup())},
			want:       false,
			wantFields: []*Field{testField2},
		},
		{
			name:       "specified item",
			target:     testGroupList1.Clone(),
			args:       args{p: PointItem(testGroupList1.ID())},
			want:       false,
			wantFields: []*Field{testField2},
		},
		{
			name:       "not found",
			target:     testGroupList1.Clone(),
			args:       args{p: PointFieldOnly("xxxxxx")},
			want:       false,
			wantFields: []*Field{testField2},
		},
		{
			name:       "empty",
			target:     &GroupList{},
			args:       args{p: PointFieldOnly(testField1.Field())},
			want:       false,
			wantFields: nil,
		},
		{
			name:       "nil",
			target:     nil,
			args:       args{p: PointFieldOnly(testField1.Field())},
			want:       false,
			wantFields: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.RemoveFields(tt.args.p))
			if tt.target != nil {
				assert.Equal(t, tt.wantFields, tt.target.Fields(nil))
			}
		})
	}
}

func TestGroupList_GroupAndFields(t *testing.T) {
	tests := []struct {
		name   string
		target *GroupList
		want   []GroupAndField
	}{
		{
			name:   "ok",
			target: testGroupList1,
			want: []GroupAndField{
				{ParentGroup: testGroupList1, Group: testGroup2, Field: testField2},
			},
		},
		{
			name:   "empty",
			target: &GroupList{},
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
			res := tt.target.GroupAndFields()
			assert.Equal(t, tt.want, res)
			for i, r := range res {
				assert.Same(t, tt.want[i].Field, r.Field)
				assert.Same(t, tt.want[i].Group, r.Group)
				assert.Same(t, tt.want[i].ParentGroup, r.ParentGroup)
			}
		})
	}
}

func TestGroupList_GuessSchema(t *testing.T) {
	tests := []struct {
		name   string
		target *GroupList
		want   *SchemaGroup
	}{
		{
			name: "ok",
			target: &GroupList{
				itemBase: itemBase{
					SchemaGroup: "aa",
				},
				groups: []*Group{
					{
						itemBase: itemBase{
							SchemaGroup: "aa",
						},
						fields: []*Field{
							{field: "a", v: NewOptionalValue(ValueTypeLatLng, nil)},
						},
					},
					{
						itemBase: itemBase{
							SchemaGroup: "aa",
						},
						fields: []*Field{
							{field: "b", v: NewOptionalValue(ValueTypeString, nil)},
						},
					},
				},
			},
			want: &SchemaGroup{
				id:   "aa",
				list: true,
				fields: []*SchemaField{
					{id: "a", propertyType: ValueTypeLatLng},
					{id: "b", propertyType: ValueTypeString},
				},
			},
		},
		{
			name:   "empty",
			target: &GroupList{},
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
