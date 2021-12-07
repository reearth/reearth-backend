package property

import (
	"context"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
)

type Item interface {
	ID() id.PropertyItemID
	SchemaGroup() id.PropertySchemaGroupID
	Schema() id.PropertySchemaID
	FieldsByLinkedDataset(id.DatasetSchemaID, id.DatasetID) []*Field
	HasLinkedField() bool
	Datasets() []id.DatasetID
	IsDatasetLinked(id.DatasetSchemaID, id.DatasetID) bool
	IsEmpty() bool
	Prune()
	MigrateSchema(context.Context, *Schema, dataset.Loader)
	ValidateSchema(*SchemaGroup) error
}

type itemBase struct {
	ID          id.PropertyItemID
	Schema      id.PropertySchemaID
	SchemaGroup id.PropertySchemaGroupID
}

func ToGroup(i Item) *Group {
	g, _ := i.(*Group)
	return g
}

func ToGroupList(i Item) *GroupList {
	g, _ := i.(*GroupList)
	return g
}

func InitItemFrom(psg *SchemaGroup) Item {
	if psg == nil {
		return nil
	}
	if psg.IsList() {
		return InitGroupListFrom(psg)
	}
	return InitGroupFrom(psg)
}
