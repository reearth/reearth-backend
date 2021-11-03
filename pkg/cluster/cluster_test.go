package cluster

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestCluster_ID(t *testing.T) {
	cid := id.NewClusterID()
	clusterA := &Cluster{
		id: cid,
	}
	tests := []struct {
		name    string
		cluster *Cluster
		want    id.ClusterID
	}{
		{
			name:    "should return cluster id",
			cluster: clusterA,
			want:    cid,
		},
		{
			name:    "should return empty if cluster is nil",
			cluster: nil,
			want:    id.ClusterID{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			got := tc.cluster.ID()
			assert.Equal(tt, tc.want, got)
		})
	}
}
func TestCluster_Name(t *testing.T) {
	clusterA := &Cluster{
		name: "clusterA",
	}
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
			name:    "should return empty if cluster is nil",
			cluster: nil,
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
	propertyId := id.NewPropertyID()
	clusterA := &Cluster{
		property: propertyId,
	}
	tests := []struct {
		name    string
		cluster *Cluster
		want    id.PropertyID
	}{
		{
			name:    "should return cluster property",
			cluster: clusterA,
			want:    propertyId,
		},
		{
			name:    "should return empty cluster property",
			cluster: nil,
			want:    id.PropertyID{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			got := tc.cluster.Property()
			assert.Equal(tt, tc.want, got)
		})
	}
}
