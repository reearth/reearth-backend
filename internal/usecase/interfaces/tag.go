package interfaces

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
)

var (
	ErrNonemptyTagGroupCannotDelete = errors.New("can't delete non-empty tag group")
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

type UpdateTagParam struct {
	Label   *string
	SceneID id.SceneID
	TagID   id.TagID
}

type Tag interface {
	Fetch(context.Context, []id.TagID, *usecase.Operator) ([]*tag.Tag, error)
	FetchByScene(context.Context, id.SceneID, *usecase.Operator) ([]*tag.Tag, error)
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
	UpdateTag(context.Context, UpdateTagParam, *usecase.Operator) (*tag.Tag, error)
	Remove(context.Context, id.TagID, *usecase.Operator) (*id.TagID, error)
}