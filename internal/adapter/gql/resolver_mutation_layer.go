package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

func (r *mutationResolver) AddLayerItem(ctx context.Context, input AddLayerItemInput) (*AddLayerItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	layer, parent, err := r.usecases.Layer.AddItem(ctx, interfaces.AddLayerItemInput{
		ParentLayerID: id.LayerID(input.ParentLayerID),
		PluginID:      &input.PluginID,
		ExtensionID:   &input.ExtensionID,
		Index:         input.Index,
		Name:          refToString(input.Name),
		LatLng:        toPropertyLatLng(input.Lat, input.Lng),
		// LinkedDatasetID: input.LinkedDatasetID,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &AddLayerItemPayload{
		Layer:       toLayerItem(layer, parent.IDRef()),
		ParentLayer: toLayerGroup(parent, nil),
		Index:       input.Index,
	}, nil
}

func (r *mutationResolver) AddLayerGroup(ctx context.Context, input AddLayerGroupInput) (*AddLayerGroupPayload, error) {
	exit := trace(ctx)
	defer exit()

	layer, parent, err := r.usecases.Layer.AddGroup(ctx, interfaces.AddLayerGroupInput{
		ParentLayerID:         id.LayerID(input.ParentLayerID),
		PluginID:              input.PluginID,
		ExtensionID:           input.ExtensionID,
		Index:                 input.Index,
		Name:                  refToString(input.Name),
		LinkedDatasetSchemaID: id.DatasetSchemaIDFromRefID(input.LinkedDatasetSchemaID),
		RepresentativeFieldId: input.RepresentativeFieldID,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &AddLayerGroupPayload{
		Layer:       toLayerGroup(layer, parent.IDRef()),
		ParentLayer: toLayerGroup(parent, nil),
		Index:       input.Index,
	}, nil
}

func (r *mutationResolver) RemoveLayer(ctx context.Context, input RemoveLayerInput) (*RemoveLayerPayload, error) {
	exit := trace(ctx)
	defer exit()

	id, layer, err := r.usecases.Layer.Remove(ctx, id.LayerID(input.LayerID), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &RemoveLayerPayload{
		LayerID:     id.ID(),
		ParentLayer: toLayerGroup(layer, nil),
	}, nil
}

func (r *mutationResolver) UpdateLayer(ctx context.Context, input UpdateLayerInput) (*UpdateLayerPayload, error) {
	exit := trace(ctx)
	defer exit()

	layer, err := r.usecases.Layer.Update(ctx, interfaces.UpdateLayerInput{
		LayerID: id.LayerID(input.LayerID),
		Name:    input.Name,
		Visible: input.Visible,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &UpdateLayerPayload{
		Layer: toLayer(layer, nil),
	}, nil
}

func (r *mutationResolver) MoveLayer(ctx context.Context, input MoveLayerInput) (*MoveLayerPayload, error) {
	exit := trace(ctx)
	defer exit()

	targetLayerID, layerGroupFrom, layerGroupTo, index, err := r.usecases.Layer.Move(ctx, interfaces.MoveLayerInput{
		LayerID:     id.LayerID(input.LayerID),
		DestLayerID: id.LayerIDFromRefID(input.DestLayerID),
		Index:       refToIndex(input.Index),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &MoveLayerPayload{
		LayerID:         targetLayerID.ID(),
		FromParentLayer: toLayerGroup(layerGroupFrom, nil),
		ToParentLayer:   toLayerGroup(layerGroupTo, nil),
		Index:           index,
	}, nil
}

func (r *mutationResolver) CreateInfobox(ctx context.Context, input CreateInfoboxInput) (*CreateInfoboxPayload, error) {
	exit := trace(ctx)
	defer exit()

	layer, err := r.usecases.Layer.CreateInfobox(ctx, id.LayerID(input.LayerID), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &CreateInfoboxPayload{
		Layer: toLayer(layer, nil),
	}, nil
}

func (r *mutationResolver) RemoveInfobox(ctx context.Context, input RemoveInfoboxInput) (*RemoveInfoboxPayload, error) {
	exit := trace(ctx)
	defer exit()

	layer, err := r.usecases.Layer.RemoveInfobox(ctx, id.LayerID(input.LayerID), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &RemoveInfoboxPayload{
		Layer: toLayer(layer, nil),
	}, nil
}

func (r *mutationResolver) AddInfoboxField(ctx context.Context, input AddInfoboxFieldInput) (*AddInfoboxFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	infoboxField, layer, err := r.usecases.Layer.AddInfoboxField(ctx, interfaces.AddInfoboxFieldParam{
		LayerID:     id.LayerID(input.LayerID),
		PluginID:    input.PluginID,
		ExtensionID: input.ExtensionID,
		Index:       input.Index,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &AddInfoboxFieldPayload{
		InfoboxField: toInfoboxField(infoboxField, layer.Scene(), nil),
		Layer:        toLayer(layer, nil),
	}, nil
}

func (r *mutationResolver) MoveInfoboxField(ctx context.Context, input MoveInfoboxFieldInput) (*MoveInfoboxFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	infoboxField, layer, index, err := r.usecases.Layer.MoveInfoboxField(ctx, interfaces.MoveInfoboxFieldParam{
		LayerID:        id.LayerID(input.LayerID),
		InfoboxFieldID: id.InfoboxFieldID(input.InfoboxFieldID),
		Index:          input.Index,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &MoveInfoboxFieldPayload{
		InfoboxFieldID: infoboxField.ID(),
		Layer:          toLayer(layer, nil),
		Index:          index,
	}, nil
}

func (r *mutationResolver) RemoveInfoboxField(ctx context.Context, input RemoveInfoboxFieldInput) (*RemoveInfoboxFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	infoboxField, layer, err := r.usecases.Layer.RemoveInfoboxField(ctx, interfaces.RemoveInfoboxFieldParam{
		LayerID:        id.LayerID(input.LayerID),
		InfoboxFieldID: id.InfoboxFieldID(input.InfoboxFieldID),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &RemoveInfoboxFieldPayload{
		InfoboxFieldID: infoboxField.ID(),
		Layer:          toLayer(layer, nil),
	}, nil
}

func (r *mutationResolver) ImportLayer(ctx context.Context, input ImportLayerInput) (*ImportLayerPayload, error) {
	exit := trace(ctx)
	defer exit()

	l, l2, err := r.usecases.Layer.ImportLayer(ctx, interfaces.ImportLayerParam{
		LayerID: id.LayerID(input.LayerID),
		File:    fromFile(&input.File),
		Format:  fromLayerEncodingFormat(input.Format),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &ImportLayerPayload{
		Layers:      toLayers(l, l2.IDRef()),
		ParentLayer: toLayerGroup(l2, nil),
	}, err
}
