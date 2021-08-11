package tag

import (
	"errors"

	"github.com/reearth/reearth-backend/pkg/id"
)

var (
	ErrEmptyLabel     = errors.New("tag label can't be empty")
	ErrInvalidSceneID = errors.New("invalid scene ID")
)

type Builder struct {
	t *tag
}

func New() *Builder {
	return &Builder{t: &tag{}}
}

func (b *Builder) Build() (*tag, error) {
	if id.ID(b.t.id).IsNil() {
		return nil, id.ErrInvalidID
	}
	if id.ID(b.t.sceneId).IsNil() {
		return nil, ErrInvalidSceneID
	}
	if b.t.label == "" {
		return nil, ErrEmptyLabel
	}
	return b.t, nil
}

func (b *Builder) ID(tid id.TagID) *Builder {
	b.t.id = tid
	return b
}

func (b *Builder) NewID() *Builder {
	b.t.id = id.NewTagID()
	return b
}

func (b *Builder) Label(l string) *Builder {
	b.t.label = l
	return b
}

func (b *Builder) Scene(sid id.SceneID) *Builder {
	b.t.sceneId = sid
	return b
}
