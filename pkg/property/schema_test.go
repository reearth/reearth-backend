package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

var (
	testSchemaID = id.MustPropertySchemaID("xx~1.0.0/aa")
	testSchema1  = NewSchema().ID(testSchemaID).Groups([]*SchemaGroup{testSchemaGroup1, testSchemaGroup2}).MustBuild()
)

func TestLinkableField_Validate(t *testing.T) {
	sid := id.MustPropertySchemaID("xx~1.0.0/aa")
	sf := NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	sg := NewSchemaGroup().ID("aaa").Schema(sid).Fields([]*SchemaField{sf}).MustBuild()

	tests := []struct {
		Name     string
		S        *Schema
		LF       LinkableFields
		Expected bool
	}{
		{
			Name:     "nil schema",
			S:        nil,
			LF:       LinkableFields{},
			Expected: false,
		},
		{
			Name: "invalid: URL",
			S:    NewSchema().ID(sid).Groups([]*SchemaGroup{sg}).MustBuild(),
			LF: LinkableFields{
				URL: &SchemaFieldPointer{
					Field: id.PropertySchemaFieldID("xx"),
				},
			},
			Expected: false,
		},
		{
			Name: "invalid: Lng",
			S:    NewSchema().ID(sid).Groups([]*SchemaGroup{sg}).MustBuild(),
			LF: LinkableFields{
				LatLng: &SchemaFieldPointer{
					Field: id.PropertySchemaFieldID("xx"),
				},
			},
			Expected: false,
		},
		{
			Name:     "empty",
			S:        NewSchema().ID(sid).Groups([]*SchemaGroup{sg}).MustBuild(),
			LF:       LinkableFields{},
			Expected: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res := tt.LF.Validate(tt.S)
			assert.Equal(t, tt.Expected, res)
		})
	}
}
