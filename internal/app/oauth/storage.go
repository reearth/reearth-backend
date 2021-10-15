package oauth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	mRand "math/rand"
	"time"

	"github.com/caos/oidc/pkg/oidc"
	"github.com/caos/oidc/pkg/op"
	"github.com/oklog/ulid"
	"github.com/reearth/reearth-backend/internal/app"
	"gopkg.in/square/go-jose.v2"
)

var (
	appConfig *app.ServerConfig
	clients   map[string]op.Client
	requests  map[string]AuthRequest
	keySet    jose.JSONWebKeySet
)

type AuthStorage struct {
	key *rsa.PrivateKey
}

func NewAuthStorage(cfg *app.ServerConfig) op.Storage {
	reader := rand.Reader
	bitSize := 2048
	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		panic(err)
	}
	appConfig = cfg

	s := &AuthStorage{
		key: key,
	}

	initData(s)

	return s
}

func initData(s *AuthStorage) {
	clients = map[string]op.Client{
		"01FH69GFQ4DFCXS5XD91JK4HZ1": &ConfClient{
			ID:              "01FH69GFQ4DFCXS5XD91JK4HZ1",
			applicationType: op.ApplicationTypeNative,
			authMethod:      oidc.AuthMethodNone,
			accessTokenType: op.AccessTokenTypeJWT,
			responseTypes:   []oidc.ResponseType{oidc.ResponseTypeCode},
			grantTypes:      []oidc.GrantType{oidc.GrantTypeCode, oidc.GrantTypeRefreshToken},
			redirectURIs:    []string{"http://localhost:3000"},
			allowedScopes:   []string{"openid", "profile", "email"},
			loginURI:        "login?id=%s",
			iDTokenLifetime: 5 * time.Minute,
			clockSkew:       0,
			devMode:         true,
		},
	}

	pubkey := s.key.Public()
	keySet = jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{Key: pubkey, Use: "sig", Algorithm: "RS256", KeyID: "1"},
		},
	}
}

func (s *AuthStorage) Health(_ context.Context) error {
	return nil
}

func (s *AuthStorage) CreateAuthRequest(_ context.Context, authReq *oidc.AuthRequest, _ string) (op.AuthRequest, error) {

	ti := time.Now()
	entropy := ulid.Monotonic(mRand.New(mRand.NewSource(ti.UnixNano())), 0)

	request := &AuthRequest{
		ID:           ulid.MustNew(ulid.Timestamp(ti), entropy).String(),
		ClientID:     authReq.ClientID,
		subject:      "",
		code:         "", // Will be set after /authorize/callback success
		state:        "state",
		scopes:       authReq.Scopes,
		ResponseType: authReq.ResponseType,
		Nonce:        authReq.Nonce,
		RedirectURI:  authReq.RedirectURI,
		createdAt:    time.Now().UTC(),
		authorizedAt: nil,
	}
	if authReq.CodeChallenge != "" {
		request.CodeChallenge = &oidc.CodeChallenge{
			Challenge: authReq.CodeChallenge,
			Method:    authReq.CodeChallengeMethod,
		}
	}

	requests[request.ID] = *request
	return request, nil
}

func (s *AuthStorage) AuthRequestByID(_ context.Context, requestID string) (op.AuthRequest, error) {

	if requestID == "" {
		return nil, errors.New("invalid id")
	}
	request, exists := requests[requestID]
	if !exists {
		return nil, errors.New("not found")
	}
	return &request, nil
}

func (s *AuthStorage) AuthRequestByCode(_ context.Context, code string) (op.AuthRequest, error) {

	if code == "" {
		return nil, errors.New("invalid code")
	}
	for _, request := range requests {
		if request.GetCode() == code {
			return &request, nil
		}
	}
	return nil, errors.New("invalid code")
}

func (s *AuthStorage) AuthRequestBySubject(_ context.Context, subject string) (op.AuthRequest, error) {

	if subject == "" {
		return nil, errors.New("invalid subject")
	}
	for _, request := range requests {
		if request.GetSubject() == subject {
			return &request, nil
		}
	}
	return nil, errors.New("invalid subject")
}

func (s *AuthStorage) SaveAuthCode(_ context.Context, requestID, code string) error {

	request, exists := requests[requestID]
	if !exists {
		return errors.New("not found")
	}

	request.code = code
	return nil
}

func (s *AuthStorage) DeleteAuthRequest(_ context.Context, requestID string) error {
	delete(clients, requestID)
	return nil
}

func (s *AuthStorage) CreateAccessToken(ctx context.Context, request op.TokenRequest) (string, time.Time, error) {
	return "id", time.Now().UTC().Add(5 * time.Hour), nil
}

