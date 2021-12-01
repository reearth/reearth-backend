package dataset

import (
	"context"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewMigrator(t *testing.T) {
	di := NewID()
	si := NewSchemaID()
	fi := NewFieldID()
	d := func(ID) *ID { return &di }
	s := func(SchemaID) *SchemaID { return &si }
	f := func(FieldID) *FieldID { return &fi }

	type args struct {
		d func(ID) *ID
		s func(SchemaID) *SchemaID
		f func(FieldID) *FieldID
	}
	tests := []struct {
		name    string
		args    args
		wantnil bool
	}{
		{
			name:    "ok",
			args:    args{d: d, s: s, f: f},
			wantnil: false,
		},
		{
			name:    "nil",
			args:    args{d: nil, s: s, f: f},
			wantnil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMigrator(tt.args.d, tt.args.s, tt.args.f)
			if !tt.wantnil {
				assert.Equal(t, tt.args.d(ID{}), m.d(ID{}))
				assert.Equal(t, tt.args.s(SchemaID{}), m.s(SchemaID{}))
				assert.Equal(t, tt.args.f(FieldID{}), m.f(FieldID{}))
			} else {
				assert.Nil(t, m)
			}
		})
	}
}

func TestMigratorFrom(t *testing.T) {
	d1 := NewID()
	d2 := NewID()
	s1 := NewSchemaID()
	s2 := NewSchemaID()
	f1 := NewFieldID()
	f2 := NewFieldID()
	d := map[ID]ID{d1: d2}
	s := map[SchemaID]SchemaID{s1: s2}
	f := map[FieldID]FieldID{f1: f2}

	type args struct {
		d map[ID]ID
		s map[SchemaID]SchemaID
		f map[FieldID]FieldID
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{d: d, s: s, f: f},
		},
		{
			name: "partial nil",
			args: args{d: d, s: nil, f: nil},
		},
		{
			name: "nil",
			args: args{d: nil, s: nil, f: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := MigratorFrom(tt.args.d, tt.args.s, tt.args.f)
			if r, ok := tt.args.d[d1]; ok {
				assert.Equal(t, &r, res.d(d1))
			} else {
				assert.Nil(t, res.d(d1))
			}
			if r, ok := tt.args.s[s1]; ok {
				assert.Equal(t, &r, res.s(s1))
			} else {
				assert.Nil(t, res.s(s1))
			}
			if r, ok := tt.args.f[f1]; ok {
				assert.Equal(t, &r, res.f(f1))
			} else {
				assert.Nil(t, res.f(f1))
			}
			assert.Nil(t, res.d(d2))
			assert.Nil(t, res.s(s2))
			assert.Nil(t, res.f(f2))
		})
	}
}

