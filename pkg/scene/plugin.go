package scene

import (
	"github.com/reearth/reearth-backend/pkg/id"
)

type Plugin struct {
	plugin   id.PluginID
	property *id.PropertyID
}

func NewPlugin(plugin id.PluginID, property *id.PropertyID) *Plugin {
	if property != nil {
		property2 := *property
		property = &property2
	}
	return &Plugin{
		plugin:   plugin,
		property: property,
	}
}

func (s Plugin) Plugin() id.PluginID {
	return s.plugin
}

func (s Plugin) Property() *id.PropertyID {
	property := s.property
	if property != nil {
		property2 := *property
		property = &property2
	}
	return property
}
