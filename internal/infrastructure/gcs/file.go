package gcs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	err1 "github.com/reearth/reearth-backend/pkg/error"
	"github.com/reearth/reearth-backend/pkg/file"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/log"
	"google.golang.org/api/iterator"
)

const (
	gcsAssetBasePath  string = "assets"
	gcsPluginBasePath string = "plugins"
	gcsMapBasePath    string = "maps"
	fileSizeLimit     int64  = 1024 * 1024 * 100 // about 100MB
)

type fileRepo struct {
	bucketName   string
	base         *url.URL
	cacheControl string
}

func NewFile(bucketName, base string, cacheControl string) (gateway.File, error) {
	if bucketName == "" {
		return nil, errors.New("bucket name is empty")
	}

	var u *url.URL
	if base == "" {
		base = fmt.Sprintf("https://storage.googleapis.com/%s", bucketName)
	}

	var err error
	u, _ = url.Parse(base)
	if err != nil {
		return nil, errors.New("invalid base URL")
	}

	return &fileRepo{
		bucketName:   bucketName,
		base:         u,
		cacheControl: cacheControl,
	}, nil
}

func (f *fileRepo) ReadAsset(ctx context.Context, name string) (io.ReadCloser, error) {
	if name == "" {
		return nil, err1.ErrNotFound
	}
	return f.read(ctx, path.Join(gcsAssetBasePath, name))
}

func (f *fileRepo) UploadAsset(ctx context.Context, file *file.File) (*url.URL, error) {
	if file == nil {
		return nil, gateway.ErrInvalidFile
	}
	if file.Size >= fileSizeLimit {
		return nil, gateway.ErrFileTooLarge
	}

	filename := path.Join(gcsAssetBasePath, id.New().String()+path.Ext(file.Path))
	u := getGCSObjectURL(f.base, filename)
	if u == nil {
		return nil, gateway.ErrInvalidFile
	}

	if err := f.upload(ctx, filename, file.Content); err != nil {
		return nil, err
	}
	return u, nil
}

func (f *fileRepo) RemoveAsset(ctx context.Context, u *url.URL) error {
	return f.delete(ctx, getGCSObjectNameFromURL(f.base, u))
}

// plugin

func (f *fileRepo) ReadPluginFile(ctx context.Context, pid id.PluginID, filename string) (io.ReadCloser, error) {
	if filename == "" {
		return nil, err1.ErrNotFound
	}
	return f.read(ctx, path.Join(gcsPluginBasePath, pid.String(), filename))
}

func (f *fileRepo) UploadPluginFile(ctx context.Context, pid id.PluginID, file *file.File) error {
	return f.upload(ctx, path.Join(gcsPluginBasePath, pid.String(), file.Path), file.Content)
}

func (f *fileRepo) RemovePlugin(ctx context.Context, pid id.PluginID) error {
	return f.deleteAll(ctx, path.Join(gcsPluginBasePath, pid.String()))
}

// built scene

func (f *fileRepo) ReadBuiltSceneFile(ctx context.Context, name string) (io.ReadCloser, error) {
	if name == "" {
		return nil, err1.ErrNotFound
	}
	return f.read(ctx, path.Join(gcsMapBasePath, name+".json"))
}

func (f *fileRepo) UploadBuiltScene(ctx context.Context, content io.Reader, name string) error {
	return f.upload(ctx, path.Join(gcsMapBasePath, name+".json"), content)
}

func (f *fileRepo) MoveBuiltScene(ctx context.Context, oldName, name string) error {
	return f.move(ctx, path.Join(gcsMapBasePath, oldName+".json"), path.Join(gcsMapBasePath, name+".json"))
}

func (f *fileRepo) RemoveBuiltScene(ctx context.Context, name string) error {
	return f.delete(ctx, path.Join(gcsMapBasePath, name+".json"))
}

// helpers

