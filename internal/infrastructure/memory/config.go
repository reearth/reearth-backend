package memory

import (
	"context"
	"sync"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/config"
)

type Config struct {
	lock sync.Mutex
	data *config.Config
}

func NewConfig() repo.Config {
	return &Config{}
}

func (r *Config) Load(_ context.Context) (*config.Config, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	return r.data, nil
}

func (r *Config) Release(_ context.Context) error {
	return nil
}

func (r *Config) Save(_ context.Context, c *config.Config) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.data = c
	return nil
}
