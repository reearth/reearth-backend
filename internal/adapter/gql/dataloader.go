package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
)

//go:generate go run github.com/vektah/dataloaden DatasetLoader github.com/reearth/reearth-backend/pkg/id.DatasetID *github.com/reearth/reearth-backend/internal/adapter/gql.Dataset
//go:generate go run github.com/vektah/dataloaden DatasetSchemaLoader github.com/reearth/reearth-backend/pkg/id.DatasetSchemaID *github.com/reearth/reearth-backend/internal/adapter/gql.DatasetSchema
//go:generate go run github.com/vektah/dataloaden LayerLoader github.com/reearth/reearth-backend/pkg/id.LayerID *github.com/reearth/reearth-backend/internal/adapter/gql.Layer
//go:generate go run github.com/vektah/dataloaden LayerGroupLoader github.com/reearth/reearth-backend/pkg/id.LayerID *github.com/reearth/reearth-backend/internal/adapter/gql.LayerGroup
//go:generate go run github.com/vektah/dataloaden LayerItemLoader github.com/reearth/reearth-backend/pkg/id.LayerID *github.com/reearth/reearth-backend/internal/adapter/gql.LayerItem
//go:generate go run github.com/vektah/dataloaden PluginLoader github.com/reearth/reearth-backend/pkg/id.PluginID *github.com/reearth/reearth-backend/internal/adapter/gql.Plugin
//go:generate go run github.com/vektah/dataloaden ProjectLoader github.com/reearth/reearth-backend/pkg/id.ProjectID *github.com/reearth/reearth-backend/internal/adapter/gql.Project
//go:generate go run github.com/vektah/dataloaden PropertyLoader github.com/reearth/reearth-backend/pkg/id.PropertyID *github.com/reearth/reearth-backend/internal/adapter/gql.Property
//go:generate go run github.com/vektah/dataloaden PropertySchemaLoader github.com/reearth/reearth-backend/pkg/id.PropertySchemaID *github.com/reearth/reearth-backend/internal/adapter/gql.PropertySchema
//go:generate go run github.com/vektah/dataloaden SceneLoader github.com/reearth/reearth-backend/pkg/id.SceneID *github.com/reearth/reearth-backend/internal/adapter/gql.Scene
//go:generate go run github.com/vektah/dataloaden TeamLoader github.com/reearth/reearth-backend/pkg/id.TeamID *github.com/reearth/reearth-backend/internal/adapter/gql.Team
//go:generate go run github.com/vektah/dataloaden UserLoader github.com/reearth/reearth-backend/pkg/id.UserID *github.com/reearth/reearth-backend/internal/adapter/gql.User

//go:generate go run github.com/reearth/reearth-backend/tools/cmd/gen -template=loader.tmpl -output=loader_gen.go -m=Dataset -m=Layer -m=Plugin -m=Project -m=Property -m=Scene -m=Team -m=User
//go:generate go run github.com/reearth/reearth-backend/tools/cmd/gen -template=loader.tmpl -output=loader_layer_item_gen.go -controller=Layer -method=FetchItem -id=LayerID -m=LayerItem
//go:generate go run github.com/reearth/reearth-backend/tools/cmd/gen -template=loader.tmpl -output=loader_layer_group_gen.go -controller=Layer -method=FetchGroup -id=LayerID -m=LayerGroup
//go:generate go run github.com/reearth/reearth-backend/tools/cmd/gen -template=loader.tmpl -output=loader_dataset_schema_gen.go -controller=Dataset -method=FetchSchema -m=DatasetSchema
//go:generate go run github.com/reearth/reearth-backend/tools/cmd/gen -template=loader.tmpl -output=loader_property_schema_gen.go -controller=Property -method=FetchSchema -m=PropertySchema

type dataLoadersKey struct{}

type DataLoaders struct {
	Dataset        DatasetDataLoader
	DatasetSchema  DatasetSchemaDataLoader
	LayerItem      LayerItemDataLoader
	LayerGroup     LayerGroupDataLoader
	Layer          LayerDataLoader
	Plugin         PluginDataLoader
	Project        ProjectDataLoader
	Property       PropertyDataLoader
	PropertySchema PropertySchemaDataLoader
	Scene          SceneDataLoader
	Team           TeamDataLoader
	User           UserDataLoader
}

func DataLoadersFromContext(ctx context.Context) *DataLoaders {
	return ctx.Value(dataLoadersKey{}).(*DataLoaders)
}

func DataLoadersKey() interface{} {
	return dataLoadersKey{}
}

func NewDataLoaders(ctx context.Context, c *Container, o *usecase.Operator) *DataLoaders {
	return &DataLoaders{
		Dataset:        newDataset(ctx, c.DatasetController, o),
		DatasetSchema:  newDatasetSchema(ctx, c.DatasetController, o),
		LayerItem:      newLayerItem(ctx, c.LayerController, o),
		LayerGroup:     newLayerGroup(ctx, c.LayerController, o),
		Layer:          newLayer(ctx, c.LayerController, o),
		Plugin:         newPlugin(ctx, c.PluginController, o),
		Project:        newProject(ctx, c.ProjectController, o),
		Property:       newProperty(ctx, c.PropertyController, o),
		PropertySchema: newPropertySchema(ctx, c.PropertyController, o),
		Scene:          newScene(ctx, c.SceneController, o),
		Team:           newTeam(ctx, c.TeamController, o),
		User:           newUser(ctx, c.UserController, o),
	}
}

func NewOrdinaryDataLoaders(ctx context.Context, c *Container, o *usecase.Operator) *DataLoaders {
	return &DataLoaders{
		Dataset:        newOrdinaryDataset(ctx, c.DatasetController, o),
		DatasetSchema:  newOrdinaryDatasetSchema(ctx, c.DatasetController, o),
		LayerItem:      newOrdinaryLayerItem(ctx, c.LayerController, o),
		LayerGroup:     newOrdinaryLayerGroup(ctx, c.LayerController, o),
		Layer:          newOrdinaryLayer(ctx, c.LayerController, o),
		Plugin:         newOrdinaryPlugin(ctx, c.PluginController, o),
		Project:        newOrdinaryProject(ctx, c.ProjectController, o),
		Property:       newOrdinaryProperty(ctx, c.PropertyController, o),
		PropertySchema: newOrdinaryPropertySchema(ctx, c.PropertyController, o),
		Scene:          newOrdinaryScene(ctx, c.SceneController, o),
		Team:           newOrdinaryTeam(ctx, c.TeamController, o),
		User:           newOrdinaryUser(ctx, c.UserController, o),
	}
}
