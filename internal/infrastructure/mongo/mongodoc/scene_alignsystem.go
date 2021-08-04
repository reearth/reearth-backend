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

type WidgetLayoutDocument struct {
	Extendable      bool
	Extended        bool
	Floating        bool
	DefaultLocation *WidgetLocationDocument
}

type WidgetAreaDocument struct {
	WidgetIDs []string
	Align     string
}

type WidgetSectionDocument struct {
	Top    WidgetAreaDocument
	Middle WidgetAreaDocument
	Bottom WidgetAreaDocument
}

type WidgetZoneDocument struct {
	Left   WidgetSectionDocument
	Center WidgetSectionDocument
	Right  WidgetSectionDocument
}

type SceneAlignSystemDocument struct {
	Inner WidgetZoneDocument
	Outer WidgetZoneDocument
}

func NewWidgetAlignSystem(was *scene.WidgetAlignSystem) *SceneAlignSystemDocument {
	widgetAlignDoc := SceneAlignSystemDocument{Inner: buildWidgetZone(was, "inner"), Outer: buildWidgetZone(was, "outer")}
	return &widgetAlignDoc
}

func (*SceneAlignSystemDocument) ToModelAlignSystem(d SceneAlignSystemDocument) *scene.WidgetAlignSystem {
	was := scene.NewWidgetAlignSystem()

	// Inner Left
	was.AddAll(
		stringsToWidgetIDs(d.Inner.Left.Top.WidgetIDs),
		d.Inner.Left.Top.Align,
		&scene.WidgetLocation{Zone: "inner", Section: "left", Area: "top"})
	was.AddAll(
		stringsToWidgetIDs(d.Inner.Left.Middle.WidgetIDs),
		d.Inner.Left.Middle.Align,
		&scene.WidgetLocation{Zone: "inner", Section: "left", Area: "middle"},
	)
	was.AddAll(
		stringsToWidgetIDs(d.Inner.Left.Bottom.WidgetIDs),
		d.Inner.Left.Bottom.Align,
		&scene.WidgetLocation{Zone: "inner", Section: "left", Area: "bottom"})
	// Inner Center
	was.AddAll(
		stringsToWidgetIDs(d.Inner.Center.Top.WidgetIDs),
		d.Inner.Center.Top.Align,
		&scene.WidgetLocation{Zone: "inner", Section: "center", Area: "top"})
	was.AddAll(
		stringsToWidgetIDs(d.Inner.Center.Middle.WidgetIDs),
		d.Inner.Center.Middle.Align,
		&scene.WidgetLocation{Zone: "inner", Section: "center", Area: "middle"})
	was.AddAll(
		stringsToWidgetIDs(d.Inner.Center.Bottom.WidgetIDs),
		d.Inner.Center.Bottom.Align,
		&scene.WidgetLocation{Zone: "inner", Section: "center", Area: "bottom"})
	// Inner Right
	was.AddAll(
		stringsToWidgetIDs(d.Inner.Right.Top.WidgetIDs),
		d.Inner.Right.Top.Align,
		&scene.WidgetLocation{Zone: "inner", Section: "right", Area: "top"})
	was.AddAll(
		stringsToWidgetIDs(d.Inner.Right.Middle.WidgetIDs),
		d.Inner.Right.Middle.Align,
		&scene.WidgetLocation{Zone: "inner", Section: "right", Area: "middle"})
	was.AddAll(
		stringsToWidgetIDs(d.Inner.Right.Bottom.WidgetIDs),
		d.Inner.Right.Bottom.Align,
		&scene.WidgetLocation{Zone: "inner", Section: "right", Area: "bottom"})
	// Outer Left
	was.AddAll(
		stringsToWidgetIDs(d.Outer.Left.Top.WidgetIDs),
		d.Outer.Left.Top.Align,
		&scene.WidgetLocation{Zone: "outer", Section: "left", Area: "top"})
	was.AddAll(
		stringsToWidgetIDs(d.Outer.Left.Middle.WidgetIDs),
		d.Outer.Left.Middle.Align,
		&scene.WidgetLocation{Zone: "outer", Section: "left", Area: "middle"})
	was.AddAll(
		stringsToWidgetIDs(d.Outer.Left.Bottom.WidgetIDs),
		d.Outer.Left.Bottom.Align,
		&scene.WidgetLocation{Zone: "outer", Section: "left", Area: "bottom"})
	// Outer Center
	was.AddAll(
		stringsToWidgetIDs(d.Outer.Center.Top.WidgetIDs),
		d.Outer.Center.Top.Align,
		&scene.WidgetLocation{Zone: "outer", Section: "center", Area: "top"})
	was.AddAll(
		stringsToWidgetIDs(d.Outer.Center.Middle.WidgetIDs),
		d.Outer.Center.Middle.Align,
		&scene.WidgetLocation{Zone: "outer", Section: "center", Area: "middle"})
	was.AddAll(
		stringsToWidgetIDs(d.Outer.Center.Bottom.WidgetIDs),
		d.Outer.Center.Bottom.Align,
		&scene.WidgetLocation{Zone: "outer", Section: "center", Area: "bottom"})
	// Outer Right
	was.AddAll(
		stringsToWidgetIDs(d.Outer.Right.Top.WidgetIDs),
		d.Outer.Right.Top.Align,
		&scene.WidgetLocation{Zone: "outer", Section: "right", Area: "top"})
	was.AddAll(
		stringsToWidgetIDs(d.Outer.Right.Middle.WidgetIDs),
		d.Outer.Right.Middle.Align,
		&scene.WidgetLocation{Zone: "outer", Section: "right", Area: "middle"})
	was.AddAll(
		stringsToWidgetIDs(d.Outer.Right.Bottom.WidgetIDs),
		d.Outer.Right.Bottom.Align,
		&scene.WidgetLocation{Zone: "outer", Section: "right", Area: "bottom"})
	return was
}

func widgetIDsToStrings(wids []id.WidgetID) []string {
	if wids == nil {
		return nil
	}
	docids := make([]string, 0, len(wids))
	for _, wid := range wids {
		docids = append(docids, wid.String())
	}
	return docids
}

func stringsToWidgetIDs(wids []string) []id.WidgetID {
	if wids == nil {
		return nil
	}
	var docids []id.WidgetID
	for _, wid := range wids {
		nid, err := id.WidgetIDFrom(wid)
		if err != nil {
			continue
		}
		docids = append(docids, nid)
	}
	return docids
}

func buildWidgetZone(sas *scene.WidgetAlignSystem, z string) WidgetZoneDocument {
	return WidgetZoneDocument{
		Left:   buildWidgetSection(sas, z, "left"),
		Center: buildWidgetSection(sas, z, "center"),
		Right:  buildWidgetSection(sas, z, "right"),
	}
}

func buildWidgetSection(was *scene.WidgetAlignSystem, z, s string) WidgetSectionDocument {
	return WidgetSectionDocument{
		Top: WidgetAreaDocument{
			WidgetIDs: widgetIDsToStrings(was.Area(z, s, "top").WidgetIDs()),
			Align:     *was.Area(z, s, "top").Alignment(),
		},
		Middle: WidgetAreaDocument{
			WidgetIDs: widgetIDsToStrings(was.Area(z, s, "middle").WidgetIDs()),
			Align:     *was.Area(z, s, "middle").Alignment(),
		},
		Bottom: WidgetAreaDocument{
			WidgetIDs: widgetIDsToStrings(was.Area(z, s, "bottom").WidgetIDs()),
			Align:     *was.Area(z, s, "bottom").Alignment(),
		},
	}
}
