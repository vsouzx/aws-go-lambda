package main

import (
	"context"
	"souzalambdago/config"
	"souzalambdago/factory"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if _, err := config.NewDb(); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	factory := factory.NewFactory()

	svc, ok := factory.GetService(request.HTTPMethod, request.Path)
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Route not found",
		}, nil
	}

	return svc.Execute(request)
}
