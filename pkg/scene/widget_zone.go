package scene

import "github.com/reearth/reearth-backend/pkg/id"

// WidgetZone is the structure of each layer (inner and outer) of the align system.
type WidgetZone struct {
	left   WidgetSection
	center WidgetSection
	right  WidgetSection
}

func (z *WidgetZone) Remove(wid id.WidgetID) {
	if z == nil {
		return
	}

	z.left.Remove(wid)
	z.center.Remove(wid)
	z.right.Remove(wid)
}

func (z *WidgetZone) Find(wid id.WidgetID) (*int, *WidgetArea) {
	if z == nil {
		return nil, nil
	}

	i, wa := z.left.Find(wid)
	if wa != nil && i != nil {
		return i, wa
	}
	i2, wa2 := z.center.Find(wid)
	if wa2 != nil && i2 != nil {
		return i2, wa2
	}
	i3, wa3 := z.right.Find(wid)
	if wa3 != nil && i3 != nil {
		return i3, wa3
	}
	return nil, nil
}

func (wz *WidgetZone) Section(s string) *WidgetSection {
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
