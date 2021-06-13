package sceneops

import (
	"context"
	"io"

	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
)

type PluginInstaller struct {
	PluginRepo           repo.Plugin
	PluginRepositoryRepo gateway.PluginRepository
	PropertySchemaRepo   repo.PropertySchema
}

func (s PluginInstaller) InstallPluginFromRepository(ctx context.Context, source io.Reader) error {

	m, err := manifest.Parse(source)
	if err != nil {
		return err
	}

	// save
	if m.Plugin != nil {
		err = s.PluginRepo.Save(ctx, m.Plugin)
		if err != nil {
			return err
		}
	}

	if m.Schema != nil {
		err = s.PropertySchemaRepo.Save(ctx, m.Schema)
		if err != nil {
			return err
		}
	}

	if m.ExtensionSchema != nil && len(m.ExtensionSchema) > 0 {
		err = s.PropertySchemaRepo.SaveAll(ctx, m.ExtensionSchema)
		if err != nil {
			return err
		}
	}

	// 	// Download and extract plugin files to storage
	// 	data, err := i.pluginRepositoryRepo.Data(inp.Name, inp.Version)
	// 	if err != nil {
	// 		i.output.Upload(nil, err1.ErrInternalBy(err))
	// 		return
	// 	}

	// 	_, err = i.fileRepo.UploadAndExtractPluginFiles(data, plugin)
	// 	if err != nil {
	// 		i.output.Upload(nil, err1.ErrInternalBy(err))
	// 		return
	// 	}

	// 	return nil
	// }

	// // UploadPlugin _
	// func (s PluginInstaller) UploadPlugin(reader io.Reader) error {
	// 	panic("not implemented")

	// 	manifest, err := s.PluginRepositoryRepo.Manifest(inp.Name, inp.Version)
	// 	if err != nil {
	// 		i.output.Upload(nil, err)
	// 		return
	// 	}

	// 	// build plugin
	// 	plugin, err := plugin.New().
	// 		NewID().
	// 		FromManifest(manifest).
	// 		Developer(operator.User).
	// 		PluginSeries(pluginSeries.ID()).
	// 		CreatedAt(time.Now()).
	// 		Public(inp.Public).
	// 		Build()
	// 	if err != nil {
	// 		i.output.Upload(nil, err1.ErrInternalBy(err))
	// 		return
	// 	}

	// 	// save
	// 	if manifest.Schema != nil {
	// 		err = i.propertySchemaRepo.Save(manifest.Schema)
	// 		if err != nil {
	// 			i.output.Upload(nil, err1.ErrInternalBy(err))
	// 			return
	// 		}
	// 	}

	// 	for _, s := range manifest.ExtensionSchema {
	// 		err = i.propertySchemaRepo.Save(&s)
	// 		if err != nil {
	// 			i.output.Upload(nil, err1.ErrInternalBy(err))
	// 			return
	// 		}
	// 	}

	// 	err = i.pluginRepo.Save(plugin)
	// 	if err != nil {
	// 		i.output.Upload(nil, err1.ErrInternalBy(err))
	// 		return
	// 	}

	// 	// Download and extract plugin files to storage
	// 	data, err := i.pluginRepositoryRepo.Data(inp.Name, inp.Version)
	// 	if err != nil {
	// 		i.output.Upload(nil, err1.ErrInternalBy(err))
	// 		return
	// 	}

	// 	_, err = i.fileRepo.UploadAndExtractPluginFiles(data, plugin)
	// 	if err != nil {
	// 		i.output.Upload(nil, err1.ErrInternalBy(err))
	// 		return
	// 	}

	return nil
}
