package appauth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/caos/oidc/pkg/oidc"
	"github.com/caos/oidc/pkg/op"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/auth"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/user"
	"gopkg.in/square/go-jose.v2"
)

type Storage struct {
	lock             sync.Mutex
	appConfig        *StorageConfig
	getUserBySubject func(context.Context, string) (*user.User, error)
	clients          map[string]op.Client
	requests         repo.AuthRequest
	keySet           jose.JSONWebKeySet
	key              *rsa.PrivateKey
	sigKey           jose.SigningKey
}

type StorageConfig struct {
	Domain string `default:"http://localhost:8080"`
	Debug  bool
	DN     *AuthDNConfig
}

type AuthDNConfig struct {
	CommonName         string
	Organization       []string
	OrganizationalUnit []string
	Country            []string
	Province           []string
	Locality           []string
	StreetAddress      []string
	PostalCode         []string
}

var dummyName = pkix.Name{
	CommonName:         "Dummy company, INC.",
	Organization:       []string{"Dummy company, INC."},
	OrganizationalUnit: []string{"Dummy OU"},
	Country:            []string{"US"},
	Province:           []string{"Dummy"},
	Locality:           []string{"Dummy locality"},
	StreetAddress:      []string{"Dummy street"},
	PostalCode:         []string{"1"},
}

func NewAuthStorage(cfg *StorageConfig, request repo.AuthRequest, getUserBySubject func(context.Context, string) (*user.User, error)) op.Storage {

	client := auth.NewLocalClient(cfg.Debug)

	name := dummyName
	if cfg.DN != nil {
		name = pkix.Name{
			CommonName:         cfg.DN.CommonName,
			Organization:       cfg.DN.Organization,
			OrganizationalUnit: cfg.DN.OrganizationalUnit,
			Country:            cfg.DN.Country,
			Province:           cfg.DN.Province,
			Locality:           cfg.DN.Locality,
			StreetAddress:      cfg.DN.StreetAddress,
			PostalCode:         cfg.DN.PostalCode,
		}
	}

	key, sigKey, keySet := initKeys(name)

	return &Storage{
		appConfig:        cfg,
		getUserBySubject: getUserBySubject,
		requests:         request,
		key:              key,
		sigKey:           sigKey,
		keySet:           keySet,
		clients: map[string]op.Client{
			client.GetID(): client,
		},
	}
}

func initKeys(name pkix.Name) (*rsa.PrivateKey, jose.SigningKey, jose.JSONWebKeySet) {

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      name,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(100, 0, 0),
		IsCA:         true,
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, key.Public(), key)
	if err != nil {
		panic("failed to create the cert")
	}

	cert, err = x509.ParseCertificate(caBytes)
	if err != nil {
		panic("failed to create the cert")
	}

	keyID := "RE01"
	sk := jose.SigningKey{
		Algorithm: jose.RS256,
		Key:       jose.JSONWebKey{Key: key, Use: "sig", Algorithm: string(jose.RS256), KeyID: keyID, Certificates: []*x509.Certificate{cert}},
	}

	return key, sk, jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{Key: key.Public(), Use: "sig", Algorithm: string(jose.RS256), KeyID: keyID, Certificates: []*x509.Certificate{cert}},
		},
	}
}

func (s *Storage) Health(_ context.Context) error {
	return nil
}

func (s *Storage) CreateAuthRequest(ctx context.Context, authReq *oidc.AuthRequest, _ string) (op.AuthRequest, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	audiences := []string{
		s.appConfig.Domain,
	}
	if s.appConfig.Debug {
		audiences = append(audiences, "http://localhost:8080")
	}

	var cc *oidc.CodeChallenge
	if authReq.CodeChallenge != "" {
		cc = &oidc.CodeChallenge{
			Challenge: authReq.CodeChallenge,
			Method:    authReq.CodeChallengeMethod,
		}
	}
	var request = auth.New().
		NewID().
		ClientID(authReq.ClientID).
		State(authReq.State).
		ResponseType(authReq.ResponseType).
		Scopes(authReq.Scopes).
		Audiences(audiences).
		RedirectURI(authReq.RedirectURI).
		Nonce(authReq.Nonce).
		CodeChallenge(cc).
		CreatedAt(time.Now().UTC()).
		AuthorizedAt(nil).
		MustBuild()

	if err := s.requests.Save(ctx, request); err != nil {
		return nil, err
	}
	return request, nil
}

