package value

var TypeString Type = "string"

var propertyString = TypeProperty{
	I2V: func(i interface{}) (interface{}, bool) {
		if v, ok := i.(string); ok {
			return v, true
		}
		if v, ok := i.(*string); ok {
			return *v, true
		}
		return nil, false
	},
	V2I: func(v interface{}) (interface{}, bool) {
		return v, true
	},
	Validate: func(i interface{}) bool {
		_, ok := i.(string)
		return ok
	},
}

func (v *Value) ValueString() (vv string, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(string)
	return
}
