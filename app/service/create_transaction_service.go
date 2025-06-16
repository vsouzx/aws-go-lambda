package service

import (
	"souzalambdago/dto"
	"souzalambdago/repository"
	"souzalambdago/service/validation"

	"github.com/aws/aws-lambda-go/events"
)

type CreateTransactionService struct {
	repository        repository.TransactionRepositoryInterface
	payloadValidation validation.PayloadValidation
}

func NewCreatePaymentService(repository repository.TransactionRepositoryInterface) *CreateTransactionService {
	return &CreateTransactionService{
		repository:        repository,
		payloadValidation: validation.PayloadValidation{},
	}
}

func (cps *CreateTransactionService) Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var payload dto.CreateTransactionRequest
	if err := cps.payloadValidation.ValidatePayload(request.Body, &payload); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Invalid payload: " + err.Error(),
		}, nil
	}
	
	model := dto.NewTransactionModel(payload)
	
	if err := cps.repository.CreateTransaction(model); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error creating transaction: " + err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       "Transaction created successfully",
	}, nil
}
