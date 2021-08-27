package gql

import (
	"context"
)

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAsset(ctx context.Context, input CreateAssetInput) (*CreateAssetPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.AssetController.Create(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) RemoveAsset(ctx context.Context, input RemoveAssetInput) (*RemoveAssetPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.AssetController.Remove(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UpdateDatasetSchema(ctx context.Context, input UpdateDatasetSchemaInput) (*UpdateDatasetSchemaPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.UpdateDatasetSchema(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) AddDynamicDatasetSchema(ctx context.Context, input AddDynamicDatasetSchemaInput) (*AddDynamicDatasetSchemaPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.AddDynamicDatasetSchema(ctx, &input)
}

func (r *mutationResolver) AddDynamicDataset(ctx context.Context, input AddDynamicDatasetInput) (*AddDynamicDatasetPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.AddDynamicDataset(ctx, &input)
}

func (r *mutationResolver) Signup(ctx context.Context, input SignupInput) (*SignupPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.UserController.Signup(ctx, &input, getSub(ctx))
}

func (r *mutationResolver) UpdateMe(ctx context.Context, input UpdateMeInput) (*UpdateMePayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.UserController.UpdateMe(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) RemoveMyAuth(ctx context.Context, input RemoveMyAuthInput) (*UpdateMePayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.UserController.RemoveMyAuth(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) DeleteMe(ctx context.Context, input DeleteMeInput) (*DeleteMePayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.UserController.DeleteMe(ctx, input.UserID, getOperator(ctx))
}

func (r *mutationResolver) CreateTeam(ctx context.Context, input CreateTeamInput) (*CreateTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.TeamController.Create(ctx, &input, getUser(ctx))
}

func (r *mutationResolver) DeleteTeam(ctx context.Context, input DeleteTeamInput) (*DeleteTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.TeamController.Remove(ctx, input.TeamID, getOperator(ctx))
}

func (r *mutationResolver) UpdateTeam(ctx context.Context, input UpdateTeamInput) (*UpdateTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.TeamController.Update(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) AddMemberToTeam(ctx context.Context, input AddMemberToTeamInput) (*AddMemberToTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.TeamController.AddMember(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) RemoveMemberFromTeam(ctx context.Context, input RemoveMemberFromTeamInput) (*RemoveMemberFromTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.TeamController.RemoveMember(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UpdateMemberOfTeam(ctx context.Context, input UpdateMemberOfTeamInput) (*UpdateMemberOfTeamPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.TeamController.UpdateMember(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) CreateProject(ctx context.Context, input CreateProjectInput) (*ProjectPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.ProjectController.Create(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UpdateProject(ctx context.Context, input UpdateProjectInput) (*ProjectPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.ProjectController.Update(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) PublishProject(ctx context.Context, input PublishProjectInput) (*ProjectPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.ProjectController.Publish(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) DeleteProject(ctx context.Context, input DeleteProjectInput) (*DeleteProjectPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.ProjectController.Delete(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UploadPlugin(ctx context.Context, input UploadPluginInput) (*UploadPluginPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PluginController.Upload(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) CreateScene(ctx context.Context, input CreateSceneInput) (*CreateScenePayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.SceneController.Create(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) AddWidget(ctx context.Context, input AddWidgetInput) (*AddWidgetPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.SceneController.AddWidget(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UpdateWidget(ctx context.Context, input UpdateWidgetInput) (*UpdateWidgetPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.SceneController.UpdateWidget(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) RemoveWidget(ctx context.Context, input RemoveWidgetInput) (*RemoveWidgetPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.SceneController.RemoveWidget(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) InstallPlugin(ctx context.Context, input InstallPluginInput) (*InstallPluginPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.SceneController.InstallPlugin(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UninstallPlugin(ctx context.Context, input UninstallPluginInput) (*UninstallPluginPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.SceneController.UninstallPlugin(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UpgradePlugin(ctx context.Context, input UpgradePluginInput) (*UpgradePluginPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.SceneController.UpgradePlugin(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) SyncDataset(ctx context.Context, input SyncDatasetInput) (*SyncDatasetPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.Sync(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UpdatePropertyValue(ctx context.Context, input UpdatePropertyValueInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.UpdateValue(ctx,
		input.PropertyID, input.SchemaItemID, input.ItemID, input.FieldID, input.Value, input.Type, getOperator(ctx))
}

func (r *mutationResolver) UpdatePropertyValueLatLng(ctx context.Context, input UpdatePropertyValueLatLngInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.UpdateValue(ctx,
		input.PropertyID, input.SchemaItemID, input.ItemID, input.FieldID, LatLng{
			Lat: input.Lat,
			Lng: input.Lng,
		}, ValueTypeLatlng, getOperator(ctx))
}

func (r *mutationResolver) UpdatePropertyValueLatLngHeight(ctx context.Context, input UpdatePropertyValueLatLngHeightInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.UpdateValue(ctx,
		input.PropertyID, input.SchemaItemID, input.ItemID, input.FieldID, LatLngHeight{
			Lat:    input.Lat,
			Lng:    input.Lng,
			Height: input.Height,
		}, ValueTypeLatlngheight, getOperator(ctx))
}

func (r *mutationResolver) UpdatePropertyValueCamera(ctx context.Context, input UpdatePropertyValueCameraInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.UpdateValue(ctx,
		input.PropertyID, input.SchemaItemID, input.ItemID, input.FieldID, Camera{
			Lat:      input.Lat,
			Lng:      input.Lng,
			Altitude: input.Altitude,
			Heading:  input.Heading,
			Pitch:    input.Pitch,
			Roll:     input.Roll,
			Fov:      input.Fov,
		}, ValueTypeLatlngheight, getOperator(ctx))
}

func (r *mutationResolver) UpdatePropertyValueTypography(ctx context.Context, input UpdatePropertyValueTypographyInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.UpdateValue(ctx,
		input.PropertyID, input.SchemaItemID, input.ItemID, input.FieldID, Typography{
			FontFamily: input.FontFamily,
			FontSize:   input.FontSize,
			FontWeight: input.FontWeight,
			Color:      input.Color,
			TextAlign:  input.TextAlign,
			Bold:       input.Bold,
			Italic:     input.Italic,
			Underline:  input.Underline,
		}, ValueTypeLatlngheight, getOperator(ctx))
}

func (r *mutationResolver) RemovePropertyField(ctx context.Context, input RemovePropertyFieldInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.RemoveField(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UploadFileToProperty(ctx context.Context, input UploadFileToPropertyInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.UploadFile(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) LinkDatasetToPropertyValue(ctx context.Context, input LinkDatasetToPropertyValueInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.LinkValue(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UnlinkPropertyValue(ctx context.Context, input UnlinkPropertyValueInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.UnlinkValue(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) AddPropertyItem(ctx context.Context, input AddPropertyItemInput) (*PropertyItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.AddItem(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) MovePropertyItem(ctx context.Context, input MovePropertyItemInput) (*PropertyItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.MoveItem(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) RemovePropertyItem(ctx context.Context, input RemovePropertyItemInput) (*PropertyItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.RemoveItem(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UpdatePropertyItems(ctx context.Context, input UpdatePropertyItemInput) (*PropertyItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.PropertyController.UpdateItems(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) AddLayerItem(ctx context.Context, input AddLayerItemInput) (*AddLayerItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.AddItem(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) AddLayerGroup(ctx context.Context, input AddLayerGroupInput) (*AddLayerGroupPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.AddGroup(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) RemoveLayer(ctx context.Context, input RemoveLayerInput) (*RemoveLayerPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.Remove(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) UpdateLayer(ctx context.Context, input UpdateLayerInput) (*UpdateLayerPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.Update(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) MoveLayer(ctx context.Context, input MoveLayerInput) (*MoveLayerPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.Move(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) CreateInfobox(ctx context.Context, input CreateInfoboxInput) (*CreateInfoboxPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.CreateInfobox(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) RemoveInfobox(ctx context.Context, input RemoveInfoboxInput) (*RemoveInfoboxPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.RemoveInfobox(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) AddInfoboxField(ctx context.Context, input AddInfoboxFieldInput) (*AddInfoboxFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.AddInfoboxField(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) MoveInfoboxField(ctx context.Context, input MoveInfoboxFieldInput) (*MoveInfoboxFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.MoveInfoboxField(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) RemoveInfoboxField(ctx context.Context, input RemoveInfoboxFieldInput) (*RemoveInfoboxFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.RemoveInfoboxField(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) RemoveDatasetSchema(ctx context.Context, input RemoveDatasetSchemaInput) (*RemoveDatasetSchemaPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.RemoveDatasetSchema(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) AddDatasetSchema(ctx context.Context, input AddDatasetSchemaInput) (*AddDatasetSchemaPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.AddDatasetSchema(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) ImportLayer(ctx context.Context, input ImportLayerInput) (*ImportLayerPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.LayerController.ImportLayer(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) ImportDataset(ctx context.Context, input ImportDatasetInput) (*ImportDatasetPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.ImportDataset(ctx, &input, getOperator(ctx))
}

func (r *mutationResolver) ImportDatasetFromGoogleSheet(ctx context.Context, input ImportDatasetFromGoogleSheetInput) (*ImportDatasetPayload, error) {
	exit := trace(ctx)
	defer exit()

	return r.config.Controllers.DatasetController.ImportDatasetFromGoogleSheet(ctx, &input, getOperator(ctx))
}
