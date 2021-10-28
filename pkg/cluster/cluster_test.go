package cluster

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestCluster_Name(t *testing.T) {
	type fields struct {
		name     string
		property id.PropertyID
	}
	clusterA := &Cluster{
		name: "clusterA",
	}
	clusterB := &Cluster{}
	tests := []struct {
		name    string
		cluster *Cluster
		want    string
	}{
		{
			name:    "should return cluster name",
			cluster: clusterA,
			want:    "clusterA",
		},
		{
			name:    "should return empty if name is nil",
			cluster: clusterB,
			want:    "",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			got := tc.cluster.Name()
			assert.Equal(tt, tc.want, got)
		})
	}
}
func TestCluster_Property(t *testing.T) {
	type fields struct {
		name     string
		property id.PropertyID
	}
	propertyId := id.NewPropertyID()
	clusterA := &Cluster{
		property: propertyId,
	}
	clusterB := &Cluster{}
	tests := []struct {
		name    string
		cluster *Cluster
		want    bool
	}{
		{
			name:    "should be true if it returns cluster property",
			cluster: clusterA,
			want:    true,
		},
		{
			name:    "should be false if it returns cluster",
			cluster: clusterB,
			want:    false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			// got := tc.cluster.Name()
			// assert.Equal(tt,tc.want,)
		})
	}
}
