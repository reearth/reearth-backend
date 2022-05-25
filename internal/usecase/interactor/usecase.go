package interactor

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
)

type uc struct {
	tx             bool
	readableTeams  id.TeamIDList
	writableTeams  id.TeamIDList
	readableScenes id.SceneIDList
	writableScenes id.SceneIDList
}

func Usecase() *uc {
	return &uc{}
}

func (u *uc) CanReadTeams(ids ...id.TeamID) *uc {
	u.readableTeams = ids
	return u
}

func (u *uc) CanWriteTeams(ids ...id.TeamID) *uc {
	u.writableTeams = ids
	return u
}

func (u *uc) CanReadScenes(ids ...id.SceneID) *uc {
	u.readableScenes = ids
	return u
}

func (u *uc) CanWriteSceens(ids ...id.SceneID) *uc {
	u.writableScenes = ids
	return u
}

func (u *uc) Transaction() *uc {
	u.tx = true
	return u
}

func Run0(ctx context.Context, op *usecase.Operator, r *repo.Container, e *uc, f func(context.Context) error) (err error) {
	_, _, _, err = Run3(
		ctx, op, r, e,
		func(ctx context.Context) (_, _, _ any, err error) {
			err = f(ctx)
			return
		})
	return
}

func Run1[A any](ctx context.Context, op *usecase.Operator, r *repo.Container, e *uc, f func(context.Context) (A, error)) (a A, err error) {
	a, _, _, err = Run3(
		ctx, op, r, e,
		func(ctx context.Context) (a A, _, _ any, err error) {
			a, err = f(ctx)
			return
		})
	return
}

func Run2[A, B any](ctx context.Context, op *usecase.Operator, r *repo.Container, e *uc, f func(context.Context) (A, B, error)) (a A, b B, err error) {
	a, b, _, err = Run3(
		ctx, op, r, e,
		func(ctx context.Context) (a A, b B, _ any, err error) {
			a, b, err = f(ctx)
			return
		})
	return
}

func Run3[A, B, C any](ctx context.Context, op *usecase.Operator, r *repo.Container, e *uc, f func(context.Context) (A, B, C, error)) (_ A, _ B, _ C, err error) {
	if err = e.checkPermission(op); err != nil {
		return
	}

	if e.tx && r.Transaction != nil {
		tx, err2 := r.Transaction.Begin()
		if err2 != nil {
			err = err2
			return
		}
		defer func() {
			if err == nil {
				tx.Commit()
			}
			if err2 := tx.End(ctx); err == nil && err2 != nil {
				err = err2
			}
		}()
	}

	return f(ctx)
}

func (u *uc) checkPermission(op *usecase.Operator) error {
	ok := true
	if u.readableTeams != nil {
		ok = op.IsReadableTeam(u.readableTeams...)
	}
	if ok && u.writableTeams != nil {
		ok = op.IsWritableTeam(u.writableTeams...)
	}
	if ok && u.readableScenes != nil {
		ok = op.IsReadableScene(u.readableScenes...)
	}
	if ok && u.writableScenes != nil {
		ok = op.IsWritableScene(u.writableScenes...)
	}
	if !ok {
		return interfaces.ErrOperationDenied
	}
	return nil
}
