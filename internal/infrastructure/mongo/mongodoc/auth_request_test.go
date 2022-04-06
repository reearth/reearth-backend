package mongodoc

import (
	"testing"
	"time"

	"github.com/caos/oidc/pkg/oidc"

	"github.com/reearth/reearth-backend/pkg/auth"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthRequest(t *testing.T) {

	req := auth.NewRequest().
		NewID().
		ClientID("client id").
		State("state").
		ResponseType("response type").
		Scopes([]string{"scope"}).
		Audiences([]string{"audience"}).
		RedirectURI("redirect uri").
		Nonce("nonce").
		CodeChallenge(&oidc.CodeChallenge{
			Challenge: "challenge",
			Method:    "S256",
		}).
		CreatedAt(time.Now().UTC()).
		AuthorizedAt(nil).
		MustBuild()

	tests := []struct {
		name   string
		target *auth.Request
		want   *AuthRequestDocument
		want1  string
	}{
		{
			name:   "new auth request",
			target: req,
			want: &AuthRequestDocument{
				ID:           req.GetID(),
				ClientID:     req.GetClientID(),
				Subject:      req.GetSubject(),
				Code:         req.GetCode(),
				State:        req.GetState(),
				ResponseType: string(req.GetResponseType()),
				Scopes:       req.GetScopes(),
				Audiences:    req.GetAudience(),
				RedirectURI:  req.GetRedirectURI(),
				Nonce:        req.GetNonce(),
				CodeChallenge: &CodeChallengeDocument{
					Challenge: "challenge",
					Method:    "S256",
				},
				CreatedAt:    req.CreatedAt(),
				AuthorizedAt: req.AuthorizedAt(),
			},
			want1: req.GetID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1 := NewAuthRequest(tt.target)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

/*
func TestAuthRequestDocument_Model(t *testing.T) {


	id, _ := id.AuthRequestIDFrom("01f2r7kg1fvvffp0gmexgy5hxy")
	mockTime := time.Now().UTC()

	authRequest, _ := auth.NewRequest().
		ID(id).
		ClientID("client id").
		Subject("subject").
		Code("code").
		State("state").
		ResponseType("response type").
		Scopes([]string{"openid", "profile", "email"}).
		Audiences([]string{"audiences"}).
		RedirectURI("redirect uri").
		Nonce("nonce").
		CreatedAt(mockTime).
		CodeChallenge(&oidc.CodeChallenge{

			Challenge: "challenge",
			Method:    "S256",
		}).Build()

	tests := []struct {
		name    string
		target  *AuthRequestDocument
		want    *auth.Request
		wantErr bool
	}{
		{
			name: "auth request document model",
			target: &AuthRequestDocument{
				ID:           authRequest.GetID(),
				ClientID:     authRequest.GetClientID(),
				Subject:      authRequest.GetSubject(),
				Code:         authRequest.GetCode(),
				State:        authRequest.GetState(),
				ResponseType: "response type",
				Scopes:       authRequest.GetScopes(),
				Audiences:    authRequest.GetAudience(),
				RedirectURI:  authRequest.GetRedirectURI(),
				Nonce:        authRequest.GetNonce(),
				CodeChallenge: &CodeChallengeDocument{
					Challenge: "challenge",
					Method:    "S256",
				},
				CreatedAt: authRequest.CreatedAt(),
			},
			want: authRequest,

			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := &AuthRequestDocument{
				ID:            tt.target.ID,
				ClientID:      tt.target.ClientID,
				Subject:       tt.target.Subject,
				Code:          tt.target.Code,
				State:         tt.target.State,
				ResponseType:  tt.target.ResponseType,
				Scopes:        tt.target.Scopes,
				Audiences:     tt.target.Audiences,
				RedirectURI:   tt.target.RedirectURI,
				Nonce:         tt.target.Nonce,
				CodeChallenge: tt.target.CodeChallenge,
				CreatedAt:     tt.target.CreatedAt,
				AuthorizedAt:  tt.target.AuthorizedAt,
			}
			got, err := d.Model()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

*/
