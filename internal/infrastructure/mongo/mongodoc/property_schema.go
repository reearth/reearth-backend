package mongodoc

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
)

type PropertySchemaDocument struct {
	ID             string
	Version        int
	Groups         []*PropertySchemaGroupDocument
	LinkableFields *PropertyLinkableFieldsDocument
}

type PropertySchemaFieldDocument struct {
	ID           string
	Type         string
	Name         map[string]string
	Description  map[string]string
	Prefix       string
	Suffix       string
	DefaultValue interface{}
	UI           *string
	Min          *float64
	Max          *float64
	Choices      []PropertySchemaFieldChoiceDocument
}

type PropertySchemaFieldChoiceDocument struct {
	Key   string
	Label map[string]string
}

type PropertyLinkableFieldsDocument struct {
	LatLng *PropertyPointerDocument
	URL    *PropertyPointerDocument
}

type PropertyPointerDocument struct {
	SchemaGroupID *string
	ItemID        *string
	FieldID       *string
}

type PropertyConditonDocument struct {
	Field string
	Type  string
	Value interface{}
}

type PropertySchemaGroupDocument struct {
	ID            string
	Fields        []*PropertySchemaFieldDocument
	List          bool
	IsAvailableIf *PropertyConditonDocument
	Title         map[string]string
}

type PropertySchemaConsumer struct {
	Rows []*property.Schema
}

func (c *PropertySchemaConsumer) Consume(raw bson.Raw) error {
	if raw == nil {
		return nil
	}

	var doc PropertySchemaDocument
	if err := bson.Unmarshal(raw, &doc); err != nil {
		return err
	}
	propertySchema, err := doc.Model()
	if err != nil {
		return err
	}
	c.Rows = append(c.Rows, propertySchema)
	return nil
}

func NewPropertySchemaField(f *property.SchemaField) *PropertySchemaFieldDocument {
	if f == nil {
		return nil
	}

	field := &PropertySchemaFieldDocument{
		ID:           string(f.ID()),
		Name:         f.Title(),
		Suffix:       f.Suffix(),
		Prefix:       f.Prefix(),
		Description:  f.Description(),
		Type:         string(f.Type()),
		DefaultValue: f.DefaultValue().Value(),
		UI:           f.UI().StringRef(),
		Min:          f.Min(),
		Max:          f.Max(),
	}
	if choices := f.Choices(); choices != nil {
		field.Choices = make([]PropertySchemaFieldChoiceDocument, 0, len(choices))
		for _, c := range choices {
			field.Choices = append(field.Choices, PropertySchemaFieldChoiceDocument{
				Key:   c.Key,
				Label: c.Title,
			})
		}
	}
	return field
}

func NewPropertySchema(m *property.Schema) (*PropertySchemaDocument, string) {
	if m == nil {
		return nil, ""
	}

	pgroups := m.Groups()
	groups := make([]*PropertySchemaGroupDocument, 0, len(pgroups))
	for _, f := range pgroups {
		groups = append(groups, newPropertySchemaGroup(f))
	}

	id := m.ID().String()
	return &PropertySchemaDocument{
		ID:             id,
		Version:        m.Version(),
		Groups:         groups,
		LinkableFields: ToDocPropertyLinkableFields(m.LinkableFields()),
	}, id
}

func NewPropertySchemas(ps []*property.Schema) ([]interface{}, []string) {
	if ps == nil {
		return nil, nil
	}

	res := make([]interface{}, 0, len(ps))
	ids := make([]string, 0, len(ps))
	for _, d := range ps {
		if d == nil {
			continue
		}
		r, id := NewPropertySchema(d)
		res = append(res, r)
		ids = append(ids, id)
	}
	return res, ids
}

func ToModelPropertySchemaField(f *PropertySchemaFieldDocument) (*property.SchemaField, error) {
	if f == nil {
		return nil, nil
	}

	var choices []property.SchemaFieldChoice
	if f.Choices != nil {
		choices = make([]property.SchemaFieldChoice, 0, len(f.Choices))
		for _, c := range f.Choices {
			choices = append(choices, property.SchemaFieldChoice{
				Key:   c.Key,
				Title: c.Label,
			})
		}
	}

	vt := property.ValueType(f.Type)
	return property.NewSchemaField().
		ID(id.PropertySchemaFieldID(f.ID)).
		Type(vt).
		Name(f.Name).
		Description(f.Description).
		Prefix(f.Prefix).
		Suffix(f.Suffix).
		DefaultValue(vt.ValueFromUnsafe(f.DefaultValue)).
		UIRef(property.SchemaFieldUIFromRef(f.UI)).
		MinRef(f.Min).
		MaxRef(f.Max).
		Choices(choices).
		Build()
}

