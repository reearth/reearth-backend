package merging

import (
	"github.com/reearth/reearth-backend/pkg/layer"
	"github.com/reearth/reearth-backend/pkg/property"
)

var (
	_ SealedLayer = &SealedLayerGroup{} // must implement SealedLayer
	_ SealedLayer = &SealedLayerItem{}  // must implement SealedLayer
)

type SealedLayer interface {
	Common() *SealedLayerCommon
	Flatten() []*SealedLayerItem
}

type SealedLayerGroup struct {
	SealedLayerCommon
	Children []SealedLayer
}

type SealedLayerItem struct {
	SealedLayerCommon
}

type SealedLayerCommon struct {
	layer.Merged
	Property *property.Sealed
	Infobox  *SealedInfobox
}

type SealedInfobox struct {
	layer.MergedInfobox
	Property *property.Sealed
	Fields   []*SealedInfoboxField
}

type SealedInfoboxField struct {
	layer.MergedInfoboxField
	Property *property.Sealed
}

func (l *SealedLayerGroup) Common() *SealedLayerCommon {
	if l == nil {
		return nil
	}
	return &l.SealedLayerCommon
}

func (l *SealedLayerGroup) Flatten() []*SealedLayerItem {
	if l == nil {
		return nil
	}
	layers := []*SealedLayerItem{}
	for _, c := range l.Children {
		layers = append(layers, c.Flatten()...)
	}
	return layers
}

func (l *SealedLayerItem) Common() *SealedLayerCommon {
	if l == nil {
		return nil
	}
	return &l.SealedLayerCommon
}

func (l *SealedLayerItem) Flatten() []*SealedLayerItem {
	if l == nil {
		return nil
	}
	return []*SealedLayerItem{l}
}