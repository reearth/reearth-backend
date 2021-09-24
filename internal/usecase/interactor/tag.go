package interactor

import (
	"context"
	"errors"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
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
		if len(tags.Tags()) == 0 {
			return nil, errors.New("can't delete non-empty tag group")
		}
	}

	if {

	}

	l,err:=i.layerRepo.FindByTag(ctx,tagID)
	if l!=nil {
		l.
	}
}
