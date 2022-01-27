package app

import (
	"context"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"github.com/reearth/reearth-backend/pkg/log"
)

type contextKey string

const (
	debugUserHeader            = "X-Reearth-Debug-User"
	contextAuth0Sub contextKey = "auth0Sub"
	contextUser     contextKey = "reearth_user"
)

// Validate the access token and inject the user clams into ctx
func jwtEchoMiddleware(cfg *ServerConfig) echo.MiddlewareFunc {

	issuerURL, err := url.Parse("https://" + cfg.Config.Auth0.Domain)
	if err != nil {
		log.Fatalf("failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 60*time.Minute)

	// Set up the validator.
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		cfg.Config.AuthClient.ISS[0],
		cfg.Config.AuthClient.AUD,
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	return echo.WrapMiddleware(middleware.CheckJWT)
}

// load claim from ctx and inject the user sub into ctx
func parseJwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()

			rawClaims := ctx.Value(jwtmiddleware.ContextKey{})
			if claims, ok := rawClaims.(*validator.ValidatedClaims); ok {

				// attach sub and access token to context
				ctx = context.WithValue(ctx, contextAuth0Sub, claims.RegisteredClaims.Subject)

				/* this claim is sent from auth0
				if user, ok := claims["https://reearth.io/user_id"].(string); ok {
					ctx = context.WithValue(ctx, contextUser, user)
				}*/

				// this one is not used!
				// ctx = context.WithValue(ctx, contextAuth0AccessToken, userProfile.Raw)
			}

			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}
