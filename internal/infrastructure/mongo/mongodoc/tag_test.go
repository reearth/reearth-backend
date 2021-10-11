package mongodoc

import (
	"reflect"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNewTag(t *testing.T) {
	sid := id.NewSceneID()
	dssid := id.NewDatasetSchemaID()
	dsid := id.NewDatasetID()
	dssfid := id.NewDatasetSchemaFieldID()
	ti, _ := tag.NewItem().
		NewID().
		Label("Item").
		Scene(sid).
		LinkedDatasetFieldID(dssfid.Ref()).
		LinkedDatasetID(dsid.Ref()).
		LinkedDatasetSchemaID(dssid.Ref()).
		Build()
	tg, _ := tag.NewGroup().
		NewID().
		Label("group").
		Tags(tag.NewListFromTags([]id.TagID{ti.ID()})).
		Scene(sid).
		Build()
	type args struct {
		t tag.Tag
	}
	tests := []struct {
		name  string
		args  args
		want  *TagDocument
		want1 string
	}{
		{
			name: "New tag group",
			args: args{
				t: tg,
			},
			want: &TagDocument{
				ID:    tg.ID().String(),
				Label: "group",
				Scene: sid.ID().String(),
				Item:  nil,
				Group: &TagGroupDocument{Tags: []string{ti.ID().String()}},
			},
			want1: tg.ID().String(),
		},
		{
			name: "New tag item",
			args: args{
				t: ti,
			},
			want: &TagDocument{
				ID:    ti.ID().String(),
				Label: "Item",
				Scene: sid.ID().String(),
				Item: &TagItemDocument{
					LinkedDatasetFieldID:  dssfid.RefString(),
					LinkedDatasetID:       dsid.RefString(),
					LinkedDatasetSchemaID: dssid.RefString(),
				},
				Group: nil,
			},
			want1: ti.ID().String(),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			got, got1 := NewTag(tc.args.t)
			assert.Equal(tt, tc.want1, got1)
			assert.Equal(tt, tc.want, got)
		})
	}
}

func TestNewTags(t *testing.T) {
	sid := id.NewSceneID()
	ti, _ := tag.NewItem().
		NewID().
		Label("Item").
		Scene(sid).
		Build()
	tg, _ := tag.NewGroup().
		NewID().
		Label("group").
		Tags(tag.NewListFromTags([]id.TagID{ti.ID()})).
		Scene(sid).
		Build()
	tgi := tag.Tag(tg)
	type args struct {
		tags []*tag.Tag
	}
	tests := []struct {
		name  string
		args  args
		want  []interface{}
		want1 []string
	}{
		{
			name: "new tags",
			args: args{
				tags: []*tag.Tag{
					&tgi,
				},
			},
			want: []interface{}{
				&TagDocument{
					ID:    tg.ID().String(),
					Label: "group",
					Scene: sid.ID().String(),
					Item:  nil,
					Group: &TagGroupDocument{Tags: []string{ti.ID().String()}},
				},
			},
			want1: []string{tgi.ID().String()},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			got, got1 := NewTags(tc.args.tags)
			assert.Equal(tt, tc.want, got)
			assert.Equal(tt, tc.want1, got1)
		})
	}
}

