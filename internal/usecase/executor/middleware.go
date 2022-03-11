package executor

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
)

func Transaction(r repo.Transaction) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context) (err error) {
			if r == nil {
				return next(ctx)
			}

			tx, err := r.Begin()
			if err != nil {
				return err
			}
			defer func() {
				if err2 := tx.End(ctx); err == nil && err2 != nil {
					err = err2
				}
			}()
			if err = next(ctx); err == nil {
				tx.Commit()
			}
			return
		}
	}
}

func CanReadTeam(s id.TeamID, op *usecase.Operator) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context) error {
			if err := op.CanReadTeam(s); err != nil {
				return err
			}
			return next(ctx)
		}
	}
}

func CanWriteTeam(s id.TeamID, op *usecase.Operator) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context) error {
			if err := op.CanWriteTeam(s); err != nil {
				return err
			}
			return next(ctx)
		}
	}
}

func CanReadScene(s id.SceneID, op *usecase.Operator) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context) error {
			if err := op.CanReadScene(s); err != nil {
				return err
			}
			return next(ctx)
		}
	}
}

func CanWriteScene(s id.SceneID, op *usecase.Operator) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context) error {
			if err := op.CanWriteScene(s); err != nil {
				return err
			}
			return next(ctx)
		}
	}
}
