package dataset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraphPointer(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	type args struct {
		p []*Pointer
	}
	tests := []struct {
		name string
		args args
		want *GraphPointer
	}{
		{
			name: "ok",
			args: args{
				p: []*Pointer{PointAt(d, s, f), nil, PointAtField(s, f)},
			},
			want: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f)},
			},
		},
		{
			name: "invalid",
			args: args{
				p: []*Pointer{PointAtField(s, f)},
			},
			want: nil,
		},
		{
			name: "empty",
			args: args{
				p: []*Pointer{},
			},
			want: nil,
		},
		{
			name: "nil",
			args: args{
				p: nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewGraphPointer(tt.args.p))
		})
	}
}

func TestGraphPointer_Clone(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   *GraphPointer
	}{
		{
			name: "ok",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAt(d, s, f)},
			},
			want: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAt(d, s, f)},
			},
		},
		{
			name:   "empty",
			target: &GraphPointer{},
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
			res := tt.target.Clone()
			assert.Equal(t, tt.want, res)
			if tt.want != nil {
				assert.NotSame(t, tt.target, res)
			}
		})
	}
}

func TestGraphPointer_WithDataset(t *testing.T) {
	d := NewID()
	d2 := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	type args struct {
		ds ID
	}
	tests := []struct {
		name   string
		target *GraphPointer
		args   args
		want   *GraphPointer
	}{
		{
			name: "ok",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAt(d, s, f)},
			},
			args: args{
				ds: d2,
			},
			want: &GraphPointer{
				[]*Pointer{PointAt(d2, s, f), PointAt(d, s, f)},
			},
		},
		{
			name:   "empty",
			target: &GraphPointer{},
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
			assert.Equal(t, tt.want, tt.target.WithDataset(tt.args.ds))
		})
	}
}

func TestGraphPointer_IsEmpty(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   bool
	}{
		{
			name: "present",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f)},
			},
			want: false,
		},
		{
			name:   "empty",
			target: &GraphPointer{},
			want:   true,
		},
		{
			name:   "nil",
			target: nil,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.IsEmpty())
		})
	}
}

func TestGraphPointer_IsLinkedFully(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   bool
	}{
		{
			name: "ok",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f)},
			},
			want: true,
		},
		{
			name: "false",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f)},
			},
			want: false,
		},
		{
			name:   "empty",
			target: &GraphPointer{},
			want:   false,
		},
		{
			name:   "nil",
			target: nil,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.IsLinkedFully())
		})
	}
}

func TestGraphPointer_Len(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   int
	}{
		{
			name: "1",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f)},
			},
			want: 1,
		},
		{
			name: "2",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f)},
			},
			want: 2,
		},
		{
			name:   "empty",
			target: &GraphPointer{},
			want:   0,
		},
		{
			name:   "nil",
			target: nil,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Len())
		})
	}
}

func TestGraphPointer_First(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   *Pointer
	}{
		{
			name: "ok",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f)},
			},
			want: PointAt(d, s, f),
		},
		{
			name:   "empty",
			target: &GraphPointer{},
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
			assert.Equal(t, tt.want, tt.target.First())
		})
	}
}

func TestGraphPointer_Last(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   *Pointer
	}{
		{
			name: "ok",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f)},
			},
			want: PointAtField(s, f),
		},
		{
			name:   "empty",
			target: &GraphPointer{},
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
			assert.Equal(t, tt.want, tt.target.Last())
		})
	}
}

func TestGraphPointer_Pointers(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   []*Pointer
	}{
		{
			name: "ok",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f)},
			},
			want: []*Pointer{PointAt(d, s, f), PointAtField(s, f)},
		},
		{
			name:   "empty",
			target: &GraphPointer{},
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
			res := tt.target.Pointers()
			assert.Equal(t, tt.want, res)
			if len(res) > 0 {
				res2 := append([]*Pointer{}, res...)
				res[0] = nil
				assert.Equal(t, res2, tt.target.Pointers()) // result not changed
			}
		})
	}
}

