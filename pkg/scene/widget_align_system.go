package scene

import (
	"github.com/reearth/reearth-backend/pkg/id"
)

// WidgetAlignSystem is the layout structure of any enabled widgets that will be displayed over the scene
type WidgetAlignSystem struct {
	inner WidgetZone
	outer WidgetZone
}

// WidgetZone is the structure of each layer (inner and outer) of the align system
type WidgetZone struct {
	left   WidgetSection
	center WidgetSection
	right  WidgetSection
}

// WidgetSection is the structure of each section of the align system
type WidgetSection struct {
	top    WidgetArea
	middle WidgetArea
	bottom WidgetArea
}

// WidgetArea has the widgets and alignment information found in each part area of a section
type WidgetArea struct {
	widgetIds []*id.WidgetID
	align     string
}

type Location struct {
	Zone    string
	Section string
	Area    string
}

// NewWidgetAlignSystem returns a new widget align system
func NewWidgetAlignSystem() *WidgetAlignSystem {
	return &WidgetAlignSystem{}
}

// WidgetZone will return a specific zone in the align system
func (was *WidgetAlignSystem) WidgetZone(zone string) *WidgetZone {
	if was == nil {
		return nil
	}
	switch zone {
	case "inner":
		return &was.inner
	case "outer":
		return &was.outer
	}
	return nil
}

// WidgetSection will return a specific section in the align system
func (was *WidgetAlignSystem) WidgetSection(zone, section string) *WidgetSection {
	if was == nil {
		return nil
	}

	z := was.WidgetZone(zone)

	switch section {
	case "left":
		return &z.left
	case "center":
		return &z.center
	case "right":
		return &z.right
	}
	return nil
}

// WidgetArea will return a specific area in the align system
func (was *WidgetAlignSystem) WidgetArea(zone, section, area string) *WidgetArea {
	if was == nil {
		return nil
	}

	s := was.WidgetSection(zone, section)

	switch area {
	case "top":
		return &s.top
	case "middle":
		return &s.middle
	case "bottom":
		return &s.bottom
	}
	return nil
}

// WidgetIds will return a slice of widget ids from a specific area
func (was *WidgetAlignSystem) WidgetIds(z, s, a string) []*id.WidgetID {
	area := was.WidgetArea(z, s, a)
	return area.widgetIds
}

// Alignment will return the alignment of a specific area
func (was *WidgetAlignSystem) Alignment(z, s, a string) *string {
	area := was.WidgetArea(z, s, a)
	return &area.align
}

// Add a widget to the align system
func (was *WidgetAlignSystem) Add(wid *id.WidgetID, l *Location) {
	if was == nil {
		return
	}
	a := was.WidgetArea(l.Zone, l.Section, l.Area)
	nIds := append(a.widgetIds, wid)
	a.widgetIds = nIds
}

// Remove a widget from the align system
func (was *WidgetAlignSystem) Remove(wid *id.WidgetID, l *Location) {
	if was == nil {
		return
	}

	a := was.WidgetArea(l.Zone, l.Section, l.Area)

	var nwid []*id.WidgetID
	for _, w := range a.widgetIds {
		if w.ID() != wid.ID() {
			nwid = append(nwid, w)
		}
	}
	a.widgetIds = nwid
}

// Update a widget in the align system
func (was *WidgetAlignSystem) Update(wid *id.WidgetID, l, newL *Location, index, newIndex *int, align *string) {
	if was == nil && wid == nil && l == nil {
		return
	}

	a := was.WidgetArea(l.Zone, l.Section, l.Area)

	if align != nil {
		a.align = *align
	}

	if index != nil && newIndex != nil {
		moveInt(a.widgetIds, *index, *newIndex)
	}
	if newL != nil {
		was.Remove(wid, l)
		was.Add(wid, newL)
	}
}

// moveInt moves a widget's index
func moveInt(array []*id.WidgetID, srcIndex int, dstIndex int) []*id.WidgetID {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

// insertInt is used in moveInt to add the widgetID to a new position(index)
func insertInt(array []*id.WidgetID, value *id.WidgetID, index int) []*id.WidgetID {
	return append(array[:index], append([]*id.WidgetID{value}, array[index:]...)...)
}

// removeInt is used in moveInt to remove the widgetID from original position(index)
func removeInt(array []*id.WidgetID, index int) []*id.WidgetID {
	return append(array[:index], array[index+1:]...)
}

// WidgetAreaFrom will add a slice of widget ids to a specific area of an align system
func (was *WidgetAlignSystem) WidgetAreaFrom(wids []*id.WidgetID, align, z, s, a string) {
	if was == nil {
		return
	}
	wa := was.WidgetArea(z, s, a)
	wa.widgetIds = wids
}
