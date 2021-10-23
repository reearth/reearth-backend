package gqlmodel

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
)

func ToTagItem(ti *tag.Item) *TagItem {
	if ti == nil {
		return nil
	}
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
	if tg == nil {
		return nil
	}
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

func ToTag(t *tag.Tag) Tag {
	if t == nil {
		return nil
	}
	tt := *t
	switch ta := tt.(type) {
	case *tag.Item:
		return ToTagItem(ta)
	case *tag.Group:
		return ToTagGroup(ta)
	}
	return nil
}
