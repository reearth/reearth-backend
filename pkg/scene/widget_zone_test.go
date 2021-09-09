package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestWidgetZone_Find(t *testing.T) {
	wid := id.NewWidgetID()

	wz := NewWidgetZone()
	wz.left.top.widgetIds = append(wz.left.top.widgetIds, wid)
	e := wz.left.top

	testCases := []struct {
		Name     string
		Input    id.WidgetID
		WZ       *WidgetZone
		Expected *WidgetArea
	}{
		{
			Name:     "Find the location of a widgetID and return the WidgetArea",
			Input:    wid,
			WZ:       wz,
			Expected: &e,
		},
		{
			Name:     "Return nil if no Zone",
			Input:    wid,
			WZ:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			_, res := tc.WZ.Find(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetZone_Section(t *testing.T) {
	wid := id.NewWidgetID()

	wz := NewWidgetZone()
	wz.left.top.widgetIds = append(wz.left.top.widgetIds, wid)
	e := wz.left

	testCases := []struct {
		Name     string
		Input    WidgetSectionType
		WZ       *WidgetZone
		Expected *WidgetSection
	}{
		{
			Name:     "Find the location of a widgetID and return the WidgetArea",
			Input:    "left",
			WZ:       wz,
			Expected: &e,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WZ.Section(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
