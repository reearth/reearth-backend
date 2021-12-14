package property

import (
	"context"
	"testing"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

var (
	testProperty1 = New().NewID().Schema(testSchema1.ID()).Scene(id.NewSceneID()).Items([]Item{testGroup1, testGroupList1}).MustBuild()
)

func TestProperty_MigrateSchema(t *testing.T) {
	sceneID := id.NewSceneID()
	oldSchema := id.MustPropertySchemaID("hoge~1.0.0/test")
	newSchema := id.MustPropertySchemaID("hoge~1.0.0/test2")
	schemaField1ID := id.PropertySchemaFieldID("a")
	schemaField2ID := id.PropertySchemaFieldID("b")
	schemaField3ID := id.PropertySchemaFieldID("c")
	schemaField4ID := id.PropertySchemaFieldID("d")
	schemaField5ID := id.PropertySchemaFieldID("e")
	schemaField6ID := id.PropertySchemaFieldID("f")
	schemaField7ID := id.PropertySchemaFieldID("g")
	schemaField8ID := id.PropertySchemaFieldID("h")
	schemaGroupID := id.PropertySchemaGroupID("i")
	datasetID := id.NewDatasetID()
	datasetSchemaID := id.NewDatasetSchemaID()
	datasetFieldID := id.NewDatasetSchemaFieldID()

	schemaField1, _ := NewSchemaField().ID(schemaField1ID).Type(ValueTypeString).Build()
	schemaField2, _ := NewSchemaField().ID(schemaField2ID).Type(ValueTypeNumber).Min(0).Max(100).Build()
	schemaField3, _ := NewSchemaField().ID(schemaField3ID).Type(ValueTypeNumber).Min(0).Max(100).Build()
	schemaField4, _ := NewSchemaField().ID(schemaField4ID).Type(ValueTypeString).Choices([]SchemaFieldChoice{
		{Title: i18n.StringFrom("x"), Key: "x"},
		{Title: i18n.StringFrom("y"), Key: "y"},
	}).Build()
	schemaField5, _ := NewSchemaField().ID(schemaField5ID).Type(ValueTypeString).Build()
	schemaField6, _ := NewSchemaField().ID(schemaField6ID).Type(ValueTypeNumber).Build()
	schemaField7, _ := NewSchemaField().ID(schemaField7ID).Type(ValueTypeNumber).Build()
	schemaFields := []*SchemaField{
		schemaField1,
		schemaField2,
		schemaField3,
		schemaField4,
		schemaField5,
		schemaField6,
		schemaField7,
	}
	schemaGroups := []*SchemaGroup{
		NewSchemaGroup().ID(schemaGroupID).Fields(schemaFields).MustBuild(),
	}

	fields := []*Field{
		// should remain
		NewField().Field(schemaField1ID).
			Value(OptionalValueFrom(ValueTypeString.ValueFrom("foobar"))).
			Build(),
		// should be removed because of max
		NewField().Field(schemaField2ID).
			Value(OptionalValueFrom(ValueTypeNumber.ValueFrom(101))).
			Build(),
		// should remain
		NewField().Field(schemaField3ID).
			Value(OptionalValueFrom(ValueTypeNumber.ValueFrom(1))).
			Build(),
		// should be removed because of choices
		NewField().Field(schemaField4ID).
			Value(OptionalValueFrom(ValueTypeString.ValueFrom("z"))).
			Build(),
		// should remain
		NewField().Field(schemaField5ID).
			Value(NewOptionalValue(ValueTypeString, nil)).
			Link(dataset.NewGraphPointer([]*dataset.Pointer{
				dataset.PointAt(datasetID, datasetSchemaID, datasetFieldID),
			})).
			Build(),
		// should be removed because of linked dataset field value type
		NewField().Field(schemaField6ID).
			Value(NewOptionalValue(ValueTypeNumber, nil)).
			Link(dataset.NewGraphPointer([]*dataset.Pointer{
				dataset.PointAt(datasetID, datasetSchemaID, datasetFieldID),
			})).
			Build(),
		// should be removed because of type
		NewField().Field(schemaField7ID).
			Value(OptionalValueFrom(ValueTypeString.ValueFrom("hogehoge"))).
			Build(),
		// should be removed because of not existing field
		NewField().Field(schemaField8ID).
			Value(OptionalValueFrom(ValueTypeString.ValueFrom("hogehoge"))).
			Build(),
	}
	items := []Item{
		NewGroup().NewID().Schema(schemaGroupID).Fields(fields).MustBuild(),
	}

	datasetFields := []*dataset.Field{
		dataset.NewField(datasetFieldID, dataset.ValueTypeString.ValueFrom("a"), ""),
	}

	schema, _ := NewSchema().ID(newSchema).Groups(schemaGroups).Build()
	property, _ := New().NewID().Scene(sceneID).Schema(oldSchema).Items(items).Build()
	ds, _ := dataset.New().ID(datasetID).Schema(datasetSchemaID).Scene(sceneID).Fields(datasetFields).Build()

	property.MigrateSchema(context.Background(), schema, dataset.LoaderFrom([]*dataset.Dataset{ds}))

	newGroup := ToGroup(property.ItemBySchema(schemaGroupID))
	newFields := newGroup.Fields(nil)

	assert.Equal(t, schema.ID(), property.Schema())
	assert.Equal(t, 1, len(property.Items()))
	assert.Equal(t, 3, len(newFields))
	assert.NotNil(t, newGroup.Field(schemaField1ID))
	assert.NotNil(t, newGroup.Field(schemaField3ID))
	assert.NotNil(t, newGroup.Field(schemaField5ID))
}

func TestProperty_GetOrCreateItem(t *testing.T) {
	sceneID := id.NewSceneID()
	sid := id.MustPropertySchemaID("hoge~1.0.0/test")
	sf1id := id.PropertySchemaFieldID("a")
	sf2id := id.PropertySchemaFieldID("b")
	sg1id := id.PropertySchemaGroupID("c")
	sg2id := id.PropertySchemaGroupID("d")

	sf1 := NewSchemaField().ID(sf1id).Type(ValueTypeString).MustBuild()
	sg1 := NewSchemaGroup().ID(sg1id).Fields([]*SchemaField{sf1}).MustBuild()
	sf2 := NewSchemaField().ID(sf2id).Type(ValueTypeString).MustBuild()
	sg2 := NewSchemaGroup().ID(sg2id).Fields([]*SchemaField{sf2}).IsList(true).MustBuild()
	s := NewSchema().ID(sid).Groups([]*SchemaGroup{sg1, sg2}).MustBuild()

	p := New().NewID().Scene(sceneID).Schema(sid).MustBuild()

	// group
	assert.Nil(t, p.ItemBySchema(sg1id))
	assert.Equal(t, []Item{}, p.Items())

	i, gl := p.GetOrCreateItem(s, PointItemBySchema(sg1id))
	assert.Nil(t, gl)
	assert.NotNil(t, i)
	assert.Equal(t, sg1id, i.SchemaGroup())
	assert.Equal(t, i, ToGroup(p.ItemBySchema(sg1id)))
	assert.Equal(t, []Item{i}, p.Items())

	i2, gl := p.GetOrCreateItem(s, PointItemBySchema(sg1id))
	assert.Nil(t, gl)
	assert.NotNil(t, i2)
	assert.Equal(t, i, i2)
	assert.Equal(t, i2, ToGroup(p.ItemBySchema(sg1id)))
	assert.Equal(t, []Item{i2}, p.Items())

	// group list
	assert.Nil(t, p.ItemBySchema(sg2id))

	i3, gl := p.GetOrCreateItem(s, PointItemBySchema(sg2id))
	assert.Nil(t, gl)
	assert.Equal(t, sg2id, i3.SchemaGroup())
	assert.Equal(t, i3, ToGroupList(p.ItemBySchema(sg2id)))
	assert.Equal(t, []Item{i, i3}, p.Items())

	i4, gl := p.GetOrCreateItem(s, PointItemBySchema(sg2id))
	assert.Nil(t, gl)
	assert.Equal(t, i3, i4)
	assert.Equal(t, i4, ToGroupList(p.ItemBySchema(sg2id)))
	assert.Equal(t, []Item{i2, i4}, p.Items())
}

func TestProperty_GetOrCreateField(t *testing.T) {
	sceneID := id.NewSceneID()
	sid := id.MustPropertySchemaID("hoge~1.0.0/test")
	sf1id := id.PropertySchemaFieldID("a")
	sf2id := id.PropertySchemaFieldID("b")
	sg1id := id.PropertySchemaGroupID("c")
	sg2id := id.PropertySchemaGroupID("d")

	sf1 := NewSchemaField().ID(sf1id).Type(ValueTypeString).MustBuild()
	sg1 := NewSchemaGroup().ID(sg1id).Fields([]*SchemaField{sf1}).MustBuild()
	sf2 := NewSchemaField().ID(sf2id).Type(ValueTypeString).MustBuild()
	sg2 := NewSchemaGroup().ID(sg2id).Fields([]*SchemaField{sf2}).IsList(true).MustBuild()
	s := NewSchema().ID(sid).Groups([]*SchemaGroup{sg1, sg2}).MustBuild()

	p := New().NewID().Scene(sceneID).Schema(sid).MustBuild()

	// field and group will be created
	assert.Nil(t, p.ItemBySchema(sg1id))
	assert.Equal(t, []Item{}, p.Items())

	f, _, _, created := p.GetOrCreateField(s, PointFieldBySchemaGroup(sg1id, sf1id))
	assert.NotNil(t, f)
	assert.True(t, created)
	assert.Equal(t, sf1id, f.Field())
	i := ToGroup(p.ItemBySchema(sg1id))
	assert.Equal(t, sg1id, i.SchemaGroup())
	assert.Equal(t, []*Field{f}, i.Fields(nil))
	field, _, _ := p.Field(PointFieldBySchemaGroup(sg1id, sf1id))
	assert.Equal(t, f, field)

	f2, _, _, created := p.GetOrCreateField(s, PointFieldBySchemaGroup(sg1id, sf1id))
	assert.NotNil(t, f2)
	assert.False(t, created)
	assert.Equal(t, f, f2)
	i2 := ToGroup(p.ItemBySchema(sg1id))
	assert.Equal(t, i, i2)
	field, _, _ = p.Field(PointFieldBySchemaGroup(sg1id, sf1id))
	assert.Equal(t, f2, field)

	// field will not be created if field is incorrect
	f3, _, _, _ := p.GetOrCreateField(s, PointFieldBySchemaGroup(sg1id, sf2id))
	assert.Nil(t, f3)

	// field and group list will not be created
	assert.Nil(t, p.ItemBySchema(sg2id))
	f4, _, _, _ := p.GetOrCreateField(s, PointFieldBySchemaGroup(sg1id, sf2id))
	assert.Nil(t, f4)
	assert.Nil(t, p.ItemBySchema(sg2id))
	assert.Equal(t, []Item{i}, p.Items())
}

func TestAddListItem(t *testing.T) {
	sceneID := id.NewSceneID()
	sid := id.MustPropertySchemaID("hoge~1.0.0/test")
	sfid := id.PropertySchemaFieldID("a")
	sgid := id.PropertySchemaGroupID("b")
	sf := NewSchemaField().ID(sfid).Type(ValueTypeString).MustBuild()
	sg := NewSchemaGroup().ID(sgid).Fields([]*SchemaField{sf}).IsList(true).MustBuild()
	ps := NewSchema().ID(sid).Groups([]*SchemaGroup{sg}).MustBuild()
	p := New().NewID().Scene(sceneID).Schema(sid).MustBuild()

	item, _ := p.AddListItem(ps, PointItemBySchema(sgid), nil)
	assert.Equal(t, sgid, item.SchemaGroup())
	_, list := p.ListItem(PointItemBySchema(sgid))
	assert.Equal(t, sgid, list.SchemaGroup())
	assert.Equal(t, []*Group{item}, list.Groups())

	index := 0
	item2, _ := p.AddListItem(ps, PointItem(list.ID()), &index)
	assert.Equal(t, sgid, item2.SchemaGroup())
	assert.Equal(t, []*Group{item2, item}, list.Groups())
}

func TestMoveListItem(t *testing.T) {
	sceneID := id.NewSceneID()
	sid := id.MustPropertySchemaID("hoge~1.0.0/test")
	sgid := id.PropertySchemaGroupID("b")
	g1 := NewGroup().NewID().Schema(sgid).MustBuild()
	g2 := NewGroup().NewID().Schema(sgid).MustBuild()
	gl := NewGroupList().NewID().Schema(sgid).Groups([]*Group{g1, g2}).MustBuild()
	p := New().NewID().Scene(sceneID).Schema(sid).Items([]Item{gl}).MustBuild()

	assert.Equal(t, []*Group{g1, g2}, gl.Groups())
	i, _ := p.MoveListItem(PointItem(g1.ID()), 1)
	assert.Equal(t, g1, i)
	assert.Equal(t, []*Group{g2, g1}, gl.Groups())
}

func TestRemoveListItem(t *testing.T) {
	sceneID := id.NewSceneID()
	sid := id.MustPropertySchemaID("hoge~1.0.0/test")
	sgid := id.PropertySchemaGroupID("b")
	g1 := NewGroup().NewID().Schema(sgid).MustBuild()
	g2 := NewGroup().NewID().Schema(sgid).MustBuild()
	gl := NewGroupList().NewID().Schema(sgid).Groups([]*Group{g1, g2}).MustBuild()
	p := New().NewID().Scene(sceneID).Schema(sid).Items([]Item{gl}).MustBuild()

	assert.Equal(t, []*Group{g1, g2}, gl.Groups())
	ok := p.RemoveListItem(PointItem(g1.ID()))
	assert.True(t, ok)
	assert.Equal(t, []*Group{g2}, gl.Groups())
	assert.Equal(t, 1, len(p.Items()))

	ok = p.RemoveListItem(PointItem(g2.ID()))
	assert.True(t, ok)
	assert.Equal(t, []*Group{}, gl.Groups())
	assert.Equal(t, 0, len(p.Items()))
}

func TestProperty_Clone(t *testing.T) {
	tests := []struct {
		name   string
		target *Property
		n      bool
	}{
		{
			name:   "ok",
			target: testProperty1,
		},
		{
			name:   "nil",
			target: nil,
			n:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.target.Clone()
			if tt.n {
				assert.Nil(t, tt.target)
			} else {
				assert.Equal(t, tt.target, res)
				assert.NotSame(t, tt.target, res)
			}
		})
	}
}

