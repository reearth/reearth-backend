package interactor

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
	"github.com/reearth/reearth-backend/pkg/plugin/pluginpack"
	"github.com/reearth/reearth-backend/pkg/plugin/repourl"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/reearth/reearth-backend/pkg/scene"
)

var pluginPackageSizeLimit int64 = 10 * 1024 * 1024 // 10MB

func (i *Plugin) Upload(ctx context.Context, r io.Reader, sid id.SceneID, operator *usecase.Operator) (_ *plugin.Plugin, _ *scene.Scene, err error) {
	if err := i.CanWriteScene(ctx, sid, operator); err != nil {
		return nil, nil, err
	}

	p, err := pluginpack.PackageFromZip(r, &sid, pluginPackageSizeLimit)
	if err != nil {
		return nil, nil, &rerror.Error{
			Label:    interfaces.ErrInvalidPluginPackage,
			Err:      err,
			Separate: true,
		}
	}

	return i.upload(ctx, p, sid, operator)
}

func (i *Plugin) UploadFromRemote(ctx context.Context, u *url.URL, sid id.SceneID, operator *usecase.Operator) (_ *plugin.Plugin, _ *scene.Scene, err error) {
	if err := i.CanWriteScene(ctx, sid, operator); err != nil {
		return nil, nil, err
	}

	ru, err := repourl.New(u)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ru.ArchiveURL().String(), nil)
	if err != nil {
		return nil, nil, interfaces.ErrInvalidPluginPackage
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, interfaces.ErrInvalidPluginPackage
	}
	if res.StatusCode != 200 {
		return nil, nil, interfaces.ErrInvalidPluginPackage
	}

	p, err := pluginpack.PackageFromZip(res.Body, &sid, pluginPackageSizeLimit)
	if err != nil {
		_ = res.Body.Close()
		return nil, nil, interfaces.ErrInvalidPluginPackage
	}

	_ = res.Body.Close()
	return i.upload(ctx, p, sid, operator)
}

func (i *Plugin) upload(ctx context.Context, p *pluginpack.Package, sid id.SceneID, operator *usecase.Operator) (_ *plugin.Plugin, _ *scene.Scene, err error) {
	if err := i.CanWriteScene(ctx, sid, operator); err != nil {
		return nil, nil, err
	}

	s, err := i.sceneRepo.FindByID(ctx, sid, operator.WritableTeams)
	if err != nil {
		return nil, nil, err
	}

	tx, err := i.transaction.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	newpid := p.Manifest.Plugin.ID()
	oldpid := s.Plugins().PluginByName(newpid.Name()).PluginRef()
	var oldp *plugin.Plugin
	if oldpid != nil {
		oldp, err = i.pluginRepo.FindByID(ctx, *oldpid, []id.SceneID{sid})
		if err != nil {
			return nil, nil, err
		}
	}

	// new (oldpid == nil): upload files, save plugin and properties -> install
	// same (oldpid.Equal(newpid)): delete old files -> upload files, save plugin and property schemas -> migrate
	// diff (!oldpid.Equal(newpid)): upload files, save plugin and property schemas -> migrate -> delete old files

	if oldpid != nil && oldpid.Equal(newpid) {
		// same only: delete old files
		if err := i.file.RemovePlugin(ctx, *oldpid); err != nil {
			return nil, nil, err
		}
	}

	// uploads files
	for {
		f, err := p.Files.Next()
		if err != nil {
			return nil, nil, interfaces.ErrInvalidPluginPackage
		}
		if f == nil {
			break
		}
		if err := i.file.UploadPluginFile(ctx, p.Manifest.Plugin.ID(), f); err != nil {
			return nil, nil, rerror.ErrInternalBy(err)
		}
	}

	// save plugin and property schemas
	if ps := p.Manifest.PropertySchemas(); len(ps) > 0 {
		if err := i.propertySchemaRepo.SaveAll(ctx, ps); err != nil {
			return nil, nil, err
		}
	}

	if err := i.pluginRepo.Save(ctx, p.Manifest.Plugin); err != nil {
		return nil, nil, err
	}

	if oldpid == nil {
		// new: install plugin
		if err := i.installScenePlugin(ctx, p, s); err != nil {
			return nil, nil, err
		}
	} else {
		// same, diff: migrate
		if err := i.migrateScenePlugin(ctx, p, s, oldp); err != nil {
			return nil, nil, err
		}
	}

	if oldpid != nil && !oldpid.Equal(newpid) {
		// diff only: delete old files
		if err := i.file.RemovePlugin(ctx, *oldpid); err != nil {
			return nil, nil, err
		}

		if oldpid.Scene() != nil {
			// remove old scene plugin
			if err := i.pluginRepo.Remove(ctx, *oldpid); err != nil {
				return nil, nil, err
			}
			if ps := oldp.PropertySchemas(); len(ps) > 0 {
				if err := i.propertySchemaRepo.RemoveAll(ctx, ps); err != nil {
					return nil, nil, err
				}
			}
		}
	}

	tx.Commit()
	return p.Manifest.Plugin, s, nil
}

