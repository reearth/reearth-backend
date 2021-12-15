package manifest

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/property"
)

type Diff struct {
	From                  plugin.ID
	To                    plugin.ID
	PropertySchemaDiff    property.SchemaDiff
	PropertySchemaDeleted bool
	DeletedExtensions     []DiffExtensionDeleted
	UpdatedExtensions     []DiffExtensionUpdated
}

type DiffExtensionUpdated struct {
	ExtensionID        plugin.ExtensionID
	OldType            plugin.ExtensionType
	NewType            plugin.ExtensionType
	PropertySchemaDiff property.SchemaDiff
}

type DiffExtensionDeleted struct {
	ExtensionID      plugin.ExtensionID
	PropertySchemaID id.PropertySchemaID
}

func DiffFrom(old, new Manifest) (d Diff) {
	d.From = old.Plugin.ID()
	d.To = new.Plugin.ID()

	oldsid, newsid := old.Plugin.Schema(), new.Plugin.Schema()
	if oldsid != nil && newsid == nil {
		d.PropertySchemaDiff.From = *oldsid
		d.PropertySchemaDeleted = true
	} else if oldsid != nil && newsid != nil {
		d.PropertySchemaDiff = property.SchemaDiffFrom(old.PropertySchema(*oldsid), old.PropertySchema(*newsid))
	}

	for _, e := range old.Plugin.Extensions() {
		ne := new.Plugin.Extension(e.ID())
		if ne == nil {
			d.DeletedExtensions = append(d.DeletedExtensions, DiffExtensionDeleted{
				ExtensionID:      e.ID(),
				PropertySchemaID: e.Schema(),
			})
			continue
		}

		oldps, newps := old.PropertySchema(e.Schema()), new.PropertySchema(ne.Schema())
		diff := DiffExtensionUpdated{
			ExtensionID:        e.ID(),
			OldType:            e.Type(),
			NewType:            ne.Type(),
			PropertySchemaDiff: property.SchemaDiffFrom(oldps, newps),
		}

		if diff.OldType != diff.NewType || !diff.PropertySchemaDiff.IsEmpty() {
			d.UpdatedExtensions = append(d.UpdatedExtensions, diff)
		}
	}

	return
}

func (d *Diff) IsEmpty() bool {
	return d == nil || len(d.DeletedExtensions) == 0 && len(d.UpdatedExtensions) == 0 && d.PropertySchemaDiff.IsEmpty() && !d.PropertySchemaDeleted
}
