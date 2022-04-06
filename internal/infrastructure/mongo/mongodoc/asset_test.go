package mongodoc

import (
	"testing"
	"time"

	"github.com/reearth/reearth-backend/pkg/asset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNewAsset(t *testing.T) {
	type args struct {
		asset *asset.Asset
	}

	assetID := asset.NewID()
	teamID := id.NewTeamID()
	now := time.Now()

	defer asset.MockID(assetID)()

	tests := []struct {
		name  string
		args  args
		want  *AssetDocument
		want1 string
	}{
		{
			name: "new asset",
			args: args{
				asset: asset.New().
					ID(assetID).
					CreatedAt(now).
					Team(teamID).
					Name("test").
					Size(10).
					URL("test_url").
					ContentType("application/json").
					MustBuild(),
			},
			want: &AssetDocument{
				ID:          assetID.String(),
				CreatedAt:   now,
				Team:        teamID.String(),
				Name:        "test",
				Size:        10,
				URL:         "test_url",
				ContentType: "application/json",
			},
			want1: assetID.String(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1 := NewAsset(tt.args.asset)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestAssetDocument_Model(t *testing.T) {
	now := time.Now()
	assetID := id.NewAssetID()
	teamID := id.NewTeamID()
	tests := []struct {
		name    string
		target  *AssetDocument
		want    *asset.Asset
		wantErr bool
	}{
		{
			name: "asset model",
			target: &AssetDocument{
				ID:          assetID.String(),
				CreatedAt:   now,
				Team:        teamID.String(),
				Name:        "name",
				Size:        10,
				URL:         "test",
				ContentType: "content type",
			},
			want: asset.New().
				ID(assetID).
				CreatedAt(now).
				Team(teamID).
				Name("name").
				Size(10).
				URL("test").
				ContentType("content type").
				MustBuild(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := &AssetDocument{
				ID:          tt.target.ID,
				CreatedAt:   tt.target.CreatedAt,
				Team:        tt.target.Team,
				Name:        tt.target.Name,
				Size:        tt.target.Size,
				URL:         tt.target.URL,
				ContentType: tt.target.ContentType,
			}
			got, err := d.Model()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAssetConsumer_Consume(t *testing.T) {

	type args struct {
		raw bson.Raw
	}

	tests := []struct {
		name    string
		target  *AssetConsumer
		args    args
		want    *AssetConsumer
		wantErr bool
	}{
		{
			name:    "asset consume",
			target:  &AssetConsumer{},
			args:    args{bson.Raw{11}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &AssetConsumer{
				Rows: tt.target.Rows,
			}
			err := c.Consume(tt.args.raw)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
