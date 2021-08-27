package gql

import (
	"context"
	"time"
)

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
