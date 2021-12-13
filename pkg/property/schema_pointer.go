package property

type SchemaFieldPointer struct {
	SchemaGroup SchemaGroupID
	Field       FieldID
}

func (p SchemaFieldPointer) Pointer() *Pointer {
	return PointFieldBySchemaGroup(p.SchemaGroup, p.Field)
}
