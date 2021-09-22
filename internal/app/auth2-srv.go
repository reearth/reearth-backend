package app

import (
	"io"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/go-session/session"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-oauth2/mongo.v3"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func AuthEndPoints(e *echo.Echo, r *echo.Group, cfg *ServerConfig) {

	manager := manage.NewDefaultManager()

	// token memory store
	// manager.MustTokenStorage(store.NewMemoryTokenStore())

	mongoStorage := mongo.NewTokenStore(mongo.NewConfig(
		cfg.Config.DB,
		"oauth2",
	))

	// use mongodb token store
	manager.MapTokenStorage(mongoStorage)

	// client memory store
	// TODO: Get values from Config
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		e.Logger.Errorf("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		e.Logger.Errorf("Response Error:", re.Error.Error())
	})

	r.File("/favicon.ico", "/static/favicon.ico")

	/*http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static\favicon.ico")
	})*/

	r.Any("/authorize", func(c echo.Context) error {
		w := c.Response()
		r := c.Request()

		store, err := session.Start(r.Context(), w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		if _, ok := store.Get("LoggedInUserID"); !ok {
			return c.Redirect(http.StatusFound, "/oauth/login")
		}

		err = srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return err
	})

	/*http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})*/

	r.Any("/token", func(c echo.Context) error {
		w := c.Response()
		r := c.Request()
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return err
	})

	/*http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})*/

	// login ui
	r.GET("/login", loginPage)

	// Actual login endpoint
	r.POST("/login", login)

	// register endpoints

	// log.Fatal(http.ListenAndServe(":9096", nil))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	if true {
		_ = dumpRequest(os.Stdout, "userAuthorizeHandler", r) // Ignore the error
	}
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/oauth/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}

func loginPage(ctx echo.Context) error {

	return ctx.File("static/login.html")
}

func login(ctx echo.Context) error {
	r := ctx.Request()
	w := ctx.Response()
	if true {
		_ = dumpRequest(os.Stdout, "login", r)
	}

	authStore, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	if len(username) == 0 || len(password) == 0 {
		ctx.Logger().Error("credentials are not provided")
		return nil
	}

	// TODO: check user credentials from db
	if username != "user" || password != "pass" {
		ctx.Logger().Error("wrong credentials!")
		return ctx.Redirect(http.StatusFound, "/oauth/login")
	}

	authStore.Set("LoggedInUserID", r.Form.Get("username"))
	err = authStore.Save()
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusFound, "/oauth/authorize")
}

func dumpRequest(writer io.Writer, header string, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	writer.Write([]byte("\n" + header + ": \n"))
	writer.Write(data)
	return nil
}
