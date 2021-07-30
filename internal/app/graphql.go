package app

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen-contrib/gqlopencensus"
	"github.com/99designs/gqlgen-contrib/gqlopentracing"
	graphql1 "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/reearth/reearth-backend/internal/adapter/graphql"
	infra_graphql "github.com/reearth/reearth-backend/internal/graphql"
	"github.com/reearth/reearth-backend/internal/graphql/dataloader"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

const enableDataLoaders = true

func getOperator(ctx context.Context) *usecase.Operator {
	if v := ctx.Value(infra_graphql.ContextOperator); v != nil {
		if v2, ok := v.(*usecase.Operator); ok {
			return v2
		}
	}
	return nil
}

func dataLoaderMiddleware(container *graphql.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			req := echoCtx.Request()
			ctx := req.Context()

			var dl *dataloader.DataLoaders
			if enableDataLoaders {
				dl = dataloader.NewDataLoaders(ctx, container, getOperator(ctx))
			} else {
				dl = dataloader.NewOrdinaryDataLoaders(ctx, container, getOperator(ctx))
			}

			ctx = context.WithValue(ctx, dataloader.DataLoadersKey(), dl)
			echoCtx.SetRequest(req.WithContext(ctx))
			return next(echoCtx)
		}
	}
}

func tracerMiddleware(enabled bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			if !enabled {
				return next(echoCtx)
			}
			req := echoCtx.Request()
			ctx := req.Context()
			t := &infra_graphql.Tracer{}
			echoCtx.SetRequest(req.WithContext(infra_graphql.AttachTracer(ctx, t)))
			defer t.Print()
			return next(echoCtx)
		}
	}
}

func graphqlAPI(
	ec *echo.Echo,
	r *echo.Group,
	conf *ServerConfig,
	controllers *graphql.Container,
) {
	playgroundEnabled := conf.Debug || conf.Config.Dev

	if playgroundEnabled {
		r.GET("/graphql", echo.WrapHandler(
			playground.Handler("reearth-backend", "/api/graphql"),
		))
	}

	schema := infra_graphql.NewExecutableSchema(infra_graphql.Config{
		Resolvers: infra_graphql.NewResolver(infra_graphql.ResolverConfig{
			Controllers: controllers,
			Debug:       conf.Debug,
		}),
	})

	srv := handler.NewDefaultServer(schema)
	srv.Use(gqlopentracing.Tracer{})
	srv.Use(gqlopencensus.Tracer{})
	if conf.Config.GraphQL.ComplexityLimit > 0 {
		srv.Use(extension.FixedComplexityLimit(conf.Config.GraphQL.ComplexityLimit))
	}
	if playgroundEnabled {
		srv.Use(extension.Introspection{})
	}
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(30),
	})
	srv.SetErrorPresenter(
		// show more detailed error messgage in debug mode
		func(ctx context.Context, e error) *gqlerror.Error {
			if conf.Debug {
				var ierr *rerror.ErrInternal
				if errors.As(e, &ierr) {
					if err2 := ierr.Unwrap(); err2 != nil {
						// TODO: display stacktrace with xerrors
						ec.Logger.Errorf("%+v", err2)
					}
				}
				return gqlerror.ErrorPathf(graphql1.GetFieldContext(ctx).Path(), e.Error())
			}
			return graphql1.DefaultErrorPresenter(ctx, e)
		},
	)

	r.POST("/graphql", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	}, dataLoaderMiddleware(controllers), tracerMiddleware(false))
}
