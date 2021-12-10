package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_New(t *testing.T) {
	b := New()
	assert.NotNil(t, b)
}

func TestBuilder_ID(t *testing.T) {
	pid := id.NewPropertyID()
	p := New().ID(pid).Scene(id.NewSceneID()).Schema(id.MustPropertySchemaID("xxx~1.1.1/aa")).MustBuild()
	assert.Equal(t, pid, p.ID())
}

func TestBuilder_NewID(t *testing.T) {
	p := New().NewID().Scene(id.NewSceneID()).Schema(id.MustPropertySchemaID("xxx~1.1.1/aa")).MustBuild()
	assert.False(t, p.ID().IsNil())
}

func TestBuilder_Schema(t *testing.T) {
	p := New().NewID().Scene(id.NewSceneID()).Schema(id.MustPropertySchemaID("xxx~1.1.1/aa")).MustBuild()
	assert.Equal(t, id.MustPropertySchemaID("xxx~1.1.1/aa"), p.Schema())
}

func TestBuilder_Scene(t *testing.T) {
	sid := id.NewSceneID()
	p := New().NewID().Scene(sid).Schema(id.MustPropertySchemaID("xxx~1.1.1/aa")).MustBuild()
	assert.Equal(t, sid, p.Scene())
}

func TestBuilder_Items(t *testing.T) {
	iid := id.NewPropertyItemID()
	propertySchemaField1ID := id.PropertySchemaFieldID("a")
	propertySchemaGroup1ID := id.PropertySchemaGroupID("A")

	testCases := []struct {
		Name            string
		Input, Expected []Item
	}{
		{
			Name:     "has nil item",
			Input:    []Item{nil},
			Expected: []Item{},
		},
		{
			Name: "has duplicated item",
			Input: []Item{
				NewGroup().ID(iid).Schema(propertySchemaGroup1ID).
					Fields([]*Field{
						NewField().
							Field(propertySchemaField1ID).
							Value(OptionalValueFrom(ValueTypeString.ValueFrom("xxx"))).
							Build(),
					}).MustBuild(),
				NewGroup().ID(iid).Schema(propertySchemaGroup1ID).
					Fields([]*Field{
						NewField().
							Field(propertySchemaField1ID).
							Value(OptionalValueFrom(ValueTypeString.ValueFrom("xxx"))).
							Build(),
					}).MustBuild(),
			},
			Expected: []Item{NewGroup().ID(iid).Schema(propertySchemaGroup1ID).
				Fields([]*Field{
					NewField().
						Field(propertySchemaField1ID).
						Value(OptionalValueFrom(ValueTypeString.ValueFrom("xxx"))).
						Build(),
				}).MustBuild()},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := New().NewID().
				Scene(id.NewSceneID()).
				Schema(id.MustPropertySchemaID("xxx~1.1.1/aa")).
				Items(tc.Input).
				MustBuild()
			assert.Equal(tt, tc.Expected, res.Items())
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	sid := id.NewSceneID()
	psid := MustSchemaID("test~1.0.0/a")
	pid := NewID()
	iid := NewItemID()
	propertySchemaField1ID := FieldID("a")
	propertySchemaGroup1ID := SchemaGroupID("A")

	tests := []struct {
		Name     string
		ID       ID
		Scene    id.SceneID
		Schema   SchemaID
		Items    []Item
		Err      error
		Expected *Property
	}{
		{
			Name:   "success",
			ID:     pid,
			Scene:  sid,
			Schema: psid,
			Items: []Item{
				NewGroup().ID(iid).Schema(propertySchemaGroup1ID).
					Fields([]*Field{
						NewField().
							Field(propertySchemaField1ID).
							Value(OptionalValueFrom(ValueTypeString.ValueFrom("xxx"))).
							Build(),
					}).MustBuild()},
			Expected: &Property{
				id:     pid,
				scene:  sid,
				schema: psid,
				items: []Item{
					NewGroup().ID(iid).Schema(propertySchemaGroup1ID).
						Fields([]*Field{
							NewField().
								Field(propertySchemaField1ID).
								Value(OptionalValueFrom(ValueTypeString.ValueFrom("xxx"))).
								Build(),
						}).MustBuild()},
			},
		},
		{
			Name:   "fail invalid id",
			ID:     id.PropertyID{},
			Scene:  sid,
			Schema: psid,
			Items:  nil,
			Err:    id.ErrInvalidID,
		},
		{
			Name:   "fail invalid scene",
			ID:     pid,
			Scene:  id.SceneID{},
			Schema: psid,
			Items:  nil,
			Err:    ErrInvalidSceneID,
		},
		{
			Name:   "fail invalid schema",
			ID:     pid,
			Scene:  sid,
			Schema: SchemaID{},
			Items:  nil,
			Err:    ErrInvalidPropertySchemaID,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			res, err := New().ID(tc.ID).Schema(tc.Schema).Items(tc.Items).Scene(tc.Scene).Build()
			if tc.Err == nil {
				assert.Equal(t, tc.Expected, res)
			} else {
				assert.Equal(t, tc.Err, err)
			}
		})
	}
}
