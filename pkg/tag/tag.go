package tag

import (
	"errors"

	"github.com/reearth/reearth-backend/pkg/id"
)

var (
	ErrEmptyLabel     = errors.New("tag label can't be empty")
	ErrInvalidSceneID = errors.New("invalid scene ID")
)

type tag struct {
	id      id.TagID
	label   string
	sceneId id.SceneID
}

type Tag interface {
	ID() id.TagID
	Scene() id.SceneID
	Label() string
}

func (t *tag) ID() id.TagID {
	return t.id
}

func (t *tag) Scene() id.SceneID {
	return t.sceneId
}

func (t *tag) Label() string {
	return t.label
}

func ToTagGroup(t Tag) *Group {
	if tg, ok := t.(*Group); ok {
		return tg
	}
	return nil
}

func ToTagItem(t Tag) *Item {
	if ti, ok := t.(*Item); ok {
		return ti
	}
	return nil
}
