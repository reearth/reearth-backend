package value

var TypeRef Type = "ref"

var propertyRef = TypeProperty{
	I2V:      propertyString.I2V,
	V2I:      propertyString.V2I,
	Validate: propertyString.Validate,
	// Compatible: []Type{},
}

func (v *Value) ValueRef() (vv string, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(string)
	return
}
