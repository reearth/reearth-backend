package gqlmodel

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
)

func ToTagItem(ti *tag.Item) *TagItem {
	return &TagItem{
		ID:                    ti.ID().ID(),
		SceneID:               ti.Scene().ID(),
		Label:                 ti.Label(),
		LinkedDatasetID:       ti.LinkedDatasetID().IDRef(),
		LinkedDatasetSchemaID: ti.LinkedDatasetSchemaID().IDRef(),
		LinkedDatasetFieldID:  ti.LinkedDatasetFieldID().IDRef(),
	}
}

func ToTagGroup(tg *tag.Group) *TagGroup {
	tags := tg.Tags()
	tids := tags.Tags()
	var ids []*id.ID
	for _, tid := range tids {
		if !tid.IsNil() {
			ids = append(ids, tid.IDRef())
		}
	}
	return &TagGroup{
		ID:      tg.ID().ID(),
		SceneID: tg.Scene().ID(),
		Label:   tg.Label(),
		Tags:    ids,
	}
}

func ToTag(t tag.Tag) Tag {
	if t == nil {
		return nil
	}
	switch ty := t.(type) {
	case *tag.Item:
		return ToTagItem(ty)
	case *tag.Group:
		return ToTagGroup(ty)
	}
	return nil
}
