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

func NewDataLoaders(ctx context.Context, c Container, o *usecase.Operator) DataLoaders {
	return DataLoaders{
		Dataset:        newDataset(ctx, c.Dataset, o),
		DatasetSchema:  newDatasetSchema(ctx, c.Dataset, o),
		LayerItem:      newLayerItem(ctx, c.Layer, o),
		LayerGroup:     newLayerGroup(ctx, c.Layer, o),
		Layer:          newLayer(ctx, c.Layer, o),
		Plugin:         newPlugin(ctx, c.Plugin, o),
		Project:        newProject(ctx, c.Project, o),
		Property:       newProperty(ctx, c.Property, o),
		PropertySchema: newPropertySchema(ctx, c.Property, o),
		Scene:          newScene(ctx, c.Scene, o),
		Team:           newTeam(ctx, c.Team, o),
		User:           newUser(ctx, c.User, o),
	}
}

func NewOrdinaryDataLoaders(ctx context.Context, c Container, o *usecase.Operator) DataLoaders {
	return DataLoaders{
		Dataset:        newOrdinaryDataset(ctx, c.Dataset, o),
		DatasetSchema:  newOrdinaryDatasetSchema(ctx, c.Dataset, o),
		LayerItem:      newOrdinaryLayerItem(ctx, c.Layer, o),
		LayerGroup:     newOrdinaryLayerGroup(ctx, c.Layer, o),
		Layer:          newOrdinaryLayer(ctx, c.Layer, o),
		Plugin:         newOrdinaryPlugin(ctx, c.Plugin, o),
		Project:        newOrdinaryProject(ctx, c.Project, o),
		Property:       newOrdinaryProperty(ctx, c.Property, o),
		PropertySchema: newOrdinaryPropertySchema(ctx, c.Property, o),
		Scene:          newOrdinaryScene(ctx, c.Scene, o),
		Team:           newOrdinaryTeam(ctx, c.Team, o),
		User:           newOrdinaryUser(ctx, c.User, o),
	}
}
