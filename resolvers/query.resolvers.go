package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/hsedjame/meetups-app/graph/generated"
	"github.com/hsedjame/meetups-app/models"
)

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {

	var result []*models.Meetup

	meetups, err := r.MeetupRepo.GetMeetups(ctx)

	if err == nil {
		for _, meetup := range meetups {
			result = append(result, &models.Meetup{
				ID: meetup.ID.String(),
				Name: meetup.Name,
				Description: meetup.Description,
				UserId: meetup.UserId,
			})
		}
	}
	return result, err
}


// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
