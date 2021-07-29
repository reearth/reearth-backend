package fs

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	err1 "github.com/reearth/reearth-backend/pkg/error"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
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
		return nil, err1.ErrInternalBy(err)
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
	return err1.ErrInternalBy(errors.New("read only"))
}