func TestMigrator_MigrateAndValidateGraphPointer(t *testing.T) {
	d1 := NewID()
	d2 := NewID()
	s1 := NewSchemaID()
	s2 := NewSchemaID()
	f1 := NewFieldID()
	f2 := NewFieldID()
	d := func(i ID) *ID {
		if i == d1 {
			return &d2
		}
		return &i
	}
	s := func(i SchemaID) *SchemaID {
		if i == s1 {
			return &s2
		}
		return &i
	}
	f := func(i FieldID) *FieldID {
		if i == f1 {
			return &f2
		}
		return &i
	}
	l := GraphLoaderFromMap(Map{
		d2: New().ID(d2).Schema(s2).Scene(id.NewSceneID()).Fields([]*Field{
			NewField(f2, ValueTypeString.ValueFrom("foo"), ""),
		}).MustBuild(),
	})

	type args struct {
		ctx context.Context
		p   *GraphPointer
		l   GraphLoader
	}
	tests := []struct {
		name    string
		target  *Migrator
		args    args
		want    *GraphPointer
		wantErr error
	}{
		{
			name:   "ok",
			target: &Migrator{d: d, s: s, f: f},
			args: args{
				ctx: context.Background(),
				p: &GraphPointer{
					pointers: []*Pointer{
						{dataset: &d1, schema: s1, field: f1},
					},
				},
				l: l,
			},
			want: &GraphPointer{
				pointers: []*Pointer{
					{dataset: &d2, schema: s2, field: f2},
				},
			},
			wantErr: nil,
		},
		{
			name:   "invalid",
			target: &Migrator{d: nil, s: s, f: f},
			args: args{
				ctx: context.Background(),
				p: &GraphPointer{
					pointers: []*Pointer{
						{dataset: &d1, schema: s1, field: f1},
					},
				},
				l: l,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name:   "nil pointer",
			target: &Migrator{d: d, s: s, f: f},
			args: args{
				ctx: context.Background(),
				p:   nil,
				l:   l,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name:   "empty",
			target: &Migrator{},
			args: args{
				ctx: context.Background(),
				p: &GraphPointer{
					pointers: []*Pointer{
						{dataset: &d1, schema: s1, field: f1},
					},
				},
				l: l,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name:   "nil",
			target: nil,
			args: args{
				ctx: context.Background(),
				p: &GraphPointer{
					pointers: []*Pointer{
						{dataset: &d1, schema: s1, field: f1},
						{dataset: nil, schema: s2, field: f1},
					},
				},
				l: l,
			},
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.target.MigrateAndValidateGraphPointer(tt.args.ctx, tt.args.p, tt.args.l)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.wantErr, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMigrator_MigrateGraphPointer(t *testing.T) {
	d1 := NewID()
	d2 := NewID()
	s1 := NewSchemaID()
	s2 := NewSchemaID()
	f1 := NewFieldID()
	f2 := NewFieldID()
	d := func(i ID) *ID {
		if i == d1 {
			return &d2
		}
		return &i
	}
	s := func(i SchemaID) *SchemaID {
		if i == s1 {
			return &s2
		}
		return &i
	}
	f := func(i FieldID) *FieldID {
		if i == f1 {
			return &f2
		}
		return &i
	}

	type args struct {
		p *GraphPointer
	}
	tests := []struct {
		name   string
		target *Migrator
		args   args
		want   *GraphPointer
	}{
		{
			name:   "ok",
			target: &Migrator{d: d, s: s, f: f},
			args: args{
				p: &GraphPointer{
					pointers: []*Pointer{
						{dataset: &d1, schema: s1, field: f1},
						{dataset: nil, schema: s2, field: f1},
					},
				},
			},
			want: &GraphPointer{
				pointers: []*Pointer{
					{dataset: &d2, schema: s2, field: f2},
					{dataset: nil, schema: s2, field: f2},
				},
			},
		},
		{
			name:   "nil pointer",
			target: &Migrator{d: d, s: s, f: f},
			args: args{
				p: nil,
			},
			want: nil,
		},
		{
			name:   "empty",
			target: &Migrator{},
			args: args{
				p: &GraphPointer{
					pointers: []*Pointer{
						{dataset: &d1, schema: s1, field: f1},
						{dataset: nil, schema: s2, field: f1},
					},
				},
			},
			want: &GraphPointer{
				pointers: []*Pointer{
					{dataset: &d1, schema: s1, field: f1},
					{dataset: nil, schema: s2, field: f1},
				},
			},
		},
		{
			name:   "nil",
			target: nil,
			args: args{
				p: &GraphPointer{
					pointers: []*Pointer{
						{dataset: &d1, schema: s1, field: f1},
						{dataset: nil, schema: s2, field: f1},
					},
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.MigrateGraphPointer(tt.args.p)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMigrator_MigratePointer(t *testing.T) {
	d1 := NewID()
	d2 := NewID()
	s1 := NewSchemaID()
	s2 := NewSchemaID()
	f1 := NewFieldID()
	f2 := NewFieldID()
	d := func(i ID) *ID {
		if i == d1 {
			return &d2
		}
		return &i
	}
	s := func(i SchemaID) *SchemaID {
		if i == s1 {
			return &s2
		}
		return &i
	}
	f := func(i FieldID) *FieldID {
		if i == f1 {
			return &f2
		}
		return &i
	}

	type args struct {
		p *Pointer
	}
	tests := []struct {
		name   string
		target *Migrator
		args   args
		want   *Pointer
	}{
		{
			name:   "ok",
			target: &Migrator{d: d, s: s, f: f},
			args: args{
				p: &Pointer{
					dataset: &d1,
					schema:  s1,
					field:   f1,
				},
			},
			want: &Pointer{
				dataset: &d2,
				schema:  s2,
				field:   f2,
			},
		},
		{
			name:   "partial",
			target: &Migrator{d: d, s: s, f: f},
			args: args{
				p: &Pointer{
					dataset: nil,
					schema:  s2,
					field:   f1,
				},
			},
			want: &Pointer{
				dataset: nil,
				schema:  s2,
				field:   f2,
			},
		},
		{
			name:   "nil pointer",
			target: &Migrator{d: d, s: s, f: f},
			args: args{
				p: nil,
			},
			want: nil,
		},
		{
			name:   "empty",
			target: &Migrator{},
			args: args{
				p: &Pointer{
					dataset: &d1,
					schema:  s1,
					field:   f1,
				},
			},
			want: &Pointer{
				dataset: &d1,
				schema:  s1,
				field:   f1,
			},
		},
		{
			name:   "nil",
			target: nil,
			args: args{
				p: &Pointer{
					dataset: &d1,
					schema:  s1,
					field:   f1,
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.MigratePointer(tt.args.p)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMigrator_getD(t *testing.T) {
	d1 := NewID()
	d2 := NewID()
	d := func(i ID) *ID {
		if i == d1 {
			return &d2
		}
		return &i
	}

	type args struct {
		d ID
	}
	tests := []struct {
		name   string
		target *Migrator
		args   args
		want   ID
	}{
		{
			name:   "found",
			target: &Migrator{d: d},
			args:   args{d: d1},
			want:   d2,
		},
		{
			name:   "not found",
			target: &Migrator{d: d},
			args:   args{d: d2},
			want:   d2,
		},
		{
			name:   "empty",
			target: &Migrator{},
			args:   args{d: d1},
			want:   d1,
		},
		{
			name:   "nil",
			target: nil,
			args:   args{d: d1},
			want:   d1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.getD(tt.args.d)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMigrator_getS(t *testing.T) {
	s1 := NewSchemaID()
	s2 := NewSchemaID()
	s := func(i SchemaID) *SchemaID {
		if i == s1 {
			return &s2
		}
		return &i
	}

	type args struct {
		s SchemaID
	}
	tests := []struct {
		name   string
		target *Migrator
		args   args
		want   SchemaID
	}{
		{
			name:   "found",
			target: &Migrator{s: s},
			args:   args{s: s1},
			want:   s2,
		},
		{
			name:   "not found",
			target: &Migrator{s: s},
			args:   args{s: s2},
			want:   s2,
		},
		{
			name:   "empty",
			target: &Migrator{},
			args:   args{s: s1},
			want:   s1,
		},
		{
			name:   "nil",
			target: nil,
			args:   args{s: s1},
			want:   s1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.getS(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMigrator_getF(t *testing.T) {
	f1 := NewFieldID()
	f2 := NewFieldID()
	f := func(i FieldID) *FieldID {
		if i == f1 {
			return &f2
		}
		return &i
	}

	type args struct {
		f FieldID
	}
	tests := []struct {
		name   string
		target *Migrator
		args   args
		want   FieldID
	}{
		{
			name:   "found",
			target: &Migrator{f: f},
			args:   args{f: f1},
			want:   f2,
		},
		{
			name:   "not found",
			target: &Migrator{f: f},
			args:   args{f: f2},
			want:   f2,
		},
		{
			name:   "empty",
			target: &Migrator{},
			args:   args{f: f1},
			want:   f1,
		},
		{
			name:   "nil",
			target: nil,
			args:   args{f: f1},
			want:   f1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.getF(tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
