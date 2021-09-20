package layer

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/tag"
	"github.com/stretchr/testify/assert"
)

func TestGroupBuilder_Tags(t *testing.T) {
	tags := []id.TagID{id.NewTagID()}
	b := NewGroup().NewID().Tags(tag.NewListFromTags(tags)).MustBuild()
	assert.Equal(t, tags, b.Tags().Tags())
}
