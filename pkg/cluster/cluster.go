package cluster

import "github.com/reearth/reearth-backend/pkg/id"

type Cluster struct {
	id       id.ClusterID
	scene    id.SceneID
	name     string
	property id.PropertyID
}

func (c *Cluster) ID() id.ClusterID {
	if c == nil {
		return id.ClusterID{}
	}
	return c.id
}

func (c *Cluster) Scene() id.SceneID {
	if c == nil {
		return id.SceneID{}
	}
	return c.scene
}

func (c *Cluster) Name() string {
	if c == nil {
		return ""
	}
	return c.name
}

func (c *Cluster) Property() id.PropertyID {
	if c == nil {
		return id.PropertyID{}
	}
	return c.property
}
