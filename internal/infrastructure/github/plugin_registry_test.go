package github

import (
	"context"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/stretchr/testify/assert"
)

func TestNewPluginRegistry(t *testing.T) {
	d := NewPluginRegistry()
	assert.NotNil(t, d)
}

func TestPluginRegistry_FetchMetadata(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://raw.githubusercontent.com/reearth/plugins/main/plugins.json",
		httpmock.NewStringResponder(
			200,
			`[{"name": "reearth","description": "Official Plugin", "author": "reearth", "thumbnailUrl": "", "createdAt": "2021-03-16T04:19:57.592Z"}]`,
		),
	)

	d := NewPluginRegistry()
	res, err := d.Fetch(context.Background())
	tm, _ := time.Parse(time.RFC3339, "2021-03-16T04:19:57.592Z")

	assert.Equal(t, res, []*plugin.Metadata{
		{
			Name:         "reearth",
			Description:  "Official Plugin",
			Author:       "reearth",
			ThumbnailUrl: "",
			CreatedAt:    tm,
		},
	})
	assert.NoError(t, err)

	// fail: bad request
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/reearth/plugins/main/plugins.json",
		httpmock.NewStringResponder(400, `mock bad request`))
	_, err = d.Fetch(context.Background())

	assert.EqualError(t, err, "StatusCode=400")

	// fail: unable to marshal
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/reearth/plugins/main/plugins.json",
		httpmock.NewStringResponder(200, `{"hoge": "test"}`))
	_, err = d.Fetch(context.Background())

	assert.Equal(t, repo.ErrFailedToFetchDataFromPluginRegistry, err)
}
