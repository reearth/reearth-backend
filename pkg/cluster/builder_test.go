package cluster

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Build(t *testing.T) {
	propertyId := id.NewPropertyID()
	clusterId := id.NewClusterID()
	sid := id.NewSceneID()
	tests := []struct {
		name    string
		builder *Builder
		want    *Cluster
		wantErr bool
	}{
		{
			name:    "build with name and property",
			builder: New().ID(clusterId).Name("ccc").Scene(sid).Property(propertyId),
			want: &Cluster{
				id:       clusterId,
				name:     "ccc",
				property: propertyId,
				scene:    sid,
			},
			wantErr: false,
		},
		{
			name:    "build empty cluster",
			builder: New(),
			wantErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			got, err := tc.builder.Build()
			assert.Equal(tt, tc.wantErr, err != nil)
			assert.Equal(tt, tc.want, got)
		})
	}
}

func TestBuilder_NewID(t *testing.T) {
	builder := New().NewID()
	assert.False(t, builder.c.id.IsNil())
}
