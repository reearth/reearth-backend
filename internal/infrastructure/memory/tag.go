package memory

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
)

type Tag struct {
	//lock sync.Mutex
	data map[id.TagID]tag.Tag
}

func NewTag() repo.Tag {
	return &Tag{
		data: map[id.TagID]tag.Tag{},
	}
}

func (t Tag) FindByID(ctx context.Context, tagID id.TagID, ids []id.SceneID) (*tag.Tag, error) {
	panic("implement me")
}

func (t Tag) FindByIDs(ctx context.Context, ids []id.TagID, ids2 []id.SceneID) ([]*tag.Tag, error) {
	panic("implement me")
}

func (t Tag) FindItemByID(ctx context.Context, tagID id.TagID, ids []id.SceneID) (*tag.Item, error) {
	panic("implement me")
}

func (t Tag) FindItemByIDs(ctx context.Context, ids []id.TagID, ids2 []id.SceneID) ([]*tag.Item, error) {
	panic("implement me")
}

func (t Tag) FindGroupByID(ctx context.Context, tagID id.TagID, ids []id.SceneID) (*tag.Group, error) {
	panic("implement me")
}

func (t Tag) FindGroupByIDs(ctx context.Context, ids []id.TagID, ids2 []id.SceneID) ([]*tag.Group, error) {
	panic("implement me")
}

func (t Tag) FindByScene(ctx context.Context, sceneID id.SceneID) ([]*tag.Tag, error) {
	panic("implement me")
}

func (t Tag) Save(ctx context.Context, t2 tag.Tag) error {
	panic("implement me")
}

func (t Tag) SaveAll(ctx context.Context, tags []*tag.Tag) error {
	panic("implement me")
}

func (t Tag) Remove(ctx context.Context, tagID id.TagID) error {
	panic("implement me")
}

func (t Tag) RemoveAll(ctx context.Context, ids []id.TagID) error {
	panic("implement me")
}

func (t Tag) RemoveByScene(ctx context.Context, sceneID id.SceneID) error {
	panic("implement me")
}
