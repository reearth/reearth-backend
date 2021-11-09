package app

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/reearth/reearth-backend/pkg/log"
)

// TODO: move the authentication logic to infrastructure

type contextKey string

const (
	userProfileKey                     = "auth0_user"
	debugUserHeader                    = "X-Reearth-Debug-User"
	contextAuth0AccessToken contextKey = "auth0AccessToken"
	contextAuth0Sub         contextKey = "auth0Sub"
	contextUser             contextKey = "reearth_user"
)

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type Jwks interface {
	GetJwks(string) ([]JSONWebKeys, error)
}

type JwksSyncOnce struct {
	jwks []JSONWebKeys
	once sync.Once
}

func (jso *JwksSyncOnce) GetJwks(publicKeyURL string) ([]JSONWebKeys, error) {
	var err error
	jso.once.Do(func() {
		jso.jwks, err = fetchJwks(publicKeyURL)
	})

	if err != nil {
		return nil, err
	}

	return jso.jwks, nil
}

func fetchJwks(publicKeyURL string) ([]JSONWebKeys, error) {
	resp, err := http.Get(publicKeyURL)
	var res struct {
		Jwks []JSONWebKeys `json:"keys"`
	}

	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return nil, err
	}

	return res.Jwks, nil
}

func getPemCert(token *jwt.Token, publicKeyURL string, jwks Jwks) (string, error) {
	cert := ""
	keys, err := jwks.GetJwks(publicKeyURL)

	if err != nil {
		return cert, err
	}

	if token.Header["kid"] == nil {
		return "-----BEGIN CERTIFICATE-----\n" + keys[0].X5c[0] + "\n-----END CERTIFICATE-----", nil
	}

	for k := range keys {
		if token.Header["kid"] == keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

func parseJwtMiddleware(cfg *ServerConfig) echo.MiddlewareFunc {
	iss := urlFromDomain(cfg.Config.Auth0.Domain)
	aud := cfg.Config.Auth0.Audience

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()

			token := ctx.Value(userProfileKey)
			if userProfile, ok := token.(*jwt.Token); ok {
				claims := userProfile.Claims.(jwt.MapClaims)

				// Verify 'iss' claim
				checkIss := claims.VerifyIssuer(iss, false)
				if !checkIss {
					return errorResponse(c, "invalid issuer")
				}

				// Verify 'aud' claim
				if !verifyAudience(claims, aud) {
					return errorResponse(c, "invalid audience")
				}

				// attach sub and access token to context
				if sub, ok := claims["sub"].(string); ok {
					ctx = context.WithValue(ctx, contextAuth0Sub, sub)
				}
				if user, ok := claims["https://reearth.io/user_id"].(string); ok {
					ctx = context.WithValue(ctx, contextUser, user)
				}
				ctx = context.WithValue(ctx, contextAuth0AccessToken, userProfile.Raw)
			}

			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}

func jwtEchoMiddleware(jwks Jwks, cfg *ServerConfig) echo.MiddlewareFunc {
	jwksURL := urlFromDomain(cfg.Config.Auth0.Domain) + ".well-known/jwks.json"

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		CredentialsOptional: cfg.Debug,
		UserProperty:        userProfileKey,
		SigningMethod:       jwt.SigningMethodRS256,
		// Make jwtmiddleware return an error object by not writing ErrorHandler to ResponseWriter
		ErrorHandler: func(w http.ResponseWriter, req *http.Request, err string) {},
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			cert, err := getPemCert(token, jwksURL, jwks)
			if err != nil {
				log.Errorf("jwt: %s", err)
				return nil, err
			}
			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := jwtMiddleware.CheckJWT(c.Response(), c.Request())
			if err != nil {
				return errorResponse(c, err.Error())
			}
			return next(c)
		}
	}
}

func urlFromDomain(path string) string {
	if path == "" {
		return path
	}
	if !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		path = "https://" + path
	}
	if path[len(path)-1] != '/' {
		path += "/"
	}
	return path
}

// WORKAROUND: https://github.com/dgrijalva/jwt-go/pull/308 should be merged
func verifyAudience(claims jwt.MapClaims, aud string) bool {
	if aud == "" {
		return true
	}

	auds, ok := claims["aud"].([]string)
	if !ok {
		auds2, ok := claims["aud"].([]interface{})
		if ok {
			for _, a := range auds2 {
				if aa, ok := a.(string); ok {
					auds = append(auds, aa)
				}
			}
		} else {
			a, ok := claims["aud"].(string)
			if !ok || a == "" {
				return false
			}
			auds = append(auds, a)
		}
	}

	for _, a := range auds {
		if jwt.MapClaims(map[string]interface{}{"aud": a}).VerifyAudience(aud, true) {
			return true
		}
	}
	return false
}

func errorResponse(c echo.Context, err string) error {
	res := map[string]string{"error": err}
	return c.JSON(http.StatusUnauthorized, res)
}
