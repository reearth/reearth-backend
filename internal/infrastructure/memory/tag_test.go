package memory

import (
	"context"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"

	"github.com/reearth/reearth-backend/pkg/tag"
)

func TestTag(t *testing.T) {
	repo := NewTag()
	assert.NotNil(t, repo)
	ctx := context.Background()
	sid := id.NewSceneID()
	sid2 := id.NewSceneID()
	sl := []id.SceneID{sid}
	t1, _ := tag.NewItem().NewID().Scene(sid).Label("item").Build()
	tl := tag.NewListFromTags([]id.TagID{t1.ID()})
	t2, _ := tag.NewGroup().NewID().Scene(sid).Label("group").Tags(*tl).Build()
	t3, _ := tag.NewItem().NewID().Scene(sid2).Label("item2").Build()
	tti2 := tag.Tag(t3)
	ttg := tag.Tag(t2)
	tti := tag.Tag(t1)
	err := repo.Save(ctx, t1)
	assert.NoError(t, err)
	out, err := repo.FindByID(ctx, t1.ID(), sl)
	assert.NoError(t, err)
	assert.Equal(t, &tti, out)
	err = repo.SaveAll(ctx, []*tag.Tag{&ttg, &tti2})
	assert.NoError(t, err)
	out2, err := repo.FindByIDs(ctx, []id.TagID{t1.ID(), t2.ID()}, sl)
	assert.NoError(t, err)
	assert.Equal(t, []*tag.Tag{&tti, &ttg}, out2)
	out3, err := repo.FindByScene(ctx, sid2)
	assert.NoError(t, err)
	assert.Equal(t, []*tag.Tag{&tti2}, out3)
	out4, err := repo.FindGroupByID(ctx, t2.ID(), sl)
	assert.NoError(t, err)
	assert.Equal(t, t2, out4)
	out5, err := repo.FindItemByID(ctx, t1.ID(), sl)
	assert.NoError(t, err)
	assert.Equal(t, t1, out5)
	out6, err := repo.FindGroupByIDs(ctx, []id.TagID{t2.ID()}, sl)
	assert.NoError(t, err)
	assert.Equal(t, []*tag.Group{t2}, out6)
	out7, err := repo.FindItemByIDs(ctx, []id.TagID{t1.ID()}, sl)
	assert.NoError(t, err)
	assert.Equal(t, []*tag.Item{t1}, out7)
	_ = repo.Remove(ctx, t1.ID())
	out8, _ := repo.FindByID(ctx, t1.ID(), sl)
	assert.Nil(t, out8)
	_ = repo.RemoveAll(ctx, []id.TagID{t2.ID()})
	out9, _ := repo.FindByID(ctx, t2.ID(), sl)
	assert.Nil(t, out9)
	_ = repo.RemoveByScene(ctx, sid2)
	out10, _ := repo.FindByID(ctx, t3.ID(), []id.SceneID{sid2})
	assert.Nil(t, out10)

}
