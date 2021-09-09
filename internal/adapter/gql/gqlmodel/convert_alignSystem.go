package gqlmodel

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

func ToAlignSystem(sas *scene.WidgetAlignSystem) *WidgetAlignSystem {
	widgetAlignDoc := WidgetAlignSystem{
		Inner: ToWidgetZone(sas.Zone(string(scene.WidgetZoneInner))),
		Outer: ToWidgetZone(sas.Zone(string(scene.WidgetZoneOuter))),
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

func ToWidgetZone(z *scene.WidgetZone) *WidgetZone {
	if z == nil {
		return nil
	}
	return &WidgetZone{
		Left:   ToWidgetSection(z.Section(string(scene.WidgetSectionLeft))),
		Center: ToWidgetSection(z.Section(string(scene.WidgetSectionCenter))),
		Right:  ToWidgetSection(z.Section(string(scene.WidgetSectionRight))),
	}
}

func ToWidgetSection(s *scene.WidgetSection) *WidgetSection {
	if s == nil {
		return nil
	}
	return &WidgetSection{
		Top:    ToWidgetArea(s.Area(string(scene.WidgetAreaTop))),
		Middle: ToWidgetArea(s.Area(string(scene.WidgetAreaMiddle))),
		Bottom: ToWidgetArea(s.Area(string(scene.WidgetAreaBottom))),
	}
}

func ToWidgetArea(a *scene.WidgetArea) *WidgetArea {
	if a == nil {
		return nil
	}
	return &WidgetArea{
		WidgetIds: IDsFrom(a.WidgetIDs()),
		Align:     a.Alignment(),
	}
}
