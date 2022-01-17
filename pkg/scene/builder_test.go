package scene

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuilder_IDs(t *testing.T) {
	tid := NewTeamID()
	lid := NewLayerID()
	b := New().NewID().RootLayer(lid).Team(tid).MustBuild()
	assert.NotNil(t, b.ID())
	assert.Equal(t, tid, b.Team())
	assert.Equal(t, lid, b.RootLayer())
	sid := NewID()
	b2 := New().ID(sid).RootLayer(lid).Team(tid).MustBuild()
	assert.Equal(t, sid, b2.ID())
}

func TestBuilder_UpdatedAt(t *testing.T) {
	ti := time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC)
	b := New().NewID().RootLayer(NewLayerID()).Team(NewTeamID()).UpdatedAt(ti).MustBuild()
	assert.Equal(t, ti, b.UpdatedAt())
}

func TestBuilder_Property(t *testing.T) {
	pid := NewPropertyID()
	b := New().NewID().RootLayer(NewLayerID()).Team(NewTeamID()).Property(pid).MustBuild()
	assert.Equal(t, pid, b.Property())
}

func TestBuilder_Plugins(t *testing.T) {
	ps := NewPlugins([]*Plugin{
		NewPlugin(OfficialPluginID, NewPropertyID().Ref()),
	})
	b := New().NewID().RootLayer(NewLayerID()).Team(NewTeamID()).Plugins(ps).MustBuild()
	assert.Equal(t, ps, b.Plugins())
}

func TestBuilder_Project(t *testing.T) {
	pid := NewProjectID()
	b := New().NewID().RootLayer(NewLayerID()).Team(NewTeamID()).Project(pid).MustBuild()
	assert.Equal(t, pid, b.Project())
}

func TestBuilder_Widgets(t *testing.T) {
	ws := NewWidgets([]*Widget{
		MustNewWidget(NewWidgetID(), OfficialPluginID, "xxx", NewPropertyID(), true, false),
	})
	b := New().NewID().RootLayer(NewLayerID()).Team(NewTeamID()).Widgets(ws).MustBuild()
	assert.Equal(t, ws, b.Widgets())
}
func TestBuilder_WidgetAlignSystem(t *testing.T) {
	was := NewWidgetAlignSystem()
	b := New().NewID().RootLayer(NewLayerID()).Team(NewTeamID()).WidgetAlignSystem(was).MustBuild()
	assert.Equal(t, was, b.WidgetAlignSystem())
}

func TestBuilder_Build(t *testing.T) {
	tid := NewTeamID()
	sid := NewID()
	pid := NewProjectID()
	ppid := NewPropertyID()
	lid := NewLayerID()
	ws := NewWidgets([]*Widget{
		MustNewWidget(NewWidgetID(), OfficialPluginID, "xxx", ppid, true, false),
	})
	was := NewWidgetAlignSystem()
	ps := NewPlugins([]*Plugin{
		NewPlugin(OfficialPluginID, ppid.Ref()),
	})
	testCases := []struct {
		Name              string
		Id                ID
		Project           ProjectID
		Team              TeamID
		RootLayer         LayerID
		Widgets           *Widgets
		WidgetAlignSystem *WidgetAlignSystem
		Plugins           *Plugins
		UpdatedAt         time.Time
		Property          PropertyID
		Expected          struct {
			Id                ID
			Project           ProjectID
			Team              TeamID
			RootLayer         LayerID
			Widgets           *Widgets
			WidgetAlignSystem *WidgetAlignSystem
			Plugins           *Plugins
			UpdatedAt         time.Time
			Property          PropertyID
		}
		err error
	}{
		{
			Name:              "fail nil scene id",
			Id:                ID{},
			Project:           pid,
			Team:              tid,
			RootLayer:         lid,
			Widgets:           ws,
			WidgetAlignSystem: was,
			Plugins:           ps,
			UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
			Property:          ppid,
			err:               ErrInvalidID,
		},
		{
			Name:              "fail nil team id",
			Id:                sid,
			Project:           pid,
			Team:              TeamID{},
			RootLayer:         lid,
			Widgets:           ws,
			WidgetAlignSystem: was,
			Plugins:           ps,
			UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
			Property:          ppid,
			err:               ErrInvalidID,
		},
		{
			Name:              "fail nil root layer id",
			Id:                sid,
			Project:           pid,
			Team:              tid,
			RootLayer:         LayerID{},
			Widgets:           ws,
			WidgetAlignSystem: was,
			Plugins:           ps,
			UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
			Property:          ppid,
			err:               ErrInvalidID,
		},
		{
			Name:              "success build new scene",
			Id:                sid,
			Project:           pid,
			Team:              tid,
			RootLayer:         lid,
			Widgets:           ws,
			WidgetAlignSystem: was,
			Plugins:           ps,
			UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
			Property:          ppid,
			Expected: struct {
				Id                ID
				Project           ProjectID
				Team              TeamID
				RootLayer         LayerID
				Widgets           *Widgets
				WidgetAlignSystem *WidgetAlignSystem
				Plugins           *Plugins
				UpdatedAt         time.Time
				Property          PropertyID
			}{
				Id:                sid,
				Project:           pid,
				Team:              tid,
				RootLayer:         lid,
				Widgets:           ws,
				WidgetAlignSystem: was,
				Plugins:           ps,
				UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
				Property:          ppid,
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res, err := New().
				ID(tc.Id).
				Widgets(tc.Widgets).
				WidgetAlignSystem(tc.WidgetAlignSystem).
				Project(tc.Project).
				Plugins(tc.Plugins).
				Property(tc.Property).
				RootLayer(tc.RootLayer).
				Team(tc.Team).
				UpdatedAt(tc.UpdatedAt).
				Build()
			if err == nil {
				assert.Equal(tt, tc.Expected.Id, res.ID())
				assert.Equal(tt, tc.Expected.UpdatedAt, res.UpdatedAt())
				assert.Equal(tt, tc.Expected.Team, res.Team())
				assert.Equal(tt, tc.Expected.RootLayer, res.RootLayer())
				assert.Equal(tt, tc.Expected.Property, res.Property())
				assert.Equal(tt, tc.Expected.Plugins, res.Plugins())
				assert.Equal(tt, tc.Expected.Widgets, res.Widgets())
				assert.Equal(tt, tc.Expected.Project, res.Project())
			} else {
				assert.True(tt, errors.As(tc.err, &err))
			}
		})
	}
}

