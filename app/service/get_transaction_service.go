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
	responseUtil util.ResponseUtil
}

func NewGetPaymentService(repository repository.TransactionRepositoryInterface) *GetTransactionService {
	return &GetTransactionService{
		repository:   repository,
		responseUtil: util.ResponseUtil{},
	}
}

func (gps *GetTransactionService) Execute(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fromUserId, err := strconv.Atoi(request.QueryStringParameters["from_user_id"])
	if err != nil {
		return gps.responseUtil.BuildResponse(http.StatusInternalServerError, map[string]string{
			"Body": "Invalid from_user_id parameter",
		})
	}

	response, err := gps.repository.GetTransactionsByFromUser(fromUserId)
	if err != nil {
		return gps.responseUtil.BuildResponse(http.StatusInternalServerError, map[string]string{
			"Body": fmt.Sprintf("Error getting transactions from user %d\n:", fromUserId),
		})
	}

	return gps.responseUtil.BuildResponse(http.StatusOK, response)
}
