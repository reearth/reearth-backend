package repo

import (
	"github.com/reearth/reearth-backend/pkg/scene"
	"github.com/reearth/reearth-backend/pkg/user"
)

type Container struct {
	Asset          Asset
	Config         Config
	DatasetSchema  DatasetSchema
	Dataset        Dataset
	Layer          Layer
	Lock           Lock
	Plugin         Plugin
	Project        Project
	PropertySchema PropertySchema
	Property       Property
	Scene          Scene
	SceneLock      SceneLock
	Tag            Tag
	Team           Team
	Transaction    Transaction
	User           User
}

func (c Container) Filtered(team TeamFilter, scene SceneFilter) Container {
	return Container{
		Asset:          c.Asset.Filtered(team),
		Config:         c.Config,
		DatasetSchema:  c.DatasetSchema.Filtered(scene),
		Dataset:        c.Dataset.Filtered(scene),
		Layer:          c.Layer.Filtered(scene),
		Lock:           c.Lock,
		Plugin:         c.Plugin,
		Project:        c.Project.Filtered(team),
		PropertySchema: c.PropertySchema,
		Property:       c.Property.Filtered(scene),
		Scene:          c.Scene.Filtered(team),
		SceneLock:      c.SceneLock,
		Tag:            c.Tag.Filtered(scene),
		Team:           c.Team,
		Transaction:    c.Transaction,
		User:           c.User,
	}
}

type TeamFilter struct {
	Readable user.TeamIDList
	Writable user.TeamIDList
}

func (f TeamFilter) Clone() TeamFilter {
	return TeamFilter{
		Readable: f.Readable.Clone(),
		Writable: f.Writable.Clone(),
	}
}

type SceneFilter struct {
	Readable scene.IDList
	Writable scene.IDList
}

func (f SceneFilter) Clone() SceneFilter {
	return SceneFilter{
		Readable: f.Readable.Clone(),
		Writable: f.Writable.Clone(),
	}
}
