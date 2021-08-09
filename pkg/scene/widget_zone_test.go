package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestWidgetZone_Find(t *testing.T) {
	wid := id.NewWidgetID()

	was := NewWidgetAlignSystem()
	was.outer.left.top.widgetIds = append(was.outer.left.top.widgetIds, wid)
	e := was.outer.left.top

	testCases := []struct {
		Name  string
		Input struct {
			id id.WidgetID
		}
		WAS      *WidgetZone
		Expected *WidgetArea
	}{
		{
			Name: "Find the location of a widgetID and return the WidgetArea",
			Input: struct {
				id id.WidgetID
			}{wid},
			WAS:      &was.outer,
			Expected: &e,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			_, res := tc.WAS.Find(tc.Input.id)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetZone_Section(t *testing.T) {
	wid := id.NewWidgetID()

	was := NewWidgetAlignSystem()
	was.outer.left.top.widgetIds = append(was.outer.left.top.widgetIds, wid)
	e := was.outer.left

	testCases := []struct {
		Name  string
		Input struct {
			s string
		}
		WAS      *WidgetZone
		Expected *WidgetSection
	}{
		{
			Name: "Find the location of a widgetID and return the WidgetArea",
			Input: struct {
				s string
			}{"left"},
			WAS:      &was.outer,
			Expected: &e,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WAS.Section(tc.Input.s)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
