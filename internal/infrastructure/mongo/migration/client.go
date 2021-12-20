package migration

import (
	"context"
	"errors"
	"fmt"

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

type DBClient = *mongodoc.Client

type MigrationFunc = func(context.Context, DBClient) error

type Client struct {
	Client *mongodoc.Client
}

func (c Client) Migrate(ctx context.Context, cfgR repo.Config) error {
	cfg, err := cfgR.Load(ctx)
	if err != nil {
		return fmt.Errorf("Could not load auth config: %w\n", err)
	}
	defer func() {
		if err := cfgR.Release(ctx); err == nil {
			log.Errorf("Could not release config lock: %s\n", err)
		}
	}()

	nextMigrations := cfg.NextMigrations(migrationKeys())
	if len(nextMigrations) == 0 {
		return nil
	}

	for _, m := range nextMigrations {
		log.Infof("DB migration: %d\n", m)

		if err := migrations[m](ctx, c.Client); err != nil {
			var ie *rerror.ErrInternal
			if ok := errors.As(err, &ie); ok {
				err = ie.Unwrap()
			}
			return fmt.Errorf("Failed to exec migration %d: %w", m, err)
		}

		cfg.Migration = m
		if err := cfgR.Save(ctx, cfg); err != nil {
			return fmt.Errorf("Failed to save config: %w\n", err)
		}
	}

	return nil
}

func migrationKeys() []int64 {
	keys := make([]int64, 0, len(migrations))
	for k := range migrations {
		keys = append(keys, k)
	}
	return keys
}

func (c *Client) loadConfig(ctx context.Context) (*config.Config, error) {
	cfg := &config.Config{}

	if err := c.Client.Collection("config").FindOne(ctx, bson.D{}).Decode(cfg); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) || errors.Is(err, mongo.ErrNilDocument) {
			return cfg, nil
		}
		return nil, err
	}

	return cfg, nil
}

func (c *Client) saveConfig(ctx context.Context, cfg *config.Config) error {
	if cfg == nil {
		return nil
	}

	if _, err := c.Client.Collection("config").UpdateOne(ctx, bson.D{}, bson.M{
		"$set": cfg,
	}, &options.UpdateOptions{
		Upsert: &upsert,
	}); err != nil {
		return rerror.ErrInternalBy(err)
	}

	return nil
}
