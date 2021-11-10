package value

type Type string

type TypeProperty struct {
	I2V        func(interface{}) (interface{}, bool)
	V2I        func(interface{}) (interface{}, bool)
	Validate   func(interface{}) bool
	Compatible []Type
}

type TypePropertyMap = map[Type]TypeProperty

var TypeUnknown = Type("")

var defaultTypes = TypePropertyMap{
	TypeBool:         propertyBool,
	TypeCamera:       propertyCamera,
	TypeCoordinates:  propertyCoordinates,
	TypeLatLng:       propertyLatLng,
	TypeLatLngHeight: propertyLatLngHeight,
	TypeNumber:       propertyNumber,
	TypePolygon:      propertyPolygon,
	TypeRect:         propertyRect,
	TypeRef:          propertyRef,
	TypeString:       propertyString,
	TypeTypography:   propertyTypography,
	TypeURL:          propertyURL,
}

func (t Type) Default() bool {
	_, ok := defaultTypes[t]
	return ok
}

func (t Type) ValueFrom(i interface{}, p TypePropertyMap) *Value {
	if t == TypeUnknown || i == nil {
		return nil
	}

	if p != nil {
		if vt, ok := p[t]; ok && vt.I2V != nil {
			if v, ok2 := vt.I2V(i); ok2 {
				return &Value{p: p, v: v, t: t}
			}
		}
	}

	if vt, ok := defaultTypes[t]; ok && vt.I2V != nil {
		if v, ok2 := vt.I2V(i); ok2 {
			return &Value{p: p, v: v, t: t}
		}
	}

	return nil
}
