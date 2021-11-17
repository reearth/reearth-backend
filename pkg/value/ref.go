package value

import "github.com/reearth/reearth-backend/pkg/id"

var TypeRef Type = "ref"

var propertyRef = TypeProperty{
	I2V: func(i interface{}) (interface{}, bool) {
		if v, ok := i.(string); ok {
			return v, true
		}
		if v, ok := i.(*string); ok {
			return *v, true
		}
		if v, ok := i.(id.ID); ok {
			return v.String(), true
		}
		if v, ok := i.(*id.ID); ok && v != nil {
			return v.String(), true
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

func (v *Value) ValueRef() (vv string, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(string)
	return
}
