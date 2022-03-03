package manifest

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/reearth/reearth-backend/pkg/visualizer"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/minimum.yml
var minimum string
var minimumExpected = &Manifest{
	Plugin: plugin.New().ID(plugin.MustID("aaa~1.1.1")).MustBuild(),
}

//go:embed testdata/test.yml
var normal string
var normalExpected = &Manifest{
	Plugin: plugin.New().ID(plugin.MustID("aaa~1.1.1")).Name(i18n.StringFrom("bbb")).Extensions([]*plugin.Extension{
		plugin.NewExtension().ID(plugin.ExtensionID("hoge")).
			Visualizer(visualizer.VisualizerCesium).
			Type(plugin.ExtensionTypePrimitive).
			WidgetLayout(nil).
			Schema(property.MustSchemaID("aaa~1.1.1/hoge")).
			MustBuild(),
	}).MustBuild(),
	ExtensionSchema: []*property.Schema{
		property.NewSchema().ID(property.MustSchemaID("aaa~1.1.1/hoge")).Groups(property.NewSchemaGroupList([]*property.SchemaGroup{
			property.NewSchemaGroup().ID(property.SchemaGroupID("default")).
				RepresentativeField(property.FieldID("a").Ref()).
				Fields([]*property.SchemaField{
					property.NewSchemaField().ID(property.FieldID("a")).
						Type(property.ValueTypeBool).
						DefaultValue(property.ValueTypeBool.ValueFrom(true)).
						IsAvailableIf(&property.Condition{
							Field: property.FieldID("b"),
							Value: property.ValueTypeNumber.ValueFrom(1),
						}).
						MustBuild(),
					property.NewSchemaField().ID(property.FieldID("b")).
						Type(property.ValueTypeNumber).
						MustBuild(),
				}).MustBuild(),
		})).MustBuild(),
	},
}

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *Manifest
		err      error
	}{
		{
			name:     "success create simple manifest",
			input:    minimum,
			expected: minimumExpected,
			err:      nil,
		},
		{
			name:     "success create manifest",
			input:    normal,
			expected: normalExpected,
			err:      nil,
		},
		{
			name:     "fail not valid JSON",
			input:    "",
			expected: nil,
			err:      ErrFailedToParseManifest,
		},
		{
			name: "fail system manifest",
			input: `{
				"system": true,
				"id": "reearth",
				"title": "bbb",
				"version": "1.1.1"
			}`,
			expected: nil,
			err:      ErrSystemManifest,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			m, err := Parse(strings.NewReader(tc.input), nil, nil)
			if tc.err == nil {
				if !assert.NoError(t, err) {
					return
				}
				assert.Equal(t, tc.expected, m)
				return
			}
			assert.ErrorIs(t, tc.err, err)
		})
	}

}

func TestParseSystemFromBytes(t *testing.T) {
	tests := []struct {
		name, input string
		expected    *Manifest
		err         error
	}{
		{
			name:     "success create simple manifest",
			input:    minimum,
			expected: minimumExpected,
			err:      nil,
		},
		{
			name:     "success create manifest",
			input:    normal,
			expected: normalExpected,
			err:      nil,
		},
		{
			name:     "fail not valid YAML",
			input:    "--",
			expected: nil,
			err:      ErrFailedToParseManifest,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			m, err := ParseSystemFromBytes([]byte(tc.input), nil, nil)
			if tc.err == nil {
				if !assert.NoError(t, err) {
					return
				}
				assert.Equal(t, tc.expected, m)
				return
			}
			assert.ErrorIs(t, tc.err, err)
		})
	}
}

func TestMustParseSystemFromBytes(t *testing.T) {
	tests := []struct {
		name, input string
		expected    *Manifest
		fails       bool
	}{
		{
			name:     "success create simple manifest",
			input:    minimum,
			expected: minimumExpected,
			fails:    false,
		},
		{
			name:     "success create manifest",
			input:    normal,
			expected: normalExpected,
			fails:    false,
		},
		{
			name:     "fail not valid JSON",
			input:    "--",
			expected: nil,
			fails:    true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.fails {
				assert.Panics(t, func() {
					_ = MustParseSystemFromBytes([]byte(tc.input), nil, nil)
				})
				return
			}

			m := MustParseSystemFromBytes([]byte(tc.input), nil, nil)
			assert.Equal(t, m, tc.expected)
		})
	}
}
