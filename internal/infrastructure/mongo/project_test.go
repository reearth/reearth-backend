package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/openlyinc/pointy"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/project"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/stretchr/testify/assert"
)

func Test_projectRepo_CountByTeam(t *testing.T) {
	tid1 := id.NewTeamID()
	tests := []struct {
		name    string
		seeds   []*project.Project
		arg     id.TeamID
		want    int
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   []*project.Project{},
			arg:     id.NewTeamID(),
			want:    0,
			wantErr: nil,
		},
		{
			name: "0 count with project for another teams",
			seeds: []*project.Project{
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			arg:     id.NewTeamID(),
			want:    0,
			wantErr: nil,
		},
		{
			name: "1 count with single project",
			seeds: []*project.Project{
				project.New().NewID().Team(tid1).MustBuild(),
			},
			arg:     tid1,
			want:    1,
			wantErr: nil,
		},
		{
			name: "1 count with multi projects",
			seeds: []*project.Project{
				project.New().NewID().Team(tid1).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			arg:     tid1,
			want:    1,
			wantErr: nil,
		},
		{
			name: "2 count with multi projects",
			seeds: []*project.Project{
				project.New().NewID().Team(tid1).MustBuild(),
				project.New().NewID().Team(tid1).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			arg:     tid1,
			want:    2,
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.Nil(t, err)
			}

			got, err := r.CountByTeam(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_projectRepo_Filtered(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond).UTC()
	tid1 := id.NewTeamID()
	id1 := id.NewProjectID()
	id2 := id.NewProjectID()
	p1 := project.New().ID(id1).Team(tid1).UpdatedAt(now).MustBuild()
	p2 := project.New().ID(id2).Team(tid1).UpdatedAt(now).MustBuild()

	tests := []struct {
		name    string
		seeds   []*project.Project
		arg     repo.TeamFilter
		wantErr error
	}{
		{
			name: "no r/w teams operation denied",
			seeds: []*project.Project{
				p1,
				p2,
			},
			arg: repo.TeamFilter{
				Readable: []id.TeamID{},
				Writable: []id.TeamID{},
			},
			wantErr: repo.ErrOperationDenied,
		},
		{
			name: "r/w teams operation success",
			seeds: []*project.Project{
				p1,
				p2,
			},
			arg: repo.TeamFilter{
				Readable: []id.TeamID{tid1},
				Writable: []id.TeamID{tid1},
			},
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewProject(client).Filtered(tc.arg)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.ErrorIs(t, err, tc.wantErr)
			}
		})
	}
}

func Test_projectRepo_FindByID(t *testing.T) {
	id1 := id.NewProjectID()
	now := time.Now().Truncate(time.Millisecond).UTC()
	p1 := project.New().ID(id1).Team(id.NewTeamID()).UpdatedAt(now).MustBuild()
	tests := []struct {
		name    string
		seeds   []*project.Project
		arg     id.ProjectID
		want    *project.Project
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   []*project.Project{},
			arg:     id.NewProjectID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Not found",
			seeds: []*project.Project{
				project.New().NewID().MustBuild(),
			},
			arg:     id.NewProjectID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Found 1",
			seeds: []*project.Project{
				p1,
			},
			arg:     id1,
			want:    p1,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: []*project.Project{
				p1,
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			arg:     id1,
			want:    p1,
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.Nil(t, err)
			}

			got, err := r.FindByID(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_projectRepo_FindByIDs(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond).UTC()
	tid1 := id.NewTeamID()
	id1 := id.NewProjectID()
	id2 := id.NewProjectID()
	p1 := project.New().ID(id1).Team(tid1).UpdatedAt(now).MustBuild()
	p2 := project.New().ID(id2).Team(tid1).UpdatedAt(now).MustBuild()

	tests := []struct {
		name    string
		seeds   []*project.Project
		arg     id.ProjectIDList
		want    []*project.Project
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   []*project.Project{},
			arg:     []id.ProjectID{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with project for another teams",
			seeds: []*project.Project{
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			arg:     []id.ProjectID{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single project",
			seeds: []*project.Project{
				p1,
			},
			arg:     []id.ProjectID{id1},
			want:    []*project.Project{p1},
			wantErr: nil,
		},
		{
			name: "1 count with multi projects",
			seeds: []*project.Project{
				p1,
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			arg:     []id.ProjectID{id1},
			want:    []*project.Project{p1},
			wantErr: nil,
		},
		{
			name: "2 count with multi projects",
			seeds: []*project.Project{
				p1,
				p2,
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			arg:     []id.ProjectID{id1, id2},
			want:    []*project.Project{p1, p2},
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewProject(client)
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

func Test_projectRepo_FindByPublicName(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond).UTC()
	id1 := id.NewProjectID()
	p1 := project.New().
		ID(id1).
		Team(id.NewTeamID()).
		Alias("xyz123").
		PublishmentStatus(project.PublishmentStatusPublic).
		UpdatedAt(now).
		MustBuild()

	id2 := id.NewProjectID()
	p2 := project.New().
		ID(id2).
		Team(id.NewTeamID()).
		Alias("xyz321").
		PublishmentStatus(project.PublishmentStatusLimited).
		UpdatedAt(now).
		MustBuild()

	tests := []struct {
		name    string
		seeds   []*project.Project
		arg     string
		want    *project.Project
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   []*project.Project{},
			arg:     "xyz123",
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Not found",
			seeds: []*project.Project{
				project.New().NewID().Alias("abc123").MustBuild(),
			},
			arg:     "xyz123",
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "private project not found",
			seeds: []*project.Project{
				project.New().
					NewID().
					Team(id.NewTeamID()).
					Alias("xyz123").
					PublishmentStatus(project.PublishmentStatusPrivate).
					MustBuild(),
			},
			arg:     "xyz123",
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "public Found",
			seeds: []*project.Project{
				p1,
			},
			arg:     "xyz123",
			want:    p1,
			wantErr: nil,
		},
		{
			name: "linited Found",
			seeds: []*project.Project{
				p2,
			},
			arg:     "xyz321",
			want:    p2,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: []*project.Project{
				p1,
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			arg:     "xyz123",
			want:    p1,
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.Nil(t, err)
			}

			got, err := r.FindByPublicName(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_projectRepo_FindByTeam(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond).UTC()
	tid1 := id.NewTeamID()
	p1 := project.New().NewID().Team(tid1).UpdatedAt(now).MustBuild()
	p2 := project.New().NewID().Team(tid1).UpdatedAt(now).MustBuild()

	type args struct {
		tid   id.TeamID
		pInfo *usecase.Pagination
	}
	tests := []struct {
		name    string
		seeds   []*project.Project
		args    args
		want    []*project.Project
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   []*project.Project{},
			args:    args{id.NewTeamID(), nil},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with project for another teams",
			seeds: []*project.Project{
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			args:    args{id.NewTeamID(), nil},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single project",
			seeds: []*project.Project{
				p1,
			},
			args:    args{tid1, usecase.NewPagination(pointy.Int(1), nil, nil, nil)},
			want:    []*project.Project{p1},
			wantErr: nil,
		},
		{
			name: "1 count with multi projects",
			seeds: []*project.Project{
				p1,
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			args:    args{tid1, usecase.NewPagination(pointy.Int(1), nil, nil, nil)},
			want:    []*project.Project{p1},
			wantErr: nil,
		},
		{
			name: "2 count with multi projects",
			seeds: []*project.Project{
				p1,
				p2,
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			args:    args{tid1, usecase.NewPagination(pointy.Int(2), nil, nil, nil)},
			want:    []*project.Project{p1, p2},
			wantErr: nil,
		},
		{
			name: "get 1st page of 2",
			seeds: []*project.Project{
				p1,
				p2,
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			args:    args{tid1, usecase.NewPagination(pointy.Int(1), nil, nil, nil)},
			want:    []*project.Project{p1},
			wantErr: nil,
		},
		{
			name: "get last page of 2",
			seeds: []*project.Project{
				p1,
				p2,
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			args:    args{tid1, usecase.NewPagination(nil, pointy.Int(1), nil, nil)},
			want:    []*project.Project{p2},
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.Nil(t, err)
			}

			got, _, err := r.FindByTeam(ctx, tc.args.tid, tc.args.pInfo)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_projectRepo_Remove(t *testing.T) {
	id1 := id.NewProjectID()
	p1 := project.New().ID(id1).Team(id.NewTeamID()).MustBuild()
	tests := []struct {
		name    string
		seeds   []*project.Project
		arg     id.ProjectID
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   []*project.Project{},
			arg:     id.NewProjectID(),
			wantErr: nil,
		},
		{
			name: "Not found",
			seeds: []*project.Project{
				project.New().NewID().MustBuild(),
			},
			arg:     id.NewProjectID(),
			wantErr: nil,
		},
		{
			name: "Found 1",
			seeds: []*project.Project{
				p1,
			},
			arg:     id1,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: []*project.Project{
				p1,
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
				project.New().NewID().Team(id.NewTeamID()).MustBuild(),
			},
			arg:     id1,
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.Nil(t, err)
			}

			err := r.Remove(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Nil(t, err)
			_, err = r.FindByID(ctx, tc.arg)
			assert.ErrorIs(t, err, rerror.ErrNotFound)
		})
	}
}

func Test_projectRepo_Save(t *testing.T) {
	id1 := id.NewProjectID()
	p1 := project.New().ID(id1).Team(id.NewTeamID()).MustBuild()
	tests := []struct {
		name    string
		seeds   []*project.Project
		arg     id.ProjectID
		want    *project.Project
		wantErr error
	}{
		{
			name: "Saved",
			seeds: []*project.Project{
				p1,
			},
			arg:     id1,
			want:    p1,
			wantErr: nil,
		},
	}

	initDB := connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := initDB(t)

			r := NewProject(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.Nil(t, err)
			}

			got, err := r.FindByID(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want.ID(), got.ID())
		})
	}
}
