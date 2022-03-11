package executor

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
)

type Handler func(context.Context) error
type Middleware func(Handler) Handler

type Exectutor struct {
	txr repo.Transaction
}

func New(r repo.Transaction) *Exectutor {
	return &Exectutor{
		txr: r,
	}
}

func (e *Exectutor) Run(ctx context.Context, handler Handler, middleware ...Middleware) error {
	return applyMiddleware(handler, middleware...)(ctx)
}

func (e *Exectutor) Transaction() Middleware {
	return Transaction(e.txr)
}

func (*Exectutor) CanReadTeam(s id.TeamID, op *usecase.Operator) Middleware {
	return CanReadTeam(s, op)
}

func (*Exectutor) CanWriteTeam(s id.TeamID, op *usecase.Operator) Middleware {
	return CanWriteTeam(s, op)
}

func (*Exectutor) CanReadScene(s id.SceneID, op *usecase.Operator) Middleware {
	return CanReadScene(s, op)
}

func (*Exectutor) CanWriteScene(s id.SceneID, op *usecase.Operator) Middleware {
	return CanWriteScene(s, op)
}

func applyMiddleware(h Handler, middleware ...Middleware) Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}
