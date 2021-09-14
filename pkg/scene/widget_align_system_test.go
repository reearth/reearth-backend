package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewWidgetAlignSystem(t *testing.T) {
	assert.Equal(t, &WidgetAlignSystem{}, NewWidgetAlignSystem())
}

func TestWidgetAlignSystem_Zone(t *testing.T) {
	was := NewWidgetAlignSystem()
	assert.Same(t, was.inner, was.Zone(WidgetZoneInner))
	assert.NotNil(t, was.inner)
	assert.Same(t, was.outer, was.Zone(WidgetZoneOuter))
	assert.NotNil(t, was.outer)
}

func TestWidgetAlignSystem_Area(t *testing.T) {
	was := NewWidgetAlignSystem()
	assert.Same(t, was.inner.right.middle, was.Area(WidgetLocation{
		Zone:    WidgetZoneInner,
		Section: WidgetSectionRight,
		Area:    WidgetAreaMiddle,
	}))
}

func TestWidgetAlignSystem_Find(t *testing.T) {
	wid1 := id.NewWidgetID()
	wid2 := id.NewWidgetID()
	wid3 := id.NewWidgetID()
	wid4 := id.NewWidgetID()
	wid5 := id.NewWidgetID()

	testCases := []struct {
		Name      string
		Input     id.WidgetID
		Expected1 int
		Expected2 WidgetLocation
		Nil       bool
	}{
		{
			Name:      "inner",
			Input:     wid2,
			Expected1: 1,
			Expected2: WidgetLocation{Zone: WidgetZoneInner, Section: WidgetSectionLeft, Area: WidgetAreaTop},
		},
		{
			Name:      "outer",
			Input:     wid4,
			Expected1: 0,
			Expected2: WidgetLocation{Zone: WidgetZoneOuter, Section: WidgetSectionLeft, Area: WidgetAreaTop},
		},
		{
			Name:      "invalid id",
			Input:     id.NewWidgetID(),
			Expected1: -1,
			Expected2: WidgetLocation{},
		},
		{
			Name:      "Return nil if no widget section",
			Input:     wid1,
			Nil:       true,
			Expected1: -1,
			Expected2: WidgetLocation{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			if tc.Nil {
				index, location := (*WidgetAlignSystem)(nil).Find(tc.Input)
				assert.Equal(tt, tc.Expected1, index)
				assert.Equal(tt, tc.Expected2, location)
				return
			}

			was := NewWidgetAlignSystem()
			was.Zone(WidgetZoneInner).Section(WidgetSectionLeft).Area(WidgetAreaTop).AddAll([]id.WidgetID{wid1, wid2, wid3})
			was.Zone(WidgetZoneOuter).Section(WidgetSectionLeft).Area(WidgetAreaTop).AddAll([]id.WidgetID{wid4, wid5})

			index, location := was.Find(tc.Input)
			assert.Equal(tt, tc.Expected1, index)
			assert.Equal(tt, tc.Expected2, location)
		})
	}
}

func TestWidgetAlignSystem_Remove(t *testing.T) {
	wid := id.NewWidgetID()

	testCases := []struct {
		Name     string
		Zone     WidgetZoneType
		Input    id.WidgetID
		Expected []id.WidgetID
		Nil      bool
	}{
		{
			Name:     "inner: remove a widget from widget section",
			Zone:     WidgetZoneInner,
			Input:    wid,
			Expected: []id.WidgetID{},
		},
		{
			Name:     "inner: couldn't find widgetId",
			Zone:     WidgetZoneInner,
			Input:    id.NewWidgetID(),
			Expected: []id.WidgetID{wid},
		},
		{
			Name:     "outer: remove a widget from widget section",
			Zone:     WidgetZoneOuter,
			Input:    wid,
			Expected: []id.WidgetID{},
		},
		{
			Name:     "outer: couldn't find widgetId",
			Zone:     WidgetZoneOuter,
			Input:    id.NewWidgetID(),
			Expected: []id.WidgetID{wid},
		},
		{
			Name:  "nil",
			Zone:  WidgetZoneInner,
			Input: wid,
			Nil:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			if tc.Nil {
				(*WidgetZone)(nil).Remove(tc.Input)
				return
			}

			ws := NewWidgetAlignSystem()
			ws.Zone(tc.Zone).Section(WidgetSectionLeft).Area(WidgetAreaTop).Add(wid)
			ws.Remove(tc.Input)
			assert.Equal(tt, tc.Expected, ws.Zone(tc.Zone).Section(WidgetSectionLeft).Area(WidgetAreaTop).WidgetIDs())
		})
	}
}
