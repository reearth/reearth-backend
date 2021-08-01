package fs

import (
	"os"
	"path/filepath"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
	"github.com/reearth/reearth-backend/pkg/rerror"
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
			return nil, rerror.ErrNotFound
		}
		return nil, rerror.ErrInternalBy(err)
	}

	defer func() {
		_ = file.Close()
	}()

	m, err := manifest.Parse(file, nil)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}

	return m, nil
}