func TestProperty_Fields(t *testing.T) {
	type args struct {
		p *Pointer
	}
	tests := []struct {
		name   string
		target *Property
		args   args
		want   []*Field
	}{
		{
			name:   "all",
			target: testProperty1,
			args:   args{p: nil},
			want:   []*Field{testField1, testField2},
		},
		{
			name:   "specified",
			target: testProperty1,
			args:   args{p: PointFieldOnly(testField1.Field())},
			want:   []*Field{testField1},
		},
		{
			name:   "not found",
			target: testProperty1,
			args:   args{p: PointFieldOnly("xxxxxx")},
			want:   []*Field{},
		},
		{
			name:   "empty",
			target: &Property{},
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

func TestProperty_RemoveFields(t *testing.T) {
	type args struct {
		p *Pointer
	}
	tests := []struct {
		name       string
		args       args
		target     *Property
		want       bool
		wantFields []*Field
	}{
		{
			name:       "nil pointer",
			target:     testProperty1.Clone(),
			args:       args{p: nil},
			want:       false,
			wantFields: []*Field{testField1, testField2},
		},
		{
			name:       "specified",
			target:     testProperty1.Clone(),
			args:       args{p: PointFieldOnly(testField1.Field())},
			want:       true,
			wantFields: []*Field{testField2},
		},
		{
			name:       "item only",
			target:     testProperty1.Clone(),
			args:       args{p: PointItem(testGroupList1.ID())},
			want:       false,
			wantFields: []*Field{testField1, testField2},
		},
		{
			name:       "not found",
			target:     testProperty1.Clone(),
			args:       args{p: PointFieldOnly("xxxxxx")},
			want:       false,
			wantFields: []*Field{testField1, testField2},
		},
		{
			name:       "empty",
			target:     &Property{},
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
			assert.Equal(t, tt.wantFields, tt.target.Fields(nil))
		})
	}
}

func TestProperty_MoveFields(t *testing.T) {
	sg1 := SchemaGroupID("aaa")
	sg2 := SchemaGroupID("bbb")
	sg3 := SchemaGroupID("ccc")
	sg4 := SchemaGroupID("ddd")
	sg5 := SchemaGroupID("eee")

	f1 := NewField().Field(FieldID("x")).Value(OptionalValueFrom(ValueTypeString.ValueFrom("aaa"))).Build()
	f2 := NewField().Field(FieldID("y")).Value(OptionalValueFrom(ValueTypeString.ValueFrom("bbb"))).Build()
	p := New().NewID().Scene(id.NewSceneID()).Schema(testSchema1.ID()).Items([]Item{
		NewGroup().NewID().Schema(sg1).Fields([]*Field{
			f1,
		}).MustBuild(),
		NewGroup().NewID().Schema(sg2).Fields([]*Field{
			// empty
		}).MustBuild(),
		NewGroupList().NewID().Schema(sg3).Groups([]*Group{
			NewGroup().NewID().Schema(sg3).Fields([]*Field{
				f2,
			}).MustBuild(),
		}).MustBuild(),
		NewGroupList().NewID().Schema(sg4).Groups([]*Group{
			NewGroup().NewID().Schema(sg4).Fields([]*Field{
				// empty
			}).MustBuild(),
		}).MustBuild(),
	}).MustBuild()

	type args struct {
		f    FieldID
		from *Pointer
		to   *Pointer
	}
	tests := []struct {
		name       string
		target     *Property
		args       args
		want       bool
		fromFields []*Field
		toFields   []*Field
	}{
		{
			name:   "group->group",
			target: p.Clone(),
			args: args{
				f:    f1.Field(),
				from: PointFieldBySchemaGroup(sg1, f1.Field()),
				to:   PointFieldBySchemaGroup(sg2, f1.Field()),
			},
			want:       true,
			fromFields: []*Field{},   // deleted
			toFields:   []*Field{f1}, // added
		},
		{
			name:   "group->new group",
			target: p.Clone(),
			args: args{
				f:    f1.Field(),
				from: PointFieldBySchemaGroup(sg1, f1.Field()),
				to:   PointFieldBySchemaGroup(sg5, f1.Field()),
			},
			want:       true,
			fromFields: []*Field{}, // deleted
			toFields:   []*Field{}, // not added
			// toFields:   []*Field{f1}, // added
		},
		{
			name:   "group->group failed",
			target: p.Clone(),
			args: args{
				f:    f2.Field(),
				from: PointFieldBySchemaGroup(sg1, f2.Field()),
				to:   PointFieldBySchemaGroup(sg2, f2.Field()),
			},
			want:       false,
			fromFields: []*Field{f1}, // not deleted
			toFields:   []*Field{},   // not added
		},
		{
			name:   "group list->group list",
			target: p.Clone(),
			args: args{
				f:    f2.Field(),
				from: PointFieldBySchemaGroup(sg3, f2.Field()),
				to:   PointFieldBySchemaGroup(sg4, f2.Field()),
			},
			want:       true,
			fromFields: []*Field{}, // deleted
			toFields:   []*Field{}, // not added
		},
		{
			name:   "group->group list",
			target: p.Clone(),
			args: args{
				f:    f1.Field(),
				from: PointFieldBySchemaGroup(sg1, f1.Field()),
				to:   PointFieldBySchemaGroup(sg4, f1.Field()),
			},
			want:       true,
			fromFields: []*Field{}, // deleted
			toFields:   []*Field{}, // not added
		},
		{
			name:   "group list->group",
			target: p.Clone(),
			args: args{
				f:    f2.Field(),
				from: PointFieldBySchemaGroup(sg3, f2.Field()),
				to:   PointFieldBySchemaGroup(sg2, f2.Field()),
			},
			want:       true,
			fromFields: []*Field{}, // deleted
			toFields:   []*Field{}, // not added
		},
		{
			name:   "empty",
			target: &Property{},
			args: args{
				f:    f1.Field(),
				from: PointFieldBySchemaGroup(sg1, f1.Field()),
				to:   PointFieldBySchemaGroup(sg2, f1.Field()),
			},
			want:       false,
			fromFields: nil,
			toFields:   nil,
		},
		{
			name: "nil",
			args: args{
				from: PointFieldBySchemaGroup(sg1, f1.Field()),
				to:   PointFieldBySchemaGroup(sg2, f1.Field()),
			},
			want:       false,
			fromFields: nil,
			toFields:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.MoveFields(tt.args.from, tt.args.to))
			assert.Equal(t, tt.fromFields, tt.target.Fields(PointItemBySchema(*tt.args.from.schemaGroup)), "from item")
			assert.Equal(t, tt.toFields, tt.target.Fields(PointItemBySchema(*tt.args.to.schemaGroup)), "to item")
		})
	}
}

