package edge

import (
	"context"
	"github.com/edgedb/edgedb-go"
)

func New(ctx context.Context, database string, user string, password string, port int) (*edgedb.Pool, error) {
	return edgedb.Connect(ctx, edgedb.Options{
		Database: database,
		Ports: []int{port},
		User: user,
		Password: password,
		MinConns: 1,
		MaxConns: 4,
	})
}
