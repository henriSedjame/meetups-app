package models

type Meetup struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	UserId string `json:"user_id"`
}
