package mongodoc

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

type WidgetLocationDocument struct {
	Zone    string
	Section string
	Area    string
}

type WidgetLayout struct {
	Extendable      bool
	Extended        bool
	Floating        bool
	DefaultLocation *WidgetLocationDocument
}

type WidgetArea struct {
	WidgetIDs []string
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
	was.AddAll(stringsToWidgetIDs(d.Inner.Left.Top.WidgetIDs), d.Inner.Left.Top.Align, "inner", "left", "top")
	was.AddAll(stringsToWidgetIDs(d.Inner.Left.Middle.WidgetIDs), d.Inner.Left.Middle.Align, "inner", "left", "middle")
	was.AddAll(stringsToWidgetIDs(d.Inner.Left.Bottom.WidgetIDs), d.Inner.Left.Bottom.Align, "inner", "left", "bottom")
	// Inner Center
	was.AddAll(stringsToWidgetIDs(d.Inner.Center.Top.WidgetIDs), d.Inner.Center.Top.Align, "inner", "center", "top")
	was.AddAll(stringsToWidgetIDs(d.Inner.Center.Middle.WidgetIDs), d.Inner.Center.Middle.Align, "inner", "center", "middle")
	was.AddAll(stringsToWidgetIDs(d.Inner.Center.Bottom.WidgetIDs), d.Inner.Center.Bottom.Align, "inner", "center", "bottom")
	// Inner Right
	was.AddAll(stringsToWidgetIDs(d.Inner.Right.Top.WidgetIDs), d.Inner.Right.Top.Align, "inner", "right", "top")
	was.AddAll(stringsToWidgetIDs(d.Inner.Right.Middle.WidgetIDs), d.Inner.Right.Middle.Align, "inner", "right", "middle")
	was.AddAll(stringsToWidgetIDs(d.Inner.Right.Bottom.WidgetIDs), d.Inner.Right.Bottom.Align, "inner", "right", "bottom")
	// Outer Left
	was.AddAll(stringsToWidgetIDs(d.Outer.Left.Top.WidgetIDs), d.Outer.Left.Top.Align, "outer", "left", "top")
	was.AddAll(stringsToWidgetIDs(d.Outer.Left.Middle.WidgetIDs), d.Outer.Left.Middle.Align, "outer", "left", "middle")
	was.AddAll(stringsToWidgetIDs(d.Outer.Left.Bottom.WidgetIDs), d.Outer.Left.Bottom.Align, "outer", "left", "bottom")
	// Outer Center
	was.AddAll(stringsToWidgetIDs(d.Outer.Center.Top.WidgetIDs), d.Outer.Center.Top.Align, "outer", "center", "top")
	was.AddAll(stringsToWidgetIDs(d.Outer.Center.Middle.WidgetIDs), d.Outer.Center.Middle.Align, "outer", "center", "middle")
	was.AddAll(stringsToWidgetIDs(d.Outer.Center.Bottom.WidgetIDs), d.Outer.Center.Bottom.Align, "outer", "center", "bottom")
	// Outer Right
	was.AddAll(stringsToWidgetIDs(d.Outer.Right.Top.WidgetIDs), d.Outer.Right.Top.Align, "outer", "right", "top")
	was.AddAll(stringsToWidgetIDs(d.Outer.Right.Middle.WidgetIDs), d.Outer.Right.Middle.Align, "outer", "right", "middle")
	was.AddAll(stringsToWidgetIDs(d.Outer.Right.Bottom.WidgetIDs), d.Outer.Right.Bottom.Align, "outer", "right", "bottom")
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
			WidgetIDs: widgetIDsToStrings(was.Area(z, s, "top").WidgetIDs()),
			Align:     *was.Area(z, s, "top").Alignment(),
		},
		Middle: WidgetArea{
			WidgetIDs: widgetIDsToStrings(was.Area(z, s, "middle").WidgetIDs()),
			Align:     *was.Area(z, s, "middle").Alignment(),
		},
		Bottom: WidgetArea{
			WidgetIDs: widgetIDsToStrings(was.Area(z, s, "bottom").WidgetIDs()),
			Align:     *was.Area(z, s, "bottom").Alignment(),
		},
	}
}
