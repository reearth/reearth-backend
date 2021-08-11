package tag

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestList_Add(t *testing.T) {
	tid := id.NewTagID()
	tl := NewList()
	tl.Add(tid)
	expected := []id.TagID{tid}
	assert.Equal(t, expected, tl.Tags())
}

func TestList_Remove(t *testing.T) {
	tid := id.NewTagID()
	tid2 := id.NewTagID()
	tags := []id.TagID{
		tid,
		tid2,
	}
	tl := NewListFromTags(tags)
	tl.Remove(tid2)
	expected := []id.TagID{tid}
	assert.Equal(t, expected, tl.Tags())
}

func TestList_Has(t *testing.T) {
	tid := id.NewTagID()
	tid2 := id.NewTagID()
	tags := []id.TagID{
		tid,
	}
	tl := NewListFromTags(tags)
	assert.True(t, tl.Has(tid))
	assert.False(t, tl.Has(tid2))
}
