package interactor

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/reearth/reearth-backend/pkg/tag"
)

type Tag struct {
	commonScene
	tagRepo     repo.Tag
	layerRepo   repo.Layer
	sceneRepo   repo.Scene
	transaction repo.Transaction
}

func NewTag(r *repo.Container) interfaces.Tag {
	return &Tag{
		commonScene: commonScene{sceneRepo: r.Scene},
		tagRepo:     r.Tag,
		layerRepo:   r.Layer,
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

func (i *Tag) AttachItemToGroup(ctx context.Context, inp interfaces.AttachItemToGroupParam, operator *usecase.Operator) (*tag.Group, error) {
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
	_, err = i.tagRepo.FindItemByID(ctx, inp.ItemID, scenes)
	if err != nil {
		return nil, err
	}

	tg, err := i.tagRepo.FindGroupByID(ctx, inp.GroupID, scenes)
	if err != nil {
		return nil, err
	}
	if !tg.Tags().Has(inp.ItemID) {
		tg.Tags().Add(inp.ItemID)
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

func (i *Tag) DetachItemFromGroup(ctx context.Context, inp interfaces.DetachItemToGroupParam, operator *usecase.Operator) (*tag.Group, error) {
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
	_, err = i.tagRepo.FindItemByID(ctx, inp.ItemID, scenes)
	if err != nil {
		return nil, err
	}

	tg, err := i.tagRepo.FindGroupByID(ctx, inp.GroupID, scenes)
	if err != nil {
		return nil, err
	}
	if tg.Tags().Has(inp.ItemID) {
		tg.Tags().Remove(inp.ItemID)
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

func (i *Tag) Remove(ctx context.Context, tagID id.TagID, operator *usecase.Operator) (*id.TagID, error) {
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

	t, err := i.tagRepo.FindByID(ctx, tagID, scenes)
	if err != nil {
		return nil, err
	}

	if group := tag.ToTagGroup(*t); group != nil {
		tags := group.Tags()
		if len(tags.Tags()) != 0 {
			return nil, interfaces.ErrNonemptyTagGroupCannotDelete
		}
	}

	if item := tag.ToTagItem(*t); item != nil {
		g, err := i.tagRepo.FindGroupByItem(ctx, item.ID(), scenes)
		if err != nil && !errors.Is(rerror.ErrNotFound, err) {
			return nil, err
		}
		if g != nil {
			g.Tags().Remove(item.ID())

			err = i.tagRepo.Save(ctx, g)
			if err != nil {
				return nil, err
			}
		}
	}

	l, err := i.layerRepo.FindByTag(ctx, tagID, scenes)
	if err != nil && !errors.Is(rerror.ErrNotFound, err) {
		return nil, err
	}

	if l != nil {
		err = l.DetachTag(tagID)
		if err != nil {
			return nil, err
		}
		err = i.layerRepo.Save(ctx, l)
		if err != nil {
			return nil, err
		}
	}

	err = i.tagRepo.Remove(ctx, tagID)
	if err != nil {
		return nil, err
	}
	return &tagID, nil
}
