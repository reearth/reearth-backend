package app

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	err1 "github.com/reearth/reearth-backend/pkg/error"
	"github.com/stretchr/testify/assert"
)

func TestPublishedAuthMiddleware(t *testing.T) {
	h := PublishedAuthMiddleware(func(ctx context.Context, name string) (interfaces.ProjectPublishedMetadata, error) {
		if name == "active" {
			return interfaces.ProjectPublishedMetadata{
				IsBasicAuthActive: true,
				BasicAuthUsername: "fooo",
				BasicAuthPassword: "baar",
			}, nil
		} else if name == "inactive" {
			return interfaces.ProjectPublishedMetadata{
				IsBasicAuthActive: false,
				BasicAuthUsername: "fooo",
				BasicAuthPassword: "baar",
			}, nil
		}
		return interfaces.ProjectPublishedMetadata{}, err1.ErrNotFound
	})(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})

	testCases := []struct {
		Name              string
		PublishedName     string
		BasicAuthUsername string
		BasicAuthPassword string
		Error             error
	}{
		{
			Name: "empty name",
		},
		{
			Name:          "not found",
			PublishedName: "aaa",
		},
		{
			Name:          "no auth",
			PublishedName: "inactive",
		},
		{
			Name:          "auth",
			PublishedName: "active",
			Error:         echo.ErrUnauthorized,
		},
		{
			Name:              "auth with invalid credentials",
			PublishedName:     "active",
			BasicAuthUsername: "aaa",
			BasicAuthPassword: "bbb",
			Error:             echo.ErrUnauthorized,
		},
		{
			Name:              "auth with valid credentials",
			PublishedName:     "active",
			BasicAuthUsername: "fooo",
			BasicAuthPassword: "baar",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			assert := assert.New(tt)
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tc.BasicAuthUsername != "" {
				req.Header.Set(echo.HeaderAuthorization, "basic "+base64.StdEncoding.EncodeToString([]byte(tc.BasicAuthUsername+":"+tc.BasicAuthPassword)))
			}
			res := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, res)
			c.SetParamNames("name")
			c.SetParamValues(tc.PublishedName)

			err := h(c)
			if tc.Error == nil {
				assert.NoError(err)
				assert.Equal(http.StatusOK, res.Code)
				assert.Equal("test", res.Body.String())
			} else {
				assert.ErrorIs(err, tc.Error)
			}
		})
	}
}

func TestPublishedData(t *testing.T) {
	h := PublishedData(func(ctx context.Context, name string) (io.Reader, error) {
		if name == "prj" {
			return strings.NewReader("aaa"), nil
		}
		return nil, err1.ErrNotFound
	})

	testCases := []struct {
		Name          string
		PublishedName string
		Error         error
	}{
		{
			Name:  "empty",
			Error: err1.ErrNotFound,
		},
		{
			Name:          "not found",
			PublishedName: "pr",
			Error:         err1.ErrNotFound,
		},
		{
			Name:          "ok",
			PublishedName: "prj",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert := assert.New(tt)
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, res)
			c.SetParamNames("name")
			c.SetParamValues(tc.PublishedName)

			err := h(c)
			if tc.Error == nil {
				assert.NoError(err)
				assert.Equal(http.StatusOK, res.Code)
				assert.Equal("application/json", res.Header().Get(echo.HeaderContentType))
				assert.Equal("aaa", res.Body.String())
			} else {
				assert.ErrorIs(err, tc.Error)
			}
		})
	}
}

func TestPublishedIndex(t *testing.T) {
	h := PublishedIndex(func(ctx context.Context, name string, url *url.URL) (string, error) {
		if name == "prj" {
			return url.String(), nil
		}
		return "", err1.ErrNotFound
	})

	testCases := []struct {
		Name          string
		PublishedName string
		Error         error
	}{
		{
			Name:  "empty",
			Error: err1.ErrNotFound,
		},
		{
			Name:          "not found",
			PublishedName: "pr",
			Error:         err1.ErrNotFound,
		},
		{
			Name:          "ok",
			PublishedName: "prj",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			assert := assert.New(tt)
			req := httptest.NewRequest(http.MethodGet, "/aaa/bbb", nil)
			res := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, res)
			c.SetParamNames("name")
			c.SetParamValues(tc.PublishedName)

			err := h(c)
			if tc.Error == nil {
				assert.NoError(err)
				assert.Equal(http.StatusOK, res.Code)
				assert.Equal("text/html; charset=UTF-8", res.Header().Get(echo.HeaderContentType))
				assert.Equal("http://example.com/aaa/bbb", res.Body.String())
			} else {
				assert.ErrorIs(err, tc.Error)
			}
		})
	}
}
