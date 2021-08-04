package interactor

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/plugin/pluginpack"
	"github.com/reearth/reearth-backend/pkg/plugin/repourl"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

var pluginPackageSizeLimit int64 = 10 * 1024 * 1024 // 10MB

func (i *Plugin) Upload(ctx context.Context, r io.Reader, sid id.SceneID, operator *usecase.Operator) (_ *plugin.Plugin, err error) {
	tx, err := i.transaction.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	if err := i.CanWriteScene(ctx, sid, operator); err != nil {
		return nil, err
	}

	p, err := pluginpack.PackageFromZip(r, &sid, pluginPackageSizeLimit)
	if err != nil {
		return nil, interfaces.ErrInvalidPluginPackage
	}

	for {
		f, err := p.Files.Next()
		if err != nil {
			return nil, interfaces.ErrInvalidPluginPackage
		}
		if f == nil {
			break
		}
		if err := i.file.UploadPluginFile(ctx, p.Manifest.Plugin.ID(), f); err != nil {
			return nil, rerror.ErrInternalBy(err)
		}
	}

	if ps := p.Manifest.PropertySchemas(); len(ps) > 0 {
		if err := i.propertySchemaRepo.SaveAll(ctx, ps); err != nil {
			return nil, err
		}
	}
	if err := i.pluginRepo.Save(ctx, p.Manifest.Plugin); err != nil {
		return nil, err
	}

	tx.Commit()
	return p.Manifest.Plugin, nil
}

func (i *Plugin) UploadFromRemote(ctx context.Context, u *url.URL, sid id.SceneID, operator *usecase.Operator) (_ *plugin.Plugin, err error) {
	ru, err := repourl.New(u)
	if err != nil {
		return nil, err
	}

	tx, err := i.transaction.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	if err := i.CanWriteScene(ctx, sid, operator); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ru.ArchiveURL().String(), nil)
	if err != nil {
		return nil, interfaces.ErrInvalidPluginPackage
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, interfaces.ErrInvalidPluginPackage
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != 200 {
		return nil, interfaces.ErrInvalidPluginPackage
	}

	p, err := pluginpack.PackageFromZip(res.Body, &sid, pluginPackageSizeLimit)
	if err != nil {
		return nil, interfaces.ErrInvalidPluginPackage
	}

	if p, err := i.pluginRepo.FindByID(ctx, p.Manifest.Plugin.ID(), []id.SceneID{sid}); err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return nil, err
	} else if p != nil {
		return nil, interfaces.ErrPluginAlreadyInstalled
	}

	for {
		f, err := p.Files.Next()
		if err != nil {
			return nil, interfaces.ErrInvalidPluginPackage
		}
		if f == nil {
			break
		}
		if err := i.file.UploadPluginFile(ctx, p.Manifest.Plugin.ID(), f); err != nil {
			return nil, rerror.ErrInternalBy(err)
		}
	}

	if ps := p.Manifest.PropertySchemas(); len(ps) > 0 {
		if err := i.propertySchemaRepo.SaveAll(ctx, ps); err != nil {
			return nil, err
		}
	}
	if err := i.pluginRepo.Save(ctx, p.Manifest.Plugin); err != nil {
		return nil, err
	}

	tx.Commit()
	return p.Manifest.Plugin, nil
}
