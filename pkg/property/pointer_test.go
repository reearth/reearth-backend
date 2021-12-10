package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	iid := id.NewPropertyItemID()
	sgid := id.PropertySchemaGroupID("foo")
	fid := id.PropertySchemaFieldID("hoge")

	var p *Pointer
	var ok bool

	p = PointItem(iid)
	i, ok := p.Item()
	assert.True(t, ok)
	assert.Equal(t, iid, i)
	_, ok = p.ItemBySchemaGroup()
	assert.False(t, ok)
	_, _, ok = p.FieldByItem()
	assert.False(t, ok)
	_, _, ok = p.FieldBySchemaGroup()
	assert.False(t, ok)

	p = PointItemBySchema(sgid)
	_, ok = p.Item()
	assert.False(t, ok)
	sg, ok := p.ItemBySchemaGroup()
	assert.True(t, ok)
	assert.Equal(t, sgid, sg)
	_, _, ok = p.FieldByItem()
	assert.False(t, ok)
	_, _, ok = p.FieldBySchemaGroup()
	assert.False(t, ok)

	p = PointFieldByItem(iid, fid)
	i, ok = p.Item()
	assert.True(t, ok)
	assert.Equal(t, iid, i)
	_, ok = p.ItemBySchemaGroup()
	assert.False(t, ok)
	i, f, ok := p.FieldByItem()
	assert.True(t, ok)
	assert.Equal(t, iid, i)
	assert.Equal(t, fid, f)
	_, _, ok = p.FieldBySchemaGroup()
	assert.False(t, ok)

	p = PointFieldBySchemaGroup(sgid, fid)
	_, ok = p.Item()
	assert.False(t, ok)
	sg, ok = p.ItemBySchemaGroup()
	assert.True(t, ok)
	assert.Equal(t, sgid, sg)
	_, _, ok = p.FieldByItem()
	assert.False(t, ok)
	sg, f, ok = p.FieldBySchemaGroup()
	assert.True(t, ok)
	assert.Equal(t, sgid, sg)
	assert.Equal(t, fid, f)

	p = PointField(&sgid, &iid, fid)
	i, ok = p.Item()
	assert.True(t, ok)
	assert.Equal(t, iid, i)
	sg, ok = p.ItemBySchemaGroup()
	assert.True(t, ok)
	assert.Equal(t, sgid, sg)
	_, _, ok = p.FieldByItem()
	assert.False(t, ok)
	_, _, ok = p.FieldBySchemaGroup()
	assert.False(t, ok)
}

