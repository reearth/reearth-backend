package tag

import (
	"errors"

	"github.com/reearth/reearth-backend/pkg/id"
)

type TagBuilder struct {
	t *Tag
}

func NewTag() *TagBuilder {
	return &TagBuilder{t: &Tag{}}
}

func (b *TagBuilder) ID(id id.TagID) *TagBuilder {
	b.t.tagBase.id = id
	return b
}

func (b *TagBuilder) NewID() *TagBuilder {
	b.t.tagBase.id = id.NewTagID()
	return b
}

func (b *TagBuilder) Label(l string) *TagBuilder {
	b.t.tagBase.label = l
	return b
}

func (b *TagBuilder) Scene(sId id.SceneID) *TagBuilder {
	b.t.tagBase.scene = sId
	return b
}

func (b *TagBuilder) LinkedDatasetSchema(dsId *id.DatasetSchemaID) *TagBuilder {
	b.t.tagBase.linkedDatasetSchema = dsId
	return b
}

func (b *TagBuilder) Build() (*Tag, error) {
	if b.t.tagBase.id.IsNil() {
		return nil, id.ErrInvalidID
	}
	if b.t.tagBase.label == "" {
		return nil, errors.New("label should not be empty")
	}
	if b.t.tagBase.scene.IsNil() {
		return nil, id.ErrInvalidID
	}
	return b.t, nil
}

func (b *TagBuilder) MustBuild() *Tag {
	tag, err := b.Build()
	if err != nil {
		panic(err)
	}
	return tag
}