func TestBuilder_MustBuild(t *testing.T) {
	tid := NewTeamID()
	sid := NewID()
	pid := NewProjectID()
	ppid := NewPropertyID()
	lid := NewLayerID()
	ws := NewWidgets([]*Widget{
		MustNewWidget(NewWidgetID(), OfficialPluginID, "xxx", ppid, true, false),
	})
	was := NewWidgetAlignSystem()
	ps := NewPlugins([]*Plugin{
		NewPlugin(OfficialPluginID, ppid.Ref()),
	})
	testCases := []struct {
		Name              string
		Id                ID
		Project           ProjectID
		Team              TeamID
		RootLayer         LayerID
		Widgets           *Widgets
		WidgetAlignSystem *WidgetAlignSystem
		Plugins           *Plugins
		UpdatedAt         time.Time
		Property          PropertyID
		Expected          struct {
			Id                ID
			Project           ProjectID
			Team              TeamID
			RootLayer         LayerID
			Widgets           *Widgets
			WidgetAlignSystem *WidgetAlignSystem
			Plugins           *Plugins
			UpdatedAt         time.Time
			Property          PropertyID
		}
		err error
	}{
		{
			Name:              "fail nil scene id",
			Id:                ID{},
			Project:           pid,
			Team:              tid,
			RootLayer:         lid,
			Widgets:           ws,
			WidgetAlignSystem: was,
			Plugins:           ps,
			UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
			Property:          ppid,
			err:               ErrInvalidID,
		},
		{
			Name:              "fail nil team id",
			Id:                sid,
			Project:           pid,
			Team:              TeamID{},
			RootLayer:         lid,
			Widgets:           ws,
			WidgetAlignSystem: was,
			Plugins:           ps,
			UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
			Property:          ppid,
			err:               ErrInvalidID,
		},
		{
			Name:              "fail nil root layer id",
			Id:                sid,
			Project:           pid,
			Team:              tid,
			RootLayer:         LayerID{},
			Widgets:           ws,
			WidgetAlignSystem: was,
			Plugins:           ps,
			UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
			Property:          ppid,
			err:               ErrInvalidID,
		},
		{
			Name:              "success build new scene",
			Id:                sid,
			Project:           pid,
			Team:              tid,
			RootLayer:         lid,
			Widgets:           ws,
			WidgetAlignSystem: was,
			Plugins:           ps,
			UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
			Property:          ppid,
			Expected: struct {
				Id                ID
				Project           ProjectID
				Team              TeamID
				RootLayer         LayerID
				Widgets           *Widgets
				WidgetAlignSystem *WidgetAlignSystem
				Plugins           *Plugins
				UpdatedAt         time.Time
				Property          PropertyID
			}{
				Id:                sid,
				Project:           pid,
				Team:              tid,
				RootLayer:         lid,
				Widgets:           ws,
				WidgetAlignSystem: was,
				Plugins:           ps,
				UpdatedAt:         time.Date(2000, 1, 1, 1, 1, 0, 0, time.UTC),
				Property:          ppid,
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			var res *Scene
			defer func() {
				if r := recover(); r == nil {
					assert.Equal(tt, tc.Expected.Id, res.ID())
					assert.Equal(tt, tc.Expected.UpdatedAt, res.UpdatedAt())
					assert.Equal(tt, tc.Expected.Team, res.Team())
					assert.Equal(tt, tc.Expected.RootLayer, res.RootLayer())
					assert.Equal(tt, tc.Expected.Property, res.Property())
					assert.Equal(tt, tc.Expected.Plugins, res.Plugins())
					assert.Equal(tt, tc.Expected.Widgets, res.Widgets())
					assert.Equal(tt, tc.Expected.WidgetAlignSystem, res.WidgetAlignSystem())
					assert.Equal(tt, tc.Expected.Project, res.Project())
				}
			}()

			res = New().
				ID(tc.Id).
				Widgets(tc.Widgets).
				WidgetAlignSystem(tc.WidgetAlignSystem).
				Project(tc.Project).
				Plugins(tc.Plugins).
				Property(tc.Property).
				RootLayer(tc.RootLayer).
				Team(tc.Team).
				UpdatedAt(tc.UpdatedAt).
				MustBuild()
		})
	}
}
