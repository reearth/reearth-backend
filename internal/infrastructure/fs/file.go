package fs

import (
	"context"
	"errors"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	err1 "github.com/reearth/reearth-backend/pkg/error"
	"github.com/reearth/reearth-backend/pkg/file"
	"github.com/reearth/reearth-backend/pkg/id"
)

type fileRepo struct {
	basePath string
	urlBase  *url.URL
}

func NewFile(basePath, urlBase string) (gateway.File, error) {
	var b *url.URL
	var err error
	b, err = url.Parse(urlBase)
	if err != nil {
		return nil, errors.New("invalid base URL")
	}

	return &fileRepo{
		basePath: basePath,
		urlBase:  b,
	}, nil
}

// asset

func (f *fileRepo) ReadAsset(ctx context.Context, filename string) (io.ReadCloser, error) {
	return f.read(ctx, filepath.Join(assetDir, filename))
}

func (f *fileRepo) UploadAsset(ctx context.Context, file *file.File) (*url.URL, error) {
	filename := id.New().String() + path.Ext(file.Path)
	if err := f.upload(ctx, filepath.Join(assetDir, filename), file.Content); err != nil {
		return nil, err
	}
	return getAssetFileURL(f.urlBase, filename), nil
}

func (f *fileRepo) RemoveAsset(ctx context.Context, u *url.URL) error {
	if u == nil {
		return gateway.ErrInvalidFile
	}
	return f.delete(ctx, filepath.Join(assetDir, getAssetFilePathFromURL(u)))
}

// plugin

func (f *fileRepo) ReadPluginFile(ctx context.Context, pid id.PluginID, filename string) (io.ReadCloser, error) {
	return f.read(ctx, filepath.Join(pluginDir, pid.String(), filename))
}

func (f *fileRepo) UploadPluginFile(ctx context.Context, pid id.PluginID, file *file.File) error {
	return f.upload(ctx, filepath.Join(pluginDir, pid.String(), file.Path), file.Content)
}

func (f *fileRepo) RemovePlugin(ctx context.Context, pid id.PluginID) error {
	return f.delete(ctx, filepath.Join(pluginDir, pid.String()))
}

// built scene

func (f *fileRepo) ReadBuiltSceneFile(ctx context.Context, name string) (io.ReadCloser, error) {
	return f.read(ctx, filepath.Join(publishedDir, name+".json"))
}

func (f *fileRepo) UploadBuiltScene(ctx context.Context, reader io.Reader, name string) error {
	return f.upload(ctx, filepath.Join(publishedDir, name+".json"), reader)
}

func (f *fileRepo) MoveBuiltScene(ctx context.Context, oldName, name string) error {
	return f.move(ctx, filepath.Join(publishedDir, oldName+".json"), filepath.Join(publishedDir, name+".json"))
}

func (f *fileRepo) RemoveBuiltScene(ctx context.Context, name string) error {
	return f.delete(ctx, filepath.Join(publishedDir, name+".json"))
}

// helpers

func (f *fileRepo) read(ctx context.Context, filename string) (io.ReadCloser, error) {
	file, err := os.Open(f.filename(filename))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err1.ErrNotFound
		}
		return nil, err1.ErrInternalBy(err)
	}
	return file, nil
}

func (f *fileRepo) upload(ctx context.Context, filename string, content io.Reader) error {
	if err := os.MkdirAll(path.Dir(f.filename(filename)), 0755); err != nil {
		return err1.ErrInternalBy(err)
	}

	dest, err := os.Create(f.filename(filename))
	if err != nil {
		return err1.ErrInternalBy(err)
	}
	defer func() {
		_ = dest.Close()
	}()

	if _, err := io.Copy(dest, content); err != nil {
		return gateway.ErrFailedToUploadFile
	}

	return nil
}

func (f *fileRepo) move(ctx context.Context, from, dest string) error {
	if from == "" || dest == "" || from == dest {
		return gateway.ErrInvalidFile
	}

	if err := os.MkdirAll(path.Dir(f.filename(dest)), 0755); err != nil {
		return err1.ErrInternalBy(err)
	}

	if err := os.Rename(f.filename(from), f.filename(dest)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return err1.ErrNotFound
		}
		return err1.ErrInternalBy(err)
	}

	return nil
}

func (f *fileRepo) delete(ctx context.Context, filename string) error {
	if err := os.Remove(f.filename(filename)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err1.ErrInternalBy(err)
	}
	return nil
}

func (f *fileRepo) filename(name string) string {
	return filepath.Join(f.basePath, filepath.Clean(name))
}

func getAssetFileURL(base *url.URL, filename string) *url.URL {
	if base == nil {
		return nil
	}

	b := *base
	b.Path = path.Join(b.Path, filename)
	return &b
}

func getAssetFilePathFromURL(u *url.URL) string {
	if u == nil {
		return ""
	}
	return path.Base(u.Path)
}
