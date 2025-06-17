package util

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type ResponseUtil struct{}

func (ru *ResponseUtil) BuildResponse(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	responseJson, err := json.Marshal(body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error marshalling response body",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(responseJson),
	}, nil
}
