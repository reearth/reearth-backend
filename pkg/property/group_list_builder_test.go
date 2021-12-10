package property

import (
	"errors"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestGroupListBuilder_Build(t *testing.T) {
	pid := id.NewPropertyItemID()
	scid := id.MustPropertySchemaID("xx~1.0.0/aa")
	groups := []*Group{NewGroup().ID(pid).MustBuild()}
	testCases := []struct {
		Name        string
		Id          id.PropertyItemID
		Schema      id.PropertySchemaID
		SchemaGroup id.PropertySchemaGroupID
		Groups      []*Group
		Expected    struct {
			Id          id.PropertyItemID
			SchemaGroup id.PropertySchemaGroupID
			Groups      []*Group
		}
		Err error
	}{
		{
			Name:        "success",
			Id:          pid,
			Schema:      scid,
			SchemaGroup: "aa",
			Groups:      groups,
			Expected: struct {
				Id          id.PropertyItemID
				SchemaGroup id.PropertySchemaGroupID
				Groups      []*Group
			}{
				Id:          pid,
				SchemaGroup: "aa",
				Groups:      groups,
			},
		},
		{
			Name: "fail invalid id",
			Err:  id.ErrInvalidID,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res, err := NewGroupList().ID(tc.Id).Schema(tc.SchemaGroup).Groups(tc.Groups).Build()
			if err == nil {
				assert.Equal(tt, tc.Expected.Id, res.ID())
				assert.Equal(tt, tc.Expected.SchemaGroup, res.SchemaGroup())
				assert.Equal(tt, tc.Expected.Groups, res.Groups())
			} else {
				assert.True(tt, errors.As(tc.Err, &err))
			}
		})
	}
}

func TestGroupListBuilder_NewID(t *testing.T) {
	b := NewGroupList().NewID().MustBuild()
	assert.NotNil(t, b.ID())
}

func TestGroupListBuilder_MustBuild(t *testing.T) {
	pid := id.NewPropertyItemID()
	scid := id.MustPropertySchemaID("xx~1.0.0/aa")
	groups := []*Group{NewGroup().ID(pid).MustBuild()}
	tests := []struct {
		Name        string
		Fails       bool
		Id          id.PropertyItemID
		Schema      id.PropertySchemaID
		SchemaGroup id.PropertySchemaGroupID
		Groups      []*Group
		Expected    struct {
			Id          id.PropertyItemID
			Schema      id.PropertySchemaID
			SchemaGroup id.PropertySchemaGroupID
			Groups      []*Group
		}
	}{
		{
			Name:        "success",
			Id:          pid,
			Schema:      scid,
			SchemaGroup: "aa",
			Groups:      groups,
			Expected: struct {
				Id          id.PropertyItemID
				Schema      id.PropertySchemaID
				SchemaGroup id.PropertySchemaGroupID
				Groups      []*Group
			}{
				Id:          pid,
				Schema:      scid,
				SchemaGroup: "aa",
				Groups:      groups,
			},
		},
		{
			Name:  "fail invalid id",
			Fails: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			if tc.Fails {
				assert.Panics(t, func() {
					_ = NewGroupList().ID(tc.Id).Schema(tc.SchemaGroup).Groups(tc.Groups).MustBuild()
				})
			} else {
				res := NewGroupList().ID(tc.Id).Schema(tc.SchemaGroup).Groups(tc.Groups).MustBuild()
				assert.Equal(t, tc.Expected.Id, res.ID())
				assert.Equal(t, tc.Expected.SchemaGroup, res.SchemaGroup())
				assert.Equal(t, tc.Expected.Groups, res.Groups())
			}

		})
	}
}

func TestInitGroupListFrom(t *testing.T) {
	testCases := []struct {
		Name        string
		SchemaGroup *SchemaGroup
		ExpectedSG  id.PropertySchemaGroupID
	}{
		{
			Name: "nil schema group",
		},
		{
			Name:        "success",
			SchemaGroup: NewSchemaGroup().ID("aa").MustBuild(),
			ExpectedSG:  "aa",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := InitGroupFrom(tc.SchemaGroup)
			assert.Equal(tt, tc.ExpectedSG, res.SchemaGroup())
		})
	}
}
