package tag

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestToTagGroup(t *testing.T) {
	tag := tag{
		id:      id.NewTagID(),
		label:   "xxx",
		sceneId: id.SceneID{},
	}
	group := ToTagGroup(&tag)
	assert.IsType(t, &Group{}, group)
}

func TestToTagItem(t *testing.T) {
	tag := tag{
		id:      id.NewTagID(),
		label:   "xxx",
		sceneId: id.SceneID{},
	}
	item := ToTagItem(&tag)
	assert.IsType(t, &Item{}, item)
}
