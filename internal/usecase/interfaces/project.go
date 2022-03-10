package interfaces

import (
	"context"
	"errors"
	"net/url"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/project"
	"github.com/reearth/reearth-backend/pkg/visualizer"
)

type CreateProjectParam struct {
	TeamID      id.TeamID
	Visualizer  visualizer.Visualizer
	Name        *string
	Description *string
	ImageURL    *url.URL
	Alias       *string
	Archived    *bool
}

type UpdateProjectParam struct {
	ID                id.ProjectID
	Name              *string
	Description       *string
	Alias             *string
	Archived          *bool
	IsBasicAuthActive *bool
	BasicAuthUsername *string
	BasicAuthPassword *string
	ImageURL          *url.URL
	PublicTitle       *string
	PublicDescription *string
	PublicImage       *string
	PublicNoIndex     *bool
	DeletePublicImage bool
	DeleteImageURL    bool
}

type PublishProjectParam struct {
	ID     id.ProjectID
	Alias  *string
	Status project.PublishmentStatus
}

var (
	ErrProjectAliasIsNotSet    error = errors.New("project alias is not set")
	ErrProjectAliasAlreadyUsed error = errors.New("project alias is already used by another project")
)

type Project interface {
	Create(context.Context, CreateProjectParam, *usecase.Operator) (*project.Project, error)
	Update(context.Context, UpdateProjectParam, *usecase.Operator) (*project.Project, error)
	Publish(context.Context, PublishProjectParam, *usecase.Operator) (*project.Project, error)
	CheckAlias(context.Context, string) (bool, error)
	Delete(context.Context, id.ProjectID, *usecase.Operator) error
}
