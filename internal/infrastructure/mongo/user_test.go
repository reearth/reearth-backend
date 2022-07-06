package mongo

import (
	"context"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/reearth/reearth-backend/pkg/user"
	"github.com/stretchr/testify/assert"
)

type TestCase[S any, A any, W any] struct {
	name    string
	seeds   S
	arg     A
	want    W
	wantErr error
}

func Test_userRepo_FindByIDs(t *testing.T) {
	tid1 := id.NewTeamID()
	id1 := id.NewUserID()
	id2 := id.NewUserID()
	u1 := user.New().ID(id1).Team(tid1).Email("test1@test.com").MustBuild()
	u2 := user.New().ID(id2).Team(tid1).Email("test2@test.com").MustBuild()

	tests := []TestCase[[]*user.User, []id.UserID, []*user.User]{
		{
			name:    "0 count in empty db",
			seeds:   []*user.User{},
			arg:     []id.UserID{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with user for another teams",
			seeds: []*user.User{
				user.New().NewID().Team(id.NewTeamID()).Email("test1@test.com").MustBuild(),
			},
			arg:     []id.UserID{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single user",
			seeds: []*user.User{
				u1,
			},
			arg:     []id.UserID{id1},
			want:    []*user.User{u1},
			wantErr: nil,
		},
		{
			name: "1 count with multi users",
			seeds: []*user.User{
				u1,
				user.New().NewID().Team(id.NewTeamID()).Email("test3@test.com").MustBuild(),
				user.New().NewID().Team(id.NewTeamID()).Email("test4@test.com").MustBuild(),
			},
			arg:     []id.UserID{id1},
			want:    []*user.User{u1},
			wantErr: nil,
		},
		{
			name: "2 count with multi users",
			seeds: []*user.User{
				u1,
				u2,
				user.New().NewID().Team(id.NewTeamID()).Email("test3@test.com").MustBuild(),
				user.New().NewID().Team(id.NewTeamID()).Email("test4@test.com").MustBuild(),
			},
			arg:     []id.UserID{id1, id2},
			want:    []*user.User{u1, u2},
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewUser(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.Nil(t, err)
			}

			got, err := r.FindByIDs(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_userRepo_FindByID(t *testing.T) {
	tid1 := id.NewTeamID()
	id1 := id.NewUserID()
	u1 := user.New().ID(id1).Team(tid1).Email("test1@test.com").MustBuild()

	testCases := []TestCase[[]*user.User, id.UserID, *user.User]{
		{
			name:    "User not found in an empty database",
			seeds:   []*user.User{},
			arg:     id.NewUserID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "User not found",
			seeds: []*user.User{
				user.New().NewID().Email("test2@test.com").MustBuild(),
			},
			arg:     id.NewUserID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Found 1 user",
			seeds: []*user.User{
				u1,
			},
			arg:     id1,
			want:    u1,
			wantErr: nil,
		},
		{
			name: "Found 1 user in many users",
			seeds: []*user.User{
				u1,
				user.New().NewID().Email("test2@test.com").Team(id.NewTeamID()).MustBuild(),
				user.New().NewID().Email("test3@test.com").Team(id.NewTeamID()).MustBuild(),
			},
			arg:     id1,
			want:    u1,
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewUser(client)
			ctx := context.Background()
			for _, s := range tc.seeds {
				err := r.Save(ctx, s)
				assert.Nil(t, err)
			}

			got, err := r.FindByID(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