func (doc *PropertySchemaDocument) Model() (*property.Schema, error) {
	if doc == nil {
		return nil, nil
	}

	pid, err := id.PropertySchemaIDFrom(doc.ID)
	if err != nil {
		return nil, err
	}

	groups := make([]*property.SchemaGroup, 0, len(doc.Groups))
	for _, g := range doc.Groups {
		g2, err := toModelPropertySchemaGroup(g, pid)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g2)
	}

	return property.NewSchema().
		ID(pid).
		Version(doc.Version).
		Groups(groups).
		LinkableFields(toModelPropertyLinkableFields(doc.LinkableFields)).
		Build()
}

func newPropertyCondition(c *property.Condition) *PropertyConditonDocument {
	if c == nil {
		return nil
	}

	return &PropertyConditonDocument{
		Field: string(c.Field),
		Type:  string(c.Value.Type()),
		Value: c.Value.Interface(),
	}
}

func toModelPropertyCondition(d *PropertyConditonDocument) *property.Condition {
	if d == nil {
		return nil
	}

	return &property.Condition{
		Field: id.PropertySchemaFieldID(d.Field),
		Value: toModelPropertyValue(d.Value, d.Type),
	}
}

func newPropertySchemaGroup(p *property.SchemaGroup) *PropertySchemaGroupDocument {
	if p == nil {
		return nil
	}

	pfields := p.Fields()
	fields := make([]*PropertySchemaFieldDocument, 0, len(pfields))
	for _, f := range pfields {
		fields = append(fields, NewPropertySchemaField(f))
	}

	return &PropertySchemaGroupDocument{
		ID:            string(p.ID()),
		List:          p.IsList(),
		IsAvailableIf: newPropertyCondition(p.IsAvailableIf()),
		Title:         p.Title(),
		Fields:        fields,
	}
}

func toModelPropertySchemaGroup(d *PropertySchemaGroupDocument, sid id.PropertySchemaID) (*property.SchemaGroup, error) {
	if d == nil {
		return nil, nil
	}

	fields := make([]*property.SchemaField, 0, len(d.Fields))
	for _, f := range d.Fields {
		field, err := ToModelPropertySchemaField(f)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}

	return property.NewSchemaGroup().
		ID(id.PropertySchemaFieldID(d.ID)).
		Schema(sid).
		IsList(d.List).
		Title(d.Title).
		IsAvailableIf(toModelPropertyCondition(d.IsAvailableIf)).
		Fields(fields).
		Build()
}

func ToDocPropertyLinkableFields(l property.LinkableFields) *PropertyLinkableFieldsDocument {
	return &PropertyLinkableFieldsDocument{
		LatLng: newDocPropertyPointer(l.LatLng),
		URL:    newDocPropertyPointer(l.URL),
	}
}

func toModelPropertyLinkableFields(l *PropertyLinkableFieldsDocument) property.LinkableFields {
	if l == nil {
		return property.LinkableFields{}
	}
	return property.LinkableFields{
		LatLng: toModelPropertyPointer(l.LatLng),
		URL:    toModelPropertyPointer(l.URL),
	}
}

func toModelPropertyPointer(p *PropertyPointerDocument) *property.Pointer {
	if p == nil {
		return nil
	}
	return property.NewPointer(
		id.PropertySchemaFieldIDFrom(p.SchemaGroupID),
		id.PropertyItemIDFromRef(p.ItemID),
		id.PropertySchemaFieldIDFrom(p.FieldID),
	)
}

func newDocPropertyPointer(p *property.Pointer) *PropertyPointerDocument {
	if p == nil {
		return nil
	}
	schemaGroupID, itemID, fieldID := p.GetAll()
	return &PropertyPointerDocument{
		SchemaGroupID: schemaGroupID.StringRef(),
		ItemID:        itemID.StringRef(),
		FieldID:       fieldID.StringRef(),
	}
}