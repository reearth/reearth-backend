package property

type SchemaDiff struct {
	Deleted     []SchemaDiffDeleted
	Moved       []SchemaDiffMoved
	TypeChanged []SchemaDiffTypeChanged
}

type SchemaDiffDeleted SchemaFieldPointer

type SchemaDiffMoved struct {
	From   SchemaFieldPointer
	To     SchemaFieldPointer
	ToList bool
}

type SchemaDiffTypeChanged struct {
	SchemaFieldPointer
	NewType ValueType
}

func SchemaDiffFrom(old, new *Schema) (d SchemaDiff) {
	if old == nil || new == nil || old == new {
		return
	}

	for _, gf := range old.Groups().GroupAndFields() {
		ngf := new.Groups().GroupAndField(gf.Field.ID())
		if ngf == nil {
			d.Deleted = append(d.Deleted, SchemaDiffDeleted(gf.SchemaFieldPointer()))
			continue
		}

		if ngf.Group.ID() != gf.Group.ID() {
			d.Moved = append(d.Moved, SchemaDiffMoved{
				From:   gf.SchemaFieldPointer(),
				To:     ngf.SchemaFieldPointer(),
				ToList: ngf.Group.IsList(),
			})
		}

		if ngf.Field.Type() != gf.Field.Type() {
			d.TypeChanged = append(d.TypeChanged, SchemaDiffTypeChanged{
				SchemaFieldPointer: ngf.SchemaFieldPointer(),
				NewType:            ngf.Field.Type(),
			})
		}
	}

	return
}

func SchemaDiffFromProperty(old *Property, new *Schema) (d SchemaDiff) {
	return SchemaDiffFrom(old.GuessSchema(), new)
}
