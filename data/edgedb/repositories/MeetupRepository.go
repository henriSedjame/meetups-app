package repositories

import (
	"context"
	"fmt"
	"github.com/edgedb/edgedb-go"
	"github.com/hsedjame/meetups-app/data/edgedb/entities"
	"github.com/hsedjame/meetups-app/data/graph"
	"github.com/hsedjame/meetups-app/models"
)

const SelectQuery = "SELECT Meetup::Meetup { %s };"
const insertQuery = "INSERT Meetup::Meetup { " +
	"name := \"%s\", " +
	"description := \"%s\", " +
	"userId := \"%s\"};"

type MeetupRepository struct {
	Pool *edgedb.Pool
}

func (repo MeetupRepository) GetMeetups(ctx context.Context) ([]entities.MeetupEntity, error) {

	predicateFunc := graph.PredicateFor(entities.MeetupEntity{})

	replaceFunc := func(name string) *graph.Replacement {
		if name == "user" {
			return &graph.Replacement{Name: "userId", HasChild: false}
		}
		return nil
	}

	query := fmt.Sprintf(SelectQuery, graph.RequestedFields(ctx, predicateFunc, replaceFunc))

	fmt.Printf("Query : %s\n", query)

	var meetups []entities.MeetupEntity
	err := repo.Pool.Query(ctx, query, &meetups)
	return meetups, err
}

func (repo MeetupRepository) CreateMeetup(ctx context.Context, meetup *models.NewMeetUp, user string) (*entities.MeetupEntity, error) {

	if uid, err := edgedb.ParseUUID(user); err != nil {
		return nil, err
	} else {

		query := fmt.Sprintf(insertQuery, meetup.Name, meetup.Description, uid)

		fmt.Printf("Query : %s\n", query)

		var meetup entities.MeetupEntity

		err :=repo.Pool.QueryOne(ctx, query, &meetup)

		return &meetup, err
	}
}
