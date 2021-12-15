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

func (i *Tag) CreateItem(ctx context.Context, inp interfaces.CreateTagItemParam, operator *usecase.Operator) (*tag.Item, *tag.Group, error) {
	tx, err := i.transaction.Begin()
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	if err := i.CanWriteScene(ctx, inp.SceneID, operator); err != nil {
		return nil, nil, interfaces.ErrOperationDenied
	}

	var parent *tag.Group
	if inp.Parent != nil {
		parent, err = i.tagRepo.FindGroupByID(ctx, *inp.Parent, []id.SceneID{inp.SceneID})
		if err != nil {
			return nil, nil, err
		}
	}

	builder := tag.NewItem().
		NewID().
		Label(inp.Label).
		Scene(inp.SceneID).
		Parent(inp.Parent)
	if inp.LinkedDatasetSchemaID != nil && inp.LinkedDatasetID != nil && inp.LinkedDatasetField != nil {
		builder = builder.
			LinkedDatasetFieldID(inp.LinkedDatasetField).
			LinkedDatasetID(inp.LinkedDatasetID).
			LinkedDatasetSchemaID(inp.LinkedDatasetSchemaID)
	}
	item, err := builder.Build()
	if err != nil {
		return nil, nil, err
	}

	if parent != nil {
		parent.Tags().Add(item.ID())
	}

	itemt := tag.Tag(item)
	tags := []*tag.Tag{&itemt}
	if parent != nil {
		parentt := tag.Tag(parent)
		tags = append(tags, &parentt)
	}
	if err := i.tagRepo.SaveAll(ctx, tags); err != nil {
		return nil, nil, err
	}

	tx.Commit()
	return item, parent, nil
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

func (i *Tag) Fetch(ctx context.Context, ids []id.TagID, operator *usecase.Operator) ([]*tag.Tag, error) {
	scenes, err := i.OnlyReadableScenes(ctx, operator)
	if err != nil {
		return nil, err
	}

	return i.tagRepo.FindByIDs(ctx, ids, scenes)
}

func (i *Tag) FetchByScene(ctx context.Context, sid id.SceneID, operator *usecase.Operator) ([]*tag.Tag, error) {
	err := i.CanReadScene(ctx, sid, operator)
	if err != nil {
		return nil, err
	}

	return i.tagRepo.FindByScene(ctx, sid)
}

func (i *Tag) FetchItem(ctx context.Context, ids []id.TagID, operator *usecase.Operator) ([]*tag.Item, error) {
	scenes, err := i.OnlyReadableScenes(ctx, operator)
	if err != nil {
		return nil, err
	}

	return i.tagRepo.FindItemByIDs(ctx, ids, scenes)
}

func (i *Tag) FetchGroup(ctx context.Context, ids []id.TagID, operator *usecase.Operator) ([]*tag.Group, error) {
	scenes, err := i.OnlyReadableScenes(ctx, operator)
	if err != nil {
		return nil, err
	}

	return i.tagRepo.FindGroupByIDs(ctx, ids, scenes)
}

func (i *Tag) FetchGroupsByLayer(ctx context.Context, lid id.LayerID, operator *usecase.Operator) ([]*tag.Group, error) {
	scenes, err := i.OnlyReadableScenes(ctx, operator)
	if err != nil {
		return nil, err
	}

	layer, err := i.layerRepo.FindByID(ctx, lid, scenes)
	if err != nil {
		return nil, err
	}

	return i.tagRepo.FindGroupByIDs(ctx, layer.Tags().Tags(), scenes)
}

func (i *Tag) FetchGroupsByScene(ctx context.Context, sid id.SceneID, operator *usecase.Operator) ([]*tag.Group, error) {
	err := i.CanReadScene(ctx, sid, operator)
	if err != nil {
		return nil, err
	}

	return i.tagRepo.FindGroupByScene(ctx, sid)
}

func (i *Tag) FetchItemsByLayer(ctx context.Context, lid id.LayerID, operator *usecase.Operator) ([]*tag.Item, error) {
	scenes, err := i.OnlyReadableScenes(ctx, operator)
	if err != nil {
		return nil, err
	}

	layer, err := i.layerRepo.FindByID(ctx, lid, scenes)
	if err != nil {
		return nil, err
	}
	return i.tagRepo.FindItemByIDs(ctx, layer.Tags().Tags(), scenes)
}

func (i *Tag) FetchItemsByScene(ctx context.Context, sid id.SceneID, operator *usecase.Operator) ([]*tag.Item, error) {
	err := i.CanReadScene(ctx, sid, operator)
	if err != nil {
		return nil, err
	}

	return i.tagRepo.FindItemByScene(ctx, sid)
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
	ti, err := i.tagRepo.FindItemByID(ctx, inp.ItemID, scenes)
	if err == nil {
		return nil, err
	}
	if ti.Parent() != nil {
		return nil, errors.New("tag is already added to the group")
	}

	tg, err := i.tagRepo.FindGroupByID(ctx, inp.GroupID, scenes)
	if err != nil {
		return nil, err
	}

	if tg.Tags().Has(inp.ItemID) {
		return nil, errors.New("tag item is already attached to the group")
	}

	tg.Tags().Add(inp.ItemID)
	ti.SetParent(tg.ID().Ref())

	tgt := tag.Tag(tg)
	tit := tag.Tag(ti)
	if err := i.tagRepo.SaveAll(ctx, []*tag.Tag{&tgt, &tit}); err != nil {
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
	ti, err := i.tagRepo.FindItemByID(ctx, inp.ItemID, scenes)
	if err != nil {
		return nil, err
	}

	tg, err := i.tagRepo.FindGroupByID(ctx, inp.GroupID, scenes)
	if err != nil {
		return nil, err
	}

	if !tg.Tags().Has(inp.ItemID) {
		return nil, errors.New("tag item is not attached to the group")
	}

	tg.Tags().Remove(inp.ItemID)
	ti.SetParent(nil)

	tgt := tag.Tag(tg)
	tit := tag.Tag(ti)
	if err := i.tagRepo.SaveAll(ctx, []*tag.Tag{&tgt, &tit}); err != nil {
		return nil, err
	}

	tx.Commit()
	return tg, nil
}

func (i *Tag) UpdateTag(ctx context.Context, inp interfaces.UpdateTagParam, operator *usecase.Operator) (*tag.Tag, error) {
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

	tg, err := i.tagRepo.FindByID(ctx, inp.TagID, []id.SceneID{inp.SceneID})
	if err != nil {
		return nil, err
	}

	if inp.Label != nil {
		tg.Rename(*inp.Label)
	}

	err = i.tagRepo.Save(ctx, tg)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return &tg, nil
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

	if group := tag.ToTagGroup(t); group != nil {
		if len(group.Tags().Tags()) != 0 {
			return nil, interfaces.ErrNonemptyTagGroupCannotDelete
		}
	}

	if item := tag.ToTagItem(t); item != nil {
		g, err := i.tagRepo.FindGroupByItem(ctx, item.ID(), scenes)
		if err != nil && !errors.Is(rerror.ErrNotFound, err) {
			return nil, err
		}
		if g != nil {
			g.Tags().Remove(item.ID())
			if err := i.tagRepo.Save(ctx, g); err != nil {
				return nil, err
			}
		}
	}

	ls, err := i.layerRepo.FindByTag(ctx, tagID, scenes)
	if err != nil && !errors.Is(rerror.ErrNotFound, err) {
		return nil, err
	}

	if len(ls) != 0 {
		for _, l := range ls.Deref() {
			if err := l.DetachTag(tagID); err != nil {
				return nil, err
			}
		}
		if err := i.layerRepo.SaveAll(ctx, ls); err != nil {
			return nil, err
		}
	}

	if err := i.tagRepo.Remove(ctx, tagID); err != nil {
		return nil, err
	}
	return &tagID, nil
}
