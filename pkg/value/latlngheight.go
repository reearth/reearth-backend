package value

import "github.com/mitchellh/mapstructure"

type LatLngHeight struct {
	Lat    float64 `json:"lat" mapstructure:"lat"`
	Lng    float64 `json:"lng" mapstructure:"lng"`
	Height float64 `json:"height" mapstructure:"height"`
}

func (l *LatLngHeight) Clone() *LatLngHeight {
	if l == nil {
		return nil
	}
	return &LatLngHeight{
		Lat:    l.Lat,
		Lng:    l.Lng,
		Height: l.Height,
	}
}

var TypeLatLngHeight Type = "latlngheight"

var propertyLatLngHeight = TypeProperty{
	I2V: func(i interface{}) (interface{}, bool) {
		if v, ok := i.(LatLngHeight); ok {
			return v, true
		}

		if v, ok := i.(*LatLngHeight); ok {
			if v != nil {
				return *v, false
			}
			return nil, false
		}

		v := LatLngHeight{}
		if err := mapstructure.Decode(i, &v); err == nil {
			return v, true
		}
		return nil, false
	},
	V2I: func(v interface{}) (interface{}, bool) {
		return v, true
	},
	Validate: func(i interface{}) bool {
		_, ok := i.(LatLngHeight)
		return ok
	},
	Compatible: []Type{},
}

func (v *Value) ValueLatLngHeight() (vv LatLngHeight, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(LatLngHeight)
	return
}
