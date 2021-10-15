package oauth

import (
	"time"

	"github.com/caos/oidc/pkg/oidc"
)

var essentialScopes = []string{"openid", "profile", "email"}

type AuthRequest struct {
	ID            string
	ClientID      string
	subject       string
	code          string
	state         string
	ResponseType  oidc.ResponseType
	scopes        []string
	RedirectURI   string
	Nonce         string
	CodeChallenge *oidc.CodeChallenge
	createdAt     time.Time
	authorizedAt  *time.Time
}

func (a *AuthRequest) GetID() string {
	return a.ID
}

func (a *AuthRequest) GetACR() string {
	return ""
}

func (a *AuthRequest) GetAMR() []string {
	return []string{
		"password",
	}
}

func (a *AuthRequest) GetAudience() []string {
	audiences := []string{
		appConfig.Config.AuthSrv.Domain,
	}

	if appConfig.Debug {
		audiences = append(audiences, "http://localhost:8080")
	}

	return audiences
}

func (a *AuthRequest) GetAuthTime() time.Time {
	return a.createdAt
}

func (a *AuthRequest) GetClientID() string {
	return a.ClientID
}

func (a *AuthRequest) GetCode() string {
	return a.code
}

func (a *AuthRequest) GetState() string {
	return a.state
}

func (a *AuthRequest) GetCodeChallenge() *oidc.CodeChallenge {
	return a.CodeChallenge
}

func (a *AuthRequest) GetNonce() string {
	return a.Nonce
}

func (a *AuthRequest) GetRedirectURI() string {
	return a.RedirectURI
}

func (a *AuthRequest) GetResponseType() oidc.ResponseType {
	return a.ResponseType
}

func (a *AuthRequest) GetScopes() []string {
	return unique(append(a.scopes, essentialScopes...))
}

func (a *AuthRequest) SetCurrentScopes(scopes []string) {
	a.scopes = unique(append(scopes, essentialScopes...))
}

func (a *AuthRequest) GetSubject() string {
	return a.subject // return "auth0|60acc23af5de37006a5d8229"
}

func (a *AuthRequest) Done() bool {
	return a.authorizedAt != nil
}

func unique(list []string) []string {
	allKeys := make(map[string]bool)
	var uniqueList []string
	for _, item := range list {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			uniqueList = append(uniqueList, item)
		}
	}
	return uniqueList
}
