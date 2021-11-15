package app

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/caos/oidc/pkg/op"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"github.com/reearth/reearth-backend/internal/app/appauth"
	"github.com/reearth/reearth-backend/internal/usecase/interactor"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"

	"github.com/golang/gddo/httputil/header"
)

func authEndPoints(ctx context.Context, e *echo.Echo, r *echo.Group, cfg *ServerConfig) {

	userUsecase := interactor.NewUser(cfg.Repos, cfg.Gateways, cfg.Config.SignupSecret)

	domain, err := url.Parse(cfg.Config.Auth.Domain)
	if err != nil {
		panic("not valid auth domain")
	}
	domain.Path = "/"

	config := &op.Config{
		Issuer:                domain.String(),
		CryptoKey:             sha256.Sum256([]byte(cfg.Config.Auth.Key)),
		GrantTypeRefreshToken: true,
	}
	storage := appauth.NewAuthStorage(&appauth.StorageConfig{
		Domain: domain.String(),
		Debug:  cfg.Debug,
	})
	handler, err := op.NewOpenIDProvider(
		ctx,
		config,
		storage,
		op.WithHttpInterceptors(jsonToFormHandler()),
		op.WithHttpInterceptors(setURLVarsHandler()),
		op.WithCustomKeysEndpoint(op.NewEndpoint(".well-known/jwks.json")),
	)
	if err != nil {
		e.Logger.Fatal(err)
	}

	router := handler.HttpHandler().(*mux.Router)

	if err := router.Walk(muxToEchoMapper(r)); err != nil {
		e.Logger.Fatal(err)
	}

	// Actual login endpoint
	r.POST("api/login", login(ctx, cfg, storage, userUsecase))

}

func setURLVarsHandler() func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/authorize/callback" {
				handler.ServeHTTP(w, r)
				return
			}

			r2 := mux.SetURLVars(r, map[string]string{"id": r.URL.Query().Get("id")})
			handler.ServeHTTP(w, r2)
		})
	}
}

func jsonToFormHandler() func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/oauth/token" {
				handler.ServeHTTP(w, r)
				return
			}

			if r.Header.Get("Content-Type") != "" {
				value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
				if value != "application/json" {
					// Content-Type header is not application/json
					handler.ServeHTTP(w, r)
					return
				}
			}

			if err := r.ParseForm(); err != nil {
				return
			}

			var result map[string]string

			if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			for key, value := range result {
				r.Form.Set(key, value)
			}

			handler.ServeHTTP(w, r)
		})
	}
}

func muxToEchoMapper(r *echo.Group) func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	return func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		methods, err := route.GetMethods()
		if err != nil {
			r.Any(path, echo.WrapHandler(route.GetHandler()))
			fmt.Println("ANY| " + path)
			return nil
		}

		for _, method := range methods {
			r.Add(method, path, echo.WrapHandler(route.GetHandler()))
			fmt.Println(method + "| " + path)
		}

		return nil
	}
}

type loginForm struct {
	Email         string `json:"username" form:"username"`
	Password      string `json:"password" form:"password"`
	AuthRequestID string `json:"id" form:"id"`
}

func login(ctx context.Context, cfg *ServerConfig, storage op.Storage, userUsecase interfaces.User) func(ctx echo.Context) error {
	return func(ec echo.Context) error {

		request := new(loginForm)
		err := ec.Bind(request)
		if err != nil {
			ec.Logger().Error("filed to parse login request")
			return err
		}

		authRequest, err := storage.AuthRequestByID(ctx, request.AuthRequestID)
		if err != nil {
			ec.Logger().Error("filed to parse login request")
			return err
		}

		if len(request.Email) == 0 || len(request.Password) == 0 {
			ec.Logger().Error("credentials are not provided")
			return ec.Redirect(http.StatusFound, redirectURL(authRequest.GetRedirectURI(), !cfg.Debug, request.AuthRequestID, "invalid login"))
		}

		// check user credentials from db
		user, err := userUsecase.GetUserByCredentials(ctx, interfaces.GetUserByCredentials{
			Email:    request.Email,
			Password: request.Password,
		})
		if err != nil {
			ec.Logger().Error("wrong credentials!")
			return ec.Redirect(http.StatusFound, redirectURL(authRequest.GetRedirectURI(), !cfg.Debug, request.AuthRequestID, "invalid login"))
		}

		// Complete the auth request && set the subject
		err = storage.(*appauth.Storage).CompleteAuthRequest(ctx, request.AuthRequestID, user.GetAuthByProvider("auth0").Sub)
		if err != nil {
			ec.Logger().Error("failed to complete the auth request !")
			return ec.Redirect(http.StatusFound, redirectURL(authRequest.GetRedirectURI(), !cfg.Debug, request.AuthRequestID, "invalid login"))
		}

		return ec.Redirect(http.StatusFound, "/authorize/callback?id="+request.AuthRequestID)
	}
}

func redirectURL(domain string, secure bool, requestID string, error string) string {
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.TrimPrefix(domain, "https://")

	schema := "http"
	if secure {
		schema = "https"
	}

	u := url.URL{
		Scheme: schema,
		Host:   domain,
		Path:   "login",
	}

	queryValues := u.Query()
	queryValues.Set("id", requestID)
	queryValues.Set("error", error)
	u.RawQuery = queryValues.Encode()

	return u.String()
}
