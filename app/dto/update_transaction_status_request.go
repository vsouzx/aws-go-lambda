package dto

type UpdateTransactionStatusRequest struct {
	Status string `json:"status" validate:"oneof=APPROVED REJECTED PENDING"`
}
