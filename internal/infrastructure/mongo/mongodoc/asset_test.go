package mongodoc

import (
	"github.com/reearth/reearth-backend/pkg/asset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestNewAsset(t *testing.T) {
	type args struct {
		asset *asset.Asset
	}

	newAsset, _ := asset.New().
		NewID().
		Team(id.NewTeamID()).
		Name("test").
		Size(10).
		URL("test_url").
		Build()

	tests := []struct {
		name  string
		args  args
		want  *AssetDocument
		want1 string
	}{
		{
			name: "new asset",
			args: args{
				asset: newAsset,
			},
			want: &AssetDocument{
				ID:          newAsset.ID().String(),
				CreatedAt:   newAsset.CreatedAt(),
				Team:        newAsset.Team().String(),
				Name:        newAsset.Name(),
				Size:        newAsset.Size(),
				URL:         newAsset.URL(),
				ContentType: newAsset.ContentType(),
			},
			want1: newAsset.ID().String(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1 := NewAsset(tt.args.asset)
			assert.Equalf(t, tt.want, got, "NewAsset(%v)", tt.args.asset)
			assert.Equalf(t, tt.want1, got1, "NewAsset(%v)", tt.args.asset)
		})
	}
}

func TestAssetDocument_Model(t *testing.T) {
	type fields struct {
		ID          string
		CreatedAt   time.Time
		Team        string
		Name        string
		Size        int64
		URL         string
		ContentType string
	}

	newFields := fields{
		ID:          "id",
		CreatedAt:   time.Time{},
		Team:        "team",
		Name:        "name",
		Size:        10,
		URL:         "test",
		ContentType: "content type",
	}

	aid, _ := id.AssetIDFrom(newFields.ID)
	tid, _ := id.TeamIDFrom(newFields.Team)

	newAsset, _ := asset.New().
		ID(aid).
		CreatedAt(newFields.CreatedAt).
		Team(tid).
		Name(newFields.Name).
		Size(newFields.Size).
		URL(newFields.URL).
		ContentType(newFields.ContentType).
		Build()

	tests := []struct {
		name    string
		fields  fields
		want    *asset.Asset
		wantErr bool
	}{
		{
			name:    "asset model",
			fields:  newFields,
			want:    newAsset,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := &AssetDocument{
				ID:          tt.fields.ID,
				CreatedAt:   tt.fields.CreatedAt,
				Team:        tt.fields.Team,
				Name:        tt.fields.Name,
				Size:        tt.fields.Size,
				URL:         tt.fields.URL,
				ContentType: tt.fields.ContentType,
			}
			got, _ := d.Model()
			assert.Equalf(t, tt.want, got, "Model()")
		})
	}
}

func TestAssetConsumer_Consume(t *testing.T) {
	type fields struct {
		Rows []*asset.Asset
	}
	type args struct {
		raw bson.Raw
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "asset consume",
			fields:  fields{},
			args:    args{bson.Raw{11}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &AssetConsumer{
				Rows: tt.fields.Rows,
			}
			out := c.Consume(tt.args.raw)
			assert.Equal(t, out == nil, tt.wantErr)
		})
	}
}
