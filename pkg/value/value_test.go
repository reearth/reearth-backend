package value

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		value *Value
		want  bool
	}{
		{
			name: "empty",
			want: true,
		},
		{
			name: "nil",
			want: true,
		},
		{
			name: "non-empty",
			value: &Value{
				t: Type("hoge"),
				v: "foo",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.value.IsEmpty())
		})
	}
}

func TestValue_Clone(t *testing.T) {
	tests := []struct {
		name    string
		value   *Value
		wantnil bool
	}{
		{
			name: "ok",
			value: &Value{
				t: TypeString,
				v: "foo",
			},
		},
		{
			name: "custom type property",
			value: &Value{
				t: Type("hoge"),
				v: "foo",
				p: TypePropertyMap{Type("hoge"): TypeProperty{
					I2V: func(i interface{}) (interface{}, bool) { return i, true },
				}},
			},
		},
		{
			name: "nil",
		},
		{
			name:    "empty",
			value:   &Value{},
			wantnil: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := tt.value
			if tt.wantnil {
				want = nil
			}
			assert.Equal(t, want, tt.value.Clone())
		})
	}
}

func TestValue_Value(t *testing.T) {
	u, _ := url.Parse("https://reearth.io")
	tests := []struct {
		name  string
		value *Value
		want  interface{}
	}{
		{
			name:  "ok",
			value: &Value{t: TypeURL, v: u},
			want:  u,
		},
		{
			name:  "empty",
			value: &Value{},
		},
		{
			name: "nil",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want == nil {
				assert.Nil(t, tt.value.Value())
			} else {
				assert.Same(t, tt.want, tt.value.Value())
			}
		})
	}
}

func TestValue_Type(t *testing.T) {
	tests := []struct {
		name  string
		value *Value
		want  Type
	}{
		{
			name:  "ok",
			value: &Value{t: TypeString},
			want:  TypeString,
		},
		{
			name:  "empty",
			value: &Value{},
			want:  TypeUnknown,
		},
		{
			name: "nil",
			want: TypeUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.value.Type())
		})
	}
}

func TestValue_TypeProperty(t *testing.T) {
	typePropertyHoge := TypeProperty{}

	tests := []struct {
		name  string
		value *Value
		want  TypeProperty
	}{
		// {
		// 	name: "default type",
		// 	value: &Value{
		// 		v: "string",
		// 		t: TypeString,
		// 	},
		// 	want: defaultTypes[TypeString],
		// },
		{
			name: "custom type",
			value: &Value{
				v: "string",
				t: Type("hoge"),
				p: TypePropertyMap{
					Type("hoge"): typePropertyHoge,
				},
			},
			want: typePropertyHoge,
		},
		{
			name:  "empty",
			value: &Value{},
			want:  TypeProperty{},
		},
		{
			name: "nil",
			want: TypeProperty{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.value.TypeProperty())
		})
	}
}

func TestValue_Interface(t *testing.T) {
	tests := []struct {
		name  string
		value *Value
		want  interface{}
	}{
		{
			name:  "string",
			value: &Value{t: TypeString, v: "hoge"},
			want:  "hoge",
		},
		{
			name:  "latlng",
			value: &Value{t: TypeLatLng, v: LatLng{Lat: 1, Lng: 2}},
			want:  LatLng{Lat: 1, Lng: 2},
		},
		{
			name: "custom",
			value: &Value{
				p: TypePropertyMap{
					Type("foo"): TypeProperty{
						V2I: func(v interface{}) (interface{}, bool) {
							return v.(string) + "bar", true
						},
					},
				},
				t: Type("foo"),
				v: "foo",
			},
			want: "foobar",
		},
		{
			name:  "empty",
			value: &Value{},
			want:  nil,
		},
		{
			name:  "nil",
			value: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.value.Interface())
		})
	}
}
