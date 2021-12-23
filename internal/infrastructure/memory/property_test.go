package memory

import (
	"context"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/stretchr/testify/assert"
)

func TestProperty_FindBySchema(t *testing.T) {
	p1 := id.NewPropertyID()
	p2 := id.NewPropertyID()
	p3 := id.NewPropertyID()
	p4 := id.NewPropertyID()
	p5 := id.NewPropertyID()
	ps1 := id.MustPropertySchemaID("a~1.0.0/a")
	ps2 := id.MustPropertySchemaID("a~1.0.0/b")
	s1 := id.NewSceneID()
	s2 := id.NewSceneID()

	type args struct {
		in0     context.Context
		schemas []id.PropertySchemaID
		s       id.SceneID
	}
	tests := []struct {
		name    string
		target  *Property
		args    args
		want    property.List
		wantErr error
	}{
		{
			name: "found",
			target: &Property{
				data: map[id.PropertyID]property.Property{
					p1: *property.New().ID(p1).Scene(s1).Schema(ps1).MustBuild(),
					p2: *property.New().ID(p2).Scene(s1).Schema(ps2).MustBuild(),
					p3: *property.New().ID(p3).Scene(s2).Schema(ps1).MustBuild(),
					p4: *property.New().ID(p4).Scene(s2).Schema(ps2).MustBuild(),
					p5: *property.New().ID(p5).Scene(s1).Schema(ps1).MustBuild(),
				},
			},
			args: args{
				in0:     nil,
				schemas: []id.PropertySchemaID{ps1, ps2},
				s:       s1,
			},
			want: property.List{
				property.New().ID(p1).Scene(s1).Schema(ps1).MustBuild(),
				property.New().ID(p2).Scene(s1).Schema(ps2).MustBuild(),
				property.New().ID(p5).Scene(s1).Schema(ps1).MustBuild(),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.target.FindBySchema(tt.args.in0, tt.args.schemas, tt.args.s)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
