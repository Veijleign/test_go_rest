package repository

import (
	"context"
	"testREST/internal/entity"
	"testREST/internal/repository/pgdb"
	"testREST/pkg/postgres"
)

type User interface {
	CreateUser(ctx context.Context, username, password string) (entity.User, error)
	GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (entity.User, error)
	GeUserById(ctx context.Context, id int) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
}

type Repositories struct {
	User
	Account
	Product
	Reservation
	Operation
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User: pgdb.NewUserRepo(pg),
	}
}
