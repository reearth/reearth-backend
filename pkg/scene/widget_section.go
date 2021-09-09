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

func NewWidgetSection() *WidgetSection {
	return &WidgetSection{}
}

func (s *WidgetSection) Add(wid id.WidgetID, area string) {
	if s == nil {
		return
	}

	switch area {
	case string(WidgetAreaTop):
		s.top.Add(wid)
	case string(WidgetAreaMiddle):
		s.middle.Add(wid)
	case string(WidgetAreaBottom):
		s.bottom.Add(wid)
	default:
		return
	}
}

func (s *WidgetSection) AddAll(wids []id.WidgetID, align, area string) {
	if s == nil {
		return
	}

	switch area {
	case string(WidgetAreaTop):
		s.top.AddAll(wids, align)
	case string(WidgetAreaMiddle):
		s.middle.AddAll(wids, align)
	case string(WidgetAreaBottom):
		s.bottom.AddAll(wids, align)
	default:
		return
	}
}

func (s *WidgetSection) Remove(wid id.WidgetID) {
	if s == nil {
		return
	}

	s.top.Remove(wid)
	s.middle.Remove(wid)
	s.bottom.Remove(wid)
}

func (s *WidgetSection) Find(wid id.WidgetID) (int, *WidgetArea) {
	if s == nil {
		return -1, nil
	}

	i, wa := s.top.Find(wid)
	if wa != nil && i != -1 {
		return i, wa
	}
	i2, wa2 := s.middle.Find(wid)
	if wa2 != nil && i2 != -1 {
		return i2, wa2
	}
	i3, wa3 := s.bottom.Find(wid)
	if wa3 != nil && i3 != -1 {
		return i3, wa3
	}
	return -1, nil
}

func (ws *WidgetSection) Area(a string) *WidgetArea {
	switch a {
	case string(WidgetAreaTop):
		return &ws.top
	case string(WidgetAreaMiddle):
		return &ws.middle
	case string(WidgetAreaBottom):
		return &ws.bottom
	}
	return nil
}
