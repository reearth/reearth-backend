package scene

import "github.com/reearth/reearth-backend/pkg/id"

// WidgetArea has the widgets and alignment information found in each part area of a section.
type WidgetArea struct {
	widgetIds []id.WidgetID
	align     WidgetAlignType
}

type WidgetAlignType string

const (
	WidgetAlignStart  WidgetAlignType = "start"
	WidgetAlignCenter WidgetAlignType = "centered"
	WidgetAlignEnd    WidgetAlignType = "end"
)

func NewWidgetArea(widgetIds []id.WidgetID, align WidgetAlignType) *WidgetArea {
	wa := &WidgetArea{}
	wa.AddAll(widgetIds)
	wa.SetAlignment(align)
	return wa
}

// WidgetIds will return a slice of widget ids from a specific area.
func (a *WidgetArea) WidgetIDs() []id.WidgetID {
	if a == nil {
		return nil
	}

	return append([]id.WidgetID{}, a.widgetIds...)
}

// Alignment will return the alignment of a specific area.
func (a *WidgetArea) Alignment() WidgetAlignType {
	if a == nil {
		return ""
	}

	return a.align
}

func (a *WidgetArea) Find(wid id.WidgetID) int {
	if a == nil {
		return -1
	}

	for i, w := range a.widgetIds {
		if w == wid {
			return i
		}
	}
	return -1
}

func (a *WidgetArea) Add(wid id.WidgetID) {
	if a == nil {
		return
	}

	if b := wid.Contains(a.widgetIds); !b {
		a.widgetIds = append(a.widgetIds, wid)
	}
}

func (a *WidgetArea) AddAll(wids []id.WidgetID) {
	if a == nil {
		return
	}

	widgetIds := make([]id.WidgetID, 0, len(wids))
	for _, w := range wids {
		if w.Contains(a.widgetIds) || w.Contains(widgetIds) {
			continue
		}
		widgetIds = append(widgetIds, w)
	}

	a.widgetIds = widgetIds
}

func (a *WidgetArea) SetAlignment(at WidgetAlignType) {
	if a == nil {
		return
	}

	if at == WidgetAlignStart || at == WidgetAlignCenter || at == WidgetAlignEnd {
		a.align = at
	} else {
		a.align = WidgetAlignStart
	}
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
