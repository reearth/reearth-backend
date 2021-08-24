package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestWidgetSection_Find(t *testing.T) {
	wid := id.NewWidgetID()

	was := NewWidgetAlignSystem()
	was.outer.left.top.widgetIds = append(was.outer.left.top.widgetIds, wid)
	e := was.outer.left.top

	testCases := []struct {
		Name     string
		Input    id.WidgetID
		WAS      *WidgetSection
		Expected *WidgetArea
	}{
		{
			Name:     "Find the location of a widgetID and return the WidgetArea",
			Input:    wid,
			WAS:      &was.outer.left,
			Expected: &e,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			_, res := tc.WAS.Find(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetSection_Area(t *testing.T) {
	wid := id.NewWidgetID()

	was := NewWidgetAlignSystem()
	was.outer.left.top.widgetIds = append(was.outer.left.top.widgetIds, wid)
	e := was.outer.left.top

	testCases := []struct {
		Name  string
		Input struct {
			s string
		}
		WAS      *WidgetSection
		Expected *WidgetArea
	}{
		{
			Name: "From a Widget Section return a specific Widget Area",
			Input: struct {
				s string
			}{"top"},
			WAS:      &was.outer.left,
			Expected: &e,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WAS.Area(tc.Input.s)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
