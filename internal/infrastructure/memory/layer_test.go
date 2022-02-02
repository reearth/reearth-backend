package memory

import (
	"context"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
	"github.com/reearth/reearth-backend/pkg/tag"
	"github.com/stretchr/testify/assert"
)

func TestLayer_FindByTag(t *testing.T) {
	ctx := context.Background()
	sid := id.NewSceneID()
	sl := []id.SceneID{sid}
	t1, _ := tag.NewItem().NewID().Scene(sid).Label("item").Build()
	tl := layer.NewTagList([]layer.Tag{layer.NewTagGroup(t1.ID(), nil)})
	lg := layer.New().NewID().Tags(tl).Scene(sid).Group().MustBuild()

	repo := Layer{
		data: layer.Map{
			lg.ID(): lg.LayerRef(),
		},
	}

	out, err := repo.FindByTag(ctx, t1.ID(), sl)
	assert.NoError(t, err)
	assert.Equal(t, layer.List{lg.LayerRef()}, out)
}
