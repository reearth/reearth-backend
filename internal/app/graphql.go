package app

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen-contrib/gqlopencensus"
	"github.com/99designs/gqlgen-contrib/gqlopentracing"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/reearth/reearth-backend/internal/adapter/gql"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

const enableDataLoaders = true

func getOperator(ctx context.Context) *usecase.Operator {
	if v := ctx.Value(gql.ContextOperator); v != nil {
		if v2, ok := v.(*usecase.Operator); ok {
			return v2
		}
	}
	return nil
}

func dataLoaderMiddleware(container *gql.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			req := echoCtx.Request()
			ctx := req.Context()

			var dl *gql.DataLoaders
			if enableDataLoaders {
				dl = gql.NewDataLoaders(ctx, container, getOperator(ctx))
			} else {
				dl = gql.NewOrdinaryDataLoaders(ctx, container, getOperator(ctx))
			}

			ctx = context.WithValue(ctx, gql.DataLoadersKey(), dl)
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
			t := &gql.Tracer{}
			echoCtx.SetRequest(req.WithContext(gql.AttachTracer(ctx, t)))
			defer t.Print()
			return next(echoCtx)
		}
	}
}

func graphqlAPI(
	ec *echo.Echo,
	r *echo.Group,
	conf *ServerConfig,
	controllers *gql.Container,
) {
	playgroundEnabled := conf.Debug || conf.Config.Dev

	if playgroundEnabled {
		r.GET("/graphql", echo.WrapHandler(
			playground.Handler("reearth-backend", "/api/graphql"),
		))
	}

	schema := gql.NewExecutableSchema(gql.Config{
		Resolvers: gql.NewResolver(gql.ResolverConfig{
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
				return gqlerror.ErrorPathf(graphql.GetFieldContext(ctx).Path(), e.Error())
			}
			return graphql.DefaultErrorPresenter(ctx, e)
		},
	)

	r.POST("/graphql", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	}, dataLoaderMiddleware(controllers), tracerMiddleware(false))
}
