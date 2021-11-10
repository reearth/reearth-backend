package value

import "encoding/json"

var TypeNumber Type = "number"

var propertyNumber = TypeProperty{
	I2V: func(i interface{}) (interface{}, bool) {
		switch v := i.(type) {
		case float64:
			return v, true
		case float32:
			return float64(v), true
		case int:
			return float64(v), true
		case int8:
			return float64(v), true
		case int16:
			return float64(v), true
		case int32:
			return float64(v), true
		case int64:
			return float64(v), true
		case uint:
			return float64(v), true
		case uint8:
			return float64(v), true
		case uint16:
			return float64(v), true
		case uint32:
			return float64(v), true
		case uint64:
			return float64(v), true
		case uintptr:
			return float64(v), true
		case json.Number:
			if f, err := v.Float64(); err == nil {
				return f, true
			}
		case *float64:
			if v != nil {
				return *v, true
			}
		case *float32:
			if v != nil {
				return float64(*v), true
			}
		case *int:
			if v != nil {
				return float64(*v), true
			}
		case *int8:
			if v != nil {
				return float64(*v), true
			}
		case *int16:
			if v != nil {
				return float64(*v), true
			}
		case *int32:
			if v != nil {
				return float64(*v), true
			}
		case *int64:
			if v != nil {
				return float64(*v), true
			}
		case *uint:
			if v != nil {
				return float64(*v), true
			}
		case *uint8:
			if v != nil {
				return float64(*v), true
			}
		case *uint16:
			if v != nil {
				return float64(*v), true
			}
		case *uint32:
			if v != nil {
				return float64(*v), true
			}
		case *uint64:
			if v != nil {
				return float64(*v), true
			}
		case *uintptr:
			if v != nil {
				return float64(*v), true
			}
		case *json.Number:
			if v != nil {
				if f, err := v.Float64(); err == nil {
					return f, true
				}
			}
		}
		return nil, false
	},
	V2I: func(v interface{}) (interface{}, bool) {
		return v, true
	},
	Validate: func(i interface{}) bool {
		_, ok := i.(float64)
		return ok
	},
	Compatible: []Type{},
}

func (v *Value) ValueNumber() (vv float64, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(float64)
	return
}