func (s *AuthStorage) CreateAccessAndRefreshTokens(ctx context.Context, request op.TokenRequest, currentRefreshToken string) (accessTokenID string, newRefreshToken string, expiration time.Time, err error) {
	return "id", "refreshToken", time.Now().UTC().Add(5 * time.Minute), nil
}

func (s *AuthStorage) TokenRequestByRefreshToken(ctx context.Context, refreshToken string) (op.RefreshTokenRequest, error) {
	return nil, errors.New("not implemented")
}

func (s *AuthStorage) TerminateSession(_ context.Context, userID, clientID string) error {
	return errors.New("not implemented")
}

func (s *AuthStorage) GetSigningKey(_ context.Context, keyCh chan<- jose.SigningKey) {
	keyCh <- jose.SigningKey{Algorithm: jose.RS256, Key: s.key}
}

func (s *AuthStorage) GetKeySet(_ context.Context) (*jose.JSONWebKeySet, error) {
	return &keySet, nil
}

func (s *AuthStorage) GetKeyByIDAndUserID(_ context.Context, kid, _ string) (*jose.JSONWebKey, error) {
	return &keySet.Key(kid)[0], nil
}

func (s *AuthStorage) GetClientByClientID(_ context.Context, clientID string) (op.Client, error) {

	if clientID == "" {
		return nil, errors.New("invalid client id")
	}

	client, exists := clients[clientID]
	if !exists {
		return nil, errors.New("not found")
	}

	return client, nil

	/*if id == "web" {
		appType = op.ApplicationTypeWeb
		authMethod = oidc.AuthMethodBasic
		accessTokenType = op.AccessTokenTypeBearer
		responseTypes = []oidc.ResponseType{oidc.ResponseTypeCode}
	} else if id == "native" {
		appType = op.ApplicationTypeNative
		authMethod = oidc.AuthMethodNone
		accessTokenType = op.AccessTokenTypeBearer
		responseTypes = []oidc.ResponseType{oidc.ResponseTypeCode}
	} else {
		appType = op.ApplicationTypeUserAgent
		authMethod = oidc.AuthMethodNone
		accessTokenType = op.AccessTokenTypeJWT
		responseTypes = []oidc.ResponseType{oidc.ResponseTypeIDToken, oidc.ResponseTypeIDTokenOnly}
	}

	return &ConfClient{ID: id, applicationType: appType, authMethod: authMethod, accessTokenType: accessTokenType, responseTypes: responseTypes, devMode: true}, nil */
}

func (s *AuthStorage) AuthorizeClientIDSecret(_ context.Context, id string, _ string) error {
	return nil
}

func (s *AuthStorage) SetUserinfoFromToken(ctx context.Context, userinfo oidc.UserInfoSetter, tokenID, subject, origin string) error {
	return s.SetUserinfoFromScopes(ctx, userinfo, "", "", []string{})
}

func (s *AuthStorage) SetUserinfoFromScopes(ctx context.Context, userinfo oidc.UserInfoSetter, subject, clientID string, scopes []string) error {

	request, err := s.AuthRequestBySubject(ctx, subject)
	if err != nil {
		return err
	}

	userinfo.SetSubject(request.GetSubject())
	userinfo.SetAddress(oidc.NewUserInfoAddress("Test 789\nPostfach 2", "", "", "", "", ""))
	userinfo.SetEmail("yk.eukarya@gmail.com", true)
	userinfo.SetPhone("0791234567", true)
	userinfo.SetName("Test")
	userinfo.AppendClaims("nickname", "test")
	userinfo.AppendClaims("updated_at", "2021-10-04T18:15:46.472Z")
	userinfo.AppendClaims("picture", "https://s.gravatar.com/avatar/170df899f275cf2d8e774f7424d46430?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fyk.png")

	return nil
}

func (s *AuthStorage) GetPrivateClaimsFromScopes(_ context.Context, _, _ string, _ []string) (map[string]interface{}, error) {
	return map[string]interface{}{"private_claim": "test"}, nil
}

func (s *AuthStorage) SetIntrospectionFromToken(ctx context.Context, introspect oidc.IntrospectionResponse, tokenID, subject, clientID string) error {
	if err := s.SetUserinfoFromScopes(ctx, introspect, subject, clientID, []string{}); err != nil {
		return err
	}
	request, err := s.AuthRequestBySubject(ctx, subject)
	if err != nil {
		return err
	}
	introspect.SetClientID(request.GetClientID())
	return nil
}

func (s *AuthStorage) ValidateJWTProfileScopes(ctx context.Context, userID string, scope []string) ([]string, error) {
	return scope, nil
}
