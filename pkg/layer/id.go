package layer

import "github.com/reearth/reearth-backend/pkg/id"

type ID = id.LayerID
type IDSet = id.LayerIDSet
type InfoboxFieldID = id.InfoboxFieldID
type InfoboxFIeldIDSet = id.InfoboxFieldIDSet
type SceneID = id.SceneID
type PluginID = id.PluginID
type PluginExtensionID = id.PluginExtensionID
type PropertyID = id.PropertyID
type PropertySchemaID = id.PropertySchemaID

var NewID = id.NewLayerID
var MustID = id.MustLayerID
var IDFrom = id.LayerIDFrom
var IDFromRef = id.LayerIDFromRef
var IDFromRefID = id.LayerIDFromRefID
var NewIDSet = id.NewLayerIDSet

var NewInfoboxFieldID = id.NewInfoboxFieldID
var MustInfoboxFieldID = id.MustInfoboxFieldID
var InfoboxFieldIDFrom = id.InfoboxFieldIDFrom
var InfoboxFieldIDFromRef = id.InfoboxFieldIDFromRef
var InfoboxFieldIDFromRefID = id.InfoboxFieldIDFromRefID
var NewInfoboxFIeldIDSet = id.NewInfoboxFieldIDSet
