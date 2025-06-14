package main

import (
	"context"
	"fmt"

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
			Path: "/hello",
			HTTPMethod: "GET",
			Body: "{\"message\": \"Hello, Lambda!\"}",
		}

	response, err := handler(context.Background(), mockRequest)
	if err != nil {	
		fmt.Println("Error:", err.Error())
	}
	println("Response:", response.Body)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, World!",
	}, nil
}

