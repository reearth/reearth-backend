package gql

import (
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
)

type Container struct {
	usecases interfaces.Container
	Asset    *AssetLoader
	Dataset  *DatasetLoader
	Layer    *LayerLoader
	Plugin   *PluginLoader
	Project  *ProjectLoader
	Property *PropertyLoader
	Scene    *SceneLoader
	Team     *TeamLoader
	User     *UserLoader
}

func NewContainer(usecases interfaces.Container) Container {
	return Container{
		usecases: usecases,
		Asset:    NewAssetLoader(usecases.Asset),
		Dataset:  NewDatasetLoader(usecases.Dataset),
		Layer:    NewLayerLoader(usecases.Layer),
		Plugin:   NewPluginLoader(usecases.Plugin),
		Project:  NewProjectLoader(usecases.Project),
		Property: NewPropertyLoader(usecases.Property),
		Scene:    NewSceneLoader(usecases.Scene),
		Team:     NewTeamLoader(usecases.Team),
		User:     NewUserLoader(usecases.User),
	}
}
