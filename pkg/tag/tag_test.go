package tag

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestTagBase_Label(t *testing.T) {
	type fields struct {
		id                  id.TagID
		label               string
		scene               id.SceneID
		linkedDatasetSchema *id.DatasetSchemaID
	}
	test := struct {
		name   string
		fields fields
		want   string
	}{
		name: "should return label",
		fields: fields{
			id:    id.NewTagID(),
			label: "label",
			scene: id.NewSceneID(),
		},
		want: "label",
	}

	t.Run(test.name, func(t *testing.T) {
		tr := &TagBase{
			id:                  test.fields.id,
			label:               test.fields.label,
			scene:               test.fields.scene,
			linkedDatasetSchema: test.fields.linkedDatasetSchema,
		}
		if got := tr.Label(); got != test.want {
			t.Errorf("TagBase.Label() = %v, want %v", got, test.want)
		}
	})
}

func TestTagBase_SetLabel(t *testing.T) {
	type fields struct {
		id                  id.TagID
		label               string
		scene               id.SceneID
		linkedDatasetSchema *id.DatasetSchemaID
	}
	type args struct {
		label string
	}
	test := struct {
		name   string
		fields fields
		args   args
	}{
		name: "should set label which passed from parameter",
		fields: fields{
			id:    id.NewTagID(),
			label: "label",
			scene: id.NewSceneID(),
		},
		args: args{label: "newLabel"},
	}
	t.Run(test.name, func(t *testing.T) {
		tr := &TagBase{
			id:                  test.fields.id,
			label:               test.fields.label,
			scene:               test.fields.scene,
			linkedDatasetSchema: test.fields.linkedDatasetSchema,
		}
		tr.SetLabel(test.args.label)
		assert.Equal(t, test.args.label, tr.label)
	})
}
