package graphql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

type SceneControllerConfig struct {
	SceneInput func() interfaces.Scene
}

type SceneController struct {
	config SceneControllerConfig
}

func NewSceneController(config SceneControllerConfig) *SceneController {
	return &SceneController{config: config}
}

func (c *SceneController) usecase() interfaces.Scene {
	if c == nil {
		return nil
	}
	return c.config.SceneInput()
}

func (c *SceneController) Create(ctx context.Context, ginput *CreateSceneInput, operator *usecase.Operator) (*CreateScenePayload, error) {
	res, err := c.usecase().Create(
		ctx,
		id.ProjectID(ginput.ProjectID),
		operator,
	)
	if err != nil {
		return nil, err
	}

	return &CreateScenePayload{Scene: toScene(res)}, nil
}

func (c *SceneController) AddWidget(ctx context.Context, ginput *AddWidgetInput, operator *usecase.Operator) (*AddWidgetPayload, error) {
	scene, widget, err := c.usecase().AddWidget(
		ctx,
		id.SceneID(ginput.SceneID),
		ginput.PluginID,
		id.PluginExtensionID(ginput.ExtensionID),
		operator,
	)
	if err != nil {
		return nil, err
	}

	return &AddWidgetPayload{Scene: toScene(scene), SceneWidget: toSceneWidget(widget)}, nil
}

func (c *SceneController) UpdateWidget(ctx context.Context, ginput *UpdateWidgetInput, operator *usecase.Operator) (*UpdateWidgetPayload, error) {
	var layout interfaces.LayoutParams
	if ginput.Layout != nil {
		layout = interfaces.LayoutParams{
			Extended: ginput.Layout.Extended,
			Index:    ginput.Layout.Index,
			Location: (*scene.WidgetLocation)(ginput.Layout.Location),
			Align:    ginput.Layout.Align,
		}
	}
	scene, widget, err := c.usecase().UpdateWidget(ctx, interfaces.UpdateWidgetParam{
		SceneID:  id.SceneID(ginput.SceneID),
		WidgetID: id.WidgetID(ginput.WidgetID),
		Enabled:  ginput.Enabled,
		Layout:   &layout,
	}, operator)
	if err != nil {
		return nil, err
	}

	return &UpdateWidgetPayload{Scene: toScene(scene), SceneWidget: toSceneWidget(widget)}, nil
}

func (c *SceneController) RemoveWidget(ctx context.Context, ginput *RemoveWidgetInput, operator *usecase.Operator) (*RemoveWidgetPayload, error) {
	scene, err := c.usecase().RemoveWidget(ctx,
		id.SceneID(ginput.SceneID),
		id.WidgetID(ginput.WidgetID),
		operator,
	)
	if err != nil {
		return nil, err
	}

	return &RemoveWidgetPayload{Scene: toScene(scene), WidgetID: ginput.WidgetID}, nil
}

func (c *SceneController) InstallPlugin(ctx context.Context, ginput *InstallPluginInput, operator *usecase.Operator) (*InstallPluginPayload, error) {
	scene, pl, pr, err := c.usecase().InstallPlugin(ctx,
		id.SceneID(ginput.SceneID),
		ginput.PluginID,
		operator,
	)
	if err != nil {
		return nil, err
	}

	return &InstallPluginPayload{Scene: toScene(scene), ScenePlugin: &ScenePlugin{
		PluginID:   pl,
		PropertyID: pr.IDRef(),
	}}, nil
}

func (c *SceneController) UninstallPlugin(ctx context.Context, ginput *UninstallPluginInput, operator *usecase.Operator) (*UninstallPluginPayload, error) {
	scene, err := c.usecase().UninstallPlugin(ctx,
		id.SceneID(ginput.SceneID),
		id.PluginID(ginput.PluginID),
		operator,
	)
	if err != nil {
		return nil, err
	}

	return &UninstallPluginPayload{Scene: toScene(scene)}, nil
}

func (c *SceneController) UpgradePlugin(ctx context.Context, ginput *UpgradePluginInput, operator *usecase.Operator) (*UpgradePluginPayload, error) {
	s, err := c.usecase().UpgradePlugin(ctx,
		id.SceneID(ginput.SceneID),
		ginput.PluginID,
		ginput.ToPluginID,
		operator,
	)
	if err != nil {
		return nil, err
	}

	return &UpgradePluginPayload{
		Scene:       toScene(s),
		ScenePlugin: toScenePlugin(s.PluginSystem().Plugin(ginput.ToPluginID)),
	}, nil
}
