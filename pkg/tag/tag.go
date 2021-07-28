package tag

import "github.com/reearth/reearth-backend/pkg/id"

type TagBase struct {
	id                  id.TagID
	label               string
	scene               id.SceneID
	linkedDatasetSchema *id.DatasetSchemaID
}

func (t *TagBase) ID() id.TagID {
	return t.id
}

func (t *TagBase) Label() string {
	return t.label
}

func (t *TagBase) Scene() id.SceneID {
	return t.scene
}

func (t *TagBase) SetLabel(label string) {
	t.label = label
}

func (t *TagBase) IsLinked() bool {
	if t == nil {
		return false
	}
	return t.linkedDatasetSchema != nil
}

func (t *TagBase) LinkedDatasetSchema() *id.DatasetSchemaID {
	if t == nil || t.linkedDatasetSchema == nil {
		return nil
	}
	id := *t.linkedDatasetSchema
	return &id
}

func (t *TagBase) LinkToDatasetSchema(ds id.DatasetSchemaID) {
	if t == nil {
		return
	}
	ds2 := ds
	t.linkedDatasetSchema = &ds2
}

func (t *TagBase) Unlink() {
	if t == nil {
		return
	}
	t.linkedDatasetSchema = nil
}
