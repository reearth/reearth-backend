package id

import (
	"encoding"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ encoding.TextMarshaler = (*PluginID)(nil)
var _ encoding.TextUnmarshaler = (*PluginID)(nil)

func TestPluginIDValidator(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "accepted name",
			input:    "1cc.1_c-d",
			expected: true,
		},
		{
			name:     "les then 100",
			input:    strings.Repeat("a", 100),
			expected: true,
		},
		{
			name:     "empty",
			input:    "",
			expected: false,
		},
		{
			name:     "spaces",
			input:    "    ",
			expected: false,
		},
		{
			name:     "contains not accepted characters",
			input:    "@bbb/aa-a_a",
			expected: false,
		},
		{
			name:     "contain space",
			input:    "bbb a",
			expected: false,
		},
		{
			name:     "contain =",
			input:    "cccd=",
			expected: false,
		},
		{
			name:     "contains reearth reserved key word",
			input:    "reearth",
			expected: false,
		},
		{
			name:     "more than 100 char",
			input:    strings.Repeat("a", 101),
			expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, validatePluginName(tc.input))
		})
	}
}

func TestPluginIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected struct {
			err    error
			result PluginID
		}
	}{
		{
			name:  "accepted name",
			input: "1cc.1_c-d#1.0.0",
			expected: struct {
				err    error
				result PluginID
			}{
				err: nil,
				result: PluginID{
					name:    "1cc.1_c-d",
					version: "1.0.0",
					sys:     false,
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, _ := PluginIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
		})
	}
}