func TestProperty_GroupAndFields(t *testing.T) {
	tests := []struct {
		name   string
		target *Property
		args   *Pointer
		want   []GroupAndField
	}{
		{
			name:   "all",
			target: testProperty1,
			args:   nil,
			want: []GroupAndField{
				{Group: testGroup1, Field: testField1},
				{ParentGroup: testGroupList1, Group: testGroup2, Field: testField2},
			},
		},
		{
			name:   "specified",
			target: testProperty1,
			args:   PointFieldBySchemaGroup(testGroupList1.SchemaGroup(), testField2.Field()),
			want: []GroupAndField{
				{ParentGroup: testGroupList1, Group: testGroup2, Field: testField2},
			},
		},
		{
			name:   "invalid field",
			target: testProperty1,
			args:   PointFieldBySchemaGroup(testGroupList1.SchemaGroup(), testField1.Field()),
			want:   []GroupAndField{},
		},
		{
			name:   "empty",
			target: &Property{},
			args:   nil,
			want:   nil,
		},
		{
			name:   "nil",
			target: &Property{},
			args:   nil,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.target.GroupAndFields(tt.args)
			assert.Equal(t, tt.want, res)
			for i, r := range res {
				assert.Same(t, tt.want[i].Field, r.Field)
				assert.Same(t, tt.want[i].Group, r.Group)
				if tt.want[i].ParentGroup != nil {
					assert.Same(t, tt.want[i].ParentGroup, r.ParentGroup)
				}
			}
		})
	}
}

