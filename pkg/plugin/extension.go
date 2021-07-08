package plugin

import (
	"errors"

	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
	"github.com/reearth/reearth-backend/pkg/visualizer"
)

// ExtensionType _
type ExtensionType string

var (
	// ErrPluginExtensionDuplicated _
	ErrPluginExtensionDuplicated error = errors.New("plugin extension duplicated")
	// ExtensionTypePrimitive _
	ExtensionTypePrimitive ExtensionType = "primitive"
	// ExtensionTypeWidget _
	ExtensionTypeWidget ExtensionType = "widget"
	// ExtensionTypeBlock _
	ExtensionTypeBlock ExtensionType = "block"
	// ExtensionTypeVisualizer _
	ExtensionTypeVisualizer ExtensionType = "visualizer"
	// ExtensionTypeInfobox _
	ExtensionTypeInfobox ExtensionType = "infobox"
)

// WidgetLayout _
type WidgetLayout struct {
	Extendable      *bool
	Extended        bool
	DefaultLocation *scene.Location
}

// Extension _
type Extension struct {
	id            id.PluginExtensionID
	extensionType ExtensionType
	name          i18n.String
	description   i18n.String
	icon          string
	schema        id.PropertySchemaID
	visualizer    visualizer.Visualizer
	widgetLayout  WidgetLayout
}

// ID _
func (w *Extension) ID() id.PluginExtensionID {
	return w.id
}

// Type _
func (w *Extension) Type() ExtensionType {
	return w.extensionType
}

// Name _
func (w *Extension) Name() i18n.String {
	return w.name.Copy()
}

// Description _
func (w *Extension) Description() i18n.String {
	return w.description.Copy()
}

// Icon _
func (w *Extension) Icon() string {
	return w.icon
}

// Schema _
func (w *Extension) Schema() id.PropertySchemaID {
	return w.schema
}

// Visualizer _
func (w *Extension) Visualizer() visualizer.Visualizer {
	return w.visualizer
}

// WidgetMetaData _
func (w *Extension) WidgetLayout() WidgetLayout {
	return w.widgetLayout
}

// Rename _
func (w *Extension) Rename(name i18n.String) {
	w.name = name.Copy()

}

// SetDescription _
func (w *Extension) SetDescription(des i18n.String) {
	w.description = des.Copy()
}
