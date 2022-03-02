package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reearth/reearth-backend/internal/adapter"
	http1 "github.com/reearth/reearth-backend/internal/adapter/http"
)

func publicAPI(
	ec *echo.Echo,
	r *echo.Group,
) {
	r.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	r.POST("/signup", func(c echo.Context) error {
		var inp http1.CreateUserInput
		if err := c.Bind(&inp); err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Errorf("failed to parse request body: %w", err)}
		}

		uc := adapter.Usecases(c.Request().Context())
		controller := http1.NewUserController(uc.User)

		output, err := controller.CreateUser(c.Request().Context(), inp)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, output)
	})

	r.GET("/published/:name", func(c echo.Context) error {
		name := c.Param("name")
		if name == "" {
			return echo.ErrNotFound
		}

		uc := adapter.Usecases(c.Request().Context())
		publishedController := http1.NewPublishedController(uc.Published)

		res, err := publishedController.Metadata(c.Request().Context(), name)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	})

	r.GET("/published_data/:name", func(c echo.Context) error {
		name := c.Param("name")
		if name == "" {
			return echo.ErrNotFound
		}

		uc := adapter.Usecases(c.Request().Context())
		publishedController := http1.NewPublishedController(uc.Published)

		r, err := publishedController.Data(c.Request().Context(), name)
		if err != nil {
			return err
		}

		return c.Stream(http.StatusOK, "application/json", r)
	})
}
