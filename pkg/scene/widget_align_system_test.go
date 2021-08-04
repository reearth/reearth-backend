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
	wid2 := id.NewWidgetID()
	loc := WidgetLocation{
		Zone:    WidgetZoneOuter,
		Section: WidgetSectionLeft,
		Area:    WidgetAreaTop,
	}
	was := NewWidgetAlignSystem()
	was2 := NewWidgetAlignSystem()
	was2.outer.left.top.widgetIds = append(was2.outer.left.top.widgetIds, wid)
	was2.outer.left.top.widgetIds = append(was2.outer.left.top.widgetIds, wid2)
	was2.outer.left.top.align = "start"

	testCases := []struct {
		Name  string
		Input struct {
			id  id.WidgetID
			id2 id.WidgetID
			loc WidgetLocation
		}
		WAS, Expected *WidgetAlignSystem
	}{
		{
			Name: "Add a widget to widget align system",
			Input: struct {
				id  id.WidgetID
				id2 id.WidgetID
				loc WidgetLocation
			}{wid, wid2, loc},
			WAS:      was,
			Expected: was2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WAS.Add(tc.Input.id, tc.Input.loc)
			tc.WAS.Add(tc.Input.id2, tc.Input.loc)
			assert.Equal(tt, tc.Expected, tc.WAS)
		})
	}
}

func TestWidgetAlignSystem_Remove(t *testing.T) {
	wid := id.NewWidgetID()
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
			id id.WidgetID
		}
		WAS, Expected *WidgetAlignSystem
	}{
		{
			Name: "Remove a widget from widget align system",
			Input: struct {
				id id.WidgetID
			}{wid},
			WAS:      was,
			Expected: was2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WAS.Remove(tc.Input.id)
			assert.Equal(tt, tc.Expected, tc.WAS)
		})
	}
}

func TestWidgetAlignSystem_Update(t *testing.T) {
	wid := id.NewWidgetID()
	align := "start"

	// for move
	oloc := WidgetLocation{"outer", "right", "middle"}
	nloc := WidgetLocation{"inner", "left", "top"}
	was := NewWidgetAlignSystem()
	was.Add(wid, oloc)

	was2 := NewWidgetAlignSystem()
	was2.Add(wid, oloc)
	was2.Remove(wid)
	was2.Add(wid, nloc)

	// for reordering
	i := 0
	wid2 := id.NewWidgetID()
	wid3 := id.NewWidgetID()
	wids := []id.WidgetID{wid2, wid3, wid}
	nwids := []id.WidgetID{wid, wid2, wid3}

	was3 := NewWidgetAlignSystem()
	was3.outer.right.middle.widgetIds = wids
	was4 := NewWidgetAlignSystem()
	was4.outer.right.middle.widgetIds = nwids
	was4.outer.right.middle.align = align

	testCases := []struct {
		Name  string
		Input struct {
			id id.WidgetID
			l  *WidgetLocation
			i  *int
			a  *string
		}
		WAS, Expected *WidgetAlignSystem
	}{
		{
			Name: "Move a widget from one location to another",
			Input: struct {
				id id.WidgetID
				l  *WidgetLocation
				i  *int
				a  *string
			}{wid, &nloc, nil, nil},
			WAS:      was,
			Expected: was2,
		},
		{
			Name: "Reorder a widget in one location",
			Input: struct {
				id id.WidgetID
				l  *WidgetLocation
				i  *int
				a  *string
			}{wid, nil, &i, &align},
			WAS:      was3,
			Expected: was4,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WAS.Update(tc.Input.id, tc.Input.l, tc.Input.i, tc.Input.a)
			assert.Equal(tt, tc.Expected, tc.WAS)
		})
	}
}

func TestWidgetAlignSystem_Zone(t *testing.T) {
	wid := id.NewWidgetID()
	loc := WidgetLocation{
		Zone:    WidgetZoneInner,
		Section: WidgetSectionCenter,
		Area:    WidgetAreaBottom,
	}
	was := NewWidgetAlignSystem()
	was.Add(wid, loc)
	testCases := []struct {
		Name     string
		Input    string
		WAS      *WidgetAlignSystem
		Expected *WidgetZone
	}{
		{
			Input:    "inner",
			WAS:      was,
			Expected: &was.inner,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WAS.Zone(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

// func TestWidgetAlignSystem_Section(t *testing.T) {
// 	wid := id.NewWidgetID()
// 	loc := WidgetLocation{
// 		Zone:    WidgetZoneInner,
// 		Section: WidgetSectionCenter,
// 		Area:    WidgetAreaBottom,
// 	}
// 	was := NewWidgetAlignSystem()
// 	was.Add(wid, &loc)
// 	testCases := []struct {
// 		Name     string
// 		Input
// 		WAS      *WidgetAlignSystem
// 		Expected *WidgetSection
// 	}{
// 		{
// 			Input:    "inner",
// 			WAS:      was,
// 			Expected: &was.inner.center,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		tc := tc
// 		t.Run(tc.Name, func(tt *testing.T) {
// 			tt.Parallel()
// 			res := tc.WAS.Section(tc.Input)
// 			assert.Equal(tt, tc.Expected, res)
// 		})
// 	}
// }
