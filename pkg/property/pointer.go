package property

import "github.com/reearth/reearth-backend/pkg/id"

// Pointer is a pointer to a field and an item in properties and schemas
type Pointer struct {
	schemaGroup *id.PropertySchemaGroupID
	item        *id.PropertyItemID
	field       *id.PropertySchemaFieldID
}

// NewPointer creates a new Pointer.
func NewPointer(sg *id.PropertySchemaGroupID, i *id.PropertyItemID, f *id.PropertySchemaFieldID) *Pointer {
	if sg == nil && i == nil && f == nil {
		return nil
	}
	return &Pointer{
		schemaGroup: sg.CopyRef(),
		item:        i.CopyRef(),
		field:       f.CopyRef(),
	}
}

// PointToEverything creates a new Pointer pointing to all items and fields.
func PointToEverything() *Pointer {
	return &Pointer{}
}

// PointField creates a new Pointer pointing to the field.
func PointField(sg *id.PropertySchemaGroupID, i *id.PropertyItemID, f id.PropertySchemaFieldID) *Pointer {
	return &Pointer{
		schemaGroup: sg.CopyRef(),
		item:        i.CopyRef(),
		field:       &f,
	}
}

// PointField creates a new Pointer pointing to the field in property schemas.
func PointFieldOnly(fid id.PropertySchemaFieldID) *Pointer {
	return &Pointer{
		field: &fid,
	}
}

// PointItemBySchema creates a new Pointer pointing to the schema item in property schemas.
func PointItemBySchema(sg id.PropertySchemaGroupID) *Pointer {
	return &Pointer{
		schemaGroup: &sg,
	}
}

// PointItem creates a new Pointer pointing to the item in properties.
func PointItem(i id.PropertyItemID) *Pointer {
	return &Pointer{
		item: &i,
	}
}

// PointFieldBySchemaGroup creates a new Pointer pointing to the field of the schema field in properties.
func PointFieldBySchemaGroup(sg id.PropertySchemaGroupID, f id.PropertySchemaFieldID) *Pointer {
	return &Pointer{
		schemaGroup: &sg,
		field:       &f,
	}
}

// PointFieldByItem creates a new Pointer pointing to the field of the item in properties.
func PointFieldByItem(i id.PropertyItemID, f id.PropertySchemaFieldID) *Pointer {
	return &Pointer{
		item:  &i,
		field: &f,
	}
}

func (p *Pointer) Clone() *Pointer {
	if p == nil {
		return nil
	}
	return &Pointer{
		field:       p.field.CopyRef(),
		item:        p.item.CopyRef(),
		schemaGroup: p.schemaGroup.CopyRef(),
	}
}

func (p *Pointer) ItemBySchemaGroupAndItem() (i id.PropertySchemaGroupID, i2 id.PropertyItemID, ok bool) {
	if p == nil || p.schemaGroup == nil || p.item == nil {
		ok = false
		return
	}
	i = *p.schemaGroup
	i2 = *p.item
	ok = true
	return
}

func (p *Pointer) ItemBySchemaGroup() (i id.PropertySchemaGroupID, ok bool) {
	if p == nil || p.schemaGroup == nil {
		ok = false
		return
	}
	i = *p.schemaGroup
	ok = true
	return
}

func (p *Pointer) Item() (i id.PropertyItemID, ok bool) {
	if p == nil || p.item == nil {
		ok = false
		return
	}
	i = *p.item
	ok = true
	return
}

func (p *Pointer) ItemRef() *id.PropertyItemID {
	i, ok := p.Item()
	if !ok {
		return nil
	}
	return i.Ref()
}

func (p *Pointer) FieldByItem() (i id.PropertyItemID, f id.PropertySchemaFieldID, ok bool) {
	if p == nil || p.item == nil || p.schemaGroup != nil || p.field == nil {
		ok = false
		return
	}
	i = *p.item
	f = *p.field
	ok = true
	return
}

func (p *Pointer) FieldBySchemaGroup() (sg id.PropertySchemaGroupID, f id.PropertySchemaFieldID, ok bool) {
	if p == nil || p.schemaGroup == nil || p.item != nil || p.field == nil {
		ok = false
		return
	}
	sg = *p.schemaGroup
	f = *p.field
	ok = true
	return
}

func (p *Pointer) Field() (f id.PropertySchemaFieldID, ok bool) {
	if p == nil || p.field == nil {
		ok = false
		return
	}
	f = *p.field
	ok = true
	return
}

func (p *Pointer) FieldRef() *id.PropertySchemaFieldID {
	f, ok := p.Field()
	if !ok {
		return nil
	}
	return f.Ref()
}

func (p *Pointer) FieldOnly() (f FieldID, ok bool) {
	if p == nil || p.field == nil || p.item != nil || p.schemaGroup != nil {
		ok = false
		return
	}
	f = *p.field
	ok = true
	return
}

func (p *Pointer) FieldOnlyRef() *FieldID {
	f, ok := p.FieldOnly()
	if !ok {
		return nil
	}
	return f.Ref()
}

func (p *Pointer) FieldIfItemIs(sg SchemaGroupID, i ItemID) (f FieldID, ok bool) {
	if p == nil || p.field == nil || !p.TestItem(sg, i) {
		ok = false
		return
	}
	f = *p.field
	ok = true
	return
}

func (p *Pointer) FieldIfItemIsRef(sg SchemaGroupID, i ItemID) *FieldID {
	f, ok := p.FieldIfItemIs(sg, i)
	if !ok {
		return nil
	}
	return f.Ref()
}

func (p *Pointer) Test(sg SchemaGroupID, i ItemID, f FieldID) bool {
	return p.TestItem(sg, i) && p.TestField(f)
}

func (p *Pointer) TestItem(sg SchemaGroupID, i ItemID) bool {
	return p.TestSchemaGroup(sg) && (p.item == nil || *p.item == i)
}

func (p *Pointer) TestSchemaGroup(sg SchemaGroupID) bool {
	return p != nil && (p.schemaGroup == nil || *p.schemaGroup == sg)
}

func (p *Pointer) TestField(f FieldID) bool {
	return p != nil && (p.field == nil || *p.field == f)
}

func (p *Pointer) GetAll() (sg *id.PropertySchemaGroupID, i *id.PropertyItemID, f *id.PropertySchemaFieldID) {
	if p == nil {
		return
	}
	sg = p.schemaGroup.CopyRef()
	i = p.item.CopyRef()
	f = p.field.CopyRef()
	return
}
