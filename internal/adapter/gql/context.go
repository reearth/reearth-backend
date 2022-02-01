package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/user"
)

type ContextKey string

const (
	ContextUser        ContextKey = "user"
	ContextOperator    ContextKey = "operator"
	ContextSub         ContextKey = "sub"
	contextUsecases    ContextKey = "usecases"
	contextLoaders     ContextKey = "loaders"
	contextDataLoaders ContextKey = "dataloaders"
)

func getUser(ctx context.Context) *user.User {
	if v := ctx.Value(ContextUser); v != nil {
		if u, ok := v.(*user.User); ok {
			return u
		}
	}
	return nil
}

func getLang(ctx context.Context, lang *string) string {
	if lang != nil && *lang != "" {
		return *lang
	}

	u := getUser(ctx)
	if u == nil {
		return "en" // default language
	}

	l := u.Lang()
	if l.IsRoot() {
		return "en" // default language
	}

	return l.String()
}

func getOperator(ctx context.Context) *usecase.Operator {
	if v := ctx.Value(ContextOperator); v != nil {
		if v2, ok := v.(*usecase.Operator); ok {
			return v2
		}
	}
	return nil
}

func getSub(ctx context.Context) string {
	if v := ctx.Value(ContextSub); v != nil {
		if v2, ok := v.(string); ok {
			return v2
		}
	}
	return ""
}

func AttachUsecases(ctx context.Context, u *interfaces.Container, enableDataLoaders bool) context.Context {
	ctx = context.WithValue(ctx, contextUsecases, u)
	loaders := NewLoaders(u)
	ctx = context.WithValue(ctx, contextLoaders, loaders)
	ctx = context.WithValue(ctx, contextDataLoaders, loaders.DataLoadersWith(ctx, enableDataLoaders))
	return ctx
}

func usecases(ctx context.Context) *interfaces.Container {
	if v := ctx.Value(contextUsecases); v != nil {
		if v2, ok := v.(*interfaces.Container); ok {
			return v2
		}
	}
	return nil
}

func loaders(ctx context.Context) *Loaders {
	if v := ctx.Value(contextLoaders); v != nil {
		if v2, ok := v.(*Loaders); ok {
			return v2
		}
	}
	return nil
}

func dataLoaders(ctx context.Context) *DataLoaders {
	if v := ctx.Value(contextLoaders); v != nil {
		if v2, ok := v.(*DataLoaders); ok {
			return v2
		}
	}
	return nil
}
