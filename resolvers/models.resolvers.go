package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/hsedjame/meetups-app/data"
	"github.com/hsedjame/meetups-app/graph/generated"
	"github.com/hsedjame/meetups-app/models"
)

func (r *meetupResolver) User(ctx context.Context, meetup *models.Meetup) (*models.User, error) {

	if user, err := r.UserRepo.GetUserById(ctx, meetup.UserId); err != nil {
		return nil, err
	} else {
		return &models.User{
			ID: user.ID.String(),
			Email: user.Email,
			Username: user.Username,
		}, nil
	}
}

func (r *userResolver) Meetups(ctx context.Context, user *models.User) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	for _, meetup := range data.Meetups {
		if meetup.UserId == user.ID {
			meetups = append(meetups, meetup)
		}
	}

	return meetups, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
