package app

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/reearth/reearth-backend/pkg/log"
)

const configPrefix = "reearth"

type Config struct {
	Port         string `default:"8080" envconfig:"PORT"`
	Dev          bool
	DB           string `default:"mongodb://localhost"`
	Auth0        Auth0Config
	Auth         AuthConfig
	GraphQL      GraphQLConfig
	Published    PublishedConfig
	GCPProject   string `envconfig:"GOOGLE_CLOUD_PROJECT"`
	Profiler     string
	Tracer       string
	TracerSample float64
	GCS          GCSConfig
	AssetBaseURL string `default:"http://localhost:8080/assets"`
	Origins      []string
	Web          WebConfig
	SignupSecret string
}

type Auth0Config struct {
	Domain       string
	Audience     string
	ClientID     string
	ClientSecret string
	WebClientID  string
}

type AuthConfig struct {
	Domain string `default:"http://localhost:8080"`
	Key    string
	Pkix   AuthPkixConfig
}

type AuthPkixConfig struct {
	Organization  string `default:"Company, INC."`
	Country       string `default:"US"`
	Province      string `default:""`
	Locality      string `default:"San Francisco"`
	StreetAddress string `default:"Golden Gate Bridge"`
	PostalCode    string `default:"94016"`
}

type GraphQLConfig struct {
	ComplexityLimit int `default:"6000"`
}

type PublishedConfig struct {
	IndexURL *url.URL
}

type GCSConfig struct {
	BucketName              string
	PublicationCacheControl string
}

func ReadConfig(debug bool) (*Config, error) {
	// load .env
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		return nil, err
	} else if err == nil {
		log.Infof("config: .env loaded")
	}

	var c Config
	err := envconfig.Process(configPrefix, &c)

	if debug {
		c.Dev = true
	}

	return &c, err
}

func (c Config) Print() string {
	s := fmt.Sprintf("%+v", c)
	for _, secret := range []string{c.DB, c.Auth0.ClientSecret} {
		if secret == "" {
			continue
		}
		s = strings.ReplaceAll(s, secret, "***")
	}
	return s
}
