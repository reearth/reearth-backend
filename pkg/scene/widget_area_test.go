package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestWidgetArea_Find(t *testing.T) {
	wid := id.NewWidgetID()

	testCases := []struct {
		Name     string
		WA       *WidgetArea
		Input    id.WidgetID
		Expected int
	}{
		{
			Name:     "Return WidgetArea if contains widget id",
			WA:       NewWidgetArea([]id.WidgetID{wid}, WidgetAlignStart),
			Input:    wid,
			Expected: 0,
		},
		{
			Name:     "Return nil if doesn't contain widget id",
			WA:       NewWidgetArea([]id.WidgetID{}, WidgetAlignStart),
			Input:    wid,
			Expected: -1,
		},
		{
			Name:     "Return nil if WidgetArea is nil",
			WA:       nil,
			Input:    wid,
			Expected: -1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			index := tc.WA.Find(tc.Input)
			assert.Equal(tt, tc.Expected, index)
		})
	}
}

func TestWidgetArea_Add(t *testing.T) {
	wid1 := id.NewWidgetID()
	wid2 := id.NewWidgetID()

	testCases := []struct {
		Name     string
		Nil      bool
		Input    id.WidgetID
		Expected []id.WidgetID
	}{
		{
			Name:     "add a widget id",
			Input:    wid2,
			Expected: []id.WidgetID{wid1, wid2},
		},
		{
			Name:     "add a widget id but already exists",
			Input:    wid1,
			Expected: []id.WidgetID{wid1},
		},
		{
			Name: "nil widget area",
			Nil:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			if tc.Nil {
				(*WidgetArea)(nil).Add(wid1)
				return
			}

			wa := NewWidgetArea([]id.WidgetID{wid1}, WidgetAlignStart)
			wa.Add(tc.Input)
			assert.Equal(tt, tc.Expected, wa.WidgetIDs())
		})
	}
}

func TestWidgetArea_AddAll(t *testing.T) {
	wid1 := id.NewWidgetID()
	wid2 := id.NewWidgetID()

	testCases := []struct {
		Name     string
		Nil      bool
		Input    []id.WidgetID
		Expected []id.WidgetID
	}{
		{
			Name:     "add widget ids",
			Input:    []id.WidgetID{wid1, wid2},
			Expected: []id.WidgetID{wid1, wid2},
		},
		{
			Name:     "add widget ids but duplicated",
			Input:    []id.WidgetID{wid1, wid1, wid2},
			Expected: []id.WidgetID{wid1, wid2},
		},
		{
			Name: "nil widget area",
			Nil:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			if tc.Nil {
				(*WidgetArea)(nil).AddAll(nil)
				return
			}

			wa := NewWidgetArea(nil, WidgetAlignStart)
			wa.AddAll(tc.Input)
			assert.Equal(tt, tc.Expected, wa.WidgetIDs())
		})
	}
}

func TestWidgetArea_Remove(t *testing.T) {
	wid := id.NewWidgetID()
	wa := NewWidgetArea(nil, WidgetAlignType(""))
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
	wa := NewWidgetArea(nil, WidgetAlignStart)
	wa.widgetIds = append(wa.widgetIds, wid)
	res := wa.WidgetIDs()
	assert.Equal(t, wa.widgetIds, res)
}

func TestWidgetArea_Alignment(t *testing.T) {
	wa := NewWidgetArea(nil, WidgetAlignEnd)
	assert.Equal(t, WidgetAlignEnd, wa.Alignment())
}
