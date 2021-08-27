package gql

import (
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
)

type Container struct {
	usecases interfaces.Container
	Asset    *AssetController
	Dataset  *DatasetController
	Layer    *LayerController
	Plugin   *PluginController
	Project  *ProjectController
	Property *PropertyController
	Scene    *SceneController
	Team     *TeamController
	User     *UserController
}

type ContainerConfig struct {
	SignupSecret string
}

func NewContainer(usecases interfaces.Container) Container {
	return Container{
		usecases: usecases,
		Asset:    NewAssetController(usecases.Asset),
		Dataset:  NewDatasetController(usecases.Dataset),
		Layer:    NewLayerController(usecases.Layer),
		Plugin:   NewPluginController(usecases.Plugin),
		Project:  NewProjectController(usecases.Project),
		Property: NewPropertyController(usecases.Property),
		Scene:    NewSceneController(usecases.Scene),
		Team:     NewTeamController(usecases.Team),
		User:     NewUserController(usecases.User),
	}
}
