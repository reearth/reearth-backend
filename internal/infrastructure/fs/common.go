package fs

import (
	"os"
	"path/filepath"

	err1 "github.com/reearth/reearth-backend/pkg/error"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
)

const (
	assetDir         = "assets"
	pluginDir        = "plugins"
	publishedDir     = "published"
	manifestFilePath = "reearth.json"
)

func readManifest(base string, pid id.PluginID) (*manifest.Manifest, error) {
	file, err := os.Open(filepath.Join(base, pluginDir, pid.String(), manifestFilePath))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err1.ErrNotFound
		}
		return nil, err1.ErrInternalBy(err)
	}

	defer func() {
		_ = file.Close()
	}()

	m, err := manifest.Parse(file, nil)
	if err != nil {
		return nil, err1.ErrInternalBy(err)
	}

	return m, nil
}
