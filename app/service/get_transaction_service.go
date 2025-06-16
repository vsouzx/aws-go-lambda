package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"souzalambdago/repository"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

type GetTransactionService struct {
	repository repository.TransactionRepositoryInterface
}

func NewGetPaymentService(repository repository.TransactionRepositoryInterface) *GetTransactionService {
	return &GetTransactionService{
		repository: repository,
	}
}

func (gps *GetTransactionService) Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fromUserId, err := strconv.Atoi(request.QueryStringParameters["from_user_id"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid from_user_id parameter",
		}, nil
	}

	response, err := gps.repository.GetTransactionsByFromUser(fromUserId)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Error getting transactions from user %d\n:", fromUserId),
		}, nil
	}

	responseJson, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       string(responseJson),
	}, nil
}
