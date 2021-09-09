package scene

import "github.com/reearth/reearth-backend/pkg/id"

// WidgetZone is the structure of each layer (inner and outer) of the align system.
type WidgetZone struct {
	left   WidgetSection
	center WidgetSection
	right  WidgetSection
}

type WidgetZoneType string

const (
	WidgetZoneInner WidgetZoneType = "inner"
	WidgetZoneOuter WidgetZoneType = "outer"
)

func NewWidgetZone() *WidgetZone {
	return &WidgetZone{}
}

func (z *WidgetZone) Add(wid id.WidgetID, section WidgetSectionType, area WidgetAreaType) {
	if z == nil {
		return
	}

	switch section {
	case WidgetSectionLeft:
		z.left.Add(wid, area)
	case WidgetSectionCenter:
		z.center.Add(wid, area)
	case WidgetSectionRight:
		z.right.Add(wid, area)
	}
}

func (z *WidgetZone) AddAll(wids []id.WidgetID, align string, section WidgetSectionType, area WidgetAreaType) {
	if z == nil {
		return
	}

	switch section {
	case WidgetSectionLeft:
		z.left.AddAll(wids, align, area)
	case WidgetSectionCenter:
		z.center.AddAll(wids, align, area)
	case WidgetSectionRight:
		z.right.AddAll(wids, align, area)
	}
}

func (z *WidgetZone) Remove(wid id.WidgetID) {
	if z == nil {
		return
	}

	z.left.Remove(wid)
	z.center.Remove(wid)
	z.right.Remove(wid)
}

func (z *WidgetZone) Find(wid id.WidgetID) (int, *WidgetArea) {
	if z == nil {
		return -1, nil
	}

	i, wa := z.left.Find(wid)
	if wa != nil && i != -1 {
		return i, wa
	}
	i2, wa2 := z.center.Find(wid)
	if wa2 != nil && i2 != -1 {
		return i2, wa2
	}
	i3, wa3 := z.right.Find(wid)
	if wa3 != nil && i3 != -1 {
		return i3, wa3
	}
	return -1, nil
}

func (wz *WidgetZone) Section(s WidgetSectionType) *WidgetSection {
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
