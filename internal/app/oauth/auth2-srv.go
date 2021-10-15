package oauth

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
	"github.com/reearth/reearth-backend/internal/app"
	"github.com/reearth/reearth-backend/internal/usecase/interactor"

	"github.com/golang/gddo/httputil/header"
)

func AuthEndPoints(e *echo.Echo, r *echo.Group, cfg *app.ServerConfig) {

	usersController := http1.NewUserController(interactor.NewUser(cfg.Repos, cfg.Gateways, cfg.Config.SignupSecret))

	ctx := context.Background()

	config := &op.Config{
		Issuer:    cfg.Config.AuthSrv.Domain,
		CryptoKey: sha256.Sum256([]byte(cfg.Config.AuthSrv.Key)),
	}
	storage := NewAuthStorage(cfg)
	handler, err := op.NewOpenIDProvider(
		ctx,
		config,
		storage,
		op.WithHttpInterceptors(jsonToFormHandler()),
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

	// login ui
	r.GET("/login", loginPage)

	// Actual login endpoint
	r.POST("/login", login(ctx, storage, usersController))

	// <-ctx.Done()
}

func jsonToFormHandler() func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.String() != "/oauth/token" {
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

			/*r.Form.Set("grant_type", req.GrantType)
			r.Form.Set("client_id", req.ClientID)
			r.Form.Set("redirect_uri", req.RedirectURI)
			r.Form.Set("code_verifier", req.CodeVerifier)
			r.Form.Set("code", req.Code)*/

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

func loginPage(ctx echo.Context) error {

	return ctx.File("static/login.html")
}

func login(ctx context.Context, storage op.Storage, usersController *http1.UserController) func(ctx echo.Context) error {
	return func(ec echo.Context) error {
		// r := ec.Request()
		// w := ec.Response()

		username := ec.FormValue("username")
		password := ec.FormValue("password")
		requestID := ec.FormValue("id")
		if len(username) == 0 || len(password) == 0 {
			ec.Logger().Error("credentials are not provided")
			return nil
		}

		// TODO: check user credentials from db
		user, err := usersController.GetUserByCredentials(ctx, http1.UserCredentialInput{
			Email:    username,
			Password: password,
		})
		if err != nil {
			ec.Logger().Error("wrong credentials!")
			return ec.Redirect(http.StatusFound, "/login")
		}

		// Complete the auth request && set the subject
		request, err := storage.AuthRequestByID(ctx, requestID)
		if err != nil {
			ec.Logger().Error("wrong credentials!")
			return ec.Redirect(http.StatusFound, "/login")
		}
		request.Complete(user.GetAuthByProvider("Auth0").Sub)

		return ec.Redirect(http.StatusFound, "/authorize/callback?id="+requestID)
	}
}
