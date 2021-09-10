package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestWidgetArea_Find(t *testing.T) {
	wid := id.NewWidgetID()

	wa := NewWidgetArea()
	wa.widgetIds = append(wa.widgetIds, wid)

	testCases := []struct {
		Name     string
		Input    id.WidgetID
		WA       *WidgetArea
		Expected *WidgetArea
	}{
		{
			Name:     "Return WidgetArea if contains widget id",
			Input:    wid,
			WA:       wa,
			Expected: wa,
		},
		{
			Name:     "Return nil if doesn't contain widget id",
			Input:    id.NewWidgetID(),
			WA:       wa,
			Expected: nil,
		},
		{
			Name:     "Return nil if WidgetArea is nil",
			Input:    id.NewWidgetID(),
			WA:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			_, res := tc.WA.Find(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetArea_Remove(t *testing.T) {
	wid := id.NewWidgetID()
	wa := NewWidgetArea()
	wa.widgetIds = append(wa.widgetIds, wid)

	testCases := []struct {
		Name         string
		Input        id.WidgetID
		WA, Expected *WidgetArea
	}{
		{
			Name:     "Remove a widget from widget area",
			Input:    wid,
			WA:       wa,
			Expected: &WidgetArea{widgetIds: []id.WidgetID{}},
		},
		{
			Name:     "Return nil if no widget area",
			Input:    wid,
			WA:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WA.Remove(tc.Input)
			assert.Equal(tt, tc.Expected, tc.WA)
		})
	}
}

func TestWidgetArea_WidgetIDs(t *testing.T) {
	wid := id.NewWidgetID()

	wa := NewWidgetArea()
	wa.widgetIds = append(wa.widgetIds, wid)

	testCases := []struct {
		Name     string
		WA       *WidgetArea
		Expected []id.WidgetID
	}{
		{
			Name:     "Return the WidgetIDs of the Widget Area",
			WA:       wa,
			Expected: wa.widgetIds,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WA.WidgetIDs()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetArea_Alignment(t *testing.T) {
	wa := NewWidgetArea()
	wa.align = "end"

	testCases := []struct {
		Name     string
		WA       *WidgetArea
		Expected *WidgetAlignType
	}{
		{
			Name:     "Return the alignment of the Widget Area",
			WA:       wa,
			Expected: &wa.align,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WA.Alignment()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