func TestPointer_Test(t *testing.T) {
	iid := NewItemID()

	type args struct {
		sg   SchemaGroupID
		i    ItemID
		f    FieldID
		want bool
	}
	tests := []struct {
		name   string
		target *Pointer
		args   []args
	}{
		{
			name:   "schema group only",
			target: &Pointer{schemaGroup: SchemaGroupID("xx").Ref()},
			args: []args{
				{sg: SchemaGroupID("xx"), i: iid, f: FieldID("a"), want: true},
				{sg: SchemaGroupID("xx"), i: iid, f: FieldID("b"), want: true},
				{sg: SchemaGroupID("yy"), i: iid, f: FieldID("a"), want: false},
			},
		},
		{
			name:   "item only",
			target: &Pointer{item: iid.Ref()},
			args: []args{
				{sg: SchemaGroupID("xx"), i: iid, f: FieldID("a"), want: true},
				{sg: SchemaGroupID("yy"), i: iid, f: FieldID("a"), want: true},
				{sg: SchemaGroupID("xx"), i: iid, f: FieldID("b"), want: true},
				{sg: SchemaGroupID("xx"), i: NewItemID(), f: FieldID("a"), want: false},
			},
		},
		{
			name:   "schema group and item",
			target: &Pointer{schemaGroup: SchemaGroupID("xx").Ref(), item: iid.Ref()},
			args: []args{
				{sg: SchemaGroupID("xx"), i: iid, f: FieldID("a"), want: true},
				{sg: SchemaGroupID("xx"), i: iid, f: FieldID("b"), want: true},
				{sg: SchemaGroupID("xx"), i: NewItemID(), f: FieldID("a"), want: false},
				{sg: SchemaGroupID("yy"), i: iid, f: FieldID("a"), want: false},
				{sg: SchemaGroupID("yy"), i: NewItemID(), f: FieldID("a"), want: false},
			},
		},
		{
			name:   "all",
			target: &Pointer{schemaGroup: SchemaGroupID("xx").Ref(), item: iid.Ref(), field: FieldID("a").Ref()},
			args: []args{
				{sg: SchemaGroupID("xx"), i: iid, f: FieldID("a"), want: true},
				{sg: SchemaGroupID("yy"), i: iid, f: FieldID("a"), want: false},
				{sg: SchemaGroupID("xx"), i: NewItemID(), f: FieldID("a"), want: false},
				{sg: SchemaGroupID("xx"), i: iid, f: FieldID("b"), want: false},
			},
		},
		{
			name:   "empty",
			target: &Pointer{},
			args: []args{
				{sg: SchemaGroupID("xx"), i: NewItemID(), f: FieldID("a"), want: true},
				{sg: SchemaGroupID("yy"), i: NewItemID(), f: FieldID("b"), want: true},
				{sg: SchemaGroupID("zz"), i: NewItemID(), f: FieldID("c"), want: true},
			},
		},
		{
			name:   "nil",
			target: nil,
			args: []args{
				{sg: SchemaGroupID("xx"), i: NewItemID(), f: FieldID("a"), want: false},
				{sg: SchemaGroupID("yy"), i: NewItemID(), f: FieldID("b"), want: false},
				{sg: SchemaGroupID("zz"), i: NewItemID(), f: FieldID("c"), want: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, a := range tt.args {
				assert.Equal(t, a.want, tt.target.Test(a.sg, a.i, a.f), "test %d", i)
			}
		})
	}
}

func TestPointer_TestItem(t *testing.T) {
	iid := NewItemID()

	type args struct {
		sg SchemaGroupID
		i  ItemID
	}
	tests := []struct {
		name   string
		target *Pointer
		args   args
		want   bool
	}{
		{
			name:   "true schema group only",
			target: &Pointer{schemaGroup: SchemaGroupID("xx").Ref()},
			args:   args{sg: SchemaGroupID("xx"), i: iid},
			want:   true,
		},
		{
			name:   "true item only",
			target: &Pointer{item: iid.Ref()},
			args:   args{sg: SchemaGroupID("xx"), i: iid},
			want:   true,
		},
		{
			name:   "true schema group and item",
			target: &Pointer{schemaGroup: SchemaGroupID("xx").Ref(), item: iid.Ref()},
			args:   args{sg: SchemaGroupID("xx"), i: iid},
			want:   true,
		},
		{
			name:   "true empty",
			target: &Pointer{},
			args:   args{sg: SchemaGroupID("xx"), i: iid},
			want:   true,
		},
		{
			name:   "false schema group only",
			target: &Pointer{schemaGroup: SchemaGroupID("xx").Ref()},
			args:   args{sg: SchemaGroupID("yy"), i: iid},
			want:   false,
		},
		{
			name:   "false item only",
			target: &Pointer{item: iid.Ref()},
			args:   args{sg: SchemaGroupID("xx"), i: NewItemID()},
			want:   false,
		},
		{
			name:   "false schema group and item",
			target: &Pointer{schemaGroup: SchemaGroupID("xx").Ref(), item: iid.Ref()},
			args:   args{sg: SchemaGroupID("xx"), i: NewItemID()},
			want:   false,
		},
		{
			name:   "false nil",
			target: nil,
			args:   args{sg: SchemaGroupID("xx"), i: iid},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.TestItem(tt.args.sg, tt.args.i))
		})
	}
}

func TestPointer_TestSchemaGroup(t *testing.T) {
	type args struct {
		sg SchemaGroupID
	}
	tests := []struct {
		name   string
		target *Pointer
		args   args
		want   bool
	}{
		{
			name:   "true",
			target: &Pointer{schemaGroup: SchemaGroupID("xx").Ref()},
			args:   args{sg: SchemaGroupID("xx")},
			want:   true,
		},
		{
			name:   "false",
			target: &Pointer{schemaGroup: SchemaGroupID("xx").Ref()},
			args:   args{sg: SchemaGroupID("yy")},
			want:   false,
		},
		{
			name:   "empty",
			target: &Pointer{},
			args:   args{sg: SchemaGroupID("xx")},
			want:   true,
		},
		{
			name:   "nil",
			target: nil,
			args:   args{sg: SchemaGroupID("xx")},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.TestSchemaGroup(tt.args.sg))
		})
	}
}

func TestPointer_TestField(t *testing.T) {
	type args struct {
		f FieldID
	}
	tests := []struct {
		name   string
		target *Pointer
		args   args
		want   bool
	}{
		{
			name:   "true",
			target: &Pointer{field: FieldID("xx").Ref()},
			args:   args{f: FieldID("xx")},
			want:   true,
		},
		{
			name:   "false",
			target: &Pointer{field: FieldID("xx").Ref()},
			args:   args{f: FieldID("yy")},
			want:   false,
		},
		{
			name:   "empty",
			target: &Pointer{},
			args:   args{f: FieldID("xx")},
			want:   true,
		},
		{
			name:   "nil",
			target: nil,
			args:   args{f: FieldID("xx")},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.TestField(tt.args.f))
		})
	}
}
