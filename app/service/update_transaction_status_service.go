package service

import (
	"fmt"
	"net/http"
	"souzalambdago/dto"
	"souzalambdago/model"
	"souzalambdago/repository"
	"souzalambdago/service/validation"
	"souzalambdago/util"

	"github.com/aws/aws-lambda-go/events"
)

type UpdateTransactionStatusService struct {
	repository        repository.TransactionRepositoryInterface
	payloadValidation validation.PayloadValidation
	responseUtil      util.ResponseUtil
}

func NewUpdateTransactionStatusService(repository repository.TransactionRepositoryInterface) *UpdateTransactionStatusService {
	return &UpdateTransactionStatusService{
		repository:        repository,
		payloadValidation: validation.PayloadValidation{},
		responseUtil:      util.ResponseUtil{},
	}
}

func (utss *UpdateTransactionStatusService) Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	transactionID := request.QueryStringParameters["transaction_id"]
	if transactionID == "" {
		return utss.responseUtil.BuildResponse(http.StatusInternalServerError, map[string]string{
			"message": "transaction_id é obrigatório",
		})
	}

	var payload dto.UpdateTransactionStatusRequest
	if err := utss.payloadValidation.ValidatePayload(request.Body, &payload); err != nil {
		return utss.responseUtil.BuildResponse(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Invalid payload:  %s", err.Error()),
		})
	}

	transaction, err := utss.repository.GetTransactionById(transactionID)
	if err != nil {
		return utss.responseUtil.BuildResponse(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("Error retrieving transaction: %s", err.Error()),
		})
	}

	if transaction == (model.Transaction{}) {
		return utss.responseUtil.BuildResponse(http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("Transaction not found: %s", transactionID),
		})
	}

	if err := utss.repository.UpdateTransactionStatus(payload.Status, transactionID); err != nil {
		return utss.responseUtil.BuildResponse(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error updating transaction status: %s", err.Error()),
		})
	}

	return utss.responseUtil.BuildResponse(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Transaction status updated successfully: %s", transactionID),
	})
}
