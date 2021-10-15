package data

import "github.com/hsedjame/meetups-app/models"

var Users = []*models.User {
	{
		ID: "1",
		Username: "Joe",
		Email: "joe@gmail.com",
	},
	{
		ID: "2",
		Username: "Mary",
		Email: "mary@gmail.com",
	},
}

var Meetups = []*models.Meetup {
	{
		ID: "1",
		Name: "Rust, new C ?",
		Description: "Rust language presentation",
		UserId: "1",
	},
	{
		ID: "2",
		Name: "Functional programming with F#",
		Description: "Talk about functional programming",
		UserId: "2",
	},
}
