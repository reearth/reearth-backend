package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

var (
	testSchemaGroup1 = NewSchemaGroup().ID("aa").Fields([]*SchemaField{testSchemaField1, testSchemaField2}).MustBuild()
	testSchemaGroup2 = NewSchemaGroup().ID("bb").Fields([]*SchemaField{testSchemaField3}).IsList(true).MustBuild()
)

func TestSchemaGroup(t *testing.T) {
	scid := id.PropertySchemaGroupID("aa")
	sid := id.MustPropertySchemaID("xx~1.0.0/aa")
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()

	testCases := []struct {
		Name     string
		G        *SchemaGroup
		Expected struct {
			GID           id.PropertySchemaGroupID
			SID           id.PropertySchemaID
			Fields        []*SchemaField
			Title         i18n.String
			IsAvailableIf *Condition
			IsList        bool
		}
	}{
		{
			Name: "nil schema group",
		},
		{
			Name: "success",
			G:    NewSchemaGroup().ID(scid).Fields([]*SchemaField{sf}).MustBuild(),
			Expected: struct {
				GID           id.PropertySchemaGroupID
				SID           id.PropertySchemaID
				Fields        []*SchemaField
				Title         i18n.String
				IsAvailableIf *Condition
				IsList        bool
			}{
				GID:    scid,
				SID:    sid,
				Fields: []*SchemaField{sf},
				Title:  nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			assert.Equal(tt, tc.Expected.GID, tc.G.ID())
			assert.Equal(tt, tc.Expected.Fields, tc.G.Fields())
			assert.Equal(tt, tc.Expected.IsList, tc.G.IsList())
			assert.Equal(tt, tc.Expected.IsAvailableIf, tc.G.IsAvailableIf())
			assert.Equal(tt, tc.Expected.Title, tc.G.Title())
		})
	}
}

func TestSchemaGroup_Field(t *testing.T) {
	scid := id.PropertySchemaGroupID("aa")
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()

	testCases := []struct {
		Name     string
		G        *SchemaGroup
		PTR      *Pointer
		Input    id.PropertySchemaFieldID
		Expected *SchemaField
	}{
		{
			Name: "nil schema group",
		},
		{
			Name:     "found",
			G:        NewSchemaGroup().ID(scid).Fields([]*SchemaField{sf}).MustBuild(),
			PTR:      NewPointer(nil, nil, sf.ID().Ref()),
			Input:    sf.ID(),
			Expected: sf,
		},
		{
			Name:  "not found",
			G:     NewSchemaGroup().ID(scid).Fields([]*SchemaField{sf}).MustBuild(),
			PTR:   NewPointer(nil, nil, id.PropertySchemaFieldID("zz").Ref()),
			Input: id.PropertySchemaFieldID("zz"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected, tc.G.Field(tc.Input))
			assert.Equal(tt, tc.Expected, tc.G.FieldByPointer(tc.PTR))
			assert.Equal(tt, tc.Expected != nil, tc.G.HasField(tc.Input))
		})
	}
}

func TestSchemaGroup_SetTitle(t *testing.T) {
	sg := NewSchemaGroup().ID(id.PropertySchemaGroupID("aa")).Fields([]*SchemaField{sf}).MustBuild()
	sg.SetTitle(i18n.StringFrom("ttt"))
	assert.Equal(t, i18n.StringFrom("ttt"), sg.Title())
}
