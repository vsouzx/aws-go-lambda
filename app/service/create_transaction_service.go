package service

import (
	"fmt"
	"net/http"
	"souzalambdago/dto"
	"souzalambdago/repository"
	"souzalambdago/service/validation"
	"souzalambdago/util"

	"github.com/aws/aws-lambda-go/events"
)

type CreateTransactionService struct {
	repository        repository.TransactionRepositoryInterface
	payloadValidation validation.PayloadValidation
	responseUtil      util.ResponseUtil
}

func NewCreatePaymentService(repository repository.TransactionRepositoryInterface) *CreateTransactionService {
	return &CreateTransactionService{
		repository:        repository,
		payloadValidation: validation.PayloadValidation{},
		responseUtil:      util.ResponseUtil{},
	}
}

func (cps *CreateTransactionService) Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var payload dto.CreateTransactionRequest
	if err := cps.payloadValidation.ValidatePayload(request.Body, &payload); err != nil {
		return cps.responseUtil.BuildResponse(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Invalid payload: %s", err.Error()),
		})
	}

	model := dto.NewTransactionModel(payload)

	if err := cps.repository.CreateTransaction(model); err != nil {
		return cps.responseUtil.BuildResponse(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Error creating transaction: %s", err.Error()),
		})
	}

	return cps.responseUtil.BuildResponse(201, map[string]string{
		"message": "Transaction created successfully",
	})
}
