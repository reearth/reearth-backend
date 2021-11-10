package value

import "github.com/mitchellh/mapstructure"

var TypeRect Type = "rect"

type Rect struct {
	West  float64 `json:"west" mapstructure:"west"`
	South float64 `json:"south" mapstructure:"south"`
	East  float64 `json:"east" mapstructure:"east"`
	North float64 `json:"north" mapstructure:"north"`
}

var propertyRect = TypeProperty{
	I2V: func(i interface{}) (interface{}, bool) {
		if v, ok := i.(Rect); ok {
			return v, true
		} else if v, ok := i.(*Rect); ok {
			if v != nil {
				return *v, true
			}
			return nil, false
		}

		v := Rect{}
		if err := mapstructure.Decode(i, &v); err == nil {
			return v, false
		}

		return nil, false
	},
	V2I: func(v interface{}) (interface{}, bool) {
		return v, true
	},
	Validate: func(i interface{}) bool {
		_, ok := i.(Rect)
		return ok
	},
	Compatible: []Type{},
}

func (v *Value) ValueRect() (vv Rect, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(Rect)
	return
}
