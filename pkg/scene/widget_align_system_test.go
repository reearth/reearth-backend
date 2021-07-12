package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewWidgetAlignSystem(t *testing.T) {
	res := NewWidgetAlignSystem()
	expected := &WidgetAlignSystem{}
	assert.Equal(t, expected, res)
}

func TestWidgetAlignSystem(t *testing.T) {
	was := NewWidgetAlignSystem()
	expected := &WidgetAlignSystem{}
	testCases := []struct {
		Name             string
		Expected, Actual *WidgetAlignSystem
	}{
		{
			Expected: expected,
			Actual:   was,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.Expected.inner, tc.Actual.inner)
			assert.Equal(tt, tc.Expected.inner.center, tc.Actual.inner.center)
			assert.Equal(tt, tc.Expected.inner.left.bottom, tc.Actual.inner.left.bottom)
		})
	}
}

func TestWidgetAlignSystem_Add(t *testing.T) {
	wid := id.NewWidgetID()
	loc := Location{"inner", "left", "top"}
	was := NewWidgetAlignSystem()
	was2 := NewWidgetAlignSystem()
	was2.inner.left.top.widgetIds = append(was2.inner.left.top.widgetIds, wid)

	testCases := []struct {
		Name  string
		Input struct {
			id  *id.WidgetID
			loc *Location
		}
		WAS, Expected *WidgetAlignSystem
	}{
		{
			Name: "Add a widget to widget align system",
			Input: struct {
				id  *id.WidgetID
				loc *Location
			}{&wid, &loc},
			WAS:      was,
			Expected: was2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WAS.Add(tc.Input.id, tc.Input.loc)
			assert.Equal(tt, tc.Expected, tc.WAS)
		})
	}
}

func TestWidgetAlignSystem_Remove(t *testing.T) {
	wid := id.NewWidgetID()
	loc := Location{"inner", "left", "top"}
	was := NewWidgetAlignSystem()
	oldWidgets := was.inner.left.top.widgetIds
	was2 := NewWidgetAlignSystem()
	for i, w := range oldWidgets {
		if w.ID().Equal(wid.ID()) {
			was2.inner.left.top.widgetIds = append(oldWidgets[:i], oldWidgets[i+1])
		}
	}
	testCases := []struct {
		Name  string
		Input struct {
			id  *id.WidgetID
			loc *Location
		}
		WAS, Expected *WidgetAlignSystem
	}{
		{
			Name: "Remove a widget from widget align system",
			Input: struct {
				id  *id.WidgetID
				loc *Location
			}{&wid, &loc},
			WAS:      was,
			Expected: was2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WAS.Remove(tc.Input.id, tc.Input.loc)
			assert.Equal(tt, tc.Expected, tc.WAS)
		})
	}
}

func TestWidgetAlignSystem_Move(t *testing.T) {
	wid := id.NewWidgetID()
	oloc := Location{"outer", "right", "middle"}
	nloc := Location{"inner", "left", "top"}

	was := NewWidgetAlignSystem()
	was.Add(&wid, &oloc)

	was2 := NewWidgetAlignSystem()
	was2.Add(&wid, &oloc)
	was2.Remove(&wid, &oloc)
	was2.Add(&wid, &nloc)

	testCases := []struct {
		Name  string
		Input struct {
			id *id.WidgetID
			ol *Location
			nl *Location
		}
		WAS, Expected *WidgetAlignSystem
	}{
		{
			Name: "Move a widget from one location to another",
			Input: struct {
				id *id.WidgetID
				ol *Location
				nl *Location
			}{&wid, &oloc, &nloc},
			WAS:      was,
			Expected: was2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WAS.Move(tc.Input.id, tc.Input.ol, tc.Input.nl)
			assert.Equal(tt, tc.Expected, tc.WAS)
		})
	}
}
func TestWidgetAlignSystem_Reorder(t *testing.T) {
	wid := id.NewWidgetID()
	wid2 := id.NewWidgetID()
	wid3 := id.NewWidgetID()
	wids := []id.WidgetID{wid2, wid3, wid}
	nwids := []id.WidgetID{wid, wid3, wid2}
	loc := Location{"outer", "right", "middle"}
	was := NewWidgetAlignSystem()
	was.outer.right.middle.widgetIds = wids

	was2 := NewWidgetAlignSystem()
	was2.outer.right.middle.widgetIds = wids
	was2.outer.right.middle.widgetIds = nwids

	testCases := []struct {
		Name  string
		Input struct {
			id    *id.WidgetID
			l     *Location
			nwids []id.WidgetID
		}
		WAS, Expected *WidgetAlignSystem
	}{
		{
			Name: "Move a widget from one location to another",
			Input: struct {
				id    *id.WidgetID
				l     *Location
				nwids []id.WidgetID
			}{&wid, &loc, nwids},
			WAS:      was,
			Expected: was2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WAS.Reorder(tc.Input.id, tc.Input.l, tc.Input.nwids)
			assert.Equal(tt, tc.Expected, tc.WAS)
		})
	}
}
