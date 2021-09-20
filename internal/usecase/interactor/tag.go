package interactor

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/pkg/id"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/tag"
)

type Tag struct {
	commonScene
	tagRepo     repo.Tag
	sceneRepo   repo.Scene
	transaction repo.Transaction
}

func NewTag(r *repo.Container) interfaces.Tag {
	return &Tag{
		commonScene: commonScene{sceneRepo: r.Scene},
		tagRepo:     r.Tag,
		sceneRepo:   r.Scene,
		transaction: r.Transaction,
	}
}

func (i *Tag) CreateItem(ctx context.Context, inp interfaces.CreateTagItemParam, operator *usecase.Operator) (*tag.Item, error) {
	tx, err := i.transaction.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	if err := i.CanWriteScene(ctx, inp.SceneID, operator); err != nil {
		return nil, interfaces.ErrOperationDenied
	}

	builder := tag.NewItem().
		NewID().
		Label(inp.Label).
		Scene(inp.SceneID)
	if inp.LinkedDatasetSchemaID != nil && inp.LinkedDatasetID != nil && inp.LinkedDatasetField != nil {
		builder = builder.
			LinkedDatasetFieldID(inp.LinkedDatasetField).
			LinkedDatasetID(inp.LinkedDatasetID).
			LinkedDatasetSchemaID(inp.LinkedDatasetSchemaID)
	}
	item, err := builder.Build()
	if err != nil {
		return nil, err
	}

	err = i.tagRepo.Save(ctx, item)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return item, nil
}

func (i *Tag) CreateGroup(ctx context.Context, inp interfaces.CreateTagGroupParam, operator *usecase.Operator) (*tag.Group, error) {
	tx, err := i.transaction.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	if err := i.CanWriteScene(ctx, inp.SceneID, operator); err != nil {
		return nil, interfaces.ErrOperationDenied
	}

	list := tag.NewListFromTags(inp.Tags)
	group, err := tag.NewGroup().
		NewID().
		Label(inp.Label).
		Scene(inp.SceneID).
		Tags(list).
		Build()

	if err != nil {
		return nil, err
	}

	err = i.tagRepo.Save(ctx, group)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return group, nil
}

func (i *Tag) AttachItemToGroup(ctx context.Context, itemID id.TagID, groupID id.TagID, operator *usecase.Operator) (*tag.Group, error) {
	tx, err := i.transaction.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	scenes, err := i.OnlyWritableScenes(ctx, operator)
	if err != nil {
		return nil, err
	}
	// make sure item exist
	_, err = i.tagRepo.FindItemByID(ctx, itemID, scenes)
	if err != nil {
		return nil, err
	}

	tg, err := i.tagRepo.FindGroupByID(ctx, groupID, scenes)
	if err != nil {
		return nil, err
	}
	if !tg.Tags().Has(itemID) {
		tg.Tags().Add(itemID)
	} else {
		return nil, errors.New("tag item is already attached to the group")
	}
	err = i.tagRepo.Save(ctx, tg)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return tg, nil
}

func (i *Tag) DetachItemFromGroup(ctx context.Context, itemID id.TagID, groupID id.TagID, operator *usecase.Operator) (*tag.Group, error) {
	tx, err := i.transaction.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	scenes, err := i.OnlyWritableScenes(ctx, operator)
	if err != nil {
		return nil, err
	}
	// make sure item exist
	_, err = i.tagRepo.FindItemByID(ctx, itemID, scenes)
	if err != nil {
		return nil, err
	}

	tg, err := i.tagRepo.FindGroupByID(ctx, groupID, scenes)
	if err != nil {
		return nil, err
	}
	if tg.Tags().Has(itemID) {
		tg.Tags().Remove(itemID)
	} else {
		return nil, errors.New("tag item is not attached to the group")
	}

	err = i.tagRepo.Save(ctx, tg)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return tg, nil
}
