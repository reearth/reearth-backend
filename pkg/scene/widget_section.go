package scene

import "github.com/reearth/reearth-backend/pkg/id"

// WidgetSection is the structure of each section of the align system.
type WidgetSection struct {
	top    WidgetArea
	middle WidgetArea
	bottom WidgetArea
}

type WidgetAreaType string

var (
	WidgetAreaTop    WidgetAreaType = "top"
	WidgetAreaMiddle WidgetAreaType = "middle"
	WidgetAreaBottom WidgetAreaType = "bottom"
)

func NewWidgetSection() *WidgetSection {
	return &WidgetSection{}
}

func (s *WidgetSection) Add(wid id.WidgetID, area WidgetAreaType) {
	if s == nil {
		return
	}

	switch area {
	case WidgetAreaTop:
		s.top.Add(wid)
	case WidgetAreaMiddle:
		s.middle.Add(wid)
	case WidgetAreaBottom:
		s.bottom.Add(wid)
	default:
		return
	}
}

func (s *WidgetSection) AddAll(wids []id.WidgetID, align WidgetAlignType, area WidgetAreaType) {
	if s == nil {
		return
	}

	switch area {
	case WidgetAreaTop:
		s.top.AddAll(wids)
	case WidgetAreaMiddle:
		s.middle.AddAll(wids)
	case WidgetAreaBottom:
		s.bottom.AddAll(wids)
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
	if i := s.top.Find(wid); i != -1 {
		return i, &s.top
	}
	if i2 := s.middle.Find(wid); i2 != -1 {
		return i2, &s.middle
	}
	if i3 := s.bottom.Find(wid); i3 != -1 {
		return i3, &s.bottom
	}
	return -1, nil
}

func (ws *WidgetSection) Area(a WidgetAreaType) *WidgetArea {
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