func TestGraphPointer_Datasets(t *testing.T) {
	d := NewID()
	d2 := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   []ID
	}{
		{
			name: "ok",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f)},
			},
			want: []ID{d},
		},
		{
			name: "ok2",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAt(d2, s, f)},
			},
			want: []ID{d, d2},
		},
		{
			name:   "empty",
			target: &GraphPointer{},
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
			assert.Equal(t, tt.want, tt.target.Datasets())
		})
	}
}

func TestGraphPointer_Schemas(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	s2 := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   []SchemaID
	}{
		{
			name: "ok",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s2, f)},
			},
			want: []SchemaID{s, s2},
		},
		{
			name:   "empty",
			target: &GraphPointer{},
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
			assert.Equal(t, tt.want, tt.target.Schemas())
		})
	}
}

func TestGraphPointer_Fields(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()
	f2 := NewFieldID()

	tests := []struct {
		name   string
		target *GraphPointer
		want   []FieldID
	}{
		{
			name: "ok",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f2)},
			},
			want: []FieldID{f, f2},
		},
		{
			name:   "empty",
			target: &GraphPointer{},
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
			assert.Equal(t, tt.want, tt.target.Fields())
		})
	}
}

func TestGraphPointer_HasDataset(t *testing.T) {
	d := NewID()
	d2 := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	type args struct {
		did ID
	}
	tests := []struct {
		name   string
		target *GraphPointer
		args   args
		want   bool
	}{
		{
			name: "found",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f)},
			},
			args: args{
				did: d,
			},
			want: true,
		},
		{
			name: "not found",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s, f)},
			},
			args: args{
				did: d2,
			},
			want: false,
		},
		{
			name:   "empty",
			target: &GraphPointer{},
			args: args{
				did: d,
			},
			want: false,
		},
		{
			name:   "nil",
			target: nil,
			args: args{
				did: d,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.target.HasDataset(tt.args.did); got != tt.want {
				assert.Equal(t, tt.want, tt.target.HasDataset(tt.args.did))
			}
		})
	}
}

func TestGraphPointer_HasSchema(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	s2 := NewSchemaID()
	s3 := NewSchemaID()
	f := NewFieldID()

	type args struct {
		dsid SchemaID
	}
	tests := []struct {
		name   string
		target *GraphPointer
		args   args
		want   bool
	}{
		{
			name: "found",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s2, f)},
			},
			args: args{
				dsid: s2,
			},
			want: true,
		},
		{
			name: "not found",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s2, f)},
			},
			args: args{
				dsid: s3,
			},
			want: false,
		},
		{
			name:   "empty",
			target: &GraphPointer{},
			args: args{
				dsid: s,
			},
			want: false,
		},
		{
			name:   "nil",
			target: nil,
			args: args{
				dsid: s,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.target.HasSchema(tt.args.dsid); got != tt.want {
				assert.Equal(t, tt.want, tt.target.HasSchema(tt.args.dsid))
			}
		})
	}
}

func TestGraphPointer_HasSchemaAndDataset(t *testing.T) {
	d := NewID()
	d2 := NewID()
	s := NewSchemaID()
	s2 := NewSchemaID()
	f := NewFieldID()
	f2 := NewFieldID()

	type args struct {
		dsid SchemaID
		did  ID
	}
	tests := []struct {
		name   string
		target *GraphPointer
		args   args
		want   bool
	}{
		{
			name: "true",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s2, f2)},
			},
			args: args{
				dsid: s,
				did:  d,
			},
			want: true,
		},
		{
			name: "false",
			target: &GraphPointer{
				[]*Pointer{PointAt(d, s, f), PointAtField(s2, f2)},
			},
			args: args{
				dsid: s2,
				did:  d2,
			},
			want: false,
		},
		{
			name:   "empty",
			target: &GraphPointer{},
			args: args{
				dsid: s,
				did:  d,
			},
			want: false,
		},
		{
			name:   "nil",
			target: nil,
			args: args{
				dsid: s,
				did:  d,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.HasSchemaAndDataset(tt.args.dsid, tt.args.did))
		})
	}
}
