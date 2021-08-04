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
	widgetIds []id.WidgetID
	align     string
}

type WidgetLocation struct {
	Zone    string
	Section string
	Area    string
}

const (
	WidgetAlignStart  = "start"
	WidgetAlignCenter = "center"
	WidgetAlignEnd    = "end"

	WidgetZoneInner = "inner"
	WidgetZoneOuter = "outer"

	WidgetSectionLeft   = "left"
	WidgetSectionCenter = "center"
	WidgetSectionRight  = "right"

	WidgetAreaTop    = "top"
	WidgetAreaMiddle = "middle"
	WidgetAreaBottom = "bottom"
)

var Areas = []string{
	WidgetAreaTop,
	WidgetAreaMiddle,
	WidgetAreaBottom,
}

var Sections = map[string][]string{
	WidgetSectionLeft:   Areas,
	WidgetSectionCenter: Areas,
	WidgetSectionRight:  Areas,
}

var Zones = map[string]map[string][]string{
	WidgetZoneInner: Sections,
	WidgetZoneOuter: Sections,
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
	case WidgetZoneInner:
		return &was.inner
	case WidgetZoneOuter:
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
	case WidgetSectionLeft:
		return &z.left
	case WidgetSectionCenter:
		return &z.center
	case WidgetSectionRight:
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
	case WidgetAreaTop:
		return &s.top
	case WidgetAreaMiddle:
		return &s.middle
	case WidgetAreaBottom:
		return &s.bottom
	}
	return nil
}

func (wz *WidgetZone) Section(s string) *WidgetSection {
	switch s {
	case WidgetSectionLeft:
		return &wz.left
	case WidgetSectionCenter:
		return &wz.center
	case WidgetSectionRight:
		return &wz.right
	}
	return nil
}

func (ws *WidgetSection) Area(a string) *WidgetArea {
	switch a {
	case WidgetAreaTop:
		return &ws.top
	case WidgetAreaMiddle:
		return &ws.middle
	case WidgetAreaBottom:
		return &ws.bottom
	}
	return nil
}

// WidgetIds will return a slice of widget ids from a specific area.
func (wa *WidgetArea) WidgetIDs() []id.WidgetID {
	return wa.widgetIds
}

// Alignment will return the alignment of a specific area.
func (wa *WidgetArea) Alignment() *string {
	return &wa.align
}

// Add a widget to the align system.
func (was *WidgetAlignSystem) Add(wid id.WidgetID, loc *WidgetLocation) {
	if was == nil {
		return
	}

	a := was.Area(loc.Zone, loc.Section, loc.Area)
	nIds := a.widgetIds

	if _, b := a.Has(wid); !b {
		nIds = append(a.widgetIds, wid)
	}
	a.widgetIds = nIds

	if a.align == "" {
		a.align = WidgetAlignStart
	}
}

// AddAll will add a slice of widget IDs and alignment to a WidgetArea
func (was *WidgetAlignSystem) AddAll(wids []id.WidgetID, align string, loc *WidgetLocation) {
	if was == nil {
		return
	}
	wa := was.Area(loc.Zone, loc.Section, loc.Area)
	wa.widgetIds = wids
	wa.align = align
}

// Remove a widget from the align system.
func (was *WidgetAlignSystem) Remove(wid id.WidgetID) {
	if was == nil {
		return
	}
	var nwids []id.WidgetID
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
		switch *align {
		case WidgetAlignStart:
			a.align = WidgetAlignStart
		case WidgetAlignCenter:
			a.align = WidgetAlignCenter
		case WidgetAlignEnd:
			a.align = WidgetAlignEnd
		default:
			a.align = WidgetAlignStart
		}
	} else {
		a.align = WidgetAlignStart
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
func moveInt(array []id.WidgetID, srcIndex int, dstIndex int) []id.WidgetID {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

// insertInt is used in moveInt to add the widgetID to a new position(index).
func insertInt(array []id.WidgetID, value id.WidgetID, index int) []id.WidgetID {
	return append(array[:index], append([]id.WidgetID{value}, array[index:]...)...)
}

// removeInt is used in moveInt to remove the widgetID from original position(index).
func removeInt(array []id.WidgetID, index int) []id.WidgetID {
	return append(array[:index], array[index+1:]...)
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
