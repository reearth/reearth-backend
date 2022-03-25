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

	newAsset := asset.New().
		NewID().
		Team(id.NewTeamID()).
		Name("test").
		Size(10).
		URL("test_url").
		MustBuild()

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
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestAssetDocument_Model(t *testing.T) {
	aid, _ := id.AssetIDFrom("01f2r7kg1fvvffp0gmexgy5hxy")
	tid, _ := id.TeamIDFrom("01f2r7kg1fvvffp0gmexgy5hxy")

	newAsset := asset.New().
		ID(aid).
		CreatedAt(time.Time{}).
		Team(tid).
		Name("name").
		Size(10).
		URL("test").
		ContentType("content type").
		MustBuild()

	tests := []struct {
		name    string
		target  *AssetDocument
		want    *asset.Asset
		wantErr bool
	}{
		{
			name: "asset model",
			target: &AssetDocument{
				ID:          "01f2r7kg1fvvffp0gmexgy5hxy",
				CreatedAt:   time.Time{},
				Team:        "01f2r7kg1fvvffp0gmexgy5hxy",
				Name:        "name",
				Size:        10,
				URL:         "test",
				ContentType: "content type",
			},
			want:    newAsset,
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