func (s *Storage) AuthRequestByID(ctx context.Context, requestID string) (op.AuthRequest, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if requestID == "" {
		return nil, errors.New("invalid id")
	}
	reqId, err := id.AuthRequestIDFrom(requestID)
	if err != nil {
		return nil, err
	}
	request, err := s.requests.FindByID(ctx, reqId)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (s *Storage) AuthRequestByCode(ctx context.Context, code string) (op.AuthRequest, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if code == "" {
		return nil, errors.New("invalid code")
	}
	return s.requests.FindByCode(ctx, code)
}

func (s *Storage) AuthRequestBySubject(ctx context.Context, subject string) (op.AuthRequest, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if subject == "" {
		return nil, errors.New("invalid subject")
	}

	return s.requests.FindBySubject(ctx, subject)
}

func (s *Storage) SaveAuthCode(ctx context.Context, requestID, code string) error {

	request, err := s.AuthRequestByID(ctx, requestID)
	if err != nil {
		return err
	}
	request2 := request.(*auth.Request)
	request2.SetCode(code)
	err = s.updateRequest(ctx, requestID, *request2)
	return err
}

func (s *Storage) DeleteAuthRequest(_ context.Context, requestID string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.clients, requestID)
	return nil
}

func (s *Storage) CreateAccessToken(_ context.Context, _ op.TokenRequest) (string, time.Time, error) {
	return "id", time.Now().UTC().Add(5 * time.Hour), nil
}

func (s *Storage) CreateAccessAndRefreshTokens(_ context.Context, request op.TokenRequest, _ string) (accessTokenID string, newRefreshToken string, expiration time.Time, err error) {
	authReq := request.(*auth.Request)
	return "id", authReq.GetID(), time.Now().UTC().Add(5 * time.Minute), nil
}

func (s *Storage) TokenRequestByRefreshToken(ctx context.Context, refreshToken string) (op.RefreshTokenRequest, error) {
	r, err := s.AuthRequestByID(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return r.(op.RefreshTokenRequest), err
}

func (s *Storage) TerminateSession(_ context.Context, _, _ string) error {
	return errors.New("not implemented")
}

func (s *Storage) GetSigningKey(_ context.Context, keyCh chan<- jose.SigningKey) {
	keyCh <- s.sigKey
}

func (s *Storage) GetKeySet(_ context.Context) (*jose.JSONWebKeySet, error) {
	return &s.keySet, nil
}

func (s *Storage) GetKeyByIDAndUserID(_ context.Context, kid, _ string) (*jose.JSONWebKey, error) {
	return &s.keySet.Key(kid)[0], nil
}

func (s *Storage) GetClientByClientID(_ context.Context, clientID string) (op.Client, error) {

	if clientID == "" {
		return nil, errors.New("invalid client id")
	}

	client, exists := s.clients[clientID]
	if !exists {
		return nil, errors.New("not found")
	}

	return client, nil
}

func (s *Storage) AuthorizeClientIDSecret(_ context.Context, _ string, _ string) error {
	return nil
}

func (s *Storage) SetUserinfoFromToken(ctx context.Context, userinfo oidc.UserInfoSetter, _, _, _ string) error {
	return s.SetUserinfoFromScopes(ctx, userinfo, "", "", []string{})
}

func (s *Storage) SetUserinfoFromScopes(ctx context.Context, userinfo oidc.UserInfoSetter, subject, _ string, _ []string) error {

	request, err := s.AuthRequestBySubject(ctx, subject)
	if err != nil {
		return err
	}

	u, err := s.getUserBySubject(ctx, subject)
	if err != nil {
		return err
	}

	userinfo.SetSubject(request.GetSubject())
	userinfo.SetEmail(u.Email(), true)
	userinfo.SetName(u.Name())
	userinfo.AppendClaims("lang", u.Lang())
	userinfo.AppendClaims("theme", u.Theme())

	return nil
}

func (s *Storage) GetPrivateClaimsFromScopes(_ context.Context, _, _ string, _ []string) (map[string]interface{}, error) {
	return map[string]interface{}{"private_claim": "test"}, nil
}

func (s *Storage) SetIntrospectionFromToken(ctx context.Context, introspect oidc.IntrospectionResponse, _, subject, clientID string) error {
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

func (s *Storage) ValidateJWTProfileScopes(_ context.Context, _ string, scope []string) ([]string, error) {
	return scope, nil
}

func (s *Storage) CompleteAuthRequest(ctx context.Context, requestId, sub string) error {
	request, err := s.AuthRequestByID(ctx, requestId)
	if err != nil {
		return err
	}
	req := request.(*auth.Request)
	req.Complete(sub)
	err = s.updateRequest(ctx, requestId, *req)
	return err
}

func (s *Storage) updateRequest(ctx context.Context, requestID string, req auth.Request) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if requestID == "" {
		return errors.New("invalid id")
	}
	reqId, err := id.AuthRequestIDFrom(requestID)
	if err != nil {
		return err
	}

	if _, err := s.requests.FindByID(ctx, reqId); err != nil {
		return err
	}

	if err := s.requests.Save(ctx, &req); err != nil {
		return err
	}

	return nil
}
