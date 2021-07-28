package interactor

import (
	"context"
	"errors"
	"io"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/plugin"
)

func (i *Plugin) Upload(ctx context.Context, r io.Reader, operator *usecase.Operator) (_ *plugin.Plugin, err error) {
	tx, err := i.transaction.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	if err := i.OnlyOperator(operator); err != nil {
		return nil, err
	}

	tx.Commit()
	return nil, errors.New("not implemented")
}
