package mongo

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/reearth/reearth-backend/internal/infrastructure/mongo/mongodoc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func connect(t *testing.T) func() (*mongodoc.Client, func()) {
	t.Helper()

	// Skip unit testing if "REEARTH_DB" is not configured
	// See details: https://github.com/reearth/reearth/issues/273
	db := os.Getenv("REEARTH_DB")
	if db == "" {
		t.SkipNow()
		return nil
	}

	c, _ := mongo.Connect(
		context.Background(),
		options.Client().
			ApplyURI(db).
			SetConnectTimeout(time.Second*10),
	)

	return func() (*mongodoc.Client, func()) {
		database, _ := uuid.New()
		databaseName := "reearth-test-" + string(database[:])
		client := mongodoc.NewClient(databaseName, c)

		return client, func() {
			_ = c.Database(databaseName).Drop(context.Background())
		}
	}
}
