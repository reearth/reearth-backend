package value

import "github.com/mitchellh/mapstructure"

type LatLng struct {
	Lat float64 `json:"lat" mapstructure:"lat"`
	Lng float64 `json:"lng" mapstructure:"lng"`
}

func (l *LatLng) Clone() *LatLng {
	if l == nil {
		return nil
	}
	return &LatLng{
		Lat: l.Lat,
		Lng: l.Lng,
	}
}

var TypeLatLng Type = "latlng"

var propertyLatLng = TypeProperty{
	I2V: func(i interface{}) (interface{}, bool) {
		if v, ok := i.(LatLng); ok {
			return v, true
		} else if v, ok := i.(*LatLng); ok {
			if v != nil {
				return *v, true
			}
			return nil, false
		}
		v := LatLng{}
		if err := mapstructure.Decode(i, &v); err != nil {
			return nil, false
		}
		return v, true
	},
	V2I: func(v interface{}) (interface{}, bool) {
		return v, true
	},
	Validate: func(i interface{}) bool {
		_, ok := i.(LatLng)
		return ok
	},
	Compatible: []Type{},
}

func (v *Value) ValueLatLng() (vv LatLng, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(LatLng)
	return
}
