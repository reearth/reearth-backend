package value

var TypeBool Type = "bool"

var propertyBool = TypeProperty{
	I2V: func(i interface{}) (interface{}, bool) {
		if v, ok := i.(bool); ok {
			return v, true
		}
		if v, ok := i.(*bool); ok && v != nil {
			return *v, true
		}
		return nil, false
	},
	V2I: func(v interface{}) (interface{}, bool) {
		return v, true
	},
	Validate: func(i interface{}) bool {
		_, ok := i.(bool)
		return ok
	},
	Compatible: []Type{},
}

func (v *Value) ValueBool() (vv bool, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(bool)
	return
}
