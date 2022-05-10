package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/reearth/reearth-backend/pkg/asset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {
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

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			client, dropDB := initDB()
			defer dropDB()

			repo := NewAsset(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.Expected.Asset)
			assert.NoError(t, err)

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

func TestRemove(t *testing.T) {
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

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			client, dropDB := initDB()
			defer dropDB()

			repo := NewAsset(client)
			ctx := context.Background()
			err := repo.Save(ctx, tc.Expected.Asset)
			assert.NoError(t, err)

			err = repo.Remove(ctx, tc.Expected.Asset.ID())
			assert.NoError(t, err)

			got, err := repo.FindByID(ctx, tc.Expected.Asset.ID())
			assert.Equal(t, err, rerror.ErrNotFound)
			assert.Nil(t, got)
		})
	}
}
