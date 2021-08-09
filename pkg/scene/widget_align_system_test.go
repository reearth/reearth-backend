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
		{
			Name: "Return nil if align system is nil",
			Input: struct {
				id  id.WidgetID
				id2 id.WidgetID
				loc WidgetLocation
			}{wid, wid2, loc},
			WAS:      nil,
			Expected: nil,
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

func TestWidgetAlignSystem_AddAll(t *testing.T) {
	wid := id.NewWidgetID()
	wid2 := id.NewWidgetID()
	wids := []id.WidgetID{wid, wid2}
	a := "center"
	loc := WidgetLocation{
		Zone:    WidgetZoneOuter,
		Section: WidgetSectionLeft,
		Area:    WidgetAreaTop,
	}
	was := NewWidgetAlignSystem()
	was2 := NewWidgetAlignSystem()
	was2.outer.left.top.widgetIds = append(was2.outer.left.top.widgetIds, wid)
	was2.outer.left.top.widgetIds = append(was2.outer.left.top.widgetIds, wid2)
	was2.outer.left.top.align = "center"

	testCases := []struct {
		Name  string
		Input struct {
			ids   []id.WidgetID
			align string
			loc   WidgetLocation
		}
		WAS, Expected *WidgetAlignSystem
	}{
		{
			Name: "Add a widget to widget align system",
			Input: struct {
				ids   []id.WidgetID
				align string
				loc   WidgetLocation
			}{wids, a, loc},
			WAS:      was,
			Expected: was2,
		},
		{
			Name: "Return nil if align system is nil",
			Input: struct {
				ids   []id.WidgetID
				align string
				loc   WidgetLocation
			}{wids, a, loc},
			WAS:      nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WAS.AddAll(tc.Input.ids, tc.Input.align, tc.Input.loc)
			assert.Equal(tt, tc.Expected, tc.WAS)
		})
	}
}

