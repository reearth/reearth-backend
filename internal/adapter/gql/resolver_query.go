package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Assets(ctx context.Context, teamID id.ID, first *int, last *int, after *usecase.Cursor, before *usecase.Cursor) (*AssetConnection, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.AssetController.FindByTeam(ctx, teamID, first, last, before, after, getOperator(ctx))
}

func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	exit := trace(ctx)
	defer exit()

	u := getUser(ctx)
	if u == nil {
		return nil, nil
	}
	return ToUser(u), nil
}

func (r *queryResolver) Node(ctx context.Context, i id.ID, typeArg NodeType) (Node, error) {
	exit := trace(ctx)
	defer exit()

	dataloaders := DataLoadersFromContext(ctx)
	switch typeArg {
	case NodeTypeDataset:
		result, err := dataloaders.Dataset.Load(id.DatasetID(i))
		if result == nil {
			return nil, nil
		}
		return result, err
	case NodeTypeDatasetSchema:
		result, err := dataloaders.DatasetSchema.Load(id.DatasetSchemaID(i))
		if result == nil {
			return nil, nil
		}
		return result, err
	case NodeTypeLayerItem:
		result, err := dataloaders.LayerItem.Load(id.LayerID(i))
		if result == nil {
			return nil, nil
		}
		return result, err
	case NodeTypeLayerGroup:
		result, err := dataloaders.LayerGroup.Load(id.LayerID(i))
		if result == nil {
			return nil, nil
		}
		return result, err
	case NodeTypeProject:
		result, err := dataloaders.Project.Load(id.ProjectID(i))
		if result == nil {
			return nil, nil
		}
		return result, err
	case NodeTypeProperty:
		result, err := dataloaders.Property.Load(id.PropertyID(i))
		if result == nil {
			return nil, nil
		}
		return result, err
	case NodeTypeScene:
		result, err := dataloaders.Scene.Load(id.SceneID(i))
		if result == nil {
			return nil, nil
		}
		return result, err
	case NodeTypeTeam:
		result, err := dataloaders.Team.Load(id.TeamID(i))
		if result == nil {
			return nil, nil
		}
		return result, err
	case NodeTypeUser:
		result, err := dataloaders.User.Load(id.UserID(i))
		if result == nil {
			return nil, nil
		}
		return result, err
	}
	return nil, nil
}

func (r *queryResolver) Nodes(ctx context.Context, ids []*id.ID, typeArg NodeType) ([]Node, error) {
	exit := trace(ctx)
	defer exit()

	dataloaders := DataLoadersFromContext(ctx)
	switch typeArg {
	case NodeTypeDataset:
		data, err := dataloaders.Dataset.LoadAll(id.DatasetIDsFromIDRef(ids))
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case NodeTypeDatasetSchema:
		data, err := dataloaders.DatasetSchema.LoadAll(id.DatasetSchemaIDsFromIDRef(ids))
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case NodeTypeLayerItem:
		data, err := dataloaders.LayerItem.LoadAll(id.LayerIDsFromIDRef(ids))
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]Node, len(data))
		for i := range data {
			nodes[i] = *data[i]
		}
		return nodes, nil
	case NodeTypeLayerGroup:
		data, err := dataloaders.LayerGroup.LoadAll(id.LayerIDsFromIDRef(ids))
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]Node, len(data))
		for i := range data {
			nodes[i] = *data[i]
		}
		return nodes, nil
	case NodeTypeProject:
		data, err := dataloaders.Project.LoadAll(id.ProjectIDsFromIDRef(ids))
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case NodeTypeProperty:
		data, err := dataloaders.Property.LoadAll(id.PropertyIDsFromIDRef(ids))
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case NodeTypeScene:
		data, err := dataloaders.Scene.LoadAll(id.SceneIDsFromIDRef(ids))
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case NodeTypeTeam:
		data, err := dataloaders.Team.LoadAll(id.TeamIDsFromIDRef(ids))
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	case NodeTypeUser:
		data, err := dataloaders.User.LoadAll(id.UserIDsFromIDRef(ids))
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}
		nodes := make([]Node, len(data))
		for i := range data {
			nodes[i] = data[i]
		}
		return nodes, nil
	default:
		return nil, nil
	}
}

