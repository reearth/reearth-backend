package tag

import "github.com/reearth/reearth-backend/pkg/id"

type TagGroup struct {
	tagBase
	tags []id.TagID
}

func (g *TagGroup) Items() []id.TagID {
	if g == nil {
		return nil
	}
	if g.tags == nil {
		return nil
	}
	return g.tags
}

func (g *TagGroup) AddItem(t id.TagID) {
	if g == nil {
		return
	}
	g.tags = append(g.tags, t)
}

func (g *TagGroup) RemoveItem(t id.TagID) {
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
