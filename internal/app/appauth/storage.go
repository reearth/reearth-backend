package appauth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"math/big"
	mrand "math/rand"
	"sync"
	"time"

	"github.com/caos/oidc/pkg/oidc"
	"github.com/caos/oidc/pkg/op"
	"github.com/oklog/ulid"
	"github.com/reearth/reearth-backend/pkg/user"
	"gopkg.in/square/go-jose.v2"
)

type Storage struct {
	lock             sync.Mutex
	appConfig        *StorageConfig
	getUserBySubject func(context.Context, string) (*user.User, error)
	clients          map[string]op.Client
	requests         map[string]AuthRequest
	keySet           jose.JSONWebKeySet
	key              *rsa.PrivateKey
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

func NewAuthStorage(cfg *StorageConfig, getUserBySubject func(context.Context, string) (*user.User, error)) op.Storage {

	client := initLocalClient(cfg.Debug)

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

	key, keySet := initKeys(name)

	return &Storage{
		appConfig:        cfg,
		getUserBySubject: getUserBySubject,
		requests:         make(map[string]AuthRequest),
		clients: map[string]op.Client{
			client.GetID(): client,
		},
		key:    key,
		keySet: keySet,
	}
}

func initKeys(name pkix.Name) (*rsa.PrivateKey, jose.JSONWebKeySet) {

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

	return key, jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{Key: key.Public(), Use: "sig", Algorithm: "RS256", KeyID: "1", Certificates: []*x509.Certificate{cert}},
		},
	}
}

func (s *Storage) Health(_ context.Context) error {
	return nil
}

func (s *Storage) CreateAuthRequest(_ context.Context, authReq *oidc.AuthRequest, _ string) (op.AuthRequest, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	ti := time.Now()
	entropy := ulid.Monotonic(mrand.New(mrand.NewSource(ti.UnixNano())), 0)

	audiences := []string{
		s.appConfig.Domain,
	}
	if s.appConfig.Debug {
		audiences = append(audiences, "http://localhost:8080")
	}

	request := &AuthRequest{
		ID:           ulid.MustNew(ulid.Timestamp(ti), entropy).String(),
		ClientID:     authReq.ClientID,
		subject:      "",
		code:         "", // Will be set after /authorize/callback success
		state:        authReq.State,
		scopes:       authReq.Scopes,
		audiences:    audiences,
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

	s.requests[request.ID] = *request
	return request, nil
}

func (s *Storage) AuthRequestByID(_ context.Context, requestID string) (op.AuthRequest, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if requestID == "" {
		return nil, errors.New("invalid id")
	}
	request, exists := s.requests[requestID]
	if !exists {
		return nil, errors.New("not found")
	}
	return &request, nil
}

func (s *Storage) AuthRequestByCode(_ context.Context, code string) (op.AuthRequest, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if code == "" {
		return nil, errors.New("invalid code")
	}
	for _, request := range s.requests {
		if request.GetCode() == code {
			return &request, nil
		}
	}
	return nil, errors.New("invalid code")
}

func (s *Storage) AuthRequestBySubject(_ context.Context, subject string) (op.AuthRequest, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if subject == "" {
		return nil, errors.New("invalid subject")
	}
	for _, request := range s.requests {
		if request.GetSubject() == subject {
			return &request, nil
		}
	}
	return nil, errors.New("invalid subject")
}

func (s *Storage) SaveAuthCode(ctx context.Context, requestID, code string) error {

	request, err := s.AuthRequestByID(ctx, requestID)
	if err != nil {
		return err
	}
	request2 := request.(*AuthRequest)
	request2.code = code
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
	authReq := request.(*AuthRequest)
	return "id", authReq.ID, time.Now().UTC().Add(5 * time.Minute), nil
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
	keyCh <- jose.SigningKey{Algorithm: jose.RS256, Key: s.key}
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
	req := request.(*AuthRequest)
	req.Complete(sub)
	err = s.updateRequest(ctx, requestId, *req)
	return err
}

func (s *Storage) updateRequest(_ context.Context, requestID string, req AuthRequest) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if requestID == "" {
		return errors.New("invalid id")
	}
	_, exists := s.requests[requestID]
	if !exists {
		return errors.New("not found")
	}

	s.requests[requestID] = req

	return nil
}
