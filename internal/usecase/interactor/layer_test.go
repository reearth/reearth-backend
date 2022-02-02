package interactor

import (
	"context"
	"testing"

	"github.com/reearth/reearth-backend/internal/infrastructure/memory"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
	"github.com/reearth/reearth-backend/pkg/scene"
	"github.com/stretchr/testify/assert"
)

func TestCreateInfobox(t *testing.T) {
	ctx := context.Background()
	db := memory.InitRepos(nil)

	s := scene.New().NewID().Team(id.NewTeamID()).Project(id.NewProjectID()).RootLayer(id.NewLayerID()).MustBuild()
	assert.NoError(t, db.Scene.Save(ctx, s))
	l := layer.NewItem().NewID().Scene(s.ID()).MustBuild()
	assert.NoError(t, db.Layer.Save(ctx, l))

	u := NewLayer(db)
	i, err := u.CreateInfobox(
		ctx,
		l.ID(), &usecase.Operator{
			WritableTeams: []id.TeamID{s.Team()},
		},
	)

	assert.NoError(t, err)
	assert.NotNil(t, i)

	l2, err := db.Layer.FindItemByID(ctx, l.ID(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, l2.Infobox())

	p, err := db.Property.FindByID(ctx, l2.Infobox().Property(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.NotNil(t, p.Schema())
}
