package tag

import "github.com/reearth/reearth-backend/pkg/id"

type tag struct {
	id      id.TagID
	label   string
	sceneId id.SceneID
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
