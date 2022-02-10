package interactor

import (
	"archive/zip"
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/reearth/reearth-backend/internal/infrastructure/fs"
	"github.com/reearth/reearth-backend/internal/infrastructure/memory"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/reearth/reearth-backend/pkg/scene"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

const mockPluginManifest = `{
  "id": "testplugin",
  "version": "1.0.1",
  "name": "testplugin",
  "extensions": [
	  {
      "id": "block",
			"type": "block",
			"schema": {
				"groups": []
			}
		}
  ]
}`

var mockPluginID = id.MustPluginID("testplugin~1.0.1")
var mockPluginFiles = map[string]string{
	"reearth.yml": mockPluginManifest,
	"block.js":    "// barfoo",
}
var mockPluginArchiveZip bytes.Buffer

func init() {
	zbuf := bytes.Buffer{}
	zw := zip.NewWriter(&zbuf)
	for p, f := range mockPluginFiles {
		w, _ := zw.Create(p)
		_, _ = w.Write([]byte(f))
	}
	_ = zw.Close()
	mockPluginArchiveZip = zbuf
}

func mockFS(files map[string]string) afero.Fs {
	mfs := afero.NewMemMapFs()
	for n, c := range files {
		f, err := mfs.Create(n)
		if err != nil {
			panic(err)
		}
		_, _ = f.Write([]byte(c))
		_ = f.Close()
	}
	return mfs
}

func TestPlugin_Upload_New(t *testing.T) {
	// upload a new plugin
	ctx := context.Background()
	team := id.NewTeamID()
	sid := id.NewSceneID()
	pid := mockPluginID.WithScene(sid.Ref())

	repos := memory.InitRepos(nil)
	mfs := mockFS(nil)
	files, err := fs.NewFile(mfs, "")
	assert.NoError(t, err)
	scene := scene.New().ID(sid).Team(team).RootLayer(id.NewLayerID()).MustBuild()
	_ = repos.Scene.Save(ctx, scene)

	uc := &Plugin{
		commonScene:        commonScene{sceneRepo: repos.Scene},
		pluginRepo:         repos.Plugin,
		propertySchemaRepo: repos.PropertySchema,
		propertyRepo:       repos.Property,
		layerRepo:          repos.Layer,
		file:               files,
		transaction:        repos.Transaction,
	}
	op := &usecase.Operator{
		ReadableTeams: []id.TeamID{team},
		WritableTeams: []id.TeamID{team},
	}

	reader := bytes.NewReader(mockPluginArchiveZip.Bytes())
	pl, s, err := uc.Upload(ctx, reader, scene.ID(), op)
	assert.NoError(t, err)
	assert.Equal(t, scene.ID(), s.ID())
	assert.Equal(t, pid, pl.ID())

	// scene
	nscene, err := repos.Scene.FindByID(ctx, scene.ID(), nil)
	assert.NoError(t, err)
	assert.True(t, nscene.Plugins().HasPlugin(pl.ID()))

	// plugin
	npl, err := repos.Plugin.FindByID(ctx, pid, []id.SceneID{scene.ID()})
	assert.NoError(t, err)
	assert.Equal(t, pid, npl.ID())

	npf, err := mfs.Open("plugins/" + pid.String() + "/block.js")
	assert.NoError(t, err)
	npfc, _ := io.ReadAll(npf)
	assert.Equal(t, "// barfoo", string(npfc))
}

func TestPlugin_Upload_SameVersion(t *testing.T) {
	// upgrade plugin to the same version
	// 1 extension is deleted -> property schema, layers, and properties of the extension should be deleted
	// old plugin files should be deleted

	ctx := context.Background()
	team := id.NewTeamID()
	sid := id.NewSceneID()
	pid := mockPluginID.WithScene(sid.Ref())
	eid := id.PluginExtensionID("marker")

	repos := memory.InitRepos(nil)
	mfs := mockFS(map[string]string{
		"plugins/" + pid.String() + "/hogehoge": "foobar",
	})
	files, err := fs.NewFile(mfs, "")
	assert.NoError(t, err)

	ps := property.NewSchema().ID(property.MustSchemaIDFromExtension(pid, eid)).MustBuild()
	pl := plugin.New().ID(pid).Extensions([]*plugin.Extension{
		plugin.NewExtension().ID(eid).Type(plugin.ExtensionTypePrimitive).Schema(ps.ID()).MustBuild(),
	}).MustBuild()

	p := property.New().NewID().Schema(ps.ID()).Scene(sid).MustBuild()
	pluginLayer := layer.NewItem().NewID().Scene(sid).Plugin(pid.Ref()).Extension(eid.Ref()).Property(p.IDRef()).MustBuild()
	rootLayer := layer.NewGroup().NewID().Scene(sid).Layers(layer.NewIDList([]layer.ID{pluginLayer.ID()})).Root(true).MustBuild()
	scene := scene.New().ID(sid).Team(team).RootLayer(rootLayer.ID()).Plugins(scene.NewPlugins([]*scene.Plugin{
		scene.NewPlugin(pid, nil),
	})).MustBuild()

	_ = repos.PropertySchema.Save(ctx, ps)
	_ = repos.Plugin.Save(ctx, pl)
	_ = repos.Property.Save(ctx, p)
	_ = repos.Layer.SaveAll(ctx, layer.List{pluginLayer.LayerRef(), rootLayer.LayerRef()})
	_ = repos.Scene.Save(ctx, scene)

	uc := &Plugin{
		commonScene:        commonScene{sceneRepo: repos.Scene},
		pluginRepo:         repos.Plugin,
		propertySchemaRepo: repos.PropertySchema,
		propertyRepo:       repos.Property,
		layerRepo:          repos.Layer,
		file:               files,
		transaction:        repos.Transaction,
	}
	op := &usecase.Operator{
		ReadableTeams: []id.TeamID{team},
		WritableTeams: []id.TeamID{team},
	}

	reader := bytes.NewReader(mockPluginArchiveZip.Bytes())
	pl, s, err := uc.Upload(ctx, reader, scene.ID(), op)

	assert.NoError(t, err)
	assert.Equal(t, scene.ID(), s.ID())
	assert.Equal(t, pid, pl.ID())

	// scene
	nscene, err := repos.Scene.FindByID(ctx, scene.ID(), nil)
	assert.NoError(t, err)
	assert.True(t, nscene.Plugins().HasPlugin(pl.ID()))

	// plugin
	npl, err := repos.Plugin.FindByID(ctx, pid, []id.SceneID{scene.ID()})
	assert.NoError(t, err)
	assert.Equal(t, pid, npl.ID())

	nlps, err := repos.PropertySchema.FindByID(ctx, ps.ID())
	assert.Nil(t, nlps) // deleted
	assert.Equal(t, rerror.ErrNotFound, err)

	_, err = mfs.Open("plugins/" + pid.String() + "/hogehoge")
	assert.True(t, os.IsNotExist(err)) // deleted

	npf, err := mfs.Open("plugins/" + pid.String() + "/block.js")
	assert.NoError(t, err)
	npfc, _ := io.ReadAll(npf)
	assert.Equal(t, "// barfoo", string(npfc))

	// layer
	nlp, err := repos.Property.FindByID(ctx, p.ID(), nil)
	assert.Nil(t, nlp) // deleted
	assert.Equal(t, rerror.ErrNotFound, err)

	nl, err := repos.Layer.FindByID(ctx, pluginLayer.ID(), nil)
	assert.Nil(t, nl) // deleted
	assert.Equal(t, rerror.ErrNotFound, err)

	nrl, err := repos.Layer.FindGroupByID(ctx, rootLayer.ID(), nil)
	assert.NoError(t, err)
	assert.Equal(t, []layer.ID{}, nrl.Layers().Layers()) // deleted
}

func TestPlugin_Upload_DiffVersion(t *testing.T) {
	// upgrade plugin to the different version
	// plugin ID of property and layers should be updated

	ctx := context.Background()
	team := id.NewTeamID()
	sid := id.NewSceneID()
	oldpid := id.MustPluginID("testplugin~1.0.0").WithScene(sid.Ref())
	pid := mockPluginID.WithScene(sid.Ref())
	eid := id.PluginExtensionID("block")
	nlpsid := id.MustPropertySchemaIDFromExtension(pid, eid)

	repos := memory.InitRepos(nil)
	mfs := mockFS(map[string]string{
		"plugins/" + oldpid.String() + "/hogehoge": "foobar",
	})
	files, err := fs.NewFile(mfs, "")
	assert.NoError(t, err)

	oldps := property.NewSchema().ID(property.MustSchemaIDFromExtension(oldpid, eid)).MustBuild()
	oldpl := plugin.New().ID(oldpid).Extensions([]*plugin.Extension{
		plugin.NewExtension().ID(eid).Type(plugin.ExtensionTypeBlock).Schema(oldps.ID()).MustBuild(),
	}).MustBuild()

	oldp := property.New().NewID().Schema(oldps.ID()).Scene(sid).MustBuild()
	pluginLayer := layer.NewItem().NewID().Scene(sid).Plugin(oldpid.Ref()).Extension(eid.Ref()).Property(oldp.IDRef()).MustBuild()
	rootLayer := layer.NewGroup().NewID().Scene(sid).Layers(layer.NewIDList([]layer.ID{pluginLayer.ID()})).Root(true).MustBuild()
	scene := scene.New().ID(sid).Team(team).RootLayer(rootLayer.ID()).Plugins(scene.NewPlugins([]*scene.Plugin{
		scene.NewPlugin(oldpid, nil),
	})).MustBuild()

	_ = repos.PropertySchema.Save(ctx, oldps)
	_ = repos.Plugin.Save(ctx, oldpl)
	_ = repos.Property.Save(ctx, oldp)
	_ = repos.Layer.SaveAll(ctx, layer.List{pluginLayer.LayerRef(), rootLayer.LayerRef()})
	_ = repos.Scene.Save(ctx, scene)

	uc := &Plugin{
		commonScene:        commonScene{sceneRepo: repos.Scene},
		pluginRepo:         repos.Plugin,
		propertySchemaRepo: repos.PropertySchema,
		propertyRepo:       repos.Property,
		layerRepo:          repos.Layer,
		file:               files,
		transaction:        repos.Transaction,
	}
	op := &usecase.Operator{
		ReadableTeams: []id.TeamID{team},
		WritableTeams: []id.TeamID{team},
	}

	reader := bytes.NewReader(mockPluginArchiveZip.Bytes())
	oldpl, s, err := uc.Upload(ctx, reader, scene.ID(), op)

	assert.NoError(t, err)
	assert.Equal(t, scene.ID(), s.ID())
	assert.Equal(t, pid, oldpl.ID())

	// scene
	nscene, err := repos.Scene.FindByID(ctx, scene.ID(), nil)
	assert.NoError(t, err)
	assert.False(t, nscene.Plugins().HasPlugin(oldpid))
	assert.True(t, nscene.Plugins().HasPlugin(pid))

	// plugin
	opl, err := repos.Plugin.FindByID(ctx, oldpid, []id.SceneID{scene.ID()})
	assert.Nil(t, opl) // deleted
	assert.Equal(t, rerror.ErrNotFound, err)

	npl, err := repos.Plugin.FindByID(ctx, pid, []id.SceneID{scene.ID()})
	assert.NoError(t, err)
	assert.Equal(t, pid, npl.ID())

	olps, err := repos.PropertySchema.FindByID(ctx, oldps.ID())
	assert.Nil(t, olps) // deleted
	assert.Equal(t, rerror.ErrNotFound, err)

	nlps, err := repos.PropertySchema.FindByID(ctx, nlpsid)
	assert.NoError(t, err)
	assert.Equal(t, nlpsid, nlps.ID())

	_, err = mfs.Open("plugins/" + oldpid.String() + "/hogehoge")
	assert.True(t, os.IsNotExist(err)) // deleted

	npf, err := mfs.Open("plugins/" + pid.String() + "/block.js")
	assert.NoError(t, err)
	npfc, _ := io.ReadAll(npf)
	assert.Equal(t, "// barfoo", string(npfc))

	// layer
	nl, err := repos.Layer.FindByID(ctx, pluginLayer.ID(), nil)
	assert.NoError(t, err)
	assert.Equal(t, pid, *nl.Plugin())
	assert.Equal(t, eid, *nl.Extension())
	assert.Equal(t, oldp.ID(), *nl.Property())

	nlp, err := repos.Property.FindByID(ctx, *nl.Property(), nil)
	assert.NoError(t, err)
	assert.Equal(t, *nl.Property(), nlp.ID())
	assert.Equal(t, nlpsid, nlp.Schema())
}
