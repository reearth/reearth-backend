package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *Resolver) Scene() SceneResolver {
	return &sceneResolver{r}
}

func (r *Resolver) ScenePlugin() ScenePluginResolver {
	return &scenePluginResolver{r}
}

func (r *Resolver) SceneWidget() SceneWidgetResolver {
	return &sceneWidgetResolver{r}
}

type sceneResolver struct{ *Resolver }

func (r *sceneResolver) Project(ctx context.Context, obj *Scene) (*Project, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Project.Load(id.ProjectID(obj.ProjectID))
}

func (r *sceneResolver) Team(ctx context.Context, obj *Scene) (*Team, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Team.Load(id.TeamID(obj.TeamID))
}

func (r *sceneResolver) Property(ctx context.Context, obj *Scene) (*Property, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Property.Load(id.PropertyID(obj.PropertyID))
}

func (r *sceneResolver) RootLayer(ctx context.Context, obj *Scene) (*LayerGroup, error) {
	exit := trace(ctx)
	defer exit()

	layer, err := DataLoadersFromContext(ctx).Layer.Load(id.LayerID(obj.RootLayerID))
	if err != nil {
		return nil, err
	}
	if layer == nil {
		return nil, nil
	}
	layerGroup, ok := (*layer).(*LayerGroup)
	if !ok {
		return nil, nil
	}
	return layerGroup, nil
}

func (r *sceneResolver) DatasetSchemas(ctx context.Context, obj *Scene, first *int, last *int, after *usecase.Cursor, before *usecase.Cursor) (*DatasetSchemaConnection, error) {
	exit := trace(ctx)
	defer exit()

	return r.controllers.Dataset.FindSchemaByScene(ctx, obj.ID, first, last, before, after, getOperator(ctx))
}

func (r *sceneResolver) LockMode(ctx context.Context, obj *Scene) (SceneLockMode, error) {
	exit := trace(ctx)
	defer exit()

	sl, err := r.controllers.Scene.FetchLock(ctx, id.SceneID(obj.ID), getOperator(ctx))
	if err != nil {
		return SceneLockModeFree, err
	}
	return *sl, nil
}

type scenePluginResolver struct{ *Resolver }

func (r *scenePluginResolver) Plugin(ctx context.Context, obj *ScenePlugin) (*Plugin, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Plugin.Load(obj.PluginID)
}
func (r *scenePluginResolver) Property(ctx context.Context, obj *ScenePlugin) (*Property, error) {
	exit := trace(ctx)
	defer exit()

	if obj.PropertyID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).Property.Load(id.PropertyID(*obj.PropertyID))
}

type sceneWidgetResolver struct{ *Resolver }

func (r *sceneWidgetResolver) Plugin(ctx context.Context, obj *SceneWidget) (*Plugin, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Plugin.Load(obj.PluginID)
}

func (r *sceneWidgetResolver) Extension(ctx context.Context, obj *SceneWidget) (*PluginExtension, error) {
	exit := trace(ctx)
	defer exit()

	plugin, err := DataLoadersFromContext(ctx).Plugin.Load(obj.PluginID)
	if err != nil {
		return nil, err
	}
	for _, e := range plugin.Extensions {
		if e.ExtensionID == obj.ExtensionID {
			return e, nil
		}
	}
	return nil, nil
}

func (r *sceneWidgetResolver) Property(ctx context.Context, obj *SceneWidget) (*Property, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Property.Load(id.PropertyID(obj.PropertyID))
}
