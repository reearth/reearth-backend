package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestWidgetArea_Find(t *testing.T) {
	wid := id.NewWidgetID()

	was := NewWidgetAlignSystem()
	was.outer.left.top.widgetIds = append(was.outer.left.top.widgetIds, wid)
	e := was.outer.left.top

	testCases := []struct {
		Name  string
		Input struct {
			id id.WidgetID
		}
		WA       *WidgetArea
		Expected *WidgetArea
	}{
		{
			Name: "Find the location of a widgetID and return the WidgetArea",
			Input: struct {
				id id.WidgetID
			}{wid},
			WA:       &was.outer.left.top,
			Expected: &e,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			_, res := tc.WA.Find(tc.Input.id)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetArea_WidgetIDs(t *testing.T) {
	wid := id.NewWidgetID()

	was := NewWidgetAlignSystem()
	was.outer.left.top.widgetIds = append(was.outer.left.top.widgetIds, wid)

	testCases := []struct {
		Name     string
		WA       *WidgetArea
		Expected []id.WidgetID
	}{
		{
			Name:     "Return the WidgetIDs of the Widget Area",
			WA:       &was.outer.left.top,
			Expected: was.outer.left.top.widgetIds,
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
	was := NewWidgetAlignSystem()
	was.outer.left.top.align = "end"

	was2 := NewWidgetAlignSystem()
	was2.Add(id.NewWidgetID(), WidgetLocation{Zone: "outer", Section: "left", Area: "top"})
	d := "start"

	testCases := []struct {
		Name     string
		WA       *WidgetArea
		Expected *string
	}{
		{
			Name:     "Return the alignment of the Widget Area",
			WA:       &was.outer.left.top,
			Expected: &was.outer.left.top.align,
		},
		{
			Name:     "Default alignment on adding widget",
			WA:       &was2.outer.left.top,
			Expected: &d,
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

func TestWidgetArea_Has(t *testing.T) {
	wid := id.NewWidgetID()
	was := NewWidgetAlignSystem()
	was.Add(wid, WidgetLocation{Zone: "outer", Section: "left", Area: "top"})

	testCases := []struct {
		Name  string
		Input struct {
			id id.WidgetID
		}
		WA       *WidgetArea
		Expected bool
	}{
		{
			Name: "Return true if Widget Area has widgetID",
			Input: struct {
				id id.WidgetID
			}{wid},
			WA:       &was.outer.left.top,
			Expected: true,
		},
		{
			Name: "Return true if Widget Area has widgetID",
			Input: struct {
				id id.WidgetID
			}{wid},
			WA:       &was.outer.left.middle,
			Expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WA.Has(wid)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