// installScenePlugin installs the plugin to the scene
func (i *Plugin) installScenePlugin(ctx context.Context, p *pluginpack.Package, s *scene.Scene) (err error) {
	var ppid *id.PropertyID
	var pp *property.Property
	if psid := p.Manifest.Plugin.Schema(); psid != nil {
		pp, err = property.New().NewID().Schema(*psid).Build()
		if err != nil {
			return err
		}
	}

	s.Plugins().Add(scene.NewPlugin(p.Manifest.Plugin.ID(), ppid))

	if pp != nil {
		if err := i.propertyRepo.Save(ctx, pp); err != nil {
			return err
		}
	}
	if err := i.sceneRepo.Save(ctx, s); err != nil {
		return err
	}
	return nil
}

func (i *Plugin) migrateScenePlugin(ctx context.Context, p *pluginpack.Package, s *scene.Scene, oldp *plugin.Plugin) (err error) {
	if oldp == nil || p.Manifest == nil {
		return nil
	}

	oldPManifest, err := i.pluginManifestFromPlugin(ctx, oldp)
	if err != nil {
		return err
	}

	diff := manifest.DiffFrom(oldPManifest, *p.Manifest)
	updatedProperties := property.List{}

	// update scene
	var spp *id.PropertyID
	if to := diff.PropertySchemaDiff.To; !to.IsNil() && diff.PropertySchemaDiff.From.IsNil() {
		// new plugin property
		p, err := property.New().NewID().Scene(s.ID()).Schema(to).Build()
		if err != nil {
			return err
		}
		spp = p.ID().Ref()
		updatedProperties = append(updatedProperties, p)
	}

	if sp := s.Plugins().Plugin(diff.From); sp != nil && sp.Property() != nil && diff.PropertySchemaDeleted {
		// plugin property should be removed
		if err := i.propertyRepo.Remove(ctx, *sp.Property()); err != nil {
			return err
		}
	}

	s.Plugins().Upgrade(diff.From, diff.To, spp, diff.PropertySchemaDeleted)

	// delete layers, blocks and widgets
	for _, e := range diff.DeletedExtensions {
		deletedProperties, err := i.deleteLayersByPluginExtension(ctx, s.ID().Ref(), diff.From, &e.ExtensionID)
		if err != nil {
			return err
		}

		if deletedProperties2, err := i.deleteBlocksByPluginExtension(ctx, s.ID().Ref(), diff.From, &e.ExtensionID); err != nil {
			return err
		} else {
			deletedProperties = append(deletedProperties, deletedProperties2...)
		}

		deletedProperties = append(deletedProperties, s.Widgets().RemoveAllByPlugin(diff.From, e.ExtensionID.Ref())...)

		if len(deletedProperties) > 0 {
			if err := i.propertyRepo.RemoveAll(ctx, deletedProperties); err != nil {
				return err
			}
		}
	}

	if err := i.sceneRepo.Save(ctx, s); err != nil {
		return err
	}

	// migrate layers
	if err := i.layerRepo.UpdatePlugin(ctx, diff.From, diff.To, []id.SceneID{s.ID()}); err != nil {
		return err
	}

	// migrate properties
	updatedPropertySchemas := diff.PropertySchmaDiffs()
	updatedPropertySchemaIDs := updatedPropertySchemas.FromSchemas()
	pl, err := i.propertyRepo.FindBySchema(ctx, updatedPropertySchemaIDs, s.ID())
	if err != nil {
		return err
	}
	for _, p := range pl {
		if e := updatedPropertySchemas.FindByFrom(p.Schema()); e != nil && e.Migrate(p) {
			updatedProperties = append(updatedProperties, p)
		}
	}
	if len(updatedProperties) > 0 {
		if err := i.propertyRepo.SaveAll(ctx, updatedProperties); err != nil {
			return err
		}
	}

	if err := i.propertyRepo.UpdateSchemaPlugin(ctx, diff.From, diff.To, s.ID()); err != nil {
		return err
	}

	// delete unused schemas and properties
	if deleted := diff.DeletedPropertySchemas(); len(deleted) > 0 {
		if err := i.propertySchemaRepo.RemoveAll(ctx, deleted); err != nil {
			return err
		}
	}

	return nil
}

