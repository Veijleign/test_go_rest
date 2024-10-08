package entity

import "time"

type Operation struct {
	Id            int       `db:"id"`
	AccountId     int       `db:"account_id"`
	Amount        int       `db:"amount"`
	OperationType string    `db:"operation_type"`
	CreatedAt     time.Time `db:"created_at"`

	ProductId   *int   `db:"product_id"`
	OrderId     *int   `db:"order_id"`
	Description string `db:"description"`
}

const (
	OperationTypeDeposit      = "deposit"
	OperationTypeWithdraw     = "withdraw"
	OperationTypeTransferFrom = "transfer_from"
	OperationTypeTransferTo   = "transfer_to"
	OperationTypeReservation  = "reservation"
	OperationTypeRevenue      = "revenue"
	OperationTypeRefund       = "refund"
)
