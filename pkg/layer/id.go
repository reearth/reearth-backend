package layer

import (
	"sort"

	"github.com/reearth/reearth-backend/pkg/id"
)

type ID = id.LayerID
type InfoboxFieldID = id.InfoboxFieldID
type TagID = id.TagID
type SceneID = id.SceneID
type PluginID = id.PluginID
type PluginExtensionID = id.PluginExtensionID
type PropertyID = id.PropertyID
type DatasetID = id.DatasetID
type DatasetSchemaID = id.DatasetSchemaID

var NewID = id.NewLayerID
var NewInfoboxFieldID = id.NewInfoboxFieldID
var NewTagID = id.NewTagID
var NewSceneID = id.NewSceneID
var NewPropertyID = id.NewPropertyID
var NewDatasetID = id.NewDatasetID
var NewDatasetSchemaID = id.NewDatasetSchemaID

var MustID = id.MustLayerID
var MustInfoboxFieldID = id.MustInfoboxFieldID
var MustTagID = id.MustTagID
var MustSceneID = id.MustSceneID
var MustPluginID = id.MustPluginID
var MustPropertyID = id.MustPropertyID
var PropertySchemaIDFromExtension = id.PropertySchemaIDFromExtension
var MustPropertySchemaIDFromExtension = id.MustPropertySchemaIDFromExtension

var IDFrom = id.LayerIDFrom
var InfoboxFieldIDFrom = id.InfoboxFieldIDFrom
var TagIDFrom = id.TagIDFrom
var SceneIDFrom = id.SceneIDFrom
var PropertyIDFrom = id.PropertyIDFrom
var DatasetIDFrom = id.DatasetIDFrom
var DatasetSchemaIDFrom = id.DatasetSchemaIDFrom

var IDFromRef = id.LayerIDFromRef
var InfoboxFieldIDFromRef = id.InfoboxFieldIDFromRef
var TagIDFromRef = id.TagIDFromRef
var SceneIDFromRef = id.SceneIDFromRef
var PropertyIDFromRef = id.PropertyIDFromRef
var DatasetIDFromRef = id.DatasetIDFromRef
var DatasetSchemaIDFromRef = id.DatasetSchemaIDFromRef

var IDFromRefID = id.LayerIDFromRefID
var InfoboxFieldIDFromRefID = id.InfoboxFieldIDFromRefID
var TagIDFromRefID = id.TagIDFromRefID
var SceneIDFromRefID = id.SceneIDFromRefID
var PropertyIDFromRefID = id.PropertyIDFromRefID
var DatasetIDFromRefID = id.DatasetIDFromRefID
var DatasetSchemaIDFromRefID = id.DatasetSchemaIDFromRefID

type IDSet = id.LayerIDSet
type InfoboxFIeldIDSet = id.InfoboxFieldIDSet
type DatasetIDSet = id.DatasetIDSet

var NewIDSet = id.NewLayerIDSet
var NewInfoboxFIeldIDSet = id.NewInfoboxFieldIDSet
var NewDatasetIDSet = id.NewDatasetIDSet

var OfficialPluginID = id.OfficialPluginID
var ErrInvalidID = id.ErrInvalidID

func sortIDs(a []ID) {
	sort.SliceStable(a, func(i, j int) bool {
		return id.ID(a[i]).Compare(id.ID(a[j])) < 0
	})
}
