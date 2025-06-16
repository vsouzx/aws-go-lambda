package service

import "github.com/aws/aws-lambda-go/events"

type GetPaymentService struct{}

func NewGetPaymentService() *GetPaymentService {
	return &GetPaymentService{}
}

func (gps *GetPaymentService)Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// This function is a placeholder for the actual service logic.
	// It should be implemented in the specific service that uses this interface.
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       "Payment returned",
	}, nil
}
