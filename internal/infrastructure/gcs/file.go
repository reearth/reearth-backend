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

func (f *fileRepo) bucket(ctx context.Context) (*storage.BucketHandle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	bucket := client.Bucket(f.bucketName)
	return bucket, err
}

func (f *fileRepo) ReadAsset(ctx context.Context, name string) (io.Reader, error) {
	if name == "" {
		return nil, err1.ErrNotFound
	}

	p := path.Join(gcsAssetBasePath, name)
	bucket, err := f.bucket(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("gcs: read asset from gs://%s/%s", f.bucketName, p)
	reader, err := bucket.Object(p).NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, err1.ErrNotFound
		}
		return nil, err1.ErrInternalBy(err)
	}
	return reader, nil
}

func (f *fileRepo) ReadPluginFile(ctx context.Context, plugin id.PluginID, name string) (io.Reader, error) {
	if name == "" {
		return nil, err1.ErrNotFound
	}

	p := path.Join(gcsPluginBasePath, plugin.Name(), plugin.Version().String(), name)
	bucket, err := f.bucket(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("gcs: read plugin from gs://%s/%s", f.bucketName, p)
	reader, err := bucket.Object(p).NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, err1.ErrNotFound
		}
		return nil, err1.ErrInternalBy(err)
	}
	return reader, nil
}

func (f *fileRepo) ReadBuiltSceneFile(ctx context.Context, name string) (io.Reader, error) {
	if name == "" {
		return nil, err1.ErrNotFound
	}

	p := path.Join(gcsMapBasePath, name+".json")
	bucket, err := f.bucket(ctx)
	if err != nil {
		return nil, err
	}

	log.Infof("gcs: read scene from gs://%s/%s", f.bucketName, p)
	reader, err := bucket.Object(p).NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, err1.ErrNotFound
		}
		return nil, err1.ErrInternalBy(err)
	}
	return reader, nil
}

func (f *fileRepo) UploadAsset(ctx context.Context, file *file.File) (*url.URL, error) {
	if file == nil {
		return nil, gateway.ErrInvalidFile
	}
	if file.Size >= fileSizeLimit {
		return nil, gateway.ErrFileTooLarge
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		return nil, err
	}

	// calc checksum
	// hasher := sha256.New()
	// tr := io.TeeReader(file.Content, hasher)
	// checksum := hex.EncodeToString(hasher.Sum(nil))

	id := id.New().String()
	filename := id + path.Ext(file.Name)
	name := path.Join(gcsAssetBasePath, filename)
	objectURL := getGCSObjectURL(f.base, name)
	if objectURL == nil {
		return nil, gateway.ErrInvalidFile
	}

	object := bucket.Object(name)
	_, err = object.Attrs(ctx)
	if !errors.Is(err, storage.ErrObjectNotExist) {
		log.Errorf("gcs: err=%+v\n", err)
		return nil, gateway.ErrFailedToUploadFile
	}

	writer := object.NewWriter(ctx)
	if _, err := io.Copy(writer, file.Content); err != nil {
		log.Errorf("gcs: err=%+v\n", err)
		return nil, gateway.ErrFailedToUploadFile
	}
	if err := writer.Close(); err != nil {
		log.Errorf("gcs: err=%+v\n", err)
		return nil, gateway.ErrFailedToUploadFile
	}

	return objectURL, nil
}

func (f *fileRepo) RemoveAsset(ctx context.Context, u *url.URL) error {
	if u == nil {
		return gateway.ErrInvalidFile
	}
	name := getGCSObjectNameFromURL(f.base, u)
	if name == "" {
		return gateway.ErrInvalidFile
	}
	bucket, err := f.bucket(ctx)
	if err != nil {
		return err1.ErrInternalBy(err)
	}
	object := bucket.Object(name)
	if err := object.Delete(ctx); err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil
		}
		return err1.ErrInternalBy(err)
	}
	return nil
}

func (f *fileRepo) RemovePlugin(ctx context.Context, pid id.PluginID) error {
	filename := path.Join(gcsMapBasePath, pid.String())
	if filename == "" {
		return gateway.ErrInvalidFile
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		return err1.ErrInternalBy(err)
	}


	object := bucket.Object(name)
	if err := object.Delete(ctx); err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil
		}
		return err1.ErrInternalBy(err)
	}
	return nil
}

func (f *fileRepo) UploadBuiltScene(ctx context.Context, reader io.Reader, name string) error {
	filename := path.Join(gcsMapBasePath, name+".json")
	bucket, err := f.bucket(ctx)
	if err != nil {
		return err
	}
	object := bucket.Object(filename)

	if err := object.Delete(ctx); err != nil && !errors.Is(err, storage.ErrObjectNotExist) {
		log.Errorf("gcs: err=%+v\n", err)
		return gateway.ErrFailedToUploadFile
	}

	writer := object.NewWriter(ctx)
	writer.ObjectAttrs.CacheControl = f.cacheControl

	if _, err := io.Copy(writer, reader); err != nil {
		log.Errorf("gcs: err=%+v\n", err)
		return gateway.ErrFailedToUploadFile
	}

	if err := writer.Close(); err != nil {
		log.Errorf("gcs: err=%+v\n", err)
		return gateway.ErrFailedToUploadFile
	}

	return nil
}

func (f *fileRepo) MoveBuiltScene(ctx context.Context, oldName, name string) error {
	oldFilename := path.Join(gcsMapBasePath, oldName+".json")
	filename := path.Join(gcsMapBasePath, name+".json")
	bucket, err := f.bucket(ctx)
	if err != nil {
		return err
	}
	object := bucket.Object(oldFilename)
	destObject := bucket.Object(filename)
	if _, err := destObject.CopierFrom(object).Run(ctx); err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return err1.ErrNotFound
		}
		return err1.ErrInternalBy(err)
	}
	if err := object.Delete(ctx); err != nil {
		return err1.ErrInternalBy(err)
	}
	return nil
}

func (f *fileRepo) RemoveBuiltScene(ctx context.Context, name string) error {
	filename := path.Join(gcsMapBasePath, name+".json")
	bucket, err := f.bucket(ctx)
	if err != nil {
		return err
	}
	object := bucket.Object(filename)
	if err := object.Delete(ctx); err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil
		}
		return err1.ErrInternalBy(err)
	}
	return nil
}

func (f *fileRepo) DeleteRecursive(ctx context.Context, path string) error {
	bucket, err := f.bucket(ctx)
	if err != nil {
		return err
	}
	it := bucket.Objects(ctx, &storage.Query{
		Prefix:    path,
	})
	for {
		oa, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err1.ErrInternalBy(err)
		}
		oa.
		// o.Delete(ctx)
	}

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