func (r *queryResolver) PropertySchema(ctx context.Context, i id.PropertySchemaID) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).PropertySchema.Load(i)
}

func (r *queryResolver) PropertySchemas(ctx context.Context, ids []*id.PropertySchemaID) ([]*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	ids2 := make([]id.PropertySchemaID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			ids2 = append(ids2, *i)
		}
	}

	data, err := DataLoadersFromContext(ctx).PropertySchema.LoadAll(ids2)
	if len(err) > 0 && err[0] != nil {
		return nil, err[0]
	}

	return data, nil
}

func (r *queryResolver) Plugin(ctx context.Context, id id.PluginID) (*Plugin, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Plugin.Load(id)
}

func (r *queryResolver) Plugins(ctx context.Context, ids []*id.PluginID) ([]*Plugin, error) {
	exit := trace(ctx)
	defer exit()

	ids2 := make([]id.PluginID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			ids2 = append(ids2, *i)
		}
	}

	data, err := DataLoadersFromContext(ctx).Plugin.LoadAll(ids2)
	if len(err) > 0 && err[0] != nil {
		return nil, err[0]
	}

	return data, nil
}

func (r *queryResolver) Layer(ctx context.Context, layerID id.ID) (Layer, error) {
	exit := trace(ctx)
	defer exit()

	dataloaders := DataLoadersFromContext(ctx)
	result, err := dataloaders.Layer.Load(id.LayerID(layerID))
	if result == nil || *result == nil {
		return nil, nil
	}
	return *result, err
}

func (r *queryResolver) Scene(ctx context.Context, projectID id.ID) (*Scene, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.SceneController.FindByProject(ctx, id.ProjectID(projectID), getOperator(ctx))
}

func (r *queryResolver) Projects(ctx context.Context, teamID id.ID, includeArchived *bool, first *int, last *int, after *usecase.Cursor, before *usecase.Cursor) (*ProjectConnection, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.ProjectController.FindByTeam(ctx, id.TeamID(teamID), first, last, before, after, getOperator(ctx))
}

func (r *queryResolver) DatasetSchemas(ctx context.Context, sceneID id.ID, first *int, last *int, after *usecase.Cursor, before *usecase.Cursor) (*DatasetSchemaConnection, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.FindSchemaByScene(ctx, sceneID, first, last, before, after, getOperator(ctx))
}

func (r *queryResolver) DynamicDatasetSchemas(ctx context.Context, sceneID id.ID) ([]*DatasetSchema, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.FindDynamicSchemasByScene(ctx, sceneID)
}

func (r *queryResolver) Datasets(ctx context.Context, datasetSchemaID id.ID, first *int, last *int, after *usecase.Cursor, before *usecase.Cursor) (*DatasetConnection, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.FindBySchema(ctx, datasetSchemaID, first, last, before, after, getOperator(ctx))
}

func (r *queryResolver) SceneLock(ctx context.Context, sceneID id.ID) (*SceneLockMode, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.SceneController.FetchLock(ctx, id.SceneID(sceneID), getOperator(ctx))
}

func (r *queryResolver) SearchUser(ctx context.Context, nameOrEmail string) (*SearchedUser, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.UserController.SearchUser(ctx, nameOrEmail, getOperator(ctx))
}

func (r *queryResolver) CheckProjectAlias(ctx context.Context, alias string) (*CheckProjectAliasPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.ProjectController.CheckAlias(ctx, alias)
}

func (r *queryResolver) InstallablePlugins(ctx context.Context) ([]*PluginMetadata, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PluginController.FetchPluginMetadata(ctx, getOperator(ctx))
}
