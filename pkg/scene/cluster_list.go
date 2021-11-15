package scene

import "github.com/reearth/reearth-backend/pkg/id"

type ClusterList struct {
	clusters []*Cluster
}

func NewClusterList() *ClusterList {
	return &ClusterList{clusters: []*Cluster{}}
}

func NewClusterListFrom(clusters []*Cluster) *ClusterList {
	return &ClusterList{clusters: clusters}
}

func (tl *ClusterList) Clusters() []*Cluster {
	if tl == nil || tl.clusters == nil {
		return nil
	}
	return append([]*Cluster{}, tl.clusters...)
}

func (tl *ClusterList) Has(tid id.ClusterID) bool {
	if tl == nil || tl.clusters == nil {
		return false
	}
	for _, cluster := range tl.clusters {
		if cluster.ID() == tid {
			return true
		}
	}
	return false
}

func (tl *ClusterList) Add(clusters ...*Cluster) {
	if tl == nil || tl.clusters == nil {
		return
	}
	tl.clusters = append(tl.clusters, clusters...)
}

func (tl *ClusterList) Update(cid id.ClusterID, name string, pid id.PropertyID) {
	if tl == nil || tl.clusters == nil || !tl.Has(cid) {
		return
	}

	for _, c := range tl.clusters {
		if c.ID() == cid {
			c.Rename(name)
			c.UpdateProperty(pid)
		}
	}
}

func (tl *ClusterList) Remove(clusters ...id.ClusterID) {
	if tl == nil || tl.clusters == nil {
		return
	}
	for i := 0; i < len(tl.clusters); i++ {
		for _, tid := range clusters {
			if tl.clusters[i].id == tid {
				tl.clusters = append(tl.clusters[:i], tl.clusters[i+1:]...)
				i--
				break
			}
		}
	}
}
