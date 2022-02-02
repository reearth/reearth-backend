package repo

import (
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

var ErrOperationDenied = interfaces.ErrOperationDenied

type Container struct {
	Asset          Asset
	Config         Config
	DatasetSchema  DatasetSchema
	Dataset        Dataset
	Layer          Layer
	Plugin         Plugin
	Project        Project
	PropertySchema PropertySchema
	Property       Property
	Scene          Scene
	Tag            Tag
	Team           Team
	User           User
	SceneLock      SceneLock
	Transaction    Transaction
	Lock           Lock
}

func (c Container) Filtered(scenes []id.SceneID, teams []id.TeamID) Container {
	return Container{
		Asset:          c.Asset.Filtered(teams),
		Config:         c.Config,
		DatasetSchema:  c.DatasetSchema.Filtered(scenes),
		Dataset:        c.Dataset.Filtered(scenes),
		Layer:          c.Layer.Filtered(scenes),
		Plugin:         c.Plugin,
		Project:        c.Project.Filtered(teams),
		PropertySchema: c.PropertySchema,
		Property:       c.Property.Filtered(scenes),
		Scene:          c.Scene.Filtered(teams),
		Tag:            c.Tag.Filtered(scenes),
		Team:           c.Team,
		User:           c.User,
		SceneLock:      c.SceneLock,
		Transaction:    c.Transaction,
		Lock:           c.Lock,
	}
}
