package appauth

import (
	"fmt"
	"time"

	"github.com/caos/oidc/pkg/oidc"
	"github.com/caos/oidc/pkg/op"
)

type ConfClient struct {
	ID                 string
	applicationType    op.ApplicationType
	authMethod         oidc.AuthMethod
	accessTokenType    op.AccessTokenType
	responseTypes      []oidc.ResponseType
	grantTypes         []oidc.GrantType
	allowedScopes      []string
	redirectURIs       []string
	logoutRedirectURIs []string
	loginURI           string
	idTokenLifetime    time.Duration
	clockSkew          time.Duration
	devMode            bool
}

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
		idTokenLifetime: 5 * time.Minute,
		clockSkew:       0,
		devMode:         devMode,
	}
}

func (c *ConfClient) GetID() string {
	return c.ID
}

func (c *ConfClient) RedirectURIs() []string {
	return c.redirectURIs
}

func (c *ConfClient) PostLogoutRedirectURIs() []string {
	return c.logoutRedirectURIs
}

func (c *ConfClient) LoginURL(id string) string {
	return fmt.Sprintf(c.loginURI, id)
}

func (c *ConfClient) ApplicationType() op.ApplicationType {
	return c.applicationType
}

func (c *ConfClient) AuthMethod() oidc.AuthMethod {
	return c.authMethod
}

func (c *ConfClient) IDTokenLifetime() time.Duration {
	return c.idTokenLifetime
}

func (c *ConfClient) AccessTokenType() op.AccessTokenType {
	return c.accessTokenType
}

func (c *ConfClient) ResponseTypes() []oidc.ResponseType {
	return c.responseTypes
}

func (c *ConfClient) GrantTypes() []oidc.GrantType {
	return c.grantTypes
}

func (c *ConfClient) DevMode() bool {
	return c.devMode
}

func (c *ConfClient) RestrictAdditionalIdTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

func (c *ConfClient) RestrictAdditionalAccessTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

func (c *ConfClient) IsScopeAllowed(scope string) bool {
	for _, clientScope := range c.allowedScopes {
		if clientScope == scope {
			return true
		}
	}
	return false
}

func (c *ConfClient) IDTokenUserinfoClaimsAssertion() bool {
	return false
}

func (c *ConfClient) ClockSkew() time.Duration {
	return c.clockSkew
}
