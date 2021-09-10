package scene

import "github.com/reearth/reearth-backend/pkg/id"

// WidgetArea has the widgets and alignment information found in each part area of a section.
type WidgetArea struct {
	widgetIds []id.WidgetID
	align     WidgetAlignType
}

type WidgetAreaType string

var (
	WidgetAreaTop    WidgetAreaType = "top"
	WidgetAreaMiddle WidgetAreaType = "middle"
	WidgetAreaBottom WidgetAreaType = "bottom"
)

func NewWidgetArea() *WidgetArea {
	return &WidgetArea{}
}

// WidgetIds will return a slice of widget ids from a specific area.
func (wa *WidgetArea) WidgetIDs() []id.WidgetID {
	return wa.widgetIds
}

// Alignment will return the alignment of a specific area.
func (wa *WidgetArea) Alignment() *WidgetAlignType {
	return &wa.align
}

func (a *WidgetArea) Add(wid id.WidgetID) {
	if a == nil {
		return
	}

	nIds := a.widgetIds

	if b := wid.Contains(a.widgetIds); !b {
		nIds = append(a.widgetIds, wid)
	}
	a.widgetIds = nIds

	if a.align == "" {
		a.align = WidgetAlignStart
	}
}

func (a *WidgetArea) AddAll(wids []id.WidgetID, align WidgetAlignType) {
	if a == nil {
		return
	}

	a.widgetIds = wids
	a.align = align
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

func (a *WidgetArea) Find(wid id.WidgetID) (int, *WidgetArea) {
	if a == nil {
		return -1, nil
	}

	for i, w := range a.widgetIds {
		if w == wid {
			return i, a
		}
	}
	return -1, nil
}
