//go:generate go run github.com/99designs/gqlgen

package gql

import (
	"errors"

	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var ErrNotImplemented = errors.New("not impleneted yet")
var ErrUnauthorized = errors.New("unauthorized")

type Resolver struct {
	usecases    interfaces.Container
	controllers Container
	debug       bool
}

func NewResolver(controllers Container, debug bool) ResolverRoot {
	return &Resolver{
		usecases:    controllers.usecases,
		controllers: controllers,
		debug:       debug,
	}
}
