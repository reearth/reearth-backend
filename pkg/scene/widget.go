package scene

import (
	"github.com/reearth/reearth-backend/pkg/id"
)

type Extendable struct {
	Horizontally *bool
	Vertically   *bool
}

type WidgetLayout struct {
	Extendable      *Extendable
	Extended        *bool
	Floating        bool
	DefaultLocation *WidgetLocation
}

type Widget struct {
	id           id.WidgetID
	plugin       id.PluginID
	extension    id.PluginExtensionID
	property     id.PropertyID
	enabled      bool
	widgetLayout *WidgetLayout
}

func NewWidget(wid id.WidgetID, plugin id.PluginID, extension id.PluginExtensionID, property id.PropertyID, enabled bool, widgetLayout *WidgetLayout) (*Widget, error) {
	if !plugin.Validate() || string(extension) == "" || id.ID(property).IsNil() {
		return nil, id.ErrInvalidID
	}

	if widgetLayout == nil {
		widgetLayout = &WidgetLayout{}
	}

	return &Widget{
		id:           wid,
		plugin:       plugin,
		extension:    extension,
		property:     property,
		enabled:      enabled,
		widgetLayout: widgetLayout,
	}, nil
}

func MustNewWidget(wid id.WidgetID, plugin id.PluginID, extension id.PluginExtensionID, property id.PropertyID, enabled bool, widgetLayout *WidgetLayout) *Widget {
	w, err := NewWidget(wid, plugin, extension, property, enabled, widgetLayout)
	if err != nil {
		panic(err)
	}
	return w
}

func (w *Widget) ID() id.WidgetID {
	return w.id
}

func (w *Widget) Plugin() id.PluginID {
	return w.plugin
}

func (w *Widget) Extension() id.PluginExtensionID {
	return w.extension
}

func (w *Widget) Property() id.PropertyID {
	return w.property
}

func (w *Widget) Enabled() bool {
	return w.enabled
}

func (w *Widget) WidgetLayout() *WidgetLayout {
	return w.widgetLayout
}

func (w *Widget) SetEnabled(enabled bool) {
	w.enabled = enabled
}

func (w *Widget) SetExtended(extended *bool) {
	w.widgetLayout.Extended = extended
}
