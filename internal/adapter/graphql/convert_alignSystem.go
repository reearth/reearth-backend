package graphql

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

func toAlignSystem(sas *scene.WidgetAlignSystem) *WidgetAlignSystem {
	widgetAlignDoc := WidgetAlignSystem{
		Inner: toWidgetZone(sas.Zone(scene.WidgetZoneInner)),
		Outer: toWidgetZone(sas.Zone(scene.WidgetZoneOuter)),
	}
	return &widgetAlignDoc
}

func IDsFrom(wids []id.WidgetID) []*id.ID {
	var nids []*id.ID
	for _, w := range wids {
		nids = append(nids, w.IDRef())
	}
	return nids
}

func toWidgetZone(z *scene.WidgetZone) *WidgetZone {
	if z == nil {
		return nil
	}
	return &WidgetZone{
		Left:   toWidgetSection(z.Section(scene.WidgetSectionLeft)),
		Center: toWidgetSection(z.Section(scene.WidgetSectionCenter)),
		Right:  toWidgetSection(z.Section(scene.WidgetSectionRight)),
	}
}

func toWidgetSection(s *scene.WidgetSection) *WidgetSection {
	if s == nil {
		return nil
	}
	return &WidgetSection{
		Top:    toWidgetArea(s.Area(scene.WidgetAreaTop)),
		Middle: toWidgetArea(s.Area(scene.WidgetAreaMiddle)),
		Bottom: toWidgetArea(s.Area(scene.WidgetAreaBottom)),
	}
}

func toWidgetArea(a *scene.WidgetArea) *WidgetArea {
	if a == nil {
		return nil
	}
	return &WidgetArea{
		WidgetIds: IDsFrom(a.WidgetIDs()),
		Align:     a.Alignment(),
	}
}
