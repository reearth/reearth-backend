package github

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/stretchr/testify/assert"
)

func TestNewPluginRegistry(t *testing.T) {
	d := NewPluginRegistry()
	assert.NotNil(t, d)
}

func TestPluginRegistry_Fetch(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/reearth/plugins/main/plugins.json",
		httpmock.NewStringResponder(200, `[{"name": "reearth","description": "Official Plugin","createdAt": "2021-03-16T04:19:57.592Z"}]`))
	d := NewPluginRegistry()
	res, err := d.Fetch(context.Background())
	tm, _ := time.Parse(time.RFC3339, "2021-03-16T04:19:57.592Z")

	assert.Equal(t, res, []*plugin.Metadata{
		{
			Name:        "reearth",
			Description: "Official Plugin",
			CreatedAt:   tm,
		},
	})
	assert.NoError(t, err)

	// fail: bad request
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/reearth/plugins/main/plugins.json",
		httpmock.NewStringResponder(400, `mock bad request`))
	_, err = d.Fetch(context.Background())
	assert.True(t, errors.As(errors.New("StatusCode=400"), &err))

	// fail: unable to marshal
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/reearth/plugins/main/plugins.json",
		httpmock.NewStringResponder(200, `{"hoge": "test"}`))
	_, err = d.Fetch(context.Background())
	assert.True(t, errors.As(errors.New("cannot unmarshal object into Go value of type []*plugin.Metadata"), &err))

}