func (i *Plugin) deleteLayersByPluginExtension(ctx context.Context, sid *id.SceneID, p id.PluginID, e *id.PluginExtensionID) ([]id.PropertyID, error) {
	var scenes []id.SceneID
	if sid != nil {
		scenes = []id.SceneID{*sid}
	}

	// delete layers
	deletedLayers := []id.LayerID{}
	layers, err := i.layerRepo.FindByPluginAndExtension(ctx, p, e, scenes)
	if err != nil {
		return nil, err
	}
	deletedLayers = append(deletedLayers, layers.IDs().Layers()...)

	parentLayers, err := i.layerRepo.FindParentsByIDs(ctx, deletedLayers, scenes)
	if err != nil {
		return nil, err
	}

	for _, p := range parentLayers {
		p.Layers().RemoveLayer(deletedLayers...)
	}
	if err := i.layerRepo.SaveAll(ctx, parentLayers.ToLayerList()); err != nil {
		return nil, err
	}
	if err := i.layerRepo.RemoveAll(ctx, deletedLayers); err != nil {
		return nil, err
	}

	return layers.Properties(), nil
}

func (i *Plugin) deleteBlocksByPluginExtension(ctx context.Context, sid *id.SceneID, p id.PluginID, e *id.PluginExtensionID) ([]id.PropertyID, error) {
	var scenes []id.SceneID
	if sid != nil {
		scenes = []id.SceneID{*sid}
	}

	layers, err := i.layerRepo.FindByPluginAndExtensionOfBlocks(ctx, p, e, scenes)
	if err != nil {
		return nil, err
	}

	var deletedProperties []id.PropertyID
	for _, l := range layers.Deref() {
		deletedProperties = append(deletedProperties, l.Infobox().RemoveAllByPlugin(p, e)...)
	}

	if err := i.layerRepo.SaveAll(ctx, layers); err != nil {
		return nil, err
	}
	return deletedProperties, nil
}

func (i *Plugin) pluginManifestFromPlugin(ctx context.Context, p *plugin.Plugin) (manifest.Manifest, error) {
	schemas, err := i.propertySchemaRepo.FindByIDs(ctx, p.PropertySchemas())
	if err != nil {
		return manifest.Manifest{}, err
	}

	var s *property.Schema
	if ps := p.Schema(); ps != nil {
		s = schemas.Find(*ps)
	}

	return manifest.Manifest{
		Plugin:          p,
		ExtensionSchema: schemas,
		Schema:          s,
	}, nil
}
