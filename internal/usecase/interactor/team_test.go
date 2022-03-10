package interactor

import (
	"context"
	"testing"

	"github.com/reearth/reearth-backend/internal/infrastructure/memory"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateTeam(t *testing.T) {
	ctx := context.Background()
	db := memory.InitRepos(nil)
	user := user.New().NewID().Team(id.NewTeamID()).MustBuild()

	team, err := NewTeam(db).Create(ctx, "team name", user.ID())
	assert.Nil(t, err)
	assert.NotNil(t, team)

	resultTeams, _ := db.Team.FindByIDs(ctx, []id.TeamID{team.ID()})
	assert.NotNil(t, resultTeams)
	assert.NotEmpty(t, resultTeams)
	assert.Equal(t, resultTeams[0].ID(), team.ID())
	assert.Equal(t, resultTeams[0].Name(), "team name")
}
