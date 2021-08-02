package scene

import (
	"github.com/reearth/reearth-backend/pkg/id"
)

// WidgetAlignSystem is the layout structure of any enabled widgets that will be displayed over the scene.
type WidgetAlignSystem struct {
	inner WidgetZone
	outer WidgetZone
}

// WidgetZone is the structure of each layer (inner and outer) of the align system.
type WidgetZone struct {
	left   WidgetSection
	center WidgetSection
	right  WidgetSection
}

// WidgetSection is the structure of each section of the align system.
type WidgetSection struct {
	top    WidgetArea
	middle WidgetArea
	bottom WidgetArea
}

// WidgetArea has the widgets and alignment information found in each part area of a section.
type WidgetArea struct {
	widgetIds []*id.WidgetID
	align     string
}

type WidgetLocation struct {
	Zone    string
	Section string
	Area    string
}

var Areas = []string{
	"top",
	"middle",
	"bottom",
}

var Sections = map[string][]string{
	"left":   Areas,
	"center": Areas,
	"right":  Areas,
}

var Zones = map[string]map[string][]string{
	"inner": Sections,
	"outer": Sections,
}

func (was *WidgetAlignSystem) FindWidgetIDLocation(wid id.WidgetID) (*int, *WidgetLocation) {
	for z, s := range Zones {
		for s2, a := range s {
			for _, a2 := range a {
				if i, h := was.Area(z, s2, a2).Has(wid); h {
					wloc := WidgetLocation{z, s2, a2}
					return i, &wloc
				}
			}
		}
	}
	return nil, nil
}

// NewWidgetAlignSystem returns a new widget align system.
func NewWidgetAlignSystem() *WidgetAlignSystem {
	return &WidgetAlignSystem{}
}

// Zone will return a specific zone in the align system.
func (was *WidgetAlignSystem) Zone(zone string) *WidgetZone {
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

// Section will return a specific section in the align system.
func (was *WidgetAlignSystem) Section(zone, section string) *WidgetSection {
	if was == nil {
		return nil
	}

	z := was.Zone(zone)

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

// Area will return a specific area in the align system.
func (was *WidgetAlignSystem) Area(zone, section, area string) *WidgetArea {
	if was == nil {
		return nil
	}

	s := was.Section(zone, section)

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

// WidgetIds will return a slice of widget ids from a specific area.
func (was *WidgetAlignSystem) WidgetIds(z, s, a string) []*id.WidgetID {
	area := was.Area(z, s, a)
	return area.widgetIds
}

// Alignment will return the alignment of a specific area.
func (was *WidgetAlignSystem) Alignment(z, s, a string) *string {
	area := was.Area(z, s, a)
	return &area.align
}

// Add a widget to the align system.
func (was *WidgetAlignSystem) Add(wid id.WidgetID, l *WidgetLocation) {
	if was == nil {
		return
	}

	a := was.Area(l.Zone, l.Section, l.Area)
	nIds := a.widgetIds

	if _, b := a.Has(wid); !b {
		nIds = append(a.widgetIds, &wid)
	}
	a.widgetIds = nIds

	if a.align == "" {
		a.align = "start"
	}
}

// Remove a widget from the align system.
func (was *WidgetAlignSystem) Remove(wid id.WidgetID) {
	if was == nil {
		return
	}
	var nwids []*id.WidgetID
	i, loc := was.FindWidgetIDLocation(wid)
	if loc != nil {
		a := was.Area(loc.Zone, loc.Section, loc.Area)
		if len(a.widgetIds) > 0 {
			nwids = append(a.widgetIds[:*i], a.widgetIds[*i+1:]...)
		}
		a.widgetIds = nwids
	}
}

// Update a widget in the align system.
func (was *WidgetAlignSystem) Update(wid id.WidgetID, l *WidgetLocation, index *int, align *string) {
	if was == nil {
		return
	}

	i, oldL := was.FindWidgetIDLocation(wid)
	a := was.Area(oldL.Zone, oldL.Section, oldL.Area)

	if align != nil {
		a.align = *align
	} else {
		a.align = "start"
	}

	if index != nil {
		moveInt(a.widgetIds, *i, *index)
	}
	if l != nil {
		was.Remove(wid)
		was.Add(wid, l)
	}
}

// moveInt moves a widget's index.
func moveInt(array []*id.WidgetID, srcIndex int, dstIndex int) []*id.WidgetID {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

// insertInt is used in moveInt to add the widgetID to a new position(index).
func insertInt(array []*id.WidgetID, value *id.WidgetID, index int) []*id.WidgetID {
	return append(array[:index], append([]*id.WidgetID{value}, array[index:]...)...)
}

// removeInt is used in moveInt to remove the widgetID from original position(index).
func removeInt(array []*id.WidgetID, index int) []*id.WidgetID {
	return append(array[:index], array[index+1:]...)
}

// WidgetAreaFrom will add a slice of widget ids to a specific area of an align system.
func (was *WidgetAlignSystem) WidgetAreaFrom(wids []*id.WidgetID, align, z, s, a string) {
	if was == nil {
		return
	}
	wa := was.Area(z, s, a)
	wa.widgetIds = wids
}

// Has will check a widget area's slice of widgetIds for the specified ID and return a bool value.
func (wa *WidgetArea) Has(wid id.WidgetID) (*int, bool) {
	if wa == nil {
		return nil, false
	}
	for i, id := range wa.widgetIds {
		if id.Equal(wid) {
			return &i, true
		}
	}
	return nil, false
}
