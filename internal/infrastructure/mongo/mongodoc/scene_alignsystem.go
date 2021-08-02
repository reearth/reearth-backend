package mongodoc

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

type WidgetLocation struct {
	Zone    string
	Section string
	Area    string
}

type WidgetLayout struct {
	Extendable      bool
	Extended        bool
	Floating        bool
	DefaultLocation *WidgetLocation
}

type WidgetArea struct {
	WidgetIds []string
	Align     string
}

type WidgetSection struct {
	Top    WidgetArea
	Middle WidgetArea
	Bottom WidgetArea
}

type WidgetZone struct {
	Left   WidgetSection
	Center WidgetSection
	Right  WidgetSection
}

type SceneAlignSystemDocument struct {
	Inner WidgetZone
	Outer WidgetZone
}

func NewWidgetAlignSystem(was *scene.WidgetAlignSystem) *SceneAlignSystemDocument {
	widgetAlignDoc := SceneAlignSystemDocument{Inner: buildWidgetZone(was, "inner"), Outer: buildWidgetZone(was, "outer")}
	return &widgetAlignDoc
}

func (*SceneAlignSystemDocument) ToModelAlignSystem(d SceneAlignSystemDocument) *scene.WidgetAlignSystem {
	was := scene.NewWidgetAlignSystem()
	// Inner Left
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Inner.Left.Top.WidgetIds), d.Inner.Left.Top.Align, "inner", "left", "top")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Inner.Left.Middle.WidgetIds), d.Inner.Left.Middle.Align, "inner", "left", "middle")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Inner.Left.Bottom.WidgetIds), d.Inner.Left.Bottom.Align, "inner", "left", "bottom")
	// Inner Center
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Inner.Center.Top.WidgetIds), d.Inner.Center.Top.Align, "inner", "center", "top")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Inner.Center.Middle.WidgetIds), d.Inner.Center.Middle.Align, "inner", "center", "middle")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Inner.Center.Bottom.WidgetIds), d.Inner.Center.Bottom.Align, "inner", "center", "bottom")
	// Inner Right
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Inner.Right.Top.WidgetIds), d.Inner.Right.Top.Align, "inner", "right", "top")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Inner.Right.Middle.WidgetIds), d.Inner.Right.Middle.Align, "inner", "right", "middle")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Inner.Right.Bottom.WidgetIds), d.Inner.Right.Bottom.Align, "inner", "right", "bottom")
	// Outer Left
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Outer.Left.Top.WidgetIds), d.Outer.Left.Top.Align, "outer", "left", "top")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Outer.Left.Middle.WidgetIds), d.Outer.Left.Middle.Align, "outer", "left", "middle")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Outer.Left.Bottom.WidgetIds), d.Outer.Left.Bottom.Align, "outer", "left", "bottom")
	// Outer Center
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Outer.Center.Top.WidgetIds), d.Outer.Center.Top.Align, "outer", "center", "top")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Outer.Center.Middle.WidgetIds), d.Outer.Center.Middle.Align, "outer", "center", "middle")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Outer.Center.Bottom.WidgetIds), d.Outer.Center.Bottom.Align, "outer", "center", "bottom")
	// Outer Right
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Outer.Right.Top.WidgetIds), d.Outer.Right.Top.Align, "outer", "right", "top")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Outer.Right.Middle.WidgetIds), d.Outer.Right.Middle.Align, "outer", "right", "middle")
	was.WidgetAreaFrom(stringsToWidgetIDs(d.Outer.Right.Bottom.WidgetIds), d.Outer.Right.Bottom.Align, "outer", "right", "bottom")
	return was
}

func widgetIDsToStrings(wids []*id.WidgetID) []string {
	if wids == nil {
		return nil
	}
	docids := make([]string, 0, len(wids))
	for _, wid := range wids {
		if wid == nil {
			continue
		}
		docids = append(docids, wid.String())
	}
	return docids
}

func stringsToWidgetIDs(wids []string) []*id.WidgetID {
	if wids == nil {
		return nil
	}
	var docids []*id.WidgetID
	for _, wid := range wids {
		nid, err := id.WidgetIDFrom(wid)
		if err != nil {
			continue
		}
		docids = append(docids, &nid)
	}
	return docids
}

func buildWidgetZone(sas *scene.WidgetAlignSystem, z string) WidgetZone {
	return WidgetZone{
		Left:   buildWidgetSection(sas, z, "left"),
		Center: buildWidgetSection(sas, z, "center"),
		Right:  buildWidgetSection(sas, z, "right"),
	}
}

func buildWidgetSection(was *scene.WidgetAlignSystem, z, s string) WidgetSection {
	return WidgetSection{
		Top: WidgetArea{
			WidgetIds: widgetIDsToStrings(was.WidgetIds(z, s, "top")),
			Align:     *was.Alignment(z, s, "top"),
		},
		Middle: WidgetArea{
			WidgetIds: widgetIDsToStrings(was.WidgetIds(z, s, "middle")),
			Align:     *was.Alignment(z, s, "middle"),
		},
		Bottom: WidgetArea{
			WidgetIds: widgetIDsToStrings(was.WidgetIds(z, s, "bottom")),
			Align:     *was.Alignment(z, s, "bottom"),
		},
	}
}
