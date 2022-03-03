package fs

import (
	"context"
	"errors"
	"path/filepath"
	"regexp"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/spf13/afero"
)

type pluginRepo struct {
	fs afero.Fs
}

func NewPlugin(fs afero.Fs) repo.Plugin {
	return &pluginRepo{
		fs: fs,
	}
}

func (r *pluginRepo) FindByID(ctx context.Context, pid id.PluginID, sids []id.SceneID) (*plugin.Plugin, error) {
	m, err := readPluginManifest(r.fs, pid)
	if err != nil {
		return nil, err
	}

	sid := m.Plugin.ID().Scene()
	if sid != nil && !sid.Contains(sids) {
		return nil, nil
	}

	return m.Plugin, nil
}

func (r *pluginRepo) FindByIDs(ctx context.Context, ids []id.PluginID, sids []id.SceneID) ([]*plugin.Plugin, error) {
	results := make([]*plugin.Plugin, 0, len(ids))
	for _, id := range ids {
		res, err := r.FindByID(ctx, id, sids)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (r *pluginRepo) Save(ctx context.Context, p *plugin.Plugin) error {
	return rerror.ErrInternalBy(errors.New("read only"))
}

func (r *pluginRepo) Remove(ctx context.Context, pid id.PluginID) error {
	return rerror.ErrInternalBy(errors.New("read only"))
}

var translationFileNameRegexp = regexp.MustCompile(`reearth_([a-zA-Z]+(?:-[a-zA-Z]+)?).yml`)

func readPluginManifest(fs afero.Fs, pid id.PluginID) (*manifest.Manifest, error) {
	base := filepath.Join(pluginDir, pid.String())
	translationMap, err := readPluginTranslation(fs, base)
	if err != nil {
		return nil, err
	}

	f, err := fs.Open(filepath.Join(base, manifestFilePath))
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}
	defer func() {
		_ = f.Close()
	}()

	m, err := manifest.Parse(f, nil, translationMap.TranslatedRef())
	if err != nil {
		return nil, err
	}

	return m, nil
}

func readPluginTranslation(fs afero.Fs, base string) (manifest.TranslationMap, error) {
	d, err := afero.ReadDir(fs, base)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}

	translationMap := manifest.TranslationMap{}
	for _, e := range d {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		lang := translationFileNameRegexp.FindStringSubmatch(name)
		if len(lang) == 0 {
			continue
		}
		langfile, err := fs.Open(filepath.Join(base, name))
		if err != nil {
			return nil, rerror.ErrInternalBy(err)
		}
		defer func() {
			_ = langfile.Close()
		}()
		t, err := manifest.ParseTranslation(langfile)
		if err != nil {
			return nil, err
		}
		translationMap[lang[1]] = t
	}

	return translationMap, nil
}
