package tag

import "github.com/reearth/reearth-backend/pkg/id"

type tagBase struct {
	id                  id.TagID
	label               string
	scene               id.SceneID
	linkedDatasetSchema *id.DatasetSchemaID
}

type Tag struct {
	tagBase tagBase
}

func (t *tagBase) ID() id.TagID {
	return t.id
}

func (t *tagBase) Label() string {
	return t.label
}

func (t *tagBase) Scene() id.SceneID {
	return t.scene
}

func (t *tagBase) SetLabel(label string) {
	t.label = label
}

func (t *tagBase) IsLinked() bool {
	if t == nil {
		return false
	}
	return t.linkedDatasetSchema != nil
}

func (t *Tag) ID() id.TagID {
	return t.tagBase.ID()
}

func (t *Tag) Label() string {
	return t.tagBase.Label()
}

func (t *Tag) Scene() id.SceneID {
	return t.tagBase.Scene()
}

func (t *Tag) SetLabel(label string) {
	t.tagBase.label = label
}

func (t *Tag) IsLinked() bool {
	if t == nil {
		return false
	}
	return t.tagBase.LinkedDatasetSchema() != nil
}

func (t *Tag) LinkedDatasetSchema() *id.DatasetSchemaID {
	if t == nil || t.tagBase.linkedDatasetSchema == nil {
		return nil
	}
	id := t.tagBase.linkedDatasetSchema
	return id
}

func (t *tagBase) LinkedDatasetSchema() *id.DatasetSchemaID {
	if t == nil || t.linkedDatasetSchema == nil {
		return nil
	}
	id := *t.linkedDatasetSchema
	return &id
}

func (t *tagBase) LinkToDatasetSchema(ds id.DatasetSchemaID) {
	if t == nil {
		return
	}
	ds2 := ds
	t.linkedDatasetSchema = &ds2
}

func (t *tagBase) Unlink() {
	if t == nil {
		return
	}
	t.linkedDatasetSchema = nil
}
