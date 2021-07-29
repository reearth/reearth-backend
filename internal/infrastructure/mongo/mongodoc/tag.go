package mongodoc

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
)

type TagDocument struct {
	ID                  string
	Label               string
	Scene               string
	LinkedDatasetSchema *string
}

type TagGroupDocument struct {
	ID                  string
	Label               string
	Scene               string
	LinkedDatasetSchema string
	Tags                []string
}

func NewTagDocument(t tag.Tag) (td *TagDocument, id string) {
	id = t.ID().String()
	return &TagDocument{
		ID:                  id,
		Label:               t.Label(),
		Scene:               t.Scene().String(),
		LinkedDatasetSchema: t.LinkedDatasetSchema().RefString(),
	}, id
}

func (d *TagDocument) Model() (*tag.Tag, error) {
	tId, err := id.TagIDFrom(d.ID)
	if err != nil {
		return nil, err
	}
	sid, err := id.SceneIDFrom(d.Scene)
	if err != nil {
		return nil, err
	}
	dsId, err := id.DatasetSchemaIDFrom(*d.LinkedDatasetSchema)

	return tag.NewTag().
		ID(tId).
		Label(d.Label).
		Scene(sid).
		LinkedDatasetSchema(&dsId).
		Build()
}
