package repo

import (
	"context"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
)

type Tag interface {
	FindGroupById(context.Context, id.TagID, id.SceneID) (tag.TagGroup, error)
	FindGroupByIds(context.Context, []id.TagID, id.SceneID) ([]tag.TagGroup, error)
	FindById(context.Context, id.TagID, id.SceneID) (tag.Tag, error)
	FindByIds(context.Context, []id.TagID, id.SceneID) ([]tag.Tag, error)
	Save(context.Context, tag.Tag) error
	SaveAll(context.Context, []id.TagID) error
}
