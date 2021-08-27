package gql

import (
	"context"
)

func (r *Resolver) PropertySchemaField() PropertySchemaFieldResolver {
	return &propertySchemaFieldResolver{r}
}

func (r *Resolver) PropertySchemaFieldChoice() PropertySchemaFieldChoiceResolver {
	return &propertySchemaFieldChoiceResolver{r}
}

func (r *Resolver) PropertyLinkableFields() PropertyLinkableFieldsResolver {
	return &propertyLinkableFieldsResolver{r}
}

func (r *Resolver) PropertySchemaGroup() PropertySchemaGroupResolver {
	return &propertySchemaGroupResolver{r}
}

type propertySchemaFieldResolver struct{ *Resolver }

func (r *propertySchemaFieldResolver) TranslatedTitle(ctx context.Context, obj *PropertySchemaField, lang *string) (string, error) {
	exit := trace(ctx)
	defer exit()

	if s, ok := obj.AllTranslatedTitle[getLang(ctx, lang)]; ok {
		return s, nil
	}
	return obj.Name, nil
}

// deprecated
func (r *propertySchemaFieldResolver) TranslatedName(ctx context.Context, obj *PropertySchemaField, lang *string) (string, error) {
	exit := trace(ctx)
	defer exit()

	if s, ok := obj.AllTranslatedName[getLang(ctx, lang)]; ok {
		return s, nil
	}
	return obj.Name, nil
}

func (r *propertySchemaFieldResolver) TranslatedDescription(ctx context.Context, obj *PropertySchemaField, lang *string) (string, error) {
	exit := trace(ctx)
	defer exit()

	if s, ok := obj.AllTranslatedDescription[getLang(ctx, lang)]; ok {
		return s, nil
	}
	return obj.Description, nil
}

type propertyLinkableFieldsResolver struct{ *Resolver }

func (r *propertyLinkableFieldsResolver) Schema(ctx context.Context, obj *PropertyLinkableFields) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
}

func (r *propertyLinkableFieldsResolver) LatlngField(ctx context.Context, obj *PropertyLinkableFields) (*PropertySchemaField, error) {
	exit := trace(ctx)
	defer exit()

	if obj.Latlng == nil {
		return nil, nil
	}
	ps, err := DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
	return ps.Field(*obj.Latlng), err
}

func (r *propertyLinkableFieldsResolver) URLField(ctx context.Context, obj *PropertyLinkableFields) (*PropertySchemaField, error) {
	exit := trace(ctx)
	defer exit()

	if obj.URL == nil {
		return nil, nil
	}
	ps, err := DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
	return ps.Field(*obj.URL), err
}

type propertySchemaGroupResolver struct{ *Resolver }

func (r *propertySchemaGroupResolver) Schema(ctx context.Context, obj *PropertySchemaGroup) (*PropertySchema, error) {
	exit := trace(ctx)
	defer exit()

	return DataLoadersFromContext(ctx).PropertySchema.Load(obj.SchemaID)
}

func (r *propertySchemaGroupResolver) TranslatedTitle(ctx context.Context, obj *PropertySchemaGroup, lang *string) (string, error) {
	if s, ok := obj.AllTranslatedTitle[getLang(ctx, lang)]; ok {
		return s, nil
	}

	if obj.Title == nil {
		return "", nil
	}
	return *obj.Title, nil
}

type propertySchemaFieldChoiceResolver struct{ *Resolver }

func (r *propertySchemaFieldChoiceResolver) TranslatedTitle(ctx context.Context, obj *PropertySchemaFieldChoice, lang *string) (string, error) {
	if s, ok := obj.AllTranslatedTitle[getLang(ctx, lang)]; ok {
		return s, nil
	}
	return obj.Label, nil
}

// deprecated
func (r *propertySchemaFieldChoiceResolver) TranslatedLabel(ctx context.Context, obj *PropertySchemaFieldChoice, lang *string) (string, error) {
	if s, ok := obj.AllTranslatedLabel[getLang(ctx, lang)]; ok {
		return s, nil
	}
	return obj.Label, nil
}