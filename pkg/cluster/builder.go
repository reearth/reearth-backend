package cluster

import "github.com/reearth/reearth-backend/pkg/id"

type Builder struct {
	c *Cluster
}

func New() *Builder {
	return &Builder{c: &Cluster{}}
}

func (b *Builder) Name(n string) *Builder {
	b.c.name = n
	return b
}

// func (b *Builder) Layers(l []id.LayerID) *Builder {
// 	b.c.layers = l
// 	return b
// }

func (b *Builder) Property(p id.PropertyID) *Builder {
	b.c.property = p
	return b
}

func (b *Builder) Build() (*Cluster, error) {
	return b.c, nil
}
