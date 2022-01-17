package gqlmodel

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
	"github.com/reearth/reearth-backend/pkg/layer/decoding"
	"github.com/stretchr/testify/assert"
)

func TestFromLayerEncodingFormat(t *testing.T) {
	type args struct {
		v LayerEncodingFormat
	}
	tests := []struct {
		name string
		args args
		want decoding.LayerEncodingFormat
	}{
		{
			name: "Kml",
			args: args{v: LayerEncodingFormatKml},
			want: decoding.LayerEncodingFormatKML,
		},
		{
			name: "Czml",
			args: args{v: LayerEncodingFormatCzml},
			want: decoding.LayerEncodingFormatCZML,
		},
		{
			name: "GeoJson",
			args: args{v: LayerEncodingFormatGeojson},
			want: decoding.LayerEncodingFormatGEOJSON,
		},
		{
			name: "Shape",
			args: args{v: LayerEncodingFormatShape},
			want: decoding.LayerEncodingFormatSHAPE,
		},
		{
			name: "Reearth",
			args: args{v: LayerEncodingFormatReearth},
			want: decoding.LayerEncodingFormatREEARTH,
		},
		{
			name: "Other",
			args: args{v: ""},
			want: decoding.LayerEncodingFormat(""),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, FromLayerEncodingFormat(tc.args.v), "FromLayerEncodingFormat(%v)", tc.args.v)
		})
	}
}

func TestToInfobox(t *testing.T) {
	type args struct {
		ib              *layer.Infobox
		parent          id.LayerID
		parentSceneID   id.SceneID
		parentDatasetID *id.DatasetID
	}
	tests := []struct {
		name string
		args args
		want *Infobox
	}{
		{
			name: "Nil IB",
			args: args{
				ib:              nil,
				parent:          id.LayerID{},
				parentSceneID:   id.SceneID{},
				parentDatasetID: nil,
			},
			want: nil,
		},
		{
			name: "Normal",
			args: args{
				ib:              &layer.Infobox{},
				parent:          id.LayerID{},
				parentSceneID:   id.SceneID{},
				parentDatasetID: nil,
			},
			want: &Infobox{
				Fields:          make([]*InfoboxField, 0),
				LinkedDatasetID: nil,
				LayerID:         id.ID{},
				Property:        nil,
				Scene:           nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToInfobox(tc.args.ib, tc.args.parent, tc.args.parentSceneID, tc.args.parentDatasetID), "ToInfobox(%v, %v, %v, %v)", tc.args.ib, tc.args.parent, tc.args.parentSceneID, tc.args.parentDatasetID)
		})
	}
}

func TestToInfoboxField(t *testing.T) {
	type args struct {
		ibf             *layer.InfoboxField
		parentSceneID   id.SceneID
		parentDatasetID *id.DatasetID
	}
	tests := []struct {
		name string
		args args
		want *InfoboxField
	}{
		{
			name: "IBF Nil",
			args: args{
				ibf:             nil,
				parentSceneID:   id.SceneID{},
				parentDatasetID: nil,
			},
			want: nil,
		},
		{
			name: "Normal",
			args: args{
				ibf:             &layer.InfoboxField{},
				parentSceneID:   id.SceneID{},
				parentDatasetID: nil,
			},
			want: &InfoboxField{
				ID:              id.ID{},
				SceneID:         id.ID{},
				PluginID:        id.PluginID{},
				ExtensionID:     "",
				PropertyID:      id.ID{},
				LinkedDatasetID: nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToInfoboxField(tc.args.ibf, tc.args.parentSceneID, tc.args.parentDatasetID), "ToInfoboxField(%v, %v, %v)", tc.args.ibf, tc.args.parentSceneID, tc.args.parentDatasetID)
		})
	}
}