func TestProperty_Cast(t *testing.T) {
	type args struct {
		ptr *Pointer
		t   ValueType
	}
	tests := []struct {
		name   string
		target *Property
		args   args
		want   bool
	}{
		{
			name:   "ok",
			target: testProperty1.Clone(),
			args: args{
				ptr: PointFieldOnly(testField2.Field()),
				t:   ValueTypeLatLngHeight,
			},
			want: true,
		},
		{
			name:   "failed to cast",
			target: testProperty1.Clone(),
			args: args{
				ptr: PointFieldOnly(testField2.Field()),
				t:   ValueTypeNumber,
			},
			want: true,
		},
		{
			name:   "not found",
			target: testProperty1.Clone(),
			args: args{
				ptr: PointFieldOnly("x"),
				t:   ValueTypeString,
			},
			want: false,
		},
		{
			name:   "empty",
			target: &Property{},
			args: args{
				ptr: PointFieldOnly(testField1.Field()),
				t:   ValueTypeString,
			},
			want: false,
		},
		{
			name:   "nil",
			target: nil,
			args: args{
				ptr: PointFieldOnly(testField1.Field()),
				t:   ValueTypeString,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Cast(tt.args.ptr, tt.args.t))
			for _, f := range tt.target.Fields(tt.args.ptr) {
				assert.Equal(t, tt.args.t, f.Type())
				assert.Nil(t, f.Links())
			}
		})
	}
}

