package gql

import (
	"context"
	"time"
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

var (
	dataLoaderWait     = 1 * time.Millisecond
	dataLoaderMaxBatch = 100
)

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

func (c Container) DataLoadersWith(ctx context.Context, enabled bool) DataLoaders {
	if enabled {
		return c.DataLoaders(ctx)
	}
	return c.OrdinaryDataLoaders(ctx)
}

func (c Container) DataLoaders(ctx context.Context) DataLoaders {
	return DataLoaders{
		Dataset:        c.Dataset.DataLoader(ctx),
		DatasetSchema:  c.Dataset.SchemaDataLoader(ctx),
		LayerItem:      c.Layer.ItemDataLoader(ctx),
		LayerGroup:     c.Layer.GroupDataLoader(ctx),
		Layer:          c.Layer.DataLoader(ctx),
		Plugin:         c.Plugin.DataLoader(ctx),
		Project:        c.Project.DataLoader(ctx),
		Property:       c.Property.DataLoader(ctx),
		PropertySchema: c.Property.SchemaDataLoader(ctx),
		Scene:          c.Scene.DataLoader(ctx),
		Team:           c.Team.DataLoader(ctx),
		User:           c.User.DataLoader(ctx),
	}
}

func (c Container) OrdinaryDataLoaders(ctx context.Context) DataLoaders {
	return DataLoaders{
		Dataset:        c.Dataset.OrdinaryDataLoader(ctx),
		DatasetSchema:  c.Dataset.SchemaOrdinaryDataLoader(ctx),
		LayerItem:      c.Layer.ItemOrdinaryDataLoader(ctx),
		LayerGroup:     c.Layer.GroupOrdinaryDataLoader(ctx),
		Layer:          c.Layer.OrdinaryDataLoader(ctx),
		Plugin:         c.Plugin.OrdinaryDataLoader(ctx),
		Project:        c.Project.OrdinaryDataLoader(ctx),
		Property:       c.Property.OrdinaryDataLoader(ctx),
		PropertySchema: c.Property.SchemaOrdinaryDataLoader(ctx),
		Scene:          c.Scene.OrdinaryDataLoader(ctx),
		Team:           c.Team.OrdinaryDataLoader(ctx),
		User:           c.User.OrdinaryDataLoader(ctx),
	}
}
