package fs

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

type pluginRepo struct {
	basePath string
}

func NewPlugin(basePath string) repo.Plugin {
	return &pluginRepo{
		basePath: basePath,
	}
}

func (r *pluginRepo) FindByID(ctx context.Context, pid id.PluginID, sids []id.SceneID) (*plugin.Plugin, error) {
	m, err := readManifest(r.basePath, pid)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
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
	return err1.ErrInternalBy(errors.New("read only"))
}