func TestToLayer(t *testing.T) {
	lid := id.NewLayerID()
	type args struct {
		l      layer.Layer
		parent *id.LayerID
	}
	tests := []struct {
		name string
		args args
		want Layer
	}{
		{
			name: "L Nil",
			args: args{
				l:      nil,
				parent: nil,
			},
			want: nil,
		},
		{
			name: "Group",
			args: args{
				l:      layer.New().ID(lid).Group().MustBuild(),
				parent: nil,
			},
			want: &LayerGroup{
				ID:                    lid.ID(),
				SceneID:               id.ID{},
				Name:                  "",
				IsVisible:             true,
				PropertyID:            nil,
				PluginID:              nil,
				ExtensionID:           nil,
				Infobox:               nil,
				ParentID:              nil,
				LinkedDatasetSchemaID: nil,
				Root:                  false,
				LayerIds:              make([]*id.ID, 0),
				Tags:                  nil,
				Parent:                nil,
				Property:              nil,
				Plugin:                nil,
				Extension:             nil,
				LinkedDatasetSchema:   nil,
				Layers:                nil,
				Scene:                 nil,
				ScenePlugin:           nil,
			},
		},
		{
			name: "Item",
			args: args{
				l:      layer.New().ID(lid).Item().MustBuild(),
				parent: nil,
			},
			want: &LayerItem{
				ID:              lid.ID(),
				SceneID:         id.ID{},
				Name:            "",
				IsVisible:       true,
				PropertyID:      nil,
				PluginID:        nil,
				ExtensionID:     nil,
				Infobox:         nil,
				ParentID:        nil,
				LinkedDatasetID: nil,
				Tags:            nil,
				Parent:          nil,
				Property:        nil,
				Plugin:          nil,
				Extension:       nil,
				LinkedDataset:   nil,
				Merged:          nil,
				Scene:           nil,
				ScenePlugin:     nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToLayer(tc.args.l, tc.args.parent), "ToLayer(%v, %v)", tc.args.l, tc.args.parent)
		})
	}
}

func TestToLayerGroup(t *testing.T) {
	lid := id.NewLayerID()
	type args struct {
		l      *layer.Group
		parent *id.LayerID
	}
	tests := []struct {
		name string
		args args
		want *LayerGroup
	}{
		{
			name: "L Nil",
			args: args{
				l:      nil,
				parent: nil,
			},
			want: nil,
		},
		{
			name: "Normal",
			args: args{
				l:      layer.NewGroup().ID(lid).MustBuild(),
				parent: nil,
			},
			want: &LayerGroup{
				ID:                    lid.ID(),
				SceneID:               id.ID{},
				Name:                  "",
				IsVisible:             true,
				PropertyID:            nil,
				PluginID:              nil,
				ExtensionID:           nil,
				Infobox:               nil,
				ParentID:              nil,
				LinkedDatasetSchemaID: nil,
				Root:                  false,
				LayerIds:              make([]*id.ID, 0),
				Tags:                  nil,
				Parent:                nil,
				Property:              nil,
				Plugin:                nil,
				Extension:             nil,
				LinkedDatasetSchema:   nil,
				Layers:                nil,
				Scene:                 nil,
				ScenePlugin:           nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToLayerGroup(tc.args.l, tc.args.parent), "ToLayerGroup(%v, %v)", tc.args.l, tc.args.parent)
		})
	}
}

func TestToLayerItem(t *testing.T) {
	lid := id.NewLayerID()
	type args struct {
		l      *layer.Item
		parent *id.LayerID
	}
	tests := []struct {
		name string
		args args
		want *LayerItem
	}{
		{
			name: "L Nil",
			args: args{
				l:      nil,
				parent: nil,
			},
			want: nil,
		},
		{
			name: "Normal",
			args: args{
				l:      layer.NewItem().ID(lid).MustBuild(),
				parent: nil,
			},
			want: &LayerItem{
				ID:              lid.ID(),
				SceneID:         id.ID{},
				Name:            "",
				IsVisible:       true,
				PropertyID:      nil,
				PluginID:        nil,
				ExtensionID:     nil,
				Infobox:         nil,
				ParentID:        nil,
				LinkedDatasetID: nil,
				Tags:            nil,
				Parent:          nil,
				Property:        nil,
				Plugin:          nil,
				Extension:       nil,
				LinkedDataset:   nil,
				Merged:          nil,
				Scene:           nil,
				ScenePlugin:     nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToLayerItem(tc.args.l, tc.args.parent), "ToLayerItem(%v, %v)", tc.args.l, tc.args.parent)
		})
	}
}

