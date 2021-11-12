package scene

import "github.com/reearth/reearth-backend/pkg/id"

type List struct {
	clusters []*Cluster
}

func NewList() *List {
	return &List{clusters: []*Cluster{}}
}

func NewListFrom(clusters []*Cluster) *List {
	return &List{clusters: clusters}
}

func (tl *List) Clusters() []*Cluster {
	if tl == nil || tl.clusters == nil {
		return nil
	}
	return append([]*Cluster{}, tl.clusters...)
}

func (tl *List) Has(tid id.ClusterID) bool {
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

func (tl *List) Add(clusters ...*Cluster) {
	if tl == nil || tl.clusters == nil {
		return
	}
	tl.clusters = append(tl.clusters, clusters...)
}

func (tl *List) Update(cid id.ClusterID, name string, pid id.PropertyID) {
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

func (tl *List) Remove(clusters ...id.ClusterID) {
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
