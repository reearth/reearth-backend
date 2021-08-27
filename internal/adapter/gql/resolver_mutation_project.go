package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/visualizer"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input CreateProjectInput) (*ProjectPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Project.Create(ctx, interfaces.CreateProjectParam{
		TeamID:      id.TeamID(input.TeamID),
		Visualizer:  visualizer.Visualizer(input.Visualizer),
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Alias:       input.Alias,
		Archived:    input.Archived,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &ProjectPayload{Project: toProject(res)}, nil
}

func (r *mutationResolver) UpdateProject(ctx context.Context, input UpdateProjectInput) (*ProjectPayload, error) {
	exit := trace(ctx)
	defer exit()

	deletePublicImage := false
	if input.DeletePublicImage != nil {
		deletePublicImage = *input.DeletePublicImage
	}

	deleteImageURL := false
	if input.DeleteImageURL != nil {
		deleteImageURL = *input.DeleteImageURL
	}

	res, err := r.usecases.Project.Update(ctx, interfaces.UpdateProjectParam{
		ID:                id.ProjectID(input.ProjectID),
		Name:              input.Name,
		Description:       input.Description,
		Alias:             input.Alias,
		ImageURL:          input.ImageURL,
		Archived:          input.Archived,
		IsBasicAuthActive: input.IsBasicAuthActive,
		BasicAuthUsername: input.BasicAuthUsername,
		BasicAuthPassword: input.BasicAuthPassword,
		PublicTitle:       input.PublicTitle,
		PublicDescription: input.PublicDescription,
		PublicImage:       input.PublicImage,
		PublicNoIndex:     input.PublicNoIndex,
		DeletePublicImage: deletePublicImage,
		DeleteImageURL:    deleteImageURL,
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &ProjectPayload{Project: toProject(res)}, nil
}

func (r *mutationResolver) PublishProject(ctx context.Context, input PublishProjectInput) (*ProjectPayload, error) {
	exit := trace(ctx)
	defer exit()

	res, err := r.usecases.Project.Publish(ctx, interfaces.PublishProjectParam{
		ID:     id.ProjectID(input.ProjectID),
		Alias:  input.Alias,
		Status: fromPublishmentStatus(input.Status),
	}, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &ProjectPayload{Project: toProject(res)}, nil
}

func (r *mutationResolver) DeleteProject(ctx context.Context, input DeleteProjectInput) (*DeleteProjectPayload, error) {
	exit := trace(ctx)
	defer exit()

	err := r.usecases.Project.Delete(ctx, id.ProjectID(input.ProjectID), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &DeleteProjectPayload{ProjectID: input.ProjectID}, nil
}
