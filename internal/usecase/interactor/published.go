package interactor

import (
	"bytes"
	"context"
	"errors"
	"html"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/cache"
	err1 "github.com/reearth/reearth-backend/pkg/error"
	"github.com/reearth/reearth-backend/pkg/log"
)

type Published struct {
	project            repo.Project
	file               gateway.File
	indexHTMLURL       *url.URL
	actualIndexHTMLURL string
	indexHTML          *cache.Cache
}

func NewPublished(project repo.Project, file gateway.File, indexHTMLURL *url.URL) interfaces.Published {
	return &Published{
		project:      project,
		file:         file,
		indexHTMLURL: indexHTMLURL,
	}
}

func (i *Published) Metadata(ctx context.Context, name string) (interfaces.ProjectPublishedMetadata, error) {
	prj, err := i.project.FindByPublicName(ctx, name)
	if err != nil || prj == nil {
		return interfaces.ProjectPublishedMetadata{}, err1.ErrNotFound
	}

	return interfaces.ProjectPublishedMetadataFrom(prj), nil
}

func (i *Published) Data(ctx context.Context, name string) (io.Reader, error) {
	r, err := i.file.ReadBuiltSceneFile(ctx, name)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (i *Published) Index(ctx context.Context, name string, u *url.URL) (string, error) {
	prj, err := i.project.FindByPublicName(ctx, name)
	if err != nil || prj == nil {
		return "", err1.ErrNotFound
	}

	md := interfaces.ProjectPublishedMetadataFrom(prj)
	i.actualIndexHTMLURL = indexURL(u, i.indexHTMLURL)

	if i.indexHTML == nil {
		i.indexHTML = cache.New(func(ctx context.Context, _ interface{}) (interface{}, error) {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, i.actualIndexHTMLURL, nil)
			if err != nil {
				return nil, err
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Errorf("published index: conn err: %s", err)
				return nil, errors.New("failed to fetch HTML")
			}
			if res.StatusCode >= 300 {
				log.Errorf("published index: status err: %d", res.StatusCode)
				return nil, errors.New("failed to fetch HTML")
			}
			defer func() {
				_ = res.Body.Close()
			}()
			str, err := io.ReadAll(res.Body)
			if err != nil {
				log.Errorf("published index: read err: %s", err)
				return "", errors.New("failed to fetch HTML")
			}
			return string(str), nil
		}, time.Hour)
	}

	indexHTML, err := i.indexHTML.Get(ctx)
	if err != nil {
		return "", err
	}

	return renderIndex(indexHTML.(string), u.String(), md), nil
}

// indexURL returns the HTML file URL from current URL and config.
func indexURL(u, indexURL *url.URL) string {
	if indexURL == nil || indexURL.String() == "" {
		u2 := *u
		u2.Path = "/published.html"
		return u2.String()
	} else if !indexURL.IsAbs() {
		u2 := *u
		u2.Path = indexURL.Path
		return u2.String()
	}
	return indexURL.String()
}

const headers = `{{if .title}}  <meta name="twitter:title" content="{{.title}}" />
  <meta property="og:title" content="{{.title}}" />{{end}}{{if .description}}
  <meta name="twitter:description" content="{{.description}}" />
  <meta property="og:description" content="{{.description}}" />{{end}}{{if .image}}
  <meta name="twitter:card" content="summary_large_image" />
  <meta name="twitter:image:src" content="{{.image}}" />
  <meta property="og:image" content="{{.image}}" />{{end}}
  <meta property="og:type" content="website" />
  <meta property="og:url" content="{{.url}}" />{{if .noindex}}
  <meta name="robots" content="noindex,nofollow" />{{end}}
`

var (
	headersTemplate = template.Must(template.New("headers").Parse(headers))
	titleRegexp     = regexp.MustCompile("<title>.+?</title>")
)

// renderIndex returns index HTML with OGP and some meta tags for the project.
func renderIndex(index, url string, d interfaces.ProjectPublishedMetadata) string {
	if d.Title != "" {
		index = titleRegexp.ReplaceAllLiteralString(index, "<title>"+html.EscapeString(d.Title)+"</title>")
	}
	var b bytes.Buffer
	_ = headersTemplate.Execute(&b,
		map[string]interface{}{
			"title":       d.Title,
			"description": d.Description,
			"image":       d.Image,
			"noindex":     d.Noindex,
			"url":         url,
		})
	return strings.Replace(index, "</head>", b.String()+"</head>", -1)
}
