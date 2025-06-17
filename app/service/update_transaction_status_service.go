package service

import (
	"fmt"
	"net/http"
	"souzalambdago/dto"
	"souzalambdago/model"
	"souzalambdago/repository"
	"souzalambdago/service/validation"

	"github.com/aws/aws-lambda-go/events"
)

type UpdateTransactionStatusService struct {
	repository        repository.TransactionRepositoryInterface
	payloadValidation validation.PayloadValidation
}

func NewUpdateTransactionStatusService(repository repository.TransactionRepositoryInterface) *UpdateTransactionStatusService {
	return &UpdateTransactionStatusService{
		repository:        repository,
		payloadValidation: validation.PayloadValidation{},
	}
}

func (gps *UpdateTransactionStatusService) Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	transactionID := request.QueryStringParameters["transaction_id"]
	if transactionID == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "transaction_id é obrigatório",
		}, nil
	}

	var payload dto.UpdateTransactionStatusRequest
	if err := gps.payloadValidation.ValidatePayload(request.Body, &payload); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Invalid payload: " + err.Error(),
		}, nil
	}

	transaction, err := gps.repository.GetTransactionById(transactionID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Error getting transaction by id: %s", transactionID),
		}, nil
	}

	if transaction == (model.Transaction{}) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       fmt.Sprintf("Transaction not found: %s", transactionID),
		}, nil
	}

	if err := gps.repository.UpdateTransactionStatus(payload.Status, transactionID); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Error updating transaction status: %s", err.Error()),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
	}, nil
}
