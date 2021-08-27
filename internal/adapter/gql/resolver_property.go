package gql

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

func (r *Resolver) Property() PropertyResolver {
	return &propertyResolver{r}
}

func (r *Resolver) PropertyField() PropertyFieldResolver {
	return &propertyFieldResolver{r}
}

func (r *Resolver) PropertyFieldLink() PropertyFieldLinkResolver {
	return &propertyFieldLinkResolver{r}
}

func (r *Resolver) MergedProperty() MergedPropertyResolver {
	return &mergedPropertyResolver{r}
}

func (r *Resolver) MergedPropertyGroup() MergedPropertyGroupResolver {
	return &mergedPropertyGroupResolver{r}
}

func (r *Resolver) MergedPropertyField() MergedPropertyFieldResolver {
	return &mergedPropertyFieldResolver{r}
}

func (r *Resolver) PropertyGroupList() PropertyGroupListResolver {
	return &propertyGroupListResolver{r}
}

func (r *Resolver) PropertyGroup() PropertyGroupResolver {
	return &propertyGroupResolver{r}
}

type propertyResolver struct{ *Resolver }

func (r *propertyResolver) Schema(ctx context.Context, obj *Property) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
}

func (r *propertyResolver) Layer(ctx context.Context, obj *Property) (Layer, error) {
	exit := trace(ctx)
	defer exit()

	l, err := r.config.Controllers.LayerController.FetchByProperty(ctx, id.PropertyID(obj.ID), getOperator(ctx))
	if err != nil || errors.Is(err, rerror.ErrNotFound) {
		return nil, nil
	}
	return l, err
}

func (r *propertyResolver) Merged(ctx context.Context, obj *Property) (*MergedProperty, error) {
	exit := trace(ctx)
	defer exit()

	l, err := r.config.Controllers.LayerController.FetchByProperty(ctx, id.PropertyID(obj.ID), getOperator(ctx))
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	li, ok := l.(*LayerItem)
	if !ok {
		return nil, nil
	}
	merged, err := r.LayerItem().Merged(ctx, li)
	if err != nil {
		return nil, err
	}
	if merged == nil {
		return nil, nil
	}
	if merged.Property != nil && merged.Property.OriginalID != nil && *merged.Property.OriginalID == obj.ID {
		return merged.Property, nil
	} else if merged.Infobox != nil && merged.Infobox.Property != nil && merged.Infobox.Property.OriginalID != nil && *merged.Infobox.Property.OriginalID == obj.ID {
		return merged.Infobox.Property, nil
	}
	return nil, nil
}

type propertyFieldResolver struct{ *Resolver }

func (r *propertyFieldResolver) Parent(ctx context.Context, obj *PropertyField) (*Property, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).Property.Load(id.PropertyID(obj.ParentID))
}

func (r *propertyFieldResolver) Schema(ctx context.Context, obj *PropertyField) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
}

func (r *propertyFieldResolver) Field(ctx context.Context, obj *PropertyField) (*PropertySchemaField, error) {
	exit := trace(ctx)
	defer exit()

	schema, err := DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
	if err != nil {
		return nil, err
	}
	return schema.Field(obj.FieldID), nil
}

func (r *propertyFieldResolver) ActualValue(ctx context.Context, obj *PropertyField) (interface{}, error) {
	exit := trace(ctx)
	defer exit()

	datasetLoader := DataLoadersFromContext(ctx).Dataset
	return actualValue(datasetLoader, obj.Value, obj.Links, false)
}

type propertyFieldLinkResolver struct{ *Resolver }

func (r *propertyFieldLinkResolver) Dataset(ctx context.Context, obj *PropertyFieldLink) (*Dataset, error) {
	exit := trace(ctx)
	defer exit()

	if obj.DatasetID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).Dataset.Load(id.DatasetID(*obj.DatasetID))
}

func (r *propertyFieldLinkResolver) DatasetField(ctx context.Context, obj *PropertyFieldLink) (*DatasetField, error) {
	exit := trace(ctx)
	defer exit()

	if obj.DatasetID == nil {
		return nil, nil
	}
	d, err := DataLoadersFromContext(ctx).Dataset.Load(id.DatasetID(*obj.DatasetID))
	if err != nil {
		return nil, err
	}
	return d.Field(obj.DatasetSchemaFieldID), nil
}

