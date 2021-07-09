package scene

import "github.com/reearth/reearth-backend/pkg/id"

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
	// position  string
	widgetIds []id.WidgetID
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

// Layout will return a widget align system if not nil
func (was *WidgetAlignSystem) WidgetAlignSystem() *WidgetAlignSystem {
	if was == nil {
		return nil
	}
	return was
}

// WidgetZone will return a specific zone in the align system
func (was *WidgetAlignSystem) WidgetZone(zone string) *WidgetZone {
	if was == nil {
		return nil
	}
	if zone == "inner" {
		return &was.inner
	} else if zone == "outer" {
		return &was.outer
	} else {
		return nil
	}
}

// WidgetSection will return a specific section in the align system
func (was *WidgetAlignSystem) WidgetSection(zone, section string) *WidgetSection {
	if was == nil {
		return nil
	}

	z := was.WidgetZone(zone)

	if section == "left" {
		return &z.left
	} else if section == "center" {
		return &z.center
	} else if section == "right" {
		return &z.right
	} else {
		return nil
	}
}

// WidgetArea will return a specific area in the align system
func (was *WidgetAlignSystem) WidgetArea(zone, section, area string) *WidgetArea {
	if was == nil {
		return nil
	}

	s := was.WidgetSection(zone, section)

	if area == "top" {
		return &s.top
	} else if area == "middle" {
		return &s.middle
	} else if area == "bottom" {
		return &s.bottom
	} else {
		return nil
	}
}

// Add a widget to the align system
func (was *WidgetAlignSystem) Add(wid *id.WidgetID, l *Location) {
	if was == nil {
		return
	}

	a := was.WidgetArea(l.Zone, l.Section, l.Area)
	id := *wid
	a.widgetIds = append(a.widgetIds, id)
}

// Remove a widget from the align system
func (was *WidgetAlignSystem) Remove(wid *id.WidgetID, l *Location) {
	if was == nil {
		return
	}

	a := was.WidgetArea(l.Zone, l.Section, l.Area)

	for i, w := range a.widgetIds {
		if w.ID().Equal(wid.ID()) {
			a.widgetIds = append(a.widgetIds[:i], a.widgetIds[i+1])
			return
		}
	}
}

// Move widget
// func (was *WidgetAlignSystem) Move(wid *id.WidgetID, oldLocation, newLocation *Location) {
// 	if was == nil {
// 		return
// 	}
// 	was.Remove(wid, oldLocation)
// 	was.Add(wid, newLocation)

// 	// old := was.WidgetArea(oldLocation.zone, oldLocation.section, oldLocation.area).widgetIds
// 	// for i, w := range old {
// 	// 	if w.ID().Equal(wid.ID()) {
// 	// 		old = append(old[:i], old[i+1])
// 	// 	}
// 	// }
// }

// // Reorder widgets in an area
// // func(was *WidgetAlignSystem) Reorder(wid *id.WidgetID, oldIndex, newIndex int){
// // 	if was == nil {
// // 		return
// // 	}

// // }
