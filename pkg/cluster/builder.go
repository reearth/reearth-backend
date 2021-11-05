package cluster

import "github.com/reearth/reearth-backend/pkg/id"

type Builder struct {
	c *Cluster
}

func New() *Builder {
	return &Builder{c: &Cluster{}}
}

func (b *Builder) ID(cid id.ClusterID) *Builder {
	b.c.id = cid
	return b
}

func (b *Builder) NewID() *Builder {
	b.c.id = id.NewClusterID()
	return b
}

func (b *Builder) Name(n string) *Builder {
	b.c.name = n
	return b
}

func (b *Builder) Scene(sid id.SceneID) *Builder {
	b.c.scene = sid
	return b
}

func (b *Builder) Property(p id.PropertyID) *Builder {
	b.c.property = p
	return b
}

func (b *Builder) Build() (*Cluster, error) {
	if id.ID(b.c.id).IsNil() {
		return nil, id.ErrInvalidID
	}
	return b.c, nil
}
