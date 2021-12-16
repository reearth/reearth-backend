package layer

import (
	"github.com/reearth/reearth-backend/pkg/id"
)

var _ Layer = &Item{}

var tags2 = []id.TagID{id.NewTagID()}
var item = Item{
	layerBase: layerBase{
		id:        id.MustLayerID(id.New().String()),
		name:      "xxx",
		visible:   false,
		plugin:    id.MustPluginID("aaa~1.1.1").Ref(),
		extension: id.PluginExtensionID("foo").Ref(),
		property:  nil,
		infobox:   nil,
		tags:      nil,
		scene:     id.SceneID{},
	},
	linkedDataset: nil,
}
