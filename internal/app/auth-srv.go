package app

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caos/oidc/pkg/op"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	http1 "github.com/reearth/reearth-backend/internal/adapter/http"
	"github.com/reearth/reearth-backend/internal/app/oauth"
	"github.com/reearth/reearth-backend/internal/usecase/interactor"

	"github.com/golang/gddo/httputil/header"
)

func AuthEndPoints(e *echo.Echo, r *echo.Group, cfg *ServerConfig) {

	usersController := http1.NewUserController(interactor.NewUser(cfg.Repos, cfg.Gateways, cfg.Config.SignupSecret))

	ctx := context.Background()

	config := &op.Config{
		Issuer:    cfg.Config.AuthSrv.Domain,
		CryptoKey: sha256.Sum256([]byte(cfg.Config.AuthSrv.Key)),
	}
	storage := oauth.NewAuthStorage(&oauth.AuthSrvConfig{
		Domain: cfg.Config.AuthSrv.Domain,
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

	err = router.Walk(muxToEchoMapper(r))
	if err != nil {
		return
	}

	// Actual login endpoint
	r.POST("api/login", login(ctx, storage, usersController))

	// <-ctx.Done()
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

			fmt.Println("JSON -> Form: " + r.URL.String())
			err := r.ParseForm()
			if err != nil {
				return
			}

			var result map[string]string

			err = json.NewDecoder(r.Body).Decode(&result)
			if err != nil {
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
		url, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		methods, err := route.GetMethods()
		if err != nil {
			r.Any(url, echo.WrapHandler(route.GetHandler()))
			fmt.Println("ANY| " + url)
			return nil
		}

		for _, method := range methods {
			r.Add(method, url, echo.WrapHandler(route.GetHandler()))
			fmt.Println(method + "| " + url)
		}

		return nil
	}
}

type Login struct {
	Email         string `json:"username"`
	Password      string `json:"password"`
	AuthRequestID string `json:"id"`
}

func login(ctx context.Context, storage op.Storage, usersController *http1.UserController) func(ctx echo.Context) error {
	return func(ec echo.Context) error {
		// r := ec.Request()
		// w := ec.Response()
		request := new(Login)
		err := ec.Bind(&request)
		if err != nil {
			ec.Logger().Error("filed to parse login request")
			return err
		}

		if len(request.Email) == 0 || len(request.Password) == 0 {
			ec.Logger().Error("credentials are not provided")
			return nil
		}

		// check user credentials from db
		user, err := usersController.GetUserByCredentials(ctx, http1.UserCredentialInput{
			Email:    request.Email,
			Password: request.Password,
		})
		if err != nil {
			ec.Logger().Error("wrong credentials!")
			return ec.Redirect(http.StatusFound, "/login")
		}

		// Complete the auth request && set the subject
		err = storage.(*oauth.AuthStorage).CompleteAuthRequest(ctx, request.AuthRequestID, user.GetAuthByProvider("auth0").Sub)
		if err != nil {
			ec.Logger().Error("failed to complete the auth request !")
			return ec.Redirect(http.StatusFound, "/login")
		}

		return ec.Redirect(http.StatusFound, "/authorize/callback?id="+request.AuthRequestID)
	}
}
