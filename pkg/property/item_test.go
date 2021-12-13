package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestInitItemFrom(t *testing.T) {
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	sg := NewSchemaGroup().ID("aa").Fields([]*SchemaField{sf}).MustBuild()
	sgl := NewSchemaGroup().ID("aa").IsList(true).Fields([]*SchemaField{sf}).MustBuild()
	iid := id.NewPropertyItemID()
	propertySchemaField1ID := id.PropertySchemaGroupID("aa")

	testCases := []struct {
		Name     string
		SG       *SchemaGroup
		Expected Item
	}{
		{
			Name: "nil psg",
		},
		{
			Name:     "init item from group",
			SG:       sg,
			Expected: NewGroup().ID(iid).Schema(propertySchemaField1ID).MustBuild(),
		},
		{
			Name:     "init item from group list",
			SG:       sgl,
			Expected: NewGroupList().ID(iid).Schema(propertySchemaField1ID).MustBuild(),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := InitItemFrom(tc.SG)
			if res != nil {
				assert.Equal(tt, tc.Expected.SchemaGroup(), res.SchemaGroup())
			} else {
				assert.Nil(tt, tc.Expected)
			}
		})
	}
}

func TestToGroup(t *testing.T) {
	iid := id.NewPropertyItemID()
	propertySchemaID := id.MustPropertySchemaID("xxx~1.1.1/aa")
	propertySchemaField1ID := id.PropertySchemaFieldID("a")
	propertySchemaGroup1ID := id.PropertySchemaGroupID("A")
	il := []Item{
		NewGroup().ID(iid).Schema(propertySchemaGroup1ID).
			Fields([]*Field{
				NewField().
					Field(propertySchemaField1ID).
					Value(OptionalValueFrom(ValueTypeString.ValueFrom("xxx"))).
					Build(),
			}).MustBuild(),
	}
	p := New().NewID().Scene(id.NewSceneID()).Items(il).Schema(propertySchemaID).MustBuild()
	g := ToGroup(p.ItemBySchema(propertySchemaGroup1ID))
	assert.Equal(t, propertySchemaGroup1ID, g.SchemaGroup())
	assert.Equal(t, iid, g.ID())
}

func TestToGroupList(t *testing.T) {
	iid := id.NewPropertyItemID()
	propertySchemaID := id.MustPropertySchemaID("xxx~1.1.1/aa")
	propertySchemaGroup1ID := id.PropertySchemaGroupID("A")
	il := []Item{
		NewGroupList().ID(iid).Schema(propertySchemaGroup1ID).MustBuild(),
	}
	p := New().NewID().Scene(id.NewSceneID()).Items(il).Schema(propertySchemaID).MustBuild()
	g := ToGroupList(p.ItemBySchema(propertySchemaGroup1ID))
	assert.Equal(t, propertySchemaGroup1ID, g.SchemaGroup())
	assert.Equal(t, iid, g.ID())
}

func TestGroupAndField_SchemaFieldPointer(t *testing.T) {
	tests := []struct {
		name   string
		target GroupAndField
		want   SchemaFieldPointer
	}{
		{
			name: "group",
			target: GroupAndField{
				ParentGroup: nil,
				Group:       testGroup1,
				Field:       testField1,
			},
			want: SchemaFieldPointer{
				SchemaGroup: testGroup1.SchemaGroup(),
				Field:       testField1.Field(),
			},
		},
		{
			name: "group list",
			target: GroupAndField{
				ParentGroup: testGroupList1,
				Group:       testGroup2,
				Field:       testField2,
			},
			want: SchemaFieldPointer{
				SchemaGroup: testGroup2.SchemaGroup(),
				Field:       testField2.Field(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.SchemaFieldPointer())
		})
	}
}
