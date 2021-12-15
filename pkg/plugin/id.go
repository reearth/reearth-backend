package plugin

import "github.com/reearth/reearth-backend/pkg/id"

type ID = id.PluginID
type ExtensionID = id.PluginExtensionID

var NewID = id.NewPluginID
var MustID = id.MustPluginID
var IDFrom = id.PluginIDFrom
var IDFromRef = id.PluginIDFromRef
var ExtensionIDFromRef = id.PluginExtensionIDFromRef
