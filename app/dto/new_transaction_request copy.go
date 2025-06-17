package dto

import (
	"souzalambdago/model"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateTransactionRequest struct {
	FromUserID  int             `json:"from_user_id" validate:"required"`
	ToUserID    int             `json:"to_user_id" validate:"required"`
	Amount      decimal.Decimal `json:"amount" validate:"required,gt=0"`
	Type        string          `json:"type" validate:"required,oneof=PIX REFUND TRANSFER"`
	Description string          `json:"description"`
}

func NewTransactionModel(request CreateTransactionRequest) model.Transaction {
	return model.Transaction{
		ID:          uuid.NewString(),
		FromUserID:  request.FromUserID,
		ToUserID:    request.ToUserID,
		Amount:      request.Amount,
		Type:        request.Type,
		Status:      "PENDING",
		Description: request.Description,
		CreatedAt:   time.Now(),
	}
}
