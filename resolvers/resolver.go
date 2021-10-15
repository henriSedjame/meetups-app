package resolvers

import "github.com/hsedjame/meetups-app/data/edgedb/repositories"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	MeetupRepo repositories.MeetupRepository
	UserRepo repositories.UserRepository
}



