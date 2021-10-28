package cluster

import "github.com/reearth/reearth-backend/pkg/id"

type Cluster struct {
	name string
	// layers   []id.LayerID
	property id.PropertyID
}

func (c *Cluster) Name() string {
	if c == nil {
		return ""
	}
	return c.name
}

// func (c *Cluster) Layers() []id.LayerID {
// 	if c == nil {
// 		return nil
// 	}
// 	res := make([]id.LayerID, 0, len(c.layers))
// 	for _, l := range c.layers {
// 		res = append(res, l)
// 	}
// 	return res
// }

func (c *Cluster) Property() id.PropertyID {
	if c == nil {
		return id.PropertyID{}
	}
	return c.property
}
