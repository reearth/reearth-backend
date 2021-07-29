package tag

import (
	"errors"

	"github.com/reearth/reearth-backend/pkg/id"
)

type TagGroupBuilder struct {
	t *TagGroup
}

func NewTagGroup() *TagGroupBuilder {
	return &TagGroupBuilder{t: &TagGroup{}}
}

func (b *TagGroupBuilder) ID(id id.TagID) *TagGroupBuilder {
	b.t.tagBase.id = id
	return b
}

func (b *TagGroupBuilder) NewID() *TagGroupBuilder {
	b.t.tagBase.id = id.NewTagID()
	return b
}

func (b *TagGroupBuilder) Label(l string) *TagGroupBuilder {
	b.t.tagBase.label = l
	return b
}

func (b *TagGroupBuilder) Scene(sId id.SceneID) *TagGroupBuilder {
	b.t.tagBase.scene = sId
	return b
}

func (b *TagGroupBuilder) LinkedDatasetSchema(dsId *id.DatasetSchemaID) *TagGroupBuilder {
	b.t.tagBase.linkedDatasetSchema = dsId
	return b
}

func (b *TagGroupBuilder) Tags(ids []id.TagID) *TagGroupBuilder {
	b.t.tags = ids
	return b
}

func (b *TagGroupBuilder) Build() (*TagGroup, error) {
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

func (b *TagGroupBuilder) MustBuild() *TagGroup {
	tagGroup, err := b.Build()
	if err != nil {
		panic(err)
	}
	return tagGroup
}
