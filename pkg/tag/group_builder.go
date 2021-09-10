package tag

import (
	"errors"

	"github.com/reearth/reearth-backend/pkg/id"
)

var ErrEmptyTagList = errors.New("tag list can't be empty")

type GroupBuilder struct {
	g *Group
}

func NewGroup() *GroupBuilder {
	return &GroupBuilder{g: &Group{}}
}

func GroupFrom(t Tag) *Group {
	li, ok := t.(*Group)
	if !ok {
		return nil
	}
	return li
}

func (b *GroupBuilder) Build() (*Group, error) {
	if id.ID(b.g.id).IsNil() {
		return nil, id.ErrInvalidID
	}
	if id.ID(b.g.sceneId).IsNil() {
		return nil, ErrInvalidSceneID
	}
	if b.g.label == "" {
		return nil, ErrEmptyLabel
	}
	if len(b.g.tags.tags) == 0 {
		return nil, ErrEmptyTagList
	}
	return b.g, nil
}

func (b *GroupBuilder) ID(tid id.TagID) *GroupBuilder {
	b.g.id = tid
	return b
}

func (b *GroupBuilder) NewID() *GroupBuilder {
	b.g.id = id.NewTagID()
	return b
}

func (b *GroupBuilder) Label(l string) *GroupBuilder {
	b.g.label = l
	return b
}

func (b *GroupBuilder) Scene(sid id.SceneID) *GroupBuilder {
	b.g.sceneId = sid
	return b
}

func (b *GroupBuilder) Tags(tl *List) *GroupBuilder {
	b.g.tags = *tl
	return b
}
