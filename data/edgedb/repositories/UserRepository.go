package repositories

import (
	"context"
	"fmt"
	"github.com/edgedb/edgedb-go"
	"github.com/hsedjame/meetups-app/data/edgedb/entities"
	"github.com/hsedjame/meetups-app/data/graph"
	"github.com/hsedjame/meetups-app/models"
)

const getByIdQuery = "SELECT Meetup::User {%s} FILTER .id=<uuid>$0;"
const insertUserQuery = "INSERT Meetup::User { " +
	"username := \"%s\", " +
	"email := \"%s\"};"

type UserRepository struct {
	Pool *edgedb.Pool
}

func (repo UserRepository) GetUserById(ctx context.Context, id string) (*entities.UserEntity, error) {
	if uuid, err := edgedb.ParseUUID(id); err != nil {
		return nil, err
	} else {
		predicateFunc := graph.PredicateFor(entities.UserEntity{})

		replaceFunc := func(name string) *graph.Replacement {
			return nil
		}

		query := fmt.Sprintf(getByIdQuery, graph.RequestedFields(ctx, predicateFunc, replaceFunc))

		fmt.Printf("Query : %s\n", query)

		user := entities.UserEntity{}

		err := repo.Pool.QueryOne(ctx, query, &user, uuid)

		return &user, err

	}
}

func (repo UserRepository) CreateUser(ctx context.Context, user models.User) error {
	query := fmt.Sprintf(insertUserQuery, user.Username, user.Email)

	fmt.Printf("Query : %s\n", query)

	var userEntity entities.UserEntity

	return repo.Pool.QueryOne(ctx, query, &userEntity)
}
