package pgdb

import (
	"testREST/pkg/postgres"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}
