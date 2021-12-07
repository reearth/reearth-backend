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

func TestSchema_Field(t *testing.T) {
	tests := []struct {
		name   string
		target *Schema
		input  id.PropertySchemaFieldID
		want   *SchemaField
	}{
		{
			name: "nil schema",
		},
		{
			name:   "found",
			target: testSchema1,
			input:  testSchemaField1.ID(),
			want:   testSchemaField1,
		},
		{
			name:   "not found",
			target: testSchema1,
			input:  id.PropertySchemaFieldID("zz"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.target.Field(tt.input))
		})
	}
}

func TestSchema_FieldByPointer(t *testing.T) {
	tests := []struct {
		name   string
		target *Schema
		input  *Pointer
		want   *SchemaField
	}{
		{
			name: "nil schema",
		},
		{
			name:   "found",
			target: testSchema1,
			input:  NewPointer(nil, nil, testSchemaField1.ID().Ref()),
			want:   testSchemaField1,
		},
		{
			name:   "not found",
			target: testSchema1,
			input:  NewPointer(nil, nil, id.PropertySchemaFieldID("zz").Ref()),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.target.FieldByPointer(tt.input))
		})
	}
}

func TestSchema_Group(t *testing.T) {
	tests := []struct {
		name   string
		target *Schema
		input  SchemaGroupID
		want   *SchemaGroup
	}{
		{
			name:   "nil schema",
			target: nil,
			input:  testSchemaGroup1.ID(),
			want:   nil,
		},
		{
			name:   "found",
			target: testSchema1,
			input:  testSchemaGroup1.ID(),
			want:   testSchemaGroup1,
		},
		{
			name:   "not found",
			target: testSchema1,
			input:  SchemaGroupID("zz"),
			want:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.target.Group(tt.input))
		})
	}
}

func TestSchema_GroupByField(t *testing.T) {
	tests := []struct {
		name   string
		target *Schema
		input  FieldID
		want   *SchemaGroup
	}{
		{
			name:   "nil schema",
			target: nil,
			input:  testSchemaField1.ID(),
			want:   nil,
		},
		{
			name:   "found",
			target: testSchema1,
			input:  testSchemaField1.ID(),
			want:   testSchemaGroup1,
		},
		{
			name:   "not found",
			target: testSchema1,
			input:  FieldID("zz"),
			want:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.target.GroupByField(tt.input))
		})
	}
}

func TestSchema_GroupAndFields(t *testing.T) {
	tests := []struct {
		name   string
		target *Schema
		want   []SchemaGroupAndField
	}{
		{
			name:   "ok",
			target: testSchema1,
			want: []SchemaGroupAndField{
				{Group: testSchemaGroup1, Field: testSchemaField1},
				{Group: testSchemaGroup1, Field: testSchemaField2},
				{Group: testSchemaGroup2, Field: testSchemaField3},
			},
		},
		{
			name:   "empty",
			target: &Schema{},
			want:   []SchemaGroupAndField{},
		},
		{
			name:   "nil",
			target: nil,
			want:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := tt.target.GroupAndFields()
			assert.Equal(t, tt.want, res)
			if len(tt.want) > 0 {
				for i, gf := range res {
					assert.Same(t, tt.want[i].Group, gf.Group)
					assert.Same(t, tt.want[i].Field, gf.Field)
				}
			}
		})
	}
}

func TestSchema_GroupAndField(t *testing.T) {
	type args struct {
		f FieldID
	}
	tests := []struct {
		name   string
		args   args
		target *Schema
		want   *SchemaGroupAndField
	}{
		{
			name:   "ok1",
			target: testSchema1,
			args:   args{f: testSchemaField1.ID()},
			want:   &SchemaGroupAndField{Group: testSchemaGroup1, Field: testSchemaField1},
		},
		{
			name:   "ok2",
			target: testSchema1,
			args:   args{f: testSchemaField2.ID()},
			want:   &SchemaGroupAndField{Group: testSchemaGroup1, Field: testSchemaField2},
		},
		{
			name:   "ok3",
			target: testSchema1,
			args:   args{f: testSchemaField3.ID()},
			want:   &SchemaGroupAndField{Group: testSchemaGroup2, Field: testSchemaField3},
		},
		{
			name:   "not found",
			target: testSchema1,
			args:   args{f: "ddd"},
			want:   nil,
		},
		{
			name:   "empty",
			target: &Schema{},
			want:   nil,
		},
		{
			name:   "nil",
			target: nil,
			want:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := tt.target.GroupAndField(tt.args.f)
			assert.Equal(t, tt.want, res)
			if tt.want != nil {
				assert.Same(t, tt.want.Group, res.Group)
				assert.Same(t, tt.want.Field, res.Field)
			}
		})
	}
}

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
			Name:     "invalid: URL",
			S:        NewSchema().ID(sid).Groups([]*SchemaGroup{sg}).MustBuild(),
			LF:       LinkableFields{URL: NewPointer(nil, nil, id.PropertySchemaFieldID("xx").Ref())},
			Expected: false,
		},
		{
			Name:     "invalid: Lng",
			S:        NewSchema().ID(sid).Groups([]*SchemaGroup{sg}).MustBuild(),
			LF:       LinkableFields{LatLng: NewPointer(nil, nil, id.PropertySchemaFieldID("xx").Ref())},
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

func TestSchemaGroupAndField_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		target SchemaGroupAndField
		want   bool
	}{
		{
			name: "present",
			target: SchemaGroupAndField{
				Group: testSchemaGroup1,
				Field: testSchemaField1,
			},
			want: false,
		},
		{
			name:   "empty",
			target: SchemaGroupAndField{},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gf := SchemaGroupAndField{
				Group: tt.target.Group,
				Field: tt.target.Field,
			}
			assert.Equal(t, tt.want, gf.IsEmpty())
		})
	}
}
