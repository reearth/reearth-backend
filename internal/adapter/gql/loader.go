package gql

import (
	"context"
	"time"

	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
)

const (
	dataLoaderWait     = 1 * time.Millisecond
	dataLoaderMaxBatch = 100
)

type Loaders struct {
	Asset    *AssetLoader
	Dataset  *DatasetLoader
	Layer    *LayerLoader
	Plugin   *PluginLoader
	Project  *ProjectLoader
	Property *PropertyLoader
	Scene    *SceneLoader
	Team     *TeamLoader
	User     *UserLoader
	Tag      *TagLoader
}

type DataLoaders struct {
	Asset          AssetDataLoader
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
	Tag            TagDataLoader
	TagItem        TagItemDataLoader
	TagGroup       TagGroupDataLoader
}

func NewLoaders(r *repo.Container, u *interfaces.Container) *Loaders {
	return &Loaders{
		Asset:    NewAssetLoader(r.Asset),
		Dataset:  NewDatasetLoader(r.Dataset, r.DatasetSchema, u.Dataset),
		Layer:    NewLayerLoader(r.Layer, u.Layer),
		Plugin:   NewPluginLoader(r.Plugin, r.PluginRegistry),
		Project:  NewProjectLoader(r.Project),
		Property: NewPropertyLoader(r.Property, r.PropertySchema, u.Property),
		Scene:    NewSceneLoader(r.Scene),
		Team:     NewTeamLoader(r.Team),
		User:     NewUserLoader(r.User),
		Tag:      NewTagLoader(r.Tag),
	}
}

func (l Loaders) DataLoadersWith(ctx context.Context, enabled bool) *DataLoaders {
	if enabled {
		return l.DataLoaders(ctx)
	}
	return l.OrdinaryDataLoaders(ctx)
}

func (l Loaders) DataLoaders(ctx context.Context) *DataLoaders {
	return &DataLoaders{
		Asset:          l.Asset.DataLoader(ctx),
		Dataset:        l.Dataset.DataLoader(ctx),
		DatasetSchema:  l.Dataset.SchemaDataLoader(ctx),
		LayerItem:      l.Layer.ItemDataLoader(ctx),
		LayerGroup:     l.Layer.GroupDataLoader(ctx),
		Layer:          l.Layer.DataLoader(ctx),
		Plugin:         l.Plugin.DataLoader(ctx),
		Project:        l.Project.DataLoader(ctx),
		Property:       l.Property.DataLoader(ctx),
		PropertySchema: l.Property.SchemaDataLoader(ctx),
		Scene:          l.Scene.DataLoader(ctx),
		Team:           l.Team.DataLoader(ctx),
		User:           l.User.DataLoader(ctx),
		Tag:            l.Tag.DataLoader(ctx),
		TagItem:        l.Tag.ItemDataLoader(ctx),
		TagGroup:       l.Tag.GroupDataLoader(ctx),
	}
}

func (l Loaders) OrdinaryDataLoaders(ctx context.Context) *DataLoaders {
	return &DataLoaders{
		Asset:          l.Asset.OrdinaryDataLoader(ctx),
		Dataset:        l.Dataset.OrdinaryDataLoader(ctx),
		DatasetSchema:  l.Dataset.SchemaOrdinaryDataLoader(ctx),
		LayerItem:      l.Layer.ItemOrdinaryDataLoader(ctx),
		LayerGroup:     l.Layer.GroupOrdinaryDataLoader(ctx),
		Layer:          l.Layer.OrdinaryDataLoader(ctx),
		Plugin:         l.Plugin.OrdinaryDataLoader(ctx),
		Project:        l.Project.OrdinaryDataLoader(ctx),
		Property:       l.Property.OrdinaryDataLoader(ctx),
		PropertySchema: l.Property.SchemaOrdinaryDataLoader(ctx),
		Scene:          l.Scene.OrdinaryDataLoader(ctx),
		Team:           l.Team.OrdinaryDataLoader(ctx),
		User:           l.User.OrdinaryDataLoader(ctx),
		Tag:            l.Tag.OrdinaryDataLoader(ctx),
		TagItem:        l.Tag.ItemDataLoader(ctx),
		TagGroup:       l.Tag.GroupDataLoader(ctx),
	}
}
