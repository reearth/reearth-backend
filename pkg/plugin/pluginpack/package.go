package pluginpack

import (
	"archive/zip"
	"bytes"
	"io"
	"path/filepath"

	"github.com/reearth/reearth-backend/pkg/file"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
)

const manfiestFilePath = "reearth.json"

type Package struct {
	Manifest *manifest.Manifest
	Files    file.Iterator
}

func PackageFromZip(r io.Reader, scene *id.SceneID, sizeLimit int64) (*Package, error) {
	b, err := io.ReadAll(io.LimitReader(r, sizeLimit))
	if err != nil {
		return nil, err
	}

	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}

	f, err := zr.Open(manfiestFilePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	m, err := manifest.Parse(f, scene)
	if err != nil {
		return nil, err
	}

	return &Package{
		Manifest: m,
		Files:    filter(file.NewZipReader(zr)),
	}, nil
}

func filter(a file.Iterator) file.Iterator {
	return file.NewFilteredIterator(a, func(p string) bool {
		return p == manfiestFilePath || filepath.Ext(p) != ".js"
	})
}
