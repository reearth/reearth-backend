package fs

import (
	"context"
	"errors"
	"io"
	"path"

	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	err1 "github.com/reearth/reearth-backend/pkg/error"
	"github.com/reearth/reearth-backend/pkg/file"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
)

type pluginRepository struct {
	basePath string
}

func NewPluginRepository(basePath string) gateway.PluginRepository {
	return &pluginRepository{
		basePath: basePath,
	}
}

func (r *pluginRepository) Data(ctx context.Context, id id.PluginID) (file.Iterator, error) {
	return r.getArchive(id)
}

func (r *pluginRepository) Manifest(ctx context.Context, id id.PluginID) (*manifest.Manifest, error) {
	archive, err := r.getArchive(id)
	if err != nil {
		return nil, err
	}

	for {
		f, err := archive.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err1.ErrInternalBy(err)
		}
		if f.Path == manifestFilePath {
			m, err := manifest.Parse(f.Content)
			if err != nil {
				return nil, err
			}
			return m, nil
		}
	}

	return nil, manifest.ErrFailedToParseManifest
}

func (r *pluginRepository) getArchive(id id.PluginID) (file.Iterator, error) {
	return NewArchive(
		path.Join(r.basePath, id.Name()+"_"+id.Version().String()),
	)
}
