package mongo

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/reearth/reearth-backend/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-backend/pkg/asset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func TestFindByID(t *testing.T) {
	// Skip unit testing if "REEARTH_DB" is not configured
	// See details: https://github.com/reearth/reearth/issues/273
	db := os.Getenv("REEARTH_DB")
	if db == "" {
		return
	}

	tests := []struct {
		Name     string
		Expected struct {
			Name  string
			Asset *asset.Asset
		}
	}{
		{
			Expected: struct {
				Name  string
				Asset *asset.Asset
			}{
				Asset: asset.New().
					NewID().
					CreatedAt(time.Now()).
					Team(id.NewTeamID()).
					Name("name").
					Size(10).
					URL("hxxps://https://reearth.io/").
					ContentType("json").
					MustBuild(),
			},
		},
	}

	c, _ := mongo.Connect(
		context.Background(),
		options.Client().
			ApplyURI(db).
			SetConnectTimeout(time.Second*10),
	)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			database, _ := uuid.New()
			client := mongodoc.NewClient(string(database[:]), c)
			repo := NewAsset(client)

			ctx := context.Background()
			err := repo.Save(ctx, tc.Expected.Asset)
			assert.NoError(t, err)

			defer func() {
				_ = c.Database(string(database[:])).Drop(ctx)
			}()

			got, err := repo.FindByID(ctx, tc.Expected.Asset.ID())
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected.Asset.ID(), got.ID())
			assert.Equal(t, tc.Expected.Asset.CreatedAt(), got.CreatedAt())
			assert.Equal(t, tc.Expected.Asset.Team(), got.Team())
			assert.Equal(t, tc.Expected.Asset.URL(), got.URL())
			assert.Equal(t, tc.Expected.Asset.Size(), got.Size())
			assert.Equal(t, tc.Expected.Asset.Name(), got.Name())
			assert.Equal(t, tc.Expected.Asset.ContentType(), got.ContentType())
		})
	}
}
