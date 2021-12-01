package dataset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointAt(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	type args struct {
		d ID
		s SchemaID
		f FieldID
	}
	tests := []struct {
		name string
		args args
		want *Pointer
	}{
		{
			name: "ok",
			args: args{
				d: d,
				s: s,
				f: f,
			},
			want: &Pointer{
				dataset: &d,
				schema:  s,
				field:   f,
			},
		},
		{
			name: "empty",
			args: args{
				d: ID{},
				s: SchemaID{},
				f: FieldID{},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, PointAt(tt.args.d, tt.args.s, tt.args.f))
		})
	}
}

func TestPointAtField(t *testing.T) {
	s := NewSchemaID()
	f := NewFieldID()

	type args struct {
		s SchemaID
		f FieldID
	}
	tests := []struct {
		name string
		args args
		want *Pointer
	}{
		{
			name: "ok",
			args: args{
				s: s,
				f: f,
			},
			want: &Pointer{
				dataset: nil,
				schema:  s,
				field:   f,
			},
		},
		{
			name: "empty",
			args: args{
				s: SchemaID{},
				f: FieldID{},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, PointAtField(tt.args.s, tt.args.f))
		})
	}
}

func TestPointer_Dataset(t *testing.T) {
	d := NewID()

	tests := []struct {
		name   string
		target *Pointer
		want   *ID
	}{
		{
			name: "ok",
			target: &Pointer{
				dataset: &d,
			},
			want: &d,
		},
		{
			name:   "empty",
			target: &Pointer{},
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
			assert.Equal(t, tt.want, tt.target.Dataset())
		})
	}
}

func TestPointer_Schema(t *testing.T) {
	s := NewSchemaID()

	tests := []struct {
		name   string
		target *Pointer
		want   SchemaID
	}{
		{
			name: "ok",
			target: &Pointer{
				schema: s,
			},
			want: s,
		},
		{
			name:   "empty",
			target: &Pointer{},
			want:   SchemaID{},
		},
		{
			name:   "nil",
			target: nil,
			want:   SchemaID{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Schema())
		})
	}
}

func TestPointer_Field(t *testing.T) {
	f := NewFieldID()

	tests := []struct {
		name   string
		target *Pointer
		want   FieldID
	}{
		{
			name: "ok",
			target: &Pointer{
				field: f,
			},
			want: f,
		},
		{
			name:   "empty",
			target: &Pointer{},
			want:   FieldID{},
		},
		{
			name:   "nil",
			target: nil,
			want:   FieldID{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Field())
		})
	}
}

func TestPointer_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		target *Pointer
		want   bool
	}{
		{
			name:   "empty",
			target: &Pointer{},
			want:   true,
		},
		{
			name:   "nil",
			target: nil,
			want:   true,
		},
		{
			name: "not empty",
			target: &Pointer{
				schema: NewSchemaID(),
				field:  NewFieldID(),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.IsEmpty())
		})
	}
}

func TestPointer_Clone(t *testing.T) {
	d := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	tests := []struct {
		name   string
		target *Pointer
		want   *Pointer
	}{
		{
			name: "ok",
			target: &Pointer{
				dataset: &d,
				schema:  s,
				field:   f,
			},
			want: &Pointer{
				dataset: &d,
				schema:  s,
				field:   f,
			},
		},
		{
			name:   "empty",
			target: &Pointer{},
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
				assert.NotSame(t, tt.want, res)
			}
		})
	}
}

func TestPointer_PointAt(t *testing.T) {
	d := NewID()
	d2 := NewID()
	s := NewSchemaID()
	f := NewFieldID()

	type args struct {
		d *ID
	}
	tests := []struct {
		name   string
		target *Pointer
		args   args
		want   *Pointer
	}{
		{
			name: "ok1",
			target: &Pointer{
				dataset: nil,
				schema:  s,
				field:   f,
			},
			args: args{
				d: &d,
			},
			want: &Pointer{
				dataset: &d,
				schema:  s,
				field:   f,
			},
		},
		{
			name: "ok2",
			target: &Pointer{
				dataset: &d,
				schema:  s,
				field:   f,
			},
			args: args{
				d: &d2,
			},
			want: &Pointer{
				dataset: &d2,
				schema:  s,
				field:   f,
			},
		},
		{
			name: "ok3",
			target: &Pointer{
				dataset: &d,
				schema:  s,
				field:   f,
			},
			args: args{
				d: nil,
			},
			want: &Pointer{
				dataset: nil,
				schema:  s,
				field:   f,
			},
		},
		{
			name:   "empty",
			target: &Pointer{},
			args: args{
				d: &d,
			},
			want: nil,
		},
		{
			name: "nil",
			args: args{
				d: &d,
			},
			target: nil,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.PointAt(tt.args.d))
		})
	}
}