func (r *propertyFieldLinkResolver) DatasetSchema(ctx context.Context, obj *PropertyFieldLink) (*DatasetSchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).DatasetSchema.Load(id.DatasetSchemaID(obj.DatasetSchemaID))
}

func (r *propertyFieldLinkResolver) DatasetSchemaField(ctx context.Context, obj *PropertyFieldLink) (*DatasetSchemaField, error) {
	exit := trace(ctx)
	defer exit()

	ds, err := DataLoadersFromContext(ctx).DatasetSchema.Load(id.DatasetSchemaID(obj.DatasetSchemaID))
	return ds.Field(obj.DatasetSchemaFieldID), err
}

type mergedPropertyResolver struct{ *Resolver }

func (r *mergedPropertyResolver) Original(ctx context.Context, obj *MergedProperty) (*Property, error) {
	exit := trace(ctx)
	defer exit()

	if obj.OriginalID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).Property.Load(id.PropertyID(*obj.OriginalID))
}

func (r *mergedPropertyResolver) Parent(ctx context.Context, obj *MergedProperty) (*Property, error) {
	exit := trace(ctx)
	defer exit()

	if obj.ParentID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).Property.Load(id.PropertyID(*obj.ParentID))
}

func (r *mergedPropertyResolver) Schema(ctx context.Context, obj *MergedProperty) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	if obj.SchemaID == nil {
		if propertyID := obj.PropertyID(); propertyID != nil {
			property, err := DataLoadersFromContext(ctx).Property.Load(id.PropertyID(*propertyID))
			if err != nil {
				return nil, err
			}
			if property == nil {
				return nil, nil
			}
			return DataLoadersFromContext(ctx).PropertySchema.Load(property.SchemaID)
		}
		return nil, nil
	}
	return DataLoadersFromContext(ctx).PropertySchema.Load(*obj.SchemaID)
}

func (r *mergedPropertyResolver) LinkedDataset(ctx context.Context, obj *MergedProperty) (*Dataset, error) {
	exit := trace(ctx)
	defer exit()

	if obj.LinkedDatasetID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).Dataset.Load(id.DatasetID(*obj.LinkedDatasetID))
}

func (r *mergedPropertyResolver) Groups(ctx context.Context, obj *MergedProperty) ([]*MergedPropertyGroup, error) {
	exit := trace(ctx)
	defer exit()

	if obj.Groups != nil {
		return obj.Groups, nil
	}
	m, err := r.config.Controllers.PropertyController.FetchMerged(ctx, obj.OriginalID, obj.ParentID, obj.LinkedDatasetID, getOperator(ctx))
	if err != nil || m == nil {
		return nil, err
	}
	return m.Groups, nil
}

type mergedPropertyGroupResolver struct{ *Resolver }

func (r *mergedPropertyGroupResolver) Original(ctx context.Context, obj *MergedPropertyGroup) (*PropertyGroup, error) {
	exit := trace(ctx)
	defer exit()

	if obj.OriginalID == nil || obj.OriginalPropertyID == nil {
		return nil, nil
	}
	p, err := DataLoadersFromContext(ctx).Property.Load(id.PropertyID(*obj.OriginalID))
	if err != nil {
		return nil, err
	}
	if i, ok := p.Item(*obj.OriginalID).(*PropertyGroup); ok {
		return i, nil
	}
	return nil, nil
}

func (r *mergedPropertyGroupResolver) Parent(ctx context.Context, obj *MergedPropertyGroup) (*PropertyGroup, error) {
	exit := trace(ctx)
	defer exit()

	if obj.ParentID == nil || obj.ParentPropertyID == nil {
		return nil, nil
	}
	p, err := DataLoadersFromContext(ctx).Property.Load(id.PropertyID(*obj.ParentID))
	if err != nil {
		return nil, err
	}
	if i, ok := p.Item(*obj.ParentID).(*PropertyGroup); ok {
		return i, nil
	}
	return nil, nil
}

func (r *mergedPropertyGroupResolver) OriginalProperty(ctx context.Context, obj *MergedPropertyGroup) (*Property, error) {
	exit := trace(ctx)
	defer exit()

	if obj.OriginalID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).Property.Load(id.PropertyID(*obj.OriginalID))
}

func (r *mergedPropertyGroupResolver) ParentProperty(ctx context.Context, obj *MergedPropertyGroup) (*Property, error) {
	exit := trace(ctx)
	defer exit()

	if obj.ParentID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).Property.Load(id.PropertyID(*obj.ParentID))
}

