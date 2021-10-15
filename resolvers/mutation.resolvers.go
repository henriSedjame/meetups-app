package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/hsedjame/meetups-app/graph/generated"
	"github.com/hsedjame/meetups-app/models"
)

func (r *mutationResolver) CreateMeetup(ctx context.Context, input *models.NewMeetUp) (*models.Meetup, error) {
	m, err := r.MeetupRepo.CreateMeetup(ctx, input, "68838a24-eeed-11eb-999a-4fbe78cdc7e4")
	return &models.Meetup{
		ID: m.ID.String(),
		Name: input.Name,
		Description: input.Description,
		UserId: m.UserId,
	}, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
