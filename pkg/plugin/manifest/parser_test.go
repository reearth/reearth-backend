package manifest

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name     string
		input    io.Reader
		expected *Manifest
		err      error
	}{
		{
			name: "success create manifest",
			input: strings.NewReader(`{
										"id": "aaa",
										"title": "bbb",
										"version": "1.1.1"
									}`),
			expected: &Manifest{
				Plugin:          plugin.New().ID(id.MustPluginID("aaa#1.1.1")).MustBuild(),
				ExtensionSchema: []*property.Schema{},
				Schema:          nil,
			},
			err: nil,
		},
		{
			name:     "fail not valid JSON",
			input:    strings.NewReader(""),
			expected: nil,
			err:      ErrFailedToParseManifest,
		},
		{
			name: "fail system manifest",
			input: strings.NewReader(`{
											"system":true,
											"id": "reearth",
											"title": "bbb",
											"version": "1.1.1"
											}`),
			expected: nil,
			err:      ErrSystemManifest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			m, err := Parse(tc.input)
			if err == nil {
				assert.Equal(t, tc.expected.Plugin.ID(), m.Plugin.ID())
				assert.Equal(t, m.Plugin.Name(), m.Plugin.Name())
			} else {
				assert.Equal(t, tc.err, err)
			}
		})
	}

}

func TestParseSystemFromBytes(t *testing.T) {
	testCases := []struct {
		name, input string
		expected    *Manifest
		err         error
	}{
		{
			name: "success create manifest",
			input: `{
						"id": "aaa",
						"title": "bbb",
						"version": "1.1.1"
									}`,
			expected: &Manifest{
				Plugin:          plugin.New().ID(id.MustPluginID("aaa#1.1.1")).MustBuild(),
				ExtensionSchema: []*property.Schema{},
				Schema:          nil,
			},
			err: nil,
		},
		{
			name:     "fail not valid YAML",
			input:    "--",
			expected: nil,
			err:      ErrFailedToParseManifest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			m, err := ParseSystemFromBytes([]byte(tc.input))
			if err == nil {
				assert.Equal(t, tc.expected.Plugin.ID(), m.Plugin.ID())
				assert.Equal(t, m.Plugin.Name(), m.Plugin.Name())
			} else {
				assert.True(t, errors.Is(tc.err, err))
			}
		})
	}
}

func TestMustParseSystemFromBytes(t *testing.T) {
	testCases := []struct {
		name, input string
		expected    *Manifest
		err         error
	}{
		{
			name: "success create manifest",
			input: `{
						"id": "aaa",
						"name": "bbb",
						"version": "1.1.1"
									}`,
			expected: &Manifest{
				Plugin:          plugin.New().ID(id.MustPluginID("aaa#1.1.1")).MustBuild(),
				ExtensionSchema: []*property.Schema{},
				Schema:          nil,
			},
			err: nil,
		},
		{
			name:     "fail not valid JSON",
			input:    "--",
			expected: nil,
			err:      ErrFailedToParseManifest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			if tc.err != nil {
				assert.PanicsWithError(tt, tc.err.Error(), func() {
					_ = MustParseSystemFromBytes([]byte(tc.input))
				})
				return
			}

			m := MustParseSystemFromBytes([]byte(tc.input))
			assert.Equal(tt, tc.expected.Plugin.ID(), m.Plugin.ID())
			assert.Equal(tt, m.Plugin.Name(), m.Plugin.Name())
		})
	}
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		name, input string
		err         bool
	}{
		{
			name: "success create manifest",
			input: `{
				"id": "aaa",
				"title": "bbb",
				"version": "1.1.1"
			}`,
			err: false,
		},
		{
			name:  "fail not valid JSON",
			input: "",
			err:   true,
		},
		{
			name: "fail invalid name type",
			input: `{
				"id": "aaa",
				"title": 123,
				"version": "1.1.1"
			}`,
			err: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			err := validate(gojsonschema.NewBytesLoader([]byte(tc.input)))
			if tc.err {
				assert.Error(tt, err)
			} else {
				assert.NoError(tt, err)
			}
		})
	}
}
