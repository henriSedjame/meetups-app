package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/edgedb/edgedb-go"
	edge "github.com/hsedjame/meetups-app/data/edgedb"
	"github.com/hsedjame/meetups-app/data/edgedb/repositories"
	"github.com/hsedjame/meetups-app/graph/generated"
	"github.com/hsedjame/meetups-app/models"
	"github.com/hsedjame/meetups-app/resolvers"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {

	// get background context
	ctx := context.Background()

	// instantiate an edge database pool
	if pool, err := edge.New(ctx, "edgedb", "edgedb", "XqgNDAkULhmYE5CKcV3H4U3f", 10702); err != nil {
		log.Fatalf("Fail to connect edgedb database.")
	} else {
		// defer the pool close action
		defer func(pool *edgedb.Pool) {
			err := pool.Close()
			if err != nil {
				log.Fatalf("Fail to close edgedb pool.")
			}
		}(pool)

		log.Printf("Connection to edgedb success %#v", pool)

		// get port env variable
		port := os.Getenv("PORT")

		// if port is empty initialize it with the default value
		if port == "" {
			port = defaultPort
		}

		// instantiate repositories
		userRepo := repositories.UserRepository{Pool: pool}
		meetupRepo := repositories.MeetupRepository{Pool: pool}

		// init databases
		initDB(ctx, userRepo)

		// initialize GraphQl server
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
			MeetupRepo: meetupRepo,
			UserRepo:   userRepo,
		}}))

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}

func initDB(ctx context.Context, repository repositories.UserRepository) {
	_ = repository.CreateUser(ctx, models.User{Username: "John", Email: "john@gmail.com"})
	_ = repository.CreateUser(ctx, models.User{Username: "Franck", Email: "franck@gmail.com"})
}
