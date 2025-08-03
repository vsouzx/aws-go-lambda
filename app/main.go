package main

import (
	"context"
	"souzalambdago/config"
	"souzalambdago/factory"
	"souzalambdago/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db := config.NewDb()
	defer db.Close()
	repository := repository.NewTransactionRepository(db)
	factory := factory.NewFactory(repository)
	service := factory.GetService(request.HTTPMethod, request.Path)
	return service.Execute(request)
}
