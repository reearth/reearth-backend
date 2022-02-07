package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	http1 "github.com/reearth/reearth-backend/internal/adapter/http"
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/reearth/reearth-backend/internal/usecase/interactor"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
)

func publicAPI(
	ec *echo.Echo,
	r *echo.Group,
	conf *Config,
	repos *repo.Container,
	gateways *gateway.Container,
) {
	controller := http1.NewUserController(
		interactor.NewUser(repos, gateways, conf.SignupSecret),
	)
	publishedController := http1.NewPublishedController(
		interactor.NewPublished(repos.Project, gateways.File, ""),
	)

	r.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	r.POST("/signup", func(c echo.Context) error {
		var inp http1.SignupInput
		if err := c.Bind(&inp); err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Errorf("failed to parse request body: %w", err)}
		}

		output, err := controller.Signup(c.Request().Context(), inp)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, output)
	})

	r.POST("/password-reset", func(c echo.Context) error {
		var inp http1.PasswordResetInput
		if err := c.Bind(&inp); err != nil {
			return err
		}

		if len(inp.Email) > 0 {
			if err := controller.StartPasswordReset(c.Request().Context(), inp); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, true)
		}

		if len(inp.Token) > 0 && len(inp.Password) > 0 {
			if err := controller.PasswordReset(c.Request().Context(), inp); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, true)
		}

		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Bad reset password request"}
	})

	r.POST("/signup/verify", func(c echo.Context) error {
		var inp http1.CreateVerificationInput
		if err := c.Bind(&inp); err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Errorf("failed to parse request body: %w", err)}
		}
		if err := controller.CreateVerification(c.Request().Context(), inp); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	r.POST("/signup/verify/:code", func(c echo.Context) error {
		code := c.Param("code")
		if len(code) == 0 {
			return echo.ErrBadRequest
		}

		output, err := controller.VerifyUser(c.Request().Context(), code)
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

		r, err := publishedController.Data(c.Request().Context(), name)
		if err != nil {
			return err
		}

		return c.Stream(http.StatusOK, "application/json", r)
	})
}
