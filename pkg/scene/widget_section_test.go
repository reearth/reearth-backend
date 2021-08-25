package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestWidgetSection_Find(t *testing.T) {
	wid := id.NewWidgetID()

	ws := NewWidgetSection()
	ws.top.widgetIds = append(ws.top.widgetIds, wid)
	e := ws.top

	testCases := []struct {
		Name     string
		Input    id.WidgetID
		WS       *WidgetSection
		Expected *WidgetArea
	}{
		{
			Name:     "Find the location of a widgetID and return the WidgetArea",
			Input:    wid,
			WS:       ws,
			Expected: &e,
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

func TestWidgetSection_Area(t *testing.T) {
	wid := id.NewWidgetID()

	ws := NewWidgetSection()
	ws.top.widgetIds = append(ws.top.widgetIds, wid)
	e := ws.top

	testCases := []struct {
		Name  string
		Input struct {
			s string
		}
		WS       *WidgetSection
		Expected *WidgetArea
	}{
		{
			Name: "From a Widget Section return a specific Widget Area",
			Input: struct {
				s string
			}{"top"},
			WS:       ws,
			Expected: &e,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WS.Area(tc.Input.s)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
