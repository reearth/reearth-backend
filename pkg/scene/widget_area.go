package scene

import "github.com/reearth/reearth-backend/pkg/id"

// WidgetArea has the widgets and alignment information found in each part area of a section.
type WidgetArea struct {
	widgetIds []id.WidgetID
	align     string
}

var (
	WidgetAreaTop    = "top"
	WidgetAreaMiddle = "middle"
	WidgetAreaBottom = "bottom"
)

// WidgetIds will return a slice of widget ids from a specific area.
func (wa *WidgetArea) WidgetIDs() []id.WidgetID {
	return wa.widgetIds
}

// Alignment will return the alignment of a specific area.
func (wa *WidgetArea) Alignment() *string {
	return &wa.align
}

func (a *WidgetArea) Remove(wid id.WidgetID) {
	if a == nil {
		return
	}

	for i, w := range a.widgetIds {
		if w == wid {
			a.widgetIds = append(a.widgetIds[:i], a.widgetIds[i+1:]...)
			return
		}
	}
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