func (f *fileRepo) bucket(ctx context.Context) (*storage.BucketHandle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	bucket := client.Bucket(f.bucketName)
	return bucket, nil
}

func (f *fileRepo) read(ctx context.Context, filename string) (io.ReadCloser, error) {
	if filename == "" {
		return nil, err1.ErrNotFound
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: read bucket err: %+v\n", err)
		return nil, err1.ErrInternalBy(err)
	}

	reader, err := bucket.Object(filename).NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, err1.ErrNotFound
		}
		log.Errorf("gcs: read err: %+v\n", err)
		return nil, err1.ErrInternalBy(err)
	}

	return reader, nil
}

func (f *fileRepo) upload(ctx context.Context, filename string, content io.Reader) error {
	if filename == "" {
		return gateway.ErrInvalidFile
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: upload bucket err: %+v\n", err)
		return err1.ErrInternalBy(err)
	}

	object := bucket.Object(filename)
	if err := object.Delete(ctx); err != nil && !errors.Is(err, storage.ErrObjectNotExist) {
		log.Errorf("gcs: upload delete err: %+v\n", err)
		return gateway.ErrFailedToUploadFile
	}

	writer := object.NewWriter(ctx)
	writer.ObjectAttrs.CacheControl = f.cacheControl

	if _, err := io.Copy(writer, content); err != nil {
		log.Errorf("gcs: upload err: %+v\n", err)
		return gateway.ErrFailedToUploadFile
	}

	if err := writer.Close(); err != nil {
		log.Errorf("gcs: upload close err: %+v\n", err)
		return gateway.ErrFailedToUploadFile
	}

	return nil
}

func (f *fileRepo) move(ctx context.Context, from, dest string) error {
	if from == "" || dest == "" || from == dest {
		return gateway.ErrInvalidFile
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: move bucket err: %+v\n", err)
		return err1.ErrInternalBy(err)
	}

	object := bucket.Object(from)
	destObject := bucket.Object(dest)
	if _, err := destObject.CopierFrom(object).Run(ctx); err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return err1.ErrNotFound
		}
		log.Errorf("gcs: move copy err: %+v\n", err)
		return err1.ErrInternalBy(err)
	}

	if err := object.Delete(ctx); err != nil {
		log.Errorf("gcs: move delete err: %+v\n", err)
		return err1.ErrInternalBy(err)
	}

	return nil
}

func (f *fileRepo) delete(ctx context.Context, filename string) error {
	if filename == "" {
		return gateway.ErrInvalidFile
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: delete bucket err: %+v\n", err)
		return err1.ErrInternalBy(err)
	}

	object := bucket.Object(filename)
	if err := object.Delete(ctx); err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil
		}

		log.Errorf("gcs: delete err: %+v\n", err)
		return err1.ErrInternalBy(err)
	}
	return nil
}

func (f *fileRepo) deleteAll(ctx context.Context, path string) error {
	if path == "" {
		return gateway.ErrInvalidFile
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: deleteAll bucket err: %+v\n", err)
		return err1.ErrInternalBy(err)
	}

	it := bucket.Objects(ctx, &storage.Query{
		Prefix: path,
	})

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Errorf("gcs: deleteAll next err: %+v\n", err)
			return err1.ErrInternalBy(err)
		}
		if err := bucket.Object(attrs.Name).Delete(ctx); err != nil {
			log.Errorf("gcs: deleteAll err: %+v\n", err)
			return err1.ErrInternalBy(err)
		}
	}
	return nil
}

func getGCSObjectURL(base *url.URL, objectName string) *url.URL {
	if base == nil {
		return nil
	}
	b := *base
	b.Path = path.Join(b.Path, objectName)
	return &b
}

func getGCSObjectNameFromURL(base, u *url.URL) string {
	if u == nil {
		return ""
	}
	bp := ""
	if base != nil {
		bp = base.Path
	}
	return strings.TrimPrefix(strings.TrimPrefix(u.Path, bp), "/")
}
