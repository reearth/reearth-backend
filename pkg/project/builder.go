package project

import (
	"net/url"
	"time"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/visualizer"
)

type Builder struct {
	p *Project
}

func New() *Builder {
	return &Builder{p: &Project{publishmentStatus: PublishmentStatusPrivate}}
}

func (b *Builder) Build() (*Project, error) {
	if id.ID(b.p.id).IsNil() {
		return nil, id.ErrInvalidID
	}
	if b.p.alias != "" && !CheckAliasPattern(b.p.alias) {
		return nil, ErrInvalidAlias
	}
	if b.p.updatedAt.IsZero() {
		b.p.updatedAt = b.p.CreatedAt()
	}
	return b.p, nil
}

func (b *Builder) MustBuild() *Project {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *Builder) ID(id id.ProjectID) *Builder {
	b.p.id = id
	return b
}

func (b *Builder) NewID() *Builder {
	b.p.id = id.ProjectID(id.New())
	return b
}

func (b *Builder) IsArchived(isArchived bool) *Builder {
	b.p.isArchived = isArchived
	return b
}

func (b *Builder) IsBasicAuthActive(isBasicAuthActive bool) *Builder {
	b.p.isBasicAuthActive = isBasicAuthActive
	return b
}

func (b *Builder) BasicAuthUsername(basicAuthUsername string) *Builder {
	b.p.basicAuthUsername = basicAuthUsername
	return b
}

func (b *Builder) BasicAuthPassword(basicAuthPassword string) *Builder {
	b.p.basicAuthPassword = basicAuthPassword
	return b
}

func (b *Builder) UpdatedAt(updatedAt time.Time) *Builder {
	b.p.updatedAt = updatedAt
	return b
}

func (b *Builder) PublishedAt(publishedAt time.Time) *Builder {
	b.p.publishedAt = publishedAt
	return b
}

func (b *Builder) Name(name string) *Builder {
	b.p.name = name
	return b
}

func (b *Builder) Description(description string) *Builder {
	b.p.description = description
	return b
}

func (b *Builder) Alias(alias string) *Builder {
	b.p.alias = alias
	return b
}

func (b *Builder) ImageURL(imageURL *url.URL) *Builder {
	if imageURL == nil {
		b.p.imageURL = nil
	} else {
		imageURL2 := *imageURL
		b.p.imageURL = &imageURL2
	}
	return b
}

func (b *Builder) PublicTitle(publicTitle string) *Builder {
	b.p.publicTitle = publicTitle
	return b
}

func (b *Builder) PublicDescription(publicDescription string) *Builder {
	b.p.publicDescription = publicDescription
	return b
}

func (b *Builder) PublicImage(publicImage string) *Builder {
	b.p.publicImage = publicImage
	return b
}

func (b *Builder) PublicNoIndex(publicNoIndex bool) *Builder {
	b.p.publicNoIndex = publicNoIndex
	return b
}

func (b *Builder) Team(team id.TeamID) *Builder {
	b.p.team = team
	return b
}

func (b *Builder) Visualizer(visualizer visualizer.Visualizer) *Builder {
	b.p.visualizer = visualizer
	return b
}

func (b *Builder) PublishmentStatus(publishmentStatus PublishmentStatus) *Builder {
	b.p.publishmentStatus = publishmentStatus
	return b
}