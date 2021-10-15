package entities

import "github.com/edgedb/edgedb-go"

type MeetupEntity struct {
	ID edgedb.UUID `edgedb:"id"`
	Name string `edgedb:"name"`
	Description string `edgedb:"description"`
	UserId string `edgedb:"userId"`
}
