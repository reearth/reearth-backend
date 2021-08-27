package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
)

type SceneController struct {
	usecase interfaces.Scene
}

func NewSceneController(usecase interfaces.Scene) *SceneController {
	return &SceneController{usecase: usecase}
}

func (c *SceneController) Fetch(ctx context.Context, ids []id.SceneID, operator *usecase.Operator) ([]*Scene, []error) {
	res, err := c.usecase.Fetch(ctx, ids, operator)
	if err != nil {
		return nil, []error{err}
	}

	scenes := make([]*Scene, 0, len(res))
	for _, scene := range res {
		scenes = append(scenes, toScene(scene))
	}
	return scenes, nil
}

func (c *SceneController) FindByProject(ctx context.Context, projectID id.ProjectID, operator *usecase.Operator) (*Scene, error) {
	res, err := c.usecase.FindByProject(ctx, projectID, operator)
	if err != nil {
		return nil, err
	}

	return toScene(res), nil
}

func (c *SceneController) FetchLock(ctx context.Context, sid id.SceneID, operator *usecase.Operator) (*SceneLockMode, error) {
	res, err := c.usecase.FetchLock(ctx, []id.SceneID{sid}, operator)
	if err != nil {
		return nil, err
	}
	if len(res) > 0 {
		return nil, nil
	}
	sl := toSceneLockMode(res[0])
	return &sl, nil
}

func (c *SceneController) FetchLockAll(ctx context.Context, sid []id.SceneID, operator *usecase.Operator) ([]SceneLockMode, []error) {
	res, err := c.usecase.FetchLock(ctx, sid, operator)
	if err != nil {
		return nil, []error{err}
	}

	res2 := make([]SceneLockMode, 0, len(res))
	for _, r := range res {
		res2 = append(res2, toSceneLockMode(r))
	}

	return res2, nil
}

// data loader

type SceneDataLoader interface {
	Load(id.SceneID) (*Scene, error)
	LoadAll([]id.SceneID) ([]*Scene, []error)
}

func (c *SceneController) DataLoader(ctx context.Context) SceneDataLoader {
	return NewSceneLoader(SceneLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.SceneID) ([]*Scene, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	})
}

func (c *SceneController) OrdinaryDataLoader(ctx context.Context) SceneDataLoader {
	return &ordinarySceneLoader{
		fetch: func(keys []id.SceneID) ([]*Scene, []error) {
			return c.Fetch(ctx, keys, getOperator(ctx))
		},
	}
}

type ordinarySceneLoader struct {
	fetch func(keys []id.SceneID) ([]*Scene, []error)
}

func (l *ordinarySceneLoader) Load(key id.SceneID) (*Scene, error) {
	res, errs := l.fetch([]id.SceneID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinarySceneLoader) LoadAll(keys []id.SceneID) ([]*Scene, []error) {
	return l.fetch(keys)
}