func (r *mergedPropertyGroupResolver) Schema(ctx context.Context, obj *MergedPropertyGroup) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	if obj.SchemaID == nil {
		if propertyID := obj.PropertyID(); propertyID != nil {
			property, err := DataLoadersFromContext(ctx).Property.Load(id.PropertyID(*propertyID))
			if err != nil {
				return nil, err
			}
			if property == nil {
				return nil, nil
			}
			return DataLoadersFromContext(ctx).PropertySchema.Load(property.SchemaID)
		}
		return nil, nil
	}
	return DataLoadersFromContext(ctx).PropertySchema.Load(*obj.SchemaID)
}

func (r *mergedPropertyGroupResolver) LinkedDataset(ctx context.Context, obj *MergedPropertyGroup) (*Dataset, error) {
	exit := trace(ctx)
	defer exit()

	if obj.LinkedDatasetID == nil {
		return nil, nil
	}
	return DataLoadersFromContext(ctx).Dataset.Load(id.DatasetID(*obj.LinkedDatasetID))
}

type mergedPropertyFieldResolver struct{ *Resolver }

func (r *mergedPropertyFieldResolver) Schema(ctx context.Context, obj *MergedPropertyField) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
}

func (r *mergedPropertyFieldResolver) Field(ctx context.Context, obj *MergedPropertyField) (*PropertySchemaField, error) {
	exit := trace(ctx)
	defer exit()

	s, err := DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
	return s.Field(obj.FieldID), err
}

func (r *mergedPropertyFieldResolver) ActualValue(ctx context.Context, obj *MergedPropertyField) (interface{}, error) {
	exit := trace(ctx)
	defer exit()

	datasetLoader := DataLoadersFromContext(ctx).Dataset
	return actualValue(datasetLoader, obj.Value, obj.Links, obj.Overridden)
}

type propertyGroupListResolver struct{ *Resolver }

func (*propertyGroupListResolver) Schema(ctx context.Context, obj *PropertyGroupList) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
}

func (*propertyGroupListResolver) SchemaGroup(ctx context.Context, obj *PropertyGroupList) (*PropertySchemaGroup, error) {
	exit := trace(ctx)
	defer exit()

	s, err := DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
	if err != nil {
		return nil, err
	}
	return s.Group(obj.SchemaGroupID), nil
}

type propertyGroupResolver struct{ *Resolver }

func (*propertyGroupResolver) Schema(ctx context.Context, obj *PropertyGroup) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
}

func (*propertyGroupResolver) SchemaGroup(ctx context.Context, obj *PropertyGroup) (*PropertySchemaGroup, error) {
	exit := trace(ctx)
	defer exit()

	s, err := DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
	if err != nil {
		return nil, err
	}
	return s.Group(obj.SchemaGroupID), nil
}

func actualValue(datasetLoader DatasetDataLoader, value interface{}, links []*PropertyFieldLink, overridden bool) (interface{}, error) {
	if len(links) == 0 || overridden {
		return &value, nil
	}
	// 先頭のリンクにしかDatasetが割り当てられていない→先頭から順々に辿っていく
	if len(links) > 1 && links[0].DatasetID != nil && links[len(links)-1].DatasetID == nil {
		dsid := *links[0].DatasetID
		for i, link := range links {
			ds, err := datasetLoader.Load(id.DatasetID(dsid))
			if err != nil {
				return nil, err
			}
			field := ds.Field(link.DatasetSchemaFieldID)
			if field != nil {
				if i == len(links)-1 {
					return &value, nil
				} else if field.Type != ValueTypeRef {
					return nil, nil
				}
				if field.Value != nil {
					val, ok := (field.Value).(id.ID)
					if ok {
						dsid = val
					} else {
						return nil, nil
					}
				} else {
					return nil, nil
				}
			}
		}
	} else if lastLink := links[len(links)-1]; lastLink.DatasetID != nil {
		// 一番最後のリンクを取得
		ds, err := datasetLoader.Load(id.DatasetID(*lastLink.DatasetID))
		if err != nil {
			return nil, err
		}
		if f := ds.Field(lastLink.DatasetSchemaFieldID); f != nil {
			return &f.Value, nil
		}
	}
	return nil, nil
}
