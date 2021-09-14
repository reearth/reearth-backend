package scene

import (
	"github.com/reearth/reearth-backend/pkg/id"
)

type WidgetLocation struct {
	Zone    WidgetZoneType
	Section WidgetSectionType
	Area    WidgetAreaType
}

// WidgetAlignSystem is the layout structure of any enabled widgets that will be displayed over the scene.
type WidgetAlignSystem struct {
	inner *WidgetZone
	outer *WidgetZone
}

type WidgetZoneType string

const (
	WidgetZoneInner WidgetZoneType = "inner"
	WidgetZoneOuter WidgetZoneType = "outer"
)

// NewWidgetAlignSystem returns a new widget align system.
func NewWidgetAlignSystem() *WidgetAlignSystem {
	return &WidgetAlignSystem{}
}

// Zone will return a specific zone in the align system.
func (was *WidgetAlignSystem) Zone(zone WidgetZoneType) *WidgetZone {
	if was == nil {
		return nil
	}
	switch zone {
	case WidgetZoneInner:
		if was.inner == nil {
			was.inner = NewWidgetZone()
		}
		return was.inner
	case WidgetZoneOuter:
		if was.outer == nil {
			was.outer = NewWidgetZone()
		}
		return was.outer
	}
	return nil
}

// Remove a widget from the align system.
func (was *WidgetAlignSystem) Remove(wid id.WidgetID) {
	if was == nil {
		return
	}

	was.inner.Remove(wid)
	was.outer.Remove(wid)
}

func (was *WidgetAlignSystem) Area(loc WidgetLocation) *WidgetArea {
	return was.Zone(loc.Zone).Section(loc.Section).Area(loc.Area)
}

func (was *WidgetAlignSystem) Find(wid id.WidgetID) (int, WidgetLocation) {
	if was == nil {
		return -1, WidgetLocation{}
	}

	if i, section, area := was.inner.Find(wid); i >= 0 {
		return i, WidgetLocation{
			Zone:    WidgetZoneInner,
			Section: section,
			Area:    area,
		}
	}
	if i, section, area := was.outer.Find(wid); i >= 0 {
		return i, WidgetLocation{
			Zone:    WidgetZoneOuter,
			Section: section,
			Area:    area,
		}
	}

	return -1, WidgetLocation{}
}

// Update a widget in the align system.
// func (was *WidgetAlignSystem) Update(wid id.WidgetID, l *WidgetLocation, index *int, align *WidgetAlignType) {
// 	if was == nil {
// 		return
// 	}

// 	i, a := was.Find(wid)
// 	if a == nil {
// 		return
// 	}

// 	if align != nil {
// 		switch *align {
// 		case WidgetAlignStart:
// 			a.align = WidgetAlignStart
// 		case WidgetAlignCenter:
// 			a.align = WidgetAlignCenter
// 		case WidgetAlignEnd:
// 			a.align = WidgetAlignEnd
// 		default:
// 			a.align = WidgetAlignStart
// 		}
// 	}

// 	if index != nil {
// 		moveInt(a.widgetIds, i, *index)
// 	}
// 	if l != nil {
// 		was.Remove(wid)
// 		// was.Add(wid, *l)
// 	}
// }

// // moveInt moves a widget's index.
// func moveInt(array []id.WidgetID, srcIndex int, dstIndex int) []id.WidgetID {
// 	value := array[srcIndex]
// 	return insertInt(removeInt(array, srcIndex), value, dstIndex)
// }

// // insertInt is used in moveInt to add the widgetID to a new position(index).
// func insertInt(array []id.WidgetID, value id.WidgetID, index int) []id.WidgetID {
// 	return append(array[:index], append([]id.WidgetID{value}, array[index:]...)...)
// }

// // removeInt is used in moveInt to remove the widgetID from original position(index).
// func removeInt(array []id.WidgetID, index int) []id.WidgetID {
// 	return append(array[:index], array[index+1:]...)
// }
