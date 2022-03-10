package app

import (
	"context"
	"crypto/subtle"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reearth/reearth-backend/internal/adapter"
	http1 "github.com/reearth/reearth-backend/internal/adapter/http"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

func Ping() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	}
}

func Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inp http1.SignupInput
		if err := c.Bind(&inp); err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Errorf("failed to parse request body: %w", err)}
		}

		uc := adapter.Usecases(c.Request().Context())
		controller := http1.NewUserController(uc.User)

		output, err := controller.Signup(c.Request().Context(), inp)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, output)
	}
}

func PasswordReset() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inp http1.PasswordResetInput
		if err := c.Bind(&inp); err != nil {
			return err
		}

		uc := adapter.Usecases(c.Request().Context())
		controller := http1.NewUserController(uc.User)

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
	}
}

func StartSignupVerify() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inp http1.CreateVerificationInput
		if err := c.Bind(&inp); err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: fmt.Errorf("failed to parse request body: %w", err)}
		}

		uc := adapter.Usecases(c.Request().Context())
		controller := http1.NewUserController(uc.User)

		if err := controller.CreateVerification(c.Request().Context(), inp); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func SignupVerify() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.Param("code")
		if len(code) == 0 {
			return echo.ErrBadRequest
		}

		uc := adapter.Usecases(c.Request().Context())
		controller := http1.NewUserController(uc.User)

		output, err := controller.VerifyUser(c.Request().Context(), code)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, output)
	}
}

func PublishedMetadata() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if name == "" {
			return rerror.ErrNotFound
		}

		contr, err := publishedController(c)
		if err != nil {
			return err
		}

		res, err := contr.Metadata(c.Request().Context(), name)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}

func PublishedData() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if name == "" {
			return rerror.ErrNotFound
		}

		contr, err := publishedController(c)
		if err != nil {
			return err
		}

		r, err := contr.Data(c.Request().Context(), name)
		if err != nil {
			return err
		}

		return c.Stream(http.StatusOK, "application/json", r)
	}
}

func PublishedIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		contr, err := publishedController(c)
		if err != nil {
			return err
		}

		index, err := contr.Index(c.Request().Context(), c.Param("name"), &url.URL{
			Scheme: "http",
			Host:   c.Request().Host,
			Path:   c.Request().URL.Path,
		})
		if err != nil {
			return err
		}
		if index == "" {
			return rerror.ErrNotFound
		}
		return c.HTML(http.StatusOK, index)
	}
}

func PublishedAuthMiddleware() echo.MiddlewareFunc {
	key := struct{}{}
	return middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(user string, password string, c echo.Context) (bool, error) {
			md, ok := c.Request().Context().Value(key).(interfaces.ProjectPublishedMetadata)
			if !ok {
				return true, echo.ErrNotFound
			}
			return !md.IsBasicAuthActive || subtle.ConstantTimeCompare([]byte(user), []byte(md.BasicAuthUsername)) == 1 && subtle.ConstantTimeCompare([]byte(password), []byte(md.BasicAuthPassword)) == 1, nil
		},
		Skipper: func(c echo.Context) bool {
			name := c.Param("name")
			if name == "" {
				return true
			}

			contr, err := publishedController(c)
			if err != nil {
				return false
			}

			md, err := contr.Metadata(c.Request().Context(), name)
			if err != nil {
				return true
			}

			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key, md)))
			return !md.IsBasicAuthActive
		},
	})
}

func publishedController(c echo.Context) (*http1.PublishedController, error) {
	uc := adapter.Usecases(c.Request().Context())
	if uc.Published == nil {
		return nil, rerror.ErrNotFound
	}
	return http1.NewPublishedController(uc.Published), nil
}
