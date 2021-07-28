package tag

import "github.com/reearth/reearth-backend/pkg/id"

type Group struct {
	TagBase
	tags []id.TagID
}

func (g *Group) Items() []id.TagID {
	if g == nil {
		return nil
	}
	if g.tags == nil {
		return nil
	}
	return g.tags
}

func (g *Group) AddItem(t id.TagID) {
	if g == nil {
		return
	}
	g.tags = append(g.tags, t)
}

func (g *Group) RemoveItem(t id.TagID) {
	if g == nil {
		return
	}
	var newTags []id.TagID
	for _, tagId := range g.tags {
		if tagId != t {
			newTags = append(newTags, t)
		}
	}
	g.tags = newTags
}
