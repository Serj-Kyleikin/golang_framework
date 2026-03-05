package handlers

import (
	"github.com/gin-gonic/gin"

	subscriptionsDB "subscriptions/Backend/db/repositories/subscriptions"
	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
)

type TotalHandler struct {
	subscriptionsRepository subscriptionsDB.SubscriptionsRepositoryInterface
}

func NewTotalHandler() *TotalHandler {
	return &TotalHandler{
		subscriptionsRepository: subscriptionsDB.Construct(),
	}
}

func (totalHandler *TotalHandler) Handle(context *gin.Context, totalInputDTO subscriptionsContracts.TotalInputDTO) (subscriptionsContracts.TotalOutputDTO, error) {

	sumTotalCost, err := totalHandler.subscriptionsRepository.SumTotalCost(context, totalInputDTO)
	if err != nil {
		return subscriptionsContracts.TotalOutputDTO{}, err
	}

	return subscriptionsContracts.TotalOutputDTO{Total: sumTotalCost}, nil
}
