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
	type args struct {
		pluginID plugin.ID
		operator *usecase.Operator
	}

	type test struct {
		name                  string
		installedScenePlugins []*scene.Plugin
		args                  args
		wantErr               error
	}

	sid := scene.NewID()
	pid := plugin.MustID("plugin~1.0.0")
	pid2 := plugin.MustID("plugin~1.0.1")
	pid3 := plugin.MustID("plugin~1.0.1").WithScene(&sid)
	pid4 := plugin.MustID("plugin~1.0.1").WithScene(scene.NewID().Ref())

	tests := []test{
		{
			name: "should install a plugin",
			args: args{
				pluginID: pid,
			},
		},
		{
			name: "should install a private plugin with property schema",
			args: args{
				pluginID: pid3,
			},
		},
		{
			name: "already installed",
			installedScenePlugins: []*scene.Plugin{
				scene.NewPlugin(pid, nil),
			},
			args: args{
				pluginID: pid,
			},
			wantErr: interfaces.ErrPluginAlreadyInstalled,
		},
		{
			name: "not found",
			args: args{
				pluginID: pid2,
			},
			wantErr: interfaces.ErrPluginNotFound,
		},
		{
			name: "diff scene",
			args: args{
				pluginID: pid4,
			},
			wantErr: interfaces.ErrPluginNotFound,
		},
		{
			name: "operation denied",
			args: args{
				operator: &usecase.Operator{},
			},
			wantErr: interfaces.ErrOperationDenied,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			ctx := context.Background()

			tid := id.NewTeamID()
			sc := scene.New().ID(sid).RootLayer(id.NewLayerID()).Team(tid).MustBuild()
			for _, p := range tt.installedScenePlugins {
				sc.Plugins().Add(p)
			}
			sr := memory.NewSceneWith(sc)

			pl := plugin.New().ID(pid).MustBuild()
			pl2 := plugin.New().ID(pid3).Schema(id.NewPropertySchemaID(pid3, "@").Ref()).MustBuild()
			pl3 := plugin.New().ID(pid4).MustBuild()
			pr := memory.NewPluginWith(pl, pl2, pl3)

			prr := memory.NewProperty()

			uc := &Scene{
				sceneRepo:    sr,
				pluginRepo:   pr,
				propertyRepo: prr,
				transaction:  memory.NewTransaction(),
			}

			o := tt.args.operator
			if o == nil {
				o = &usecase.Operator{
					WritableTeams: id.TeamIDList{tid},
				}
			}
			gotSc, gotPid, gotPrid, err := uc.InstallPlugin(ctx, sid, tt.args.pluginID, o)

			assert.Equal(tt.args.pluginID, gotPid)
			if tt.wantErr != nil {
				assert.Equal(tt.wantErr, err)
				assert.Nil(gotSc)
				assert.True(gotPrid.IsNil())
			} else {
				assert.NoError(err)
				assert.Same(sc, gotSc)
				if tt.args.pluginID.Equal(pl2.ID()) {
					assert.False(gotPid.IsNil())
					gotPr, _ := prr.FindByID(ctx, *gotPrid)
					assert.Equal(*pl2.Schema(), gotPr.Schema())
				} else {
					assert.True(gotPrid.IsNil())
				}
				assert.True(gotSc.Plugins().Has(tt.args.pluginID))
			}
		})
	}
}

func TestScene_UpgradePlugin(t *testing.T) {

}

func TestScene_UninstallPlugin(t *testing.T) {

}
