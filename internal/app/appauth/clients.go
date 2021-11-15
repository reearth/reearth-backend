package appauth

import (
	"time"

	"github.com/caos/oidc/pkg/oidc"
	"github.com/caos/oidc/pkg/op"
)

func initLocalClient(devMode bool) op.Client {
	return &ConfClient{
		ID:              "01FH69GFQ4DFCXS5XD91JK4HZ1",
		applicationType: op.ApplicationTypeWeb,
		authMethod:      oidc.AuthMethodNone,
		accessTokenType: op.AccessTokenTypeJWT,
		responseTypes:   []oidc.ResponseType{oidc.ResponseTypeCode},
		grantTypes:      []oidc.GrantType{oidc.GrantTypeCode, oidc.GrantTypeRefreshToken},
		redirectURIs:    []string{"http://localhost:3000"},
		allowedScopes:   []string{"openid", "profile", "email"},
		loginURI:        "http://localhost:3000/login?id=%s",
		iDTokenLifetime: 5 * time.Minute,
		clockSkew:       0,
		devMode:         devMode,
	}
}
