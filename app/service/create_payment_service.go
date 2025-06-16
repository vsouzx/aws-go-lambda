package service

import "github.com/aws/aws-lambda-go/events"

type CreatePaymentService struct{}

func NewCreatePaymentService() *CreatePaymentService {
	return &CreatePaymentService{}
}

func (cps *CreatePaymentService) Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// This function is a placeholder for the actual service logic.
	// It should be implemented in the specific service that uses this interface.
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       "Payement created successfully",
	}, nil
}
