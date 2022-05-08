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

	type a struct {
		id          asset.ID
		createdAt   time.Time
		team        asset.TeamID
		name        string
		size        int64
		url         string
		contentType string
	}

	tests := []struct {
		Name     string
		Expected struct {
			Name  string
			Asset *a
		}
	}{
		{
			Expected: struct {
				Name  string
				Asset *a
			}{
				Asset: &a{
					id:          id.NewAssetID(),
					createdAt:   time.Now(),
					team:        id.NewTeamID(),
					name:        "name",
					size:        10,
					url:         "hxxps://xxx.com",
					contentType: "json",
				},
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

			want, err := asset.New().
				NewID().
				Team(tc.Expected.Asset.team).
				Name(tc.Expected.Asset.name).
				Size(tc.Expected.Asset.size).
				URL(tc.Expected.Asset.url).
				CreatedAt(tc.Expected.Asset.createdAt).
				Build()

			assert.NoError(t, err)

			database, _ := uuid.New()
			client := mongodoc.NewClient(string(database[:]), c)
			repo := NewAsset(client)

			ctx := context.Background()
			err = repo.Save(ctx, want)
			assert.NoError(t, err)

			defer func() {
				_ = c.Database(string(database[:])).Drop(ctx)
			}()

			got, err := repo.FindByID(ctx, want.ID())
			assert.NoError(t, err)
			assert.Equal(t, want.ID(), got.ID())
			assert.Equal(t, want.CreatedAt(), got.CreatedAt())
			assert.Equal(t, want.Team(), got.Team())
			assert.Equal(t, want.URL(), got.URL())
			assert.Equal(t, want.Size(), got.Size())
			assert.Equal(t, want.Name(), got.Name())
			assert.Equal(t, want.ContentType(), got.ContentType())
		})
	}
}
