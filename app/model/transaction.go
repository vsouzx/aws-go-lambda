package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID          string          `db:"id"`
	FromUserID  int             `db:"from_user_id"`
	ToUserID    int             `db:"to_user_id"`
	Amount      decimal.Decimal `db:"amount"`
	Type        string          `db:"type"`
	Status      string          `db:"status"`
	Description string          `db:"description"`
	CreatedAt   time.Time       `db:"created_at"`
}
