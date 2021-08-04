package pluginpack

import (
	"archive/zip"
	"os"
	"testing"

	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
	"github.com/stretchr/testify/assert"
)

func TestPackageFromZip(t *testing.T) {
	f, err := os.Open("testdata/test.zip")
	assert.NoError(t, err)
	defer func() {
		_ = f.Close()
	}()

	expected := plugin.New().
		ID(id.MustPluginID("testplugin#1.0.1")).
		Name(i18n.StringFrom("testplugin")).
		MustBuild()

	p, err := PackageFromZip(f, nil, 1000)
	assert.NoError(t, err)
	assert.Equal(t, &manifest.Manifest{
		Plugin: expected,
	}, p.Manifest)

	var files []string
	for {
		n, err := p.Files.Next()
		assert.NoError(t, err)
		if n == nil {
			break
		}
		files = append(files, n.Path)
	}
	assert.Equal(t, []string{"index.js"}, files)
}

func TestPackageFromZip2(t *testing.T) {
	f, err := os.Open("testdata/test.zip")
	assert.NoError(t, err)
	defer func() {
		_ = f.Close()
	}()

	_, err = PackageFromZip(f, nil, 100)
	assert.ErrorIs(t, err, zip.ErrFormat)
}
