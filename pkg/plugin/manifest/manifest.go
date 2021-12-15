package manifest

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/property"
)

type Manifest struct {
	Plugin          *plugin.Plugin
	ExtensionSchema property.SchemaList
	Schema          *property.Schema
}

func (m Manifest) PropertySchemas() property.SchemaList {
	if m.Schema == nil {
		return append([]*property.Schema{}, m.ExtensionSchema...)
	}
	return append(m.ExtensionSchema, m.Schema)
}

func (m Manifest) PropertySchema(psid id.PropertySchemaID) *property.Schema {
	if psid.IsNil() {
		return nil
	}
	if m.Schema != nil && psid.Equal(m.Schema.ID()) {
		return m.Schema
	}
	return m.ExtensionSchema.Find(psid)
}
