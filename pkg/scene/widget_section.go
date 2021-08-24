package scene

import "github.com/reearth/reearth-backend/pkg/id"

// WidgetSection is the structure of each section of the align system.
type WidgetSection struct {
	top    WidgetArea
	middle WidgetArea
	bottom WidgetArea
}

const (
	WidgetSectionLeft   = "left"
	WidgetSectionCenter = "center"
	WidgetSectionRight  = "right"
)

func (s *WidgetSection) Remove(wid id.WidgetID) {
	if s == nil {
		return
	}

	s.top.Remove(wid)
	s.middle.Remove(wid)
	s.bottom.Remove(wid)
}

func (s *WidgetSection) Find(wid id.WidgetID) (*int, *WidgetArea) {
	if s == nil {
		return nil, nil
	}

	i, wa := s.top.Find(wid)
	if wa != nil && i != nil {
		return i, wa
	}
	i2, wa2 := s.middle.Find(wid)
	if wa2 != nil && i2 != nil {
		return i2, wa2
	}
	i3, wa3 := s.bottom.Find(wid)
	if wa3 != nil && i3 != nil {
		return i3, wa3
	}
	return nil, nil
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