func TestToLayerTag(t *testing.T) {
	tid := id.NewTagID()
	type args struct {
		l layer.Tag
	}
	tests := []struct {
		name string
		args args
		want LayerTag
	}{
		{
			name: "L nil",
			args: args{l: nil},
			want: nil,
		},
		{
			name: "Group",
			args: args{
				l: layer.NewTagGroup(tid, nil),
			},
			want: &LayerTagGroup{
				TagID:    tid.ID(),
				Children: make([]*LayerTagItem, 0),
				Tag:      nil,
			},
		},
		{
			name: "Item",
			args: args{
				l: layer.NewTagItem(tid),
			},
			want: &LayerTagItem{
				TagID: tid.ID(),
				Tag:   nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToLayerTag(tc.args.l), "ToLayerTag(%v)", tc.args.l)
		})
	}
}

func TestToLayerTagGroup(t *testing.T) {
	tid := id.NewTagID()
	type args struct {
		t *layer.TagGroup
	}
	tests := []struct {
		name string
		args args
		want *LayerTagGroup
	}{
		{
			name: "L nil",
			args: args{t: nil},
			want: nil,
		},
		{
			name: "Group",
			args: args{
				t: layer.NewTagGroup(tid, nil),
			},
			want: &LayerTagGroup{
				TagID:    tid.ID(),
				Children: make([]*LayerTagItem, 0),
				Tag:      nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToLayerTagGroup(tc.args.t), "ToLayerTagGroup(%v)", tc.args.t)
		})
	}
}

func TestToLayerTagItem(t *testing.T) {
	tid := id.NewTagID()
	type args struct {
		t *layer.TagItem
	}
	tests := []struct {
		name string
		args args
		want *LayerTagItem
	}{
		{
			name: "L nil",
			args: args{t: nil},
			want: nil,
		},
		{
			name: "Item",
			args: args{
				t: layer.NewTagItem(tid),
			},
			want: &LayerTagItem{
				TagID: tid.ID(),
				Tag:   nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToLayerTagItem(tc.args.t), "ToLayerTagItem(%v)", tc.args.t)
		})
	}
}

func TestToLayerTagList(t *testing.T) {
	tid := id.NewTagID()
	type args struct {
		t   *layer.TagList
		sid id.SceneID
	}
	tests := []struct {
		name string
		args args
		want []LayerTag
	}{
		{
			name: "t Nil",
			args: args{
				t:   nil,
				sid: id.SceneID{},
			},
			want: nil,
		},
		{
			name: "t empty",
			args: args{
				t:   &layer.TagList{},
				sid: id.SceneID{},
			},
			want: nil,
		},
		{
			name: "Normal",
			args: args{
				t: layer.NewTagList([]layer.Tag{
					nil,
					layer.NewTagItem(tid),
					layer.NewTagGroup(tid, nil),
				}),
				sid: id.SceneID{},
			},
			want: []LayerTag{
				&LayerTagItem{
					TagID: tid.ID(),
					Tag:   nil,
				},
				&LayerTagGroup{
					TagID:    tid.ID(),
					Children: make([]*LayerTagItem, 0),
					Tag:      nil,
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToLayerTagList(tc.args.t, tc.args.sid), "ToLayerTagList(%v, %v)", tc.args.t, tc.args.sid)
		})
	}
}

