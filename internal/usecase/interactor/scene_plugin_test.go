package interactor

import (
	"context"
	"testing"

	"github.com/reearth/reearth-backend/internal/infrastructure/memory"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/scene"
	"github.com/stretchr/testify/assert"
)

func TestScene_InstallPlugin(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	tid := id.NewTeamID()
	sid := scene.NewID()
	pid := plugin.MustID("plugin~1.0.0")
	pid2 := plugin.MustID("plugin~1.0.1")
	pid3 := plugin.MustID("plugin~1.0.1").WithScene(&sid)
	pid4 := plugin.MustID("plugin~1.0.1").WithScene(scene.NewID().Ref())
	sc := scene.New().ID(sid).RootLayer(id.NewLayerID()).Team(tid).MustBuild()
	pl := plugin.New().ID(pid).MustBuild()
	pl2 := plugin.New().ID(pid3).MustBuild()
	pl3 := plugin.New().ID(pid4).MustBuild()

	sr := memory.NewSceneWith(sc)
	pr := memory.NewPluginWith(pl, pl2, pl3)
	uc := &Scene{
		sceneRepo:   sr,
		pluginRepo:  pr,
		transaction: memory.NewTransaction(),
	}

	// normal case 1
	gotSc, gotPid, gotPrid, err := uc.InstallPlugin(ctx, sid, pid, &usecase.Operator{
		WritableTeams: id.TeamIDList{tid},
	})
	assert.NoError(err)
	assert.Same(sc, gotSc)
	assert.Equal(pid, gotPid)
	assert.True(gotPrid.IsNil())
	assert.True(gotSc.Plugins().Has(pid))

	// normal case 2: scene specific plugin
	sc.Plugins().Remove(pid)
	gotSc, gotPid, gotPrid, err = uc.InstallPlugin(ctx, sid, pid3, &usecase.Operator{
		WritableTeams: id.TeamIDList{tid},
	})
	assert.NoError(err)
	assert.Same(gotSc, sc)
	assert.Equal(gotPid, pid3)
	assert.True(gotPrid.IsNil())
	assert.True(gotSc.Plugins().Has(pid3))

	// abnormal case 1: plugin not found
	gotSc, gotPid, gotPrid, err = uc.InstallPlugin(ctx, sid, pid2, &usecase.Operator{
		WritableTeams: id.TeamIDList{tid},
	})
	assert.Equal(interfaces.ErrPluginNotFound, err)
	assert.Nil(gotSc)
	assert.Equal(pid2, gotPid)
	assert.True(gotPrid.IsNil())

	// abnormal case 2: already installed
	sc.Plugins().Remove(pid3)
	sc.Plugins().Add(scene.NewPlugin(pid, nil))
	gotSc, gotPid, gotPrid, err = uc.InstallPlugin(ctx, sid, pid, &usecase.Operator{
		WritableTeams: id.TeamIDList{tid},
	})
	assert.Equal(interfaces.ErrPluginAlreadyInstalled, err)
	assert.Nil(gotSc)
	assert.Equal(pid, gotPid)
	assert.True(gotPrid.IsNil())

	// abnormal case 3: plugin scene is different
	gotSc, gotPid, gotPrid, err = uc.InstallPlugin(ctx, sid, pid4, &usecase.Operator{
		WritableTeams: id.TeamIDList{tid},
	})
	assert.Equal(interfaces.ErrPluginNotFound, err)
	assert.Nil(gotSc)
	assert.Equal(pid4, gotPid)
	assert.True(gotPrid.IsNil())
}

func TestScene_UpgradePlugin(t *testing.T) {

}

func TestScene_UninstallPlugin(t *testing.T) {

}
