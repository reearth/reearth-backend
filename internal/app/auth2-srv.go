package app

import (
	"net/http"

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

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		e.Logger.Errorf("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		e.Logger.Errorf("Response Error:", re.Error.Error())
	})

	/*http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})*/

	r.Any("/authorize", func(c echo.Context) error {
		w := c.Response()
		r := c.Request()
		err := srv.HandleAuthorizeRequest(w, r)
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

	// log.Fatal(http.ListenAndServe(":9096", nil))
}
