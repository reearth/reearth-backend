package interfaces

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
)

type CreateTagItemParam struct {
	Label                 string
	SceneID               id.SceneID
	LinkedDatasetSchemaID *id.DatasetSchemaID
	LinkedDatasetID       *id.DatasetID
	LinkedDatasetField    *id.DatasetSchemaFieldID
}

type CreateTagGroupParam struct {
	Label   string
	SceneID id.SceneID
	Tags    []id.TagID
}

type AttachItemToGroupParam struct {
	ItemID, GroupID id.TagID
}

type DetachItemToGroupParam struct {
	ItemID, GroupID id.TagID
}

type Tag interface {
	Fetch(context.Context, []id.TagID, *usecase.Operator) ([]*tag.Tag, error)
	FetchItem(context.Context, []id.TagID, *usecase.Operator) ([]*tag.Item, error)
	FetchGroup(context.Context, []id.TagID, *usecase.Operator) ([]*tag.Group, error)
	FetchGroupsByLayer(context.Context, id.LayerID, *usecase.Operator) ([]*tag.Group, error)
	FetchGroupsByScene(context.Context, id.SceneID, *usecase.Operator) ([]*tag.Group, error)
	FetchItemsByLayer(context.Context, id.LayerID, *usecase.Operator) ([]*tag.Item, error)
	FetchItemsByScene(context.Context, id.SceneID, *usecase.Operator) ([]*tag.Item, error)
	CreateItem(context.Context, CreateTagItemParam, *usecase.Operator) (*tag.Item, error)
	CreateGroup(context.Context, CreateTagGroupParam, *usecase.Operator) (*tag.Group, error)
	AttachItemToGroup(context.Context, AttachItemToGroupParam, *usecase.Operator) (*tag.Group, error)
	DetachItemFromGroup(context.Context, DetachItemToGroupParam, *usecase.Operator) (*tag.Group, error)
}