func TestFuncConsumer_Consume(t *testing.T) {
	sid := id.NewSceneID()
	tg, _ := tag.NewGroup().
		NewID().
		Label("group").
		Scene(sid).
		Build()
	ti, _ := tag.NewItem().
		NewID().
		Label("group").
		Scene(sid).
		Build()
	doc, _ := NewTag(tg)
	doc1, _ := NewTag(ti)
	r, _ := bson.Marshal(doc)
	r1, _ := bson.Marshal(doc1)
	type fields struct {
		Rows      []*tag.Tag
		GroupRows []*tag.Group
		ItemRows  []*tag.Item
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
			name: "nil row",
			fields: fields{
				Rows:      nil,
				GroupRows: nil,
				ItemRows:  nil,
			},
			args: args{
				raw: nil,
			},
			wantErr: false,
		},
		{
			name: "consume tag group",
			fields: fields{
				Rows:      nil,
				GroupRows: nil,
				ItemRows:  nil,
			},
			args: args{
				raw: r,
			},
			wantErr: false,
		},
		{
			name: "consume tag item",
			fields: fields{
				Rows:      nil,
				GroupRows: nil,
				ItemRows:  nil,
			},
			args: args{
				raw: r1,
			},
			wantErr: false,
		},
		{
			name: "fail: unmarshal error",
			fields: fields{
				Rows:      nil,
				GroupRows: nil,
				ItemRows:  nil,
			},
			args: args{
				raw: []byte{},
			},
			wantErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			c := &TagConsumer{
				Rows:      tc.fields.Rows,
				GroupRows: tc.fields.GroupRows,
				ItemRows:  tc.fields.ItemRows,
			}
			if err := c.Consume(tc.args.raw); (err != nil) != tc.wantErr {
				t.Errorf("Consume() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestTagDocument_Model(t *testing.T) {
	sid := id.NewSceneID()
	dssid := id.NewDatasetSchemaID()
	dsid := id.NewDatasetID()
	dssfid := id.NewDatasetSchemaFieldID()
	ti, _ := tag.NewItem().
		NewID().
		Label("Item").
		Scene(sid).
		LinkedDatasetFieldID(dssfid.Ref()).
		LinkedDatasetID(dsid.Ref()).
		LinkedDatasetSchemaID(dssid.Ref()).
		Build()
	tg, _ := tag.NewGroup().
		NewID().
		Label("group").
		Tags(tag.NewListFromTags([]id.TagID{ti.ID()})).
		Scene(sid).
		Build()
	type fields struct {
		ID    string
		Label string
		Scene string
		Item  *TagItemDocument
		Group *TagGroupDocument
	}
	tests := []struct {
		name    string
		fields  fields
		want    *tag.Item
		want1   *tag.Group
		wantErr bool
	}{
		{
			name: "item model",
			fields: fields{
				ID:    ti.ID().String(),
				Label: "Item",
				Scene: sid.ID().String(),
				Item: &TagItemDocument{
					LinkedDatasetFieldID:  dssfid.RefString(),
					LinkedDatasetID:       dsid.RefString(),
					LinkedDatasetSchemaID: dssid.RefString(),
				},
				Group: nil,
			},
			want:    ti,
			want1:   nil,
			wantErr: false,
		},
		{
			name: "group model",
			fields: fields{
				ID:    tg.ID().String(),
				Label: "group",
				Scene: sid.ID().String(),
				Item:  nil,
				Group: &TagGroupDocument{Tags: []string{ti.ID().String()}},
			},
			want:    nil,
			want1:   tg,
			wantErr: false,
		},
		{
			name:    "fail: invalid tag",
			fields:  fields{},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			d := &TagDocument{
				ID:    tc.fields.ID,
				Label: tc.fields.Label,
				Scene: tc.fields.Scene,
				Item:  tc.fields.Item,
				Group: tc.fields.Group,
			}
			got, got1, err := d.Model()
			if (err != nil) != tc.wantErr {
				t.Errorf("Model() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Model() got = %v, want %v", got, tc.want)
			}
			if !reflect.DeepEqual(got1, tc.want1) {
				t.Errorf("Model() got1 = %v, want %v", got1, tc.want1)
			}
		})
	}
}

func TestTagDocument_ModelGroup(t *testing.T) {
	sid := id.NewSceneID()
	ti, _ := tag.NewItem().
		NewID().
		Label("Item").
		Scene(sid).
		Build()
	tg, _ := tag.NewGroup().
		NewID().
		Label("group").
		Tags(tag.NewListFromTags([]id.TagID{ti.ID()})).
		Scene(sid).
		Build()
	type fields struct {
		ID    string
		Label string
		Scene string
		Item  *TagItemDocument
		Group *TagGroupDocument
	}
	tests := []struct {
		name    string
		fields  fields
		want    *tag.Group
		wantErr bool
	}{
		{
			name: "invalid id",
			fields: fields{
				ID: "xxx",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid id",
			fields: fields{
				ID:    id.NewTagID().String(),
				Scene: "xxx",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid item id",
			fields: fields{
				ID:    id.NewTagID().String(),
				Scene: id.NewSceneID().String(),
				Item:  nil,
				Group: &TagGroupDocument{Tags: []string{"xxx"}},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "pass",
			fields: fields{
				ID:    tg.ID().String(),
				Label: "group",
				Scene: sid.ID().String(),
				Item:  nil,
				Group: &TagGroupDocument{Tags: []string{ti.ID().String()}},
			},
			want:    tg,
			wantErr: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			d := &TagDocument{
				ID:    tc.fields.ID,
				Label: tc.fields.Label,
				Scene: tc.fields.Scene,
				Item:  tc.fields.Item,
				Group: tc.fields.Group,
			}
			got, err := d.ModelGroup()
			if (err != nil) != tc.wantErr {
				t.Errorf("ModelGroup() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ModelGroup() got = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTagDocument_ModelItem(t *testing.T) {
	sid := id.NewSceneID()
	dssid := id.NewDatasetSchemaID()
	dsid := id.NewDatasetID()
	dssfid := id.NewDatasetSchemaFieldID()
	ti, _ := tag.NewItem().
		NewID().
		Label("Item").
		Scene(sid).
		LinkedDatasetFieldID(dssfid.Ref()).
		LinkedDatasetID(dsid.Ref()).
		LinkedDatasetSchemaID(dssid.Ref()).
		Build()
	type fields struct {
		ID    string
		Label string
		Scene string
		Item  *TagItemDocument
		Group *TagGroupDocument
	}
	tests := []struct {
		name    string
		fields  fields
		want    *tag.Item
		wantErr bool
	}{
		{
			name: "invalid id",
			fields: fields{
				ID: "xxx",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid id",
			fields: fields{
				ID:    id.NewTagID().String(),
				Scene: "xxx",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "pass",
			fields: fields{
				ID:    ti.ID().String(),
				Label: ti.Label(),
				Scene: ti.Scene().String(),
				Item: &TagItemDocument{
					LinkedDatasetFieldID:  dssfid.RefString(),
					LinkedDatasetID:       dsid.RefString(),
					LinkedDatasetSchemaID: dssid.RefString(),
				},
				Group: nil,
			},
			want:    ti,
			wantErr: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			d := &TagDocument{
				ID:    tc.fields.ID,
				Label: tc.fields.Label,
				Scene: tc.fields.Scene,
				Item:  tc.fields.Item,
				Group: tc.fields.Group,
			}
			got, err := d.ModelItem()
			if (err != nil) != tc.wantErr {
				t.Errorf("ModelItem() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ModelItem() got = %v, want %v", got, tc.want)
			}
		})
	}
}
