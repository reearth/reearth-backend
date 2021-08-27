package gql

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
)

func (r *mutationResolver) UpdatePropertyValue(ctx context.Context, input UpdatePropertyValueInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	v, ok := fromPropertyValueAndType(input.Value, input.Type)
	if !ok {
		return nil, errors.New("invalid value")
	}

	pp, pgl, pg, pf, err := r.usecases.Property.UpdateValue(ctx, interfaces.UpdatePropertyValueParam{
		PropertyID: id.PropertyID(input.PropertyID),
		Pointer:    fromPointer(input.SchemaItemID, input.ItemID, &input.FieldID),
		Value:      v,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &PropertyFieldPayload{
		Property:      toProperty(pp),
		PropertyField: toPropertyField(pf, pp, pgl, pg),
	}, nil
}

func (r *mutationResolver) UpdatePropertyValueLatLng(ctx context.Context, input UpdatePropertyValueLatLngInput) (*PropertyFieldPayload, error) {
	return nil, ErrNotImplemented
}

func (r *mutationResolver) UpdatePropertyValueLatLngHeight(ctx context.Context, input UpdatePropertyValueLatLngHeightInput) (*PropertyFieldPayload, error) {
	return nil, ErrNotImplemented
}

func (r *mutationResolver) UpdatePropertyValueCamera(ctx context.Context, input UpdatePropertyValueCameraInput) (*PropertyFieldPayload, error) {
	return nil, ErrNotImplemented
}

func (r *mutationResolver) UpdatePropertyValueTypography(ctx context.Context, input UpdatePropertyValueTypographyInput) (*PropertyFieldPayload, error) {
	return nil, ErrNotImplemented
}

func (r *mutationResolver) RemovePropertyField(ctx context.Context, input RemovePropertyFieldInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	p, err := r.usecases.Property.RemoveField(ctx, interfaces.RemovePropertyFieldParam{
		PropertyID: id.PropertyID(input.PropertyID),
		Pointer:    fromPointer(input.SchemaItemID, input.ItemID, &input.FieldID),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &PropertyFieldPayload{
		Property: toProperty(p),
	}, nil
}

func (r *mutationResolver) UploadFileToProperty(ctx context.Context, input UploadFileToPropertyInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	p, pgl, pg, pf, err := r.usecases.Property.UploadFile(ctx, interfaces.UploadFileParam{
		PropertyID: id.PropertyID(input.PropertyID),
		Pointer:    fromPointer(input.SchemaItemID, input.ItemID, &input.FieldID),
		File:       fromFile(&input.File),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &PropertyFieldPayload{
		Property:      toProperty(p),
		PropertyField: toPropertyField(pf, p, pgl, pg),
	}, nil
}

func (r *mutationResolver) LinkDatasetToPropertyValue(ctx context.Context, input LinkDatasetToPropertyValueInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	p, pgl, pg, pf, err := r.usecases.Property.LinkValue(ctx, interfaces.LinkPropertyValueParam{
		PropertyID: id.PropertyID(input.PropertyID),
		Pointer:    fromPointer(input.SchemaItemID, input.ItemID, &input.FieldID),
		Links: fromPropertyFieldLink(
			input.DatasetSchemaIds,
			input.DatasetIds,
			input.DatasetSchemaFieldIds,
		),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &PropertyFieldPayload{
		Property:      toProperty(p),
		PropertyField: toPropertyField(pf, p, pgl, pg),
	}, nil
}

func (r *mutationResolver) UnlinkPropertyValue(ctx context.Context, input UnlinkPropertyValueInput) (*PropertyFieldPayload, error) {
	exit := trace(ctx)
	defer exit()

	p, pgl, pg, pf, err := r.usecases.Property.UnlinkValue(ctx, interfaces.UnlinkPropertyValueParam{
		PropertyID: id.PropertyID(input.PropertyID),
		Pointer:    fromPointer(input.SchemaItemID, input.ItemID, &input.FieldID),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &PropertyFieldPayload{
		Property:      toProperty(p),
		PropertyField: toPropertyField(pf, p, pgl, pg),
	}, nil
}

func (r *mutationResolver) AddPropertyItem(ctx context.Context, input AddPropertyItemInput) (*PropertyItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	var v *property.Value
	if input.NameFieldType != nil {
		v, _ = fromPropertyValueAndType(input.NameFieldValue, *input.NameFieldType)
	}

	p, pgl, pi, err := r.usecases.Property.AddItem(ctx, interfaces.AddPropertyItemParam{
		PropertyID:     id.PropertyID(input.PropertyID),
		Pointer:        fromPointer(&input.SchemaItemID, nil, nil),
		Index:          input.Index,
		NameFieldValue: v,
	}, getOperator(ctx))

	if err != nil {
		return nil, err
	}

	return &PropertyItemPayload{
		Property:     toProperty(p),
		PropertyItem: toPropertyItem(pi, p, pgl),
	}, nil
}

func (r *mutationResolver) MovePropertyItem(ctx context.Context, input MovePropertyItemInput) (*PropertyItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	p, pgl, pi, err := r.usecases.Property.MoveItem(ctx, interfaces.MovePropertyItemParam{
		PropertyID: id.PropertyID(input.PropertyID),
		Pointer:    fromPointer(&input.SchemaItemID, &input.ItemID, nil),
		Index:      input.Index,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &PropertyItemPayload{
		Property:     toProperty(p),
		PropertyItem: toPropertyItem(pi, p, pgl),
	}, nil
}

func (r *mutationResolver) RemovePropertyItem(ctx context.Context, input RemovePropertyItemInput) (*PropertyItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	p, err := r.usecases.Property.RemoveItem(ctx, interfaces.RemovePropertyItemParam{
		PropertyID: id.PropertyID(input.PropertyID),
		Pointer:    fromPointer(&input.SchemaItemID, &input.ItemID, nil),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &PropertyItemPayload{
		Property: toProperty(p),
	}, nil
}

func (r *mutationResolver) UpdatePropertyItems(ctx context.Context, input UpdatePropertyItemInput) (*PropertyItemPayload, error) {
	exit := trace(ctx)
	defer exit()

	op := make([]interfaces.UpdatePropertyItemsOperationParam, 0, len(input.Operations))
	for _, o := range input.Operations {
		var v *property.Value
		if o.NameFieldType != nil {
			v, _ = fromPropertyValueAndType(o.NameFieldValue, *o.NameFieldType)
		}

		op = append(op, interfaces.UpdatePropertyItemsOperationParam{
			Operation:      fromListOperation(o.Operation),
			ItemID:         id.PropertyItemIDFromRefID(o.ItemID),
			Index:          o.Index,
			NameFieldValue: v,
		})
	}

	p, err2 := r.usecases.Property.UpdateItems(ctx, interfaces.UpdatePropertyItemsParam{
		PropertyID: id.PropertyID(input.PropertyID),
		Pointer:    fromPointer(&input.SchemaItemID, nil, nil),
		Operations: op,
	}, getOperator(ctx))
	if err2 != nil {
		return nil, err2
	}

	return &PropertyItemPayload{
		Property: toProperty(p),
	}, nil
}
