package main

import (
	"context"
	"fmt"
	"souzalambdago/factory"

	"github.com/aws/aws-lambda-go/events"
)

func main() {
	mockRequest := events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"name": "Vinicius",
		},
		Headers: map[string]string{
			"Authorization": "Bearer fake-token",
		},
		Path:       "/payment",
		HTTPMethod: "GET",
		Body:       "{\"message\": \"Hello, Lambda!\"}",
	}

	response, err := handler(context.Background(), mockRequest)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Print("Response:", response.Body)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	/*if _, err := config.NewDb(); err != nil {
	#	fmt.Printf("Failed to connect to database: %v\n", err)
	#	return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}*/
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
