package plugin

import (
	"errors"

	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
	"github.com/reearth/reearth-backend/pkg/visualizer"
)

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

type Extension struct {
	id            id.PluginExtensionID
	extensionType ExtensionType
	name          i18n.String
	description   i18n.String
	icon          string
	schema        id.PropertySchemaID
	visualizer    visualizer.Visualizer
	widgetLayout  *scene.WidgetLayout
}

func (w *Extension) ID() id.PluginExtensionID {
	return w.id
}

func (w *Extension) Type() ExtensionType {
	return w.extensionType
}

func (w *Extension) Name() i18n.String {
	return w.name.Copy()
}

func (w *Extension) Description() i18n.String {
	return w.description.Copy()
}

func (w *Extension) Icon() string {
	return w.icon
}

func (w *Extension) Schema() id.PropertySchemaID {
	return w.schema
}

func (w *Extension) Visualizer() visualizer.Visualizer {
	return w.visualizer
}

func (w *Extension) Layout() *scene.WidgetLayout {
	return w.widgetLayout
}

func (w *Extension) Rename(name i18n.String) {
	w.name = name.Copy()

}

func (w *Extension) SetDescription(des i18n.String) {
	w.description = des.Copy()
}
