package github

import (
	"context"
	"encoding/json"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/log"
	"github.com/reearth/reearth-backend/pkg/plugin"
)

type pluginRegistry struct{}

func NewPluginRegistry() repo.PluginRegistry {
	return &pluginRegistry{}
}

const source = `https://raw.githubusercontent.com/reearth/plugins/main/plugins.json`

func (d *pluginRegistry) Fetch(ctx context.Context) ([]*plugin.Metadata, error) {
	response, err := fetchURL(ctx, source)
	if err != nil {
		return nil, err
	}

	defer func() { err = response.Close() }()

	var result []*plugin.Metadata
	err = json.NewDecoder(response).Decode(&result)
	if err != nil {
		log.Errorf("plugin_registry: error: %s", err)
		return nil, repo.ErrFailedToFetchDataFromPluginRegistry
	}
	return result, nil
}
