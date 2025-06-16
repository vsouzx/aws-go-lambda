package service

import "github.com/aws/aws-lambda-go/events"

type Service interface {
	Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
