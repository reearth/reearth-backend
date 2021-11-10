package value

import "net/url"

var TypeURL Type = "url"

var propertyURL = TypeProperty{
	I2V: func(i interface{}) (interface{}, bool) {
		if v, ok := i.(url.URL); ok {
			return &v, true
		}

		if v, ok := i.(*url.URL); ok {
			if v == nil {
				return nil, false
			}
			return v, true
		}

		if v, ok := i.(string); ok {
			if u, err := url.Parse(v); err == nil {
				return u, true
			}
		}

		return nil, false
	},
	V2I: func(v interface{}) (interface{}, bool) {
		u, ok := v.(*url.URL)
		if !ok {
			return nil, false
		}
		if u == nil {
			return "", true
		}
		return u.String(), true
	},
	Validate: func(i interface{}) bool {
		_, ok := i.(*url.URL)
		return ok
	},
	Compatible: []Type{},
}

func (v *Value) ValueURL() (vv *url.URL, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.v.(*url.URL)
	return
}
