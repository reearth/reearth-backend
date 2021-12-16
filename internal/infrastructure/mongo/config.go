package mongo

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/reearth/reearth-backend/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/config"
	"github.com/reearth/reearth-backend/pkg/log"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var upsert = true

type configRepo struct {
	client *mongodoc.ClientCollection
	lockID *uuid.UUID
	lock   sync.Mutex
}

type ConfigDoc struct {
	Migration int64
	Auth      *config.Auth
	lock      *uuid.UUID
}

var (
	ErrLoadingLockedConfig = errors.New("loading locked config")
	ErrConfigNotLocked     = errors.New("trying to save not locked config")
)

func NewConfig(client *mongodoc.Client) repo.Config {
	return &configRepo{client: client.WithCollection("config")}
}

func (r *configRepo) Load(ctx context.Context) (*config.Config, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	numOfTries := 2

	for i := 1; i <= numOfTries; i++ {
		cfg, err := r.loadFromDB(ctx)
		if err == nil {
			return cfg, nil
		}
		if !errors.Is(err, ErrLoadingLockedConfig) {
			return nil, err
		}
		time.Sleep(time.Duration(rand.Intn(1000)+1000) * time.Millisecond)
	}

	log.Errorf("failed to load config after %d tries.\n", numOfTries)
	return nil, ErrLoadingLockedConfig
}

func (r *configRepo) loadFromDB(ctx context.Context) (*config.Config, error) {
	cfg := &ConfigDoc{}
	if r.lockID == nil {
		lockID := uuid.New()
		r.lockID = &lockID
	}

	if err := r.client.Collection().FindOneAndUpdate(ctx,
		bson.M{"$or": bson.A{
			bson.M{"lock": bson.M{"$exists": false}},
			bson.M{"lock": bson.M{"$eq": r.lockID}},
		}},
		bson.M{"$set": bson.M{"lock": r.lockID}},
	).Decode(cfg); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &config.Config{}, nil
		}
		return nil, rerror.ErrInternalBy(err)
	}
	if cfg.lock != nil && cfg.lock != r.lockID {
		return nil, ErrLoadingLockedConfig
	}
	return &config.Config{
		Migration: cfg.Migration,
		Auth:      cfg.Auth,
	}, nil
}

func (r *configRepo) Release(ctx context.Context) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.lockID == nil {
		return ErrConfigNotLocked
	}
	if _, err := r.client.Collection().UpdateOne(ctx,
		bson.M{"lock": bson.M{"$eq": r.lockID}},
		bson.M{"$unset": bson.M{"lock": nil}},
	); err != nil {
		return err
	}
	r.lockID = nil
	return nil
}

func (r *configRepo) Save(ctx context.Context, cfg *config.Config) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if cfg == nil {
		return nil
	}
	if r.lockID == nil {
		return ErrConfigNotLocked
	}
	if _, err := r.client.Collection().UpdateOne(ctx,
		bson.M{"lock": bson.M{"$eq": r.lockID}},
		bson.M{"$set": cfg}, &options.UpdateOptions{
			Upsert: &upsert,
		}); err != nil {
		return rerror.ErrInternalBy(err)
	}
	return nil
}
