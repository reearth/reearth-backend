package scene

import (
	"github.com/reearth/reearth-backend/pkg/id"
)

// WidgetAlignSystem is the layout structure of any enabled widgets that will be displayed over the scene.
type WidgetAlignSystem struct {
	inner WidgetZone
	outer WidgetZone
}

type WidgetLocation struct {
	Zone    string
	Section string
	Area    string
}

const (
	WidgetAlignStart  = "start"
	WidgetAlignCenter = "centered"
	WidgetAlignEnd    = "end"
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

// Add a widget to the align system.
func (was *WidgetAlignSystem) Add(wid id.WidgetID, loc WidgetLocation) {
	if was == nil {
		return
	}

	a := was.Area(loc.Zone, loc.Section, loc.Area)
	nIds := a.widgetIds

	if b := a.Has(wid); !b {
		nIds = append(a.widgetIds, wid)
	}
	a.widgetIds = nIds

	if a.align == "" {
		a.align = WidgetAlignStart
	}
}

// AddAll will add a slice of widget IDs and alignment to a WidgetArea
func (was *WidgetAlignSystem) AddAll(wids []id.WidgetID, align string, loc WidgetLocation) {
	if was == nil {
		return
	}
	wa := was.Area(loc.Zone, loc.Section, loc.Area)
	wa.widgetIds = wids
	wa.align = align
}

// Update a widget in the align system.
func (was *WidgetAlignSystem) Update(wid id.WidgetID, l *WidgetLocation, index *int, align *string) {
	if was == nil {
		return
	}

	i, a := was.FindWidgetLocation(wid)
	if a == nil {
		return
	}

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
	}

	if index != nil {
		moveInt(a.widgetIds, i, *index)
	}
	if l != nil {
		was.Remove(wid)
		was.Add(wid, *l)
	}
}

// Remove a widget from the align system.
func (was *WidgetAlignSystem) Remove(wid id.WidgetID) {
	if was == nil {
		return
	}

	was.inner.Remove(wid)
	was.outer.Remove(wid)
}

func (was *WidgetAlignSystem) FindWidgetLocation(wid id.WidgetID) (int, *WidgetArea) {
	if was == nil {
		return -1, nil
	}
	i, wa := was.inner.Find(wid)
	if wa != nil && i != -1 {
		return i, wa
	}
	i2, wa2 := was.outer.Find(wid)
	if wa2 != nil && i2 != -1 {
		return i2, wa2
	}

	return -1, nil
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
