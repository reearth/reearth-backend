package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

var (
	sf = NewSchemaField().ID("aa").Type(ValueTypeString).MustBuild()
	sg = NewSchemaGroup().ID("aa").Fields([]*SchemaField{sf}).MustBuild()
	p  = New().NewID().Scene(id.NewSceneID()).Schema(id.MustPropertySchemaID("xx~1.0.0/aa")).Items([]Item{InitItemFrom(sg)}).MustBuild()
	p2 = New().NewID().Scene(id.NewSceneID()).Schema(id.MustPropertySchemaID("xx~1.0.0/aa")).Items([]Item{InitItemFrom(sg)}).MustBuild()
)

func TestList_IDs(t *testing.T) {
	p1 := NewID()
	p2 := NewID()

	tests := []struct {
		name   string
		target List
		want   []ID
	}{
		{
			name:   "ok",
			target: List{&Property{id: p1}, &Property{id: p2}, &Property{id: p1}},
			want:   []ID{p1, p2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.IDs())
		})
	}
}

func TestList_Schemas(t *testing.T) {
	ps1 := MustSchemaID("x~1.0.0/a")
	ps2 := MustSchemaID("x~1.0.0/b")

	tests := []struct {
		name   string
		target List
		want   []SchemaID
	}{
		{
			name:   "ok",
			target: List{&Property{schema: ps1}, &Property{schema: ps2}, &Property{schema: ps1}},
			want:   []SchemaID{ps1, ps2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Schemas())
		})
	}
}

func TestList_Map(t *testing.T) {
	p1 := NewID()
	p2 := NewID()

	tests := []struct {
		name   string
		target List
		want   Map
	}{
		{
			name:   "ok",
			target: List{&Property{id: p1}, &Property{id: p2}, &Property{id: p1}},
			want: Map{
				p1: &Property{id: p1},
				p2: &Property{id: p2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Map())
		})
	}
}

func TestMap_Add(t *testing.T) {
	testCases := []struct {
		Name        string
		Input       *Property
		M, Expected Map
	}{
		{
			Name: "nil map",
		},
		{
			Name:     "add property list",
			Input:    p,
			M:        Map{},
			Expected: Map{p.ID(): p},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.M.Add(tc.Input)
			assert.Equal(tt, tc.Expected, tc.M)
			assert.Equal(tt, tc.Expected.List(), tc.M.List())
		})
	}
}

func TestMapFrom(t *testing.T) {
	m := MapFrom(p)
	e := Map{p.ID(): p}
	assert.Equal(t, e, m)
}

func TestMap_Clone(t *testing.T) {
	testCases := []struct {
		Name        string
		M, Expected Map
	}{
		{
			Name:     "nil map",
			Expected: Map{},
		},
		{
			Name:     "add property list",
			M:        Map{p.ID(): p},
			Expected: Map{p.ID(): p},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.M.Clone()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestMap_Merge(t *testing.T) {
	testCases := []struct {
		Name             string
		M1, M2, Expected Map
	}{
		{
			Name:     "nil map",
			Expected: Map{},
		},
		{
			Name:     "add property list",
			M1:       Map{p.ID(): p},
			M2:       Map{p2.ID(): p2},
			Expected: Map{p.ID(): p, p2.ID(): p2},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.M1.Merge(tc.M2)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