func TestToLayers(t *testing.T) {
	lid := id.NewLayerID()
	type args struct {
		layers layer.List
		parent *id.LayerID
	}
	tests := []struct {
		name string
		args args
		want []Layer
	}{
		{
			name: "Nil",
			args: args{
				layers: nil,
				parent: nil,
			},
			want: nil,
		},
		{
			name: "Empty",
			args: args{
				layers: []*layer.Layer{},
				parent: nil,
			},
			want: nil,
		},
		{
			name: "Normal",
			args: args{
				layers: []*layer.Layer{
					nil,
					layer.New().ID(lid).Item().MustBuild().LayerRef(),
					layer.New().ID(lid).Group().MustBuild().LayerRef(),
				},
				parent: nil,
			},
			want: []Layer{
				&LayerItem{
					ID:              lid.ID(),
					SceneID:         id.ID{},
					Name:            "",
					IsVisible:       true,
					PropertyID:      nil,
					PluginID:        nil,
					ExtensionID:     nil,
					Infobox:         nil,
					ParentID:        nil,
					LinkedDatasetID: nil,
					Tags:            nil,
					Parent:          nil,
					Property:        nil,
					Plugin:          nil,
					Extension:       nil,
					LinkedDataset:   nil,
					Merged:          nil,
					Scene:           nil,
					ScenePlugin:     nil,
				},
				&LayerGroup{
					ID:                    lid.ID(),
					SceneID:               id.ID{},
					Name:                  "",
					IsVisible:             true,
					PropertyID:            nil,
					PluginID:              nil,
					ExtensionID:           nil,
					Infobox:               nil,
					ParentID:              nil,
					LinkedDatasetSchemaID: nil,
					Root:                  false,
					LayerIds:              make([]*id.ID, 0),
					Tags:                  nil,
					Parent:                nil,
					Property:              nil,
					Plugin:                nil,
					Extension:             nil,
					LinkedDatasetSchema:   nil,
					Layers:                nil,
					Scene:                 nil,
					ScenePlugin:           nil,
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToLayers(tc.args.layers, tc.args.parent), "ToLayers(%v, %v)", tc.args.layers, tc.args.parent)
		})
	}
}

func TestToMergedInfobox(t *testing.T) {
	type args struct {
		ib      *layer.MergedInfobox
		sceneID id.SceneID
	}
	tests := []struct {
		name string
		args args
		want *MergedInfobox
	}{
		{
			name: "IB Nil",
			args: args{
				ib:      nil,
				sceneID: id.SceneID{},
			},
			want: nil,
		},
		{
			name: "Normal",
			args: args{
				ib:      layer.MergeInfobox(nil, nil, nil),
				sceneID: id.SceneID{},
			},
			want: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToMergedInfobox(tc.args.ib, tc.args.sceneID), "ToMergedInfobox(%v, %v)", tc.args.ib, tc.args.sceneID)
		})
	}
}

func TestToMergedInfoboxField(t *testing.T) {
	sid := id.NewSceneID()
	pid, _ := id.NewPluginID("test", "1.1", nil)
	ibfId := id.NewInfoboxFieldID()
	type args struct {
		ibf     *layer.MergedInfoboxField
		sceneID id.SceneID
	}
	tests := []struct {
		name string
		args args
		want *MergedInfoboxField
	}{
		{
			name: "MIB Nil",
			args: args{
				ibf:     nil,
				sceneID: id.SceneID{},
			},
			want: nil,
		},
		{
			name: "Normal",
			args: args{
				ibf: &layer.MergedInfoboxField{
					ID:        ibfId,
					Plugin:    pid,
					Extension: "",
					Property:  nil,
				},
				sceneID: sid,
			},
			want: &MergedInfoboxField{
				SceneID:     sid.ID(),
				OriginalID:  ibfId.ID(),
				PluginID:    pid,
				ExtensionID: "",
				Property:    nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToMergedInfoboxField(tc.args.ibf, tc.args.sceneID), "ToMergedInfoboxField(%v, %v)", tc.args.ibf, tc.args.sceneID)
		})
	}
}

func TestToMergedLayer(t *testing.T) {
	lid := id.NewLayerID()
	sid := id.NewSceneID()
	type args struct {
		layer *layer.Merged
	}
	tests := []struct {
		name string
		args args
		want *MergedLayer
	}{
		{
			name: "L Nil",
			args: args{nil},
			want: nil,
		},
		{
			name: "Normal",
			args: args{
				layer: &layer.Merged{
					Original:    lid,
					Parent:      nil,
					Name:        "",
					Scene:       sid,
					Property:    nil,
					Infobox:     nil,
					PluginID:    nil,
					ExtensionID: nil,
				}},
			want: &MergedLayer{
				OriginalID: lid.ID(),
				ParentID:   nil,
				SceneID:    sid.ID(),
				Property:   nil,
				Infobox:    nil,
				Original:   nil,
				Parent:     nil,
				Scene:      nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equalf(tt, tc.want, ToMergedLayer(tc.args.layer), "ToMergedLayer(%v)", tc.args.layer)
		})
	}
}
