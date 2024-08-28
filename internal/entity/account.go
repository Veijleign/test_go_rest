package entity

import (
	"time"
)

type Account struct {
	Id      int       `db:"id"`
	Balance int       `db:"balance"`
	Created time.Time `db:"created"`
}
