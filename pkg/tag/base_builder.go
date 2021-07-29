package tag

import (
	"errors"

	"github.com/reearth/reearth-backend/pkg/id"
)

type Builder struct {
	base tagBase
}

func New() *Builder {
	return &Builder{
		base: tagBase{},
	}
}

func (b *Builder) ID(id id.TagID) *Builder {
	b.base.id = id
	return b
}

func (b *Builder) Label(l string) *Builder {
	b.base.label = l
	return b
}

func (b *Builder) Scene(sId id.SceneID) *Builder {
	b.base.scene = sId
	return b
}

func (b *Builder) LinkedDatasetSchema(dsId *id.DatasetSchemaID) *Builder {
	b.base.linkedDatasetSchema = dsId
	return b
}

func (b *Builder) Build() (*tagBase, error) {
	if b.base.id.IsNil() {
		return nil, id.ErrInvalidID
	}
	if b.base.label == "" {
		return nil, errors.New("label should not be empty")
	}
	if b.base.scene.IsNil() {
		return nil, id.ErrInvalidID
	}
	return &b.base, nil
}
