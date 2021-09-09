package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestWidgetSection_Find(t *testing.T) {
	wid := id.NewWidgetID()
	wid2 := id.NewWidgetID()
	wid3 := id.NewWidgetID()

	ws := NewWidgetSection()
	ws.top.widgetIds = append(ws.top.widgetIds, wid)
	ws.middle.widgetIds = append(ws.middle.widgetIds, wid2)
	ws.bottom.widgetIds = append(ws.bottom.widgetIds, wid3)
	top := ws.top
	mid := ws.middle
	bot := ws.bottom

	testCases := []struct {
		Name     string
		Input    id.WidgetID
		WS       *WidgetSection
		Expected *WidgetArea
	}{
		{
			Name:     "Find the location (top) of a widgetID and return the WidgetArea",
			Input:    wid,
			WS:       ws,
			Expected: &top,
		},
		{
			Name:     "Find the location (middle) of a widgetID and return the WidgetArea",
			Input:    wid2,
			WS:       ws,
			Expected: &mid,
		},
		{
			Name:     "Find the location (bottom) of a widgetID and return the WidgetArea",
			Input:    wid3,
			WS:       ws,
			Expected: &bot,
		},
		{
			Name:     "Return nil if no widget section",
			Input:    wid3,
			WS:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			_, res := tc.WS.Find(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetSection_Remove(t *testing.T) {
	wid := id.NewWidgetID()
	ws := NewWidgetSection()
	ws.top.widgetIds = append(ws.top.widgetIds, wid)

	testCases := []struct {
		Name         string
		Input        id.WidgetID
		WS, Expected *WidgetSection
	}{
		{
			Name:     "Remove a widget from widget section",
			Input:    wid,
			WS:       ws,
			Expected: &WidgetSection{top: WidgetArea{widgetIds: []id.WidgetID{}}},
		},
		{
			Name:     "Return nil if no widget section",
			Input:    wid,
			WS:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WS.Remove(tc.Input)
			assert.Equal(tt, tc.Expected, tc.WS)
		})
	}
}

func TestWidgetSection_Area(t *testing.T) {
	wid := id.NewWidgetID()
	wid2 := id.NewWidgetID()

	testCases := []struct {
		Name     string
		Input    WidgetAreaType
		WS       *WidgetSection
		Expected *WidgetArea
	}{
		{
			Name:  "From a Widget Section return top Widget Area",
			Input: "top",
			Expected: &WidgetArea{
				widgetIds: []id.WidgetID{wid},
				align:     "start",
			},
		},
		{
			Name:  "From a Widget Section return middle Widget Area",
			Input: "middle",
			Expected: &WidgetArea{
				widgetIds: []id.WidgetID{wid2},
				align:     "start",
			},
		},
		{
			Name:     "From a Widget Section return bottom Widget Area",
			Input:    "bottom",
			Expected: &WidgetArea{},
		},
		{
			Name:     "Return nil when can't find Widget Area",
			Input:    "topMiddleBottom",
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			ws := NewWidgetSection()
			ws.Area(WidgetAreaTop).Add(wid)
			ws.Area(WidgetAreaMiddle).Add(wid2)
			res := ws.Area(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
