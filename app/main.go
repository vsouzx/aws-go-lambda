package main

import (
	"context"
	"fmt"
	"souzalambdago/config"
	"souzalambdago/factory"
	"souzalambdago/repository"

	"github.com/aws/aws-lambda-go/events"
)
func main() {
    mockRequest := events.APIGatewayProxyRequest{
        QueryStringParameters: map[string]string{
            "from_user_id": "5",
        },
        Headers: map[string]string{
            "Authorization": "Bearer fake-token",
        },
        Path:       "/payment",
        HTTPMethod: "GET",
        Body:       `{
            "from_user_id": 1,
            "to_user_id": 2,
            "amount": 100.50,
            "type": "TRANSFER",
            "description": "Payment for services"
        }`,
    }

    response, err := handler(context.Background(), mockRequest)
    if err != nil {
        fmt.Println("Error:", err.Error())
    }
    fmt.Println("Response:", response)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	repository := repository.NewTransactionRepository(config.NewDb())
	factory := factory.NewFactory(repository)
	service := factory.GetService(request.HTTPMethod, request.Path)
	return service.Execute(request)
}
