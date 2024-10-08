package entity

import "time"

type Reservation struct {
	Id        int       `db:"id"`
	AccountId int       `db:"account_id"`
	ProductId int       `db:"product_id"`
	OrderId   int       `db:"order_id"`
	Amount    int       `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}
