package service

import (
	"fmt"
	"net/http"
	"souzalambdago/repository"
	"souzalambdago/util"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

type GetTransactionService struct {
	repository   repository.TransactionRepositoryInterface
}

func NewGetPaymentService(repository repository.TransactionRepositoryInterface) *GetTransactionService {
	return &GetTransactionService{
		repository:   repository,
	}
}

func (gps *GetTransactionService) Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fromUserId, err := strconv.Atoi(request.QueryStringParameters["from_user_id"])
	if err != nil {
		return util.BuildResponse(http.StatusInternalServerError, map[string]string{"Body": "Invalid from_user_id parameter",})
	}

	response, err := gps.repository.GetTransactionsByFromUser(fromUserId)
	if err != nil {
		return util.BuildResponse(http.StatusInternalServerError, map[string]string{"Body": fmt.Sprintf("Error getting transactions from user %d\n:", fromUserId),})
	}

	return util.BuildResponse(http.StatusOK, response)
}
