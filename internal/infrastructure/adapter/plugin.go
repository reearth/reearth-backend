package adapter

import (
	"context"
	"errors"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

// TODO: ここで幅優先探索していくアルゴリズムを書いてmongoからビルトインの検索ロジックを除去する
type pluginRepo struct {
	readers []repo.Plugin
	writer  repo.Plugin
}

// NewPlugin generates a new repository which has fallback repositories to be used when the plugin is not found
func NewPlugin(readers []repo.Plugin, writer repo.Plugin) repo.Plugin {
	return &pluginRepo{
		readers: append([]repo.Plugin{}, readers...),
		writer:  writer,
	}
}

func (r *pluginRepo) FindByID(ctx context.Context, id id.PluginID, sids []id.SceneID) (*plugin.Plugin, error) {
	for _, re := range r.readers {
		if res, err := re.FindByID(ctx, id, sids); err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				continue
			} else {
				return nil, err
			}
		} else {
			return res, nil
		}
	}
	return nil, rerror.ErrNotFound
}

func (r *pluginRepo) FindByIDs(ctx context.Context, ids []id.PluginID, sids []id.SceneID) ([]*plugin.Plugin, error) {
	results := make([]*plugin.Plugin, 0, len(ids))
	for _, id := range ids {
		res, err := r.FindByID(ctx, id, sids)
		if err != nil && err != rerror.ErrNotFound {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (r *pluginRepo) Save(ctx context.Context, p *plugin.Plugin) error {
	if r.writer == nil {
		return errors.New("cannot write")
	}
	return r.writer.Save(ctx, p)
}

func (r *pluginRepo) Remove(ctx context.Context, p id.PluginID) error {
	if r.writer == nil {
		return errors.New("cannot write")
	}
	return r.writer.Remove(ctx, p)
}
