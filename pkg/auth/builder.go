package auth

import (
	"time"

	"github.com/caos/oidc/pkg/oidc"
	"github.com/reearth/reearth-backend/pkg/id"
)

type Builder struct {
	r *Request
}

func New() *Builder {
	return &Builder{r: &Request{}}
}

func (b *Builder) Build() (*Request, error) {
	if id.ID(b.r.id).IsNil() {
		return nil, id.ErrInvalidID
	}
	b.r.createdAt = time.Now()
	return b.r, nil
}

func (b *Builder) MustBuild() *Request {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *Builder) ID(id id.AuthRequestID) *Builder {
	b.r.id = id
	return b
}

func (b *Builder) NewID() *Builder {
	b.r.id = id.AuthRequestID(id.New())
	return b
}

func (b *Builder) ClientID(id string) *Builder {
	b.r.clientID = id
	return b
}

func (b *Builder) Subject(subject string) *Builder {
	b.r.subject = subject
	return b
}

func (b *Builder) Code(code string) *Builder {
	b.r.code = code
	return b
}

func (b *Builder) State(state string) *Builder {
	b.r.state = state
	return b
}

func (b *Builder) ResponseType(rt oidc.ResponseType) *Builder {
	b.r.responseType = rt
	return b
}

func (b *Builder) Scopes(scopes []string) *Builder {
	b.r.scopes = scopes
	return b
}

func (b *Builder) Audiences(audiences []string) *Builder {
	b.r.audiences = audiences
	return b
}

func (b *Builder) RedirectURI(redirectURI string) *Builder {
	b.r.redirectURI = redirectURI
	return b
}

func (b *Builder) Nonce(nonce string) *Builder {
	b.r.nonce = nonce
	return b
}

func (b *Builder) CodeChallenge(CodeChallenge *oidc.CodeChallenge) *Builder {
	b.r.codeChallenge = CodeChallenge
	return b
}

func (b *Builder) CreatedAt(createdAt time.Time) *Builder {
	b.r.createdAt = createdAt
	return b
}

func (b *Builder) AuthorizedAt(authorizedAt *time.Time) *Builder {
	b.r.authorizedAt = authorizedAt
	return b
}
