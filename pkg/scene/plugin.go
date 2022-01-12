package scene

import (
	"github.com/reearth/reearth-backend/pkg/id"
)

type Plugin struct {
	plugin   id.PluginID
	property *id.PropertyID
}

func NewPlugin(plugin id.PluginID, property *id.PropertyID) *Plugin {
	return &Plugin{
		plugin:   plugin,
		property: property.CopyRef(),
	}
}

func (s *Plugin) Plugin() id.PluginID {
	if s == nil {
		return id.PluginID{}
	}
	return s.plugin
}

func (s *Plugin) PluginRef() *id.PluginID {
	if s == nil {
		return nil
	}
	return s.plugin.Ref()
}

func (s *Plugin) Property() *id.PropertyID {
	if s == nil {
		return nil
	}
	return s.property.CopyRef()
}

func (s *Plugin) Clone() *Plugin {
	if s == nil {
		return nil
	}
	return &Plugin{
		plugin:   s.plugin,
		property: s.property.CopyRef(),
	}
}