func TestWidgetAlignSystem_Remove(t *testing.T) {
	wid := id.NewWidgetID()
	was := NewWidgetAlignSystem()
	was.Add(wid, WidgetLocation{Zone: "inner", Section: "left", Area: "top"})
	oldWidgets := was.inner.left.top.widgetIds
	was2 := NewWidgetAlignSystem()
	was2.Add(wid, WidgetLocation{Zone: "inner", Section: "left", Area: "top"})
	for i, w := range oldWidgets {
		if w.ID().Equal(wid.ID()) {
			was2.inner.left.top.widgetIds = append(oldWidgets[:i], oldWidgets[i+1:]...)
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
		{
			Name: "Return nil if no align system",
			Input: struct {
				id id.WidgetID
			}{wid},
			WAS:      nil,
			Expected: nil,
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
	alignStart := "start"
	alignCenter := "centered"
	alignEnd := "end"

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
	was4.outer.right.middle.align = alignStart

	was5 := NewWidgetAlignSystem()
	was5.outer.right.middle.widgetIds = wids
	was5.outer.right.middle.align = alignCenter
	was6 := NewWidgetAlignSystem()
	was6.outer.right.middle.widgetIds = wids
	was6.outer.right.middle.align = alignCenter

	was7 := NewWidgetAlignSystem()
	was7.outer.right.middle.widgetIds = wids
	was7.outer.right.middle.align = alignEnd
	was8 := NewWidgetAlignSystem()
	was8.outer.right.middle.widgetIds = wids
	was8.outer.right.middle.align = alignEnd

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
			}{wid, nil, &i, &alignStart},
			WAS:      was3,
			Expected: was4,
		},
		{
			Name: "Change a widgets alignment to center in one location",
			Input: struct {
				id id.WidgetID
				l  *WidgetLocation
				i  *int
				a  *string
			}{wid, nil, nil, &alignCenter},
			WAS:      was5,
			Expected: was6,
		},
		{
			Name: "Change a widgets alignment to end in one location",
			Input: struct {
				id id.WidgetID
				l  *WidgetLocation
				i  *int
				a  *string
			}{wid, nil, nil, &alignEnd},
			WAS:      was7,
			Expected: was8,
		},
		{
			Name: "Return without doing anything if no align system",
			Input: struct {
				id id.WidgetID
				l  *WidgetLocation
				i  *int
				a  *string
			}{wid, nil, nil, nil},
			WAS:      nil,
			Expected: nil,
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

func TestWidgetAlignSystem_FindWidgetLocation(t *testing.T) {
	wid := id.NewWidgetID()
	wid2 := id.NewWidgetID()
	wid3 := id.NewWidgetID()

	was := NewWidgetAlignSystem()
	was.outer.left.top.widgetIds = append(was.outer.left.top.widgetIds, wid)
	was.inner.left.top.widgetIds = append(was.inner.left.top.widgetIds, wid2)
	e := was.outer.left.top
	e2 := was.inner.left.top

	testCases := []struct {
		Name  string
		Input struct {
			id id.WidgetID
		}
		WAS      *WidgetAlignSystem
		Expected *WidgetArea
	}{
		{
			Name: "Find the location of a widgetID and return the WidgetArea in the Inner Widget Zone",
			Input: struct {
				id id.WidgetID
			}{wid2},
			WAS:      was,
			Expected: &e2,
		},
		{
			Name: "Find the location of a widgetID and return the WidgetArea in the Outer Widget Zone",
			Input: struct {
				id id.WidgetID
			}{wid},
			WAS:      was,
			Expected: &e,
		},
		{
			Name: "Return nil if no align system",
			Input: struct {
				id id.WidgetID
			}{wid},
			WAS:      nil,
			Expected: nil,
		},
		{
			Name: "Return nil if nothing found",
			Input: struct {
				id id.WidgetID
			}{wid3},
			WAS:      was,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			_, res := tc.WAS.FindWidgetLocation(tc.Input.id)
			assert.Equal(tt, tc.Expected, res)
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
			Name:     "Return the Widget Zone of a Widget Align System",
			Input:    "inner",
			WAS:      was,
			Expected: &was.inner,
		},
		{
			Name:     "Return nil if inputted zone doesn't exist",
			Input:    "pinner",
			WAS:      was,
			Expected: nil,
		},
		{
			Name:     "Return nil when no Widget Align System",
			Input:    "inner",
			WAS:      nil,
			Expected: nil,
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

func TestWidgetAlignSystem_Section(t *testing.T) {
	wid := id.NewWidgetID()
	loc := WidgetLocation{
		Zone:    WidgetZoneInner,
		Section: WidgetSectionCenter,
		Area:    WidgetAreaBottom,
	}
	was := NewWidgetAlignSystem()
	was.Add(wid, loc)
	testCases := []struct {
		Name           string
		Input1, Input2 string
		WAS            *WidgetAlignSystem
		Expected       *WidgetSection
	}{
		{
			Name:     "Return the Widget Section of a Widget Align System",
			Input1:   "inner",
			Input2:   "center",
			WAS:      was,
			Expected: &was.inner.center,
		},
		{
			Name:     "Return nil if Section doesn't exist",
			Input1:   "pinner",
			Input2:   "centered",
			WAS:      was,
			Expected: nil,
		},
		{
			Name:     "Return nil when no Widget Align System",
			Input1:   "inner",
			Input2:   "left",
			WAS:      nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WAS.Section(tc.Input1, tc.Input2)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetAlignSystem_Area(t *testing.T) {
	wid := id.NewWidgetID()
	loc := WidgetLocation{
		Zone:    WidgetZoneInner,
		Section: WidgetSectionCenter,
		Area:    WidgetAreaBottom,
	}
	was := NewWidgetAlignSystem()
	was.Add(wid, loc)
	testCases := []struct {
		Name string
		Input1,
		Input2,
		Input3 string
		WAS      *WidgetAlignSystem
		Expected *WidgetArea
	}{
		{
			Name:     "Return the Widget Area of a Widget Align System",
			Input1:   "inner",
			Input2:   "center",
			Input3:   "bottom",
			WAS:      was,
			Expected: &was.inner.center.bottom,
		},
		{
			Name:     "Return nil if Area doesn't exist",
			Input1:   "icnner",
			Input2:   "ceenter",
			Input3:   "bottoms",
			WAS:      was,
			Expected: nil,
		},
		{
			Name:     "Return nil when no Widget Align System",
			Input1:   "inner",
			Input2:   "center",
			Input3:   "bottom",
			WAS:      nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WAS.Area(tc.Input1, tc.Input2, tc.Input3)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
