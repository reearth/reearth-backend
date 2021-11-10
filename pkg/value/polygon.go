package value

import "github.com/mitchellh/mapstructure"

var TypePolygon Type = "polygon"

type Polygon []Coordinates

func PolygonFrom(rings [][]float64) Polygon {
	p := make([]Coordinates, 0, len(rings))
	for _, ring := range rings {
		p = append(p, CoordinatesFrom(ring))
	}
	return p
}

var propertyPolygon = TypeProperty{
	I2V: func(i interface{}) (interface{}, bool) {
		if v, ok := i.(Polygon); ok {
			return v, true
		}

		if v, ok := i.(*Polygon); ok {
			if v != nil {
				return *v, true
			}
			return nil, false
		}

		v := Polygon{}
		if err := mapstructure.Decode(i, &v); err == nil {
			return v, true
		}

		v2 := [][]float64{}
		if err := mapstructure.Decode(i, &v); err == nil {
			return PolygonFrom(v2), true
		}

		return nil, false
	},
	V2I: func(v interface{}) (interface{}, bool) {
		return v, true
	},
	Validate: func(i interface{}) bool {
		_, ok := i.(Polygon)
		return ok
	},
	Compatible: []Type{},
}

func (v *Value) ValuePolygon() (vv Polygon, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(Polygon)
	return
}
