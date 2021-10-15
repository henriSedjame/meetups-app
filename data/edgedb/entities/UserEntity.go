package entities

import "github.com/edgedb/edgedb-go"

type UserEntity struct {
	ID edgedb.UUID `edgedb:"id"`
	Username string `edgedb:"username"`
	Email string `edgedb:"email"`
}
