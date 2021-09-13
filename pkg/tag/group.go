package tag

type Group struct {
	tag
	tags List
}

func (g *Group) Tags() List {
	return g.tags
}

func (g *Group) Rename(s string) {
	g.label = s
}