func TestProperty_Prune(t *testing.T) {
	tests := []struct {
		name      string
		target    *Property
		wantRes   bool
		wantItems []Item
	}{
		{
			name: "ok",
			target: &Property{
				items: []Item{
					&Group{},
					&GroupList{},
				},
			},
			wantRes:   true,
			wantItems: []Item{},
		},
		{
			name: "not pruned",
			target: &Property{
				items: []Item{
					&Group{fields: []*Field{testField1.Clone()}},
				},
			},
			wantRes: false,
			wantItems: []Item{
				&Group{fields: []*Field{testField1.Clone()}},
			},
		},
		{
			name:      "empty",
			target:    &Property{},
			wantRes:   false,
			wantItems: nil,
		},
		{
			name:      "nil",
			target:    nil,
			wantRes:   false,
			wantItems: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantRes, tt.target.Prune())
			if tt.target != nil {
				assert.Equal(t, tt.wantItems, tt.target.items)
			}
		})
	}
}

func TestProperty_GuessSchema(t *testing.T) {
	tests := []struct {
		name   string
		target *Property
		want   *Schema
	}{
		{
			name: "ok",
			target: &Property{
				schema: MustSchemaID("x~1.1.1/a"),
				items: []Item{
					&Group{
						itemBase: itemBase{SchemaGroup: "aa"},
						fields: []*Field{
							{field: "a", v: NewOptionalValue(ValueTypeString, nil)},
						},
					},
					&GroupList{
						itemBase: itemBase{SchemaGroup: "bb"},
						groups: []*Group{
							{
								itemBase: itemBase{SchemaGroup: "aa"},
								fields: []*Field{
									{field: "b", v: NewOptionalValue(ValueTypeLatLng, nil)},
								},
							},
						},
					},
				},
			},
			want: &Schema{
				id: MustSchemaID("x~1.1.1/a"),
				groups: &SchemaGroupList{
					groups: []*SchemaGroup{
						{
							id: "aa",
							fields: []*SchemaField{
								{id: "a", propertyType: ValueTypeString},
							},
						},
						{
							id:   "bb",
							list: true,
							fields: []*SchemaField{
								{id: "b", propertyType: ValueTypeLatLng},
							},
						},
					},
				},
			},
		},
		{
			name:   "empty",
			target: &Property{},
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

func TestProperty_Field(t *testing.T) {
	tests := []struct {
		name   string
		target *Property
		args   *Pointer
		want   *Field
		want1  *GroupList
		want2  *Group
	}{
		{
			name:   "field by schema group",
			target: testProperty1,
			args:   PointFieldBySchemaGroup(testGroup1.SchemaGroup(), testField1.Field()),
			want:   testField1,
			want1:  nil,
			want2:  testGroup1,
		},
		{
			name:   "field by item",
			target: testProperty1,
			args:   PointFieldByItem(testGroup1.ID(), testField1.Field()),
			want:   testField1,
			want1:  nil,
			want2:  testGroup1,
		},
		{
			name:   "field only",
			target: testProperty1,
			args:   PointFieldOnly(testField1.Field()),
			want:   testField1,
			want1:  nil,
			want2:  testGroup1,
		},
		{
			name:   "field in list",
			target: testProperty1,
			args:   PointFieldOnly(testField2.Field()),
			want:   nil,
			want1:  nil,
			want2:  nil,
		},
		{
			name:   "not found",
			target: testProperty1,
			args:   PointFieldOnly("_"),
			want:   nil,
			want1:  nil,
			want2:  nil,
		},
		{
			name:   "empty",
			target: &Property{},
			args:   PointFieldOnly("_"),
			want:   nil,
			want1:  nil,
			want2:  nil,
		},
		{
			name:   "nil pointer",
			target: testProperty1,
			args:   nil,
			want:   nil,
			want1:  nil,
			want2:  nil,
		},
		{
			name:   "nil",
			target: nil,
			args:   PointFieldOnly("_"),
			want:   nil,
			want1:  nil,
			want2:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.target.Field(tt.args)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
		})
	}
}

func TestProperty_Item(t *testing.T) {
	tests := []struct {
		name   string
		target *Property
		args   *Pointer
		want   Item
	}{
		{
			name:   "group",
			target: testProperty1,
			args:   PointItem(testGroup1.ID()),
			want:   testGroup1,
		},
		{
			name:   "group list",
			target: testProperty1,
			args:   PointItem(testGroupList1.ID()),
			want:   testGroupList1,
		},
		{
			name:   "group by schema group",
			target: testProperty1,
			args:   PointItemBySchema(testGroup1.SchemaGroup()),
			want:   testGroup1,
		},
		{
			name:   "group list by schema group",
			target: testProperty1,
			args:   PointItemBySchema(testGroupList1.SchemaGroup()),
			want:   testGroupList1,
		},
		{
			name:   "not found",
			target: testProperty1,
			args:   PointItemBySchema("_"),
			want:   nil,
		},
		{
			name:   "field only",
			target: testProperty1,
			args:   PointFieldOnly(testField1.Field()),
			want:   nil,
		},
		{
			name:   "nil pointer",
			target: testProperty1,
			args:   nil,
			want:   nil,
		},
		{
			name:   "empty",
			target: &Property{},
			args:   PointFieldOnly(testField1.Field()),
			want:   nil,
		},
		{
			name:   "nil",
			target: nil,
			args:   PointFieldOnly(testField1.Field()),
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Item(tt.args))
		})
	}
}
