package handlers

import (
	"github.com/gin-gonic/gin"

	"subscriptions/Backend/kernel/utils"

	subscriptionsDB "subscriptions/Backend/db/repositories/subscriptions"
	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
)

type GetHandler struct {
	subscriptionsRepository subscriptionsDB.SubscriptionsRepositoryInterface
}

func NewGetHandler() *GetHandler {
	return &GetHandler{
		subscriptionsRepository: subscriptionsDB.Construct(),
	}
}

func (getHandler *GetHandler) Handle(context *gin.Context, getInputDTO subscriptionsContracts.GetInputDTO) (subscriptionsContracts.GetOutputDTO, error) {

	subscriptionModel, err := getHandler.subscriptionsRepository.GetByID(context, getInputDTO.ID)
	if err != nil {
		return subscriptionsContracts.GetOutputDTO{}, err
	}

	getOutputDTO := subscriptionsContracts.GetOutputDTO{
		ID:          subscriptionModel.ID.String(),
		ServiceName: subscriptionModel.ServiceName,
		Price:       subscriptionModel.Price,
		UserID:      subscriptionModel.UserID.String(),
		StartDate:   utils.FormatMonthUTC(subscriptionModel.StartDate),
	}

	if subscriptionModel.EndDate != nil {
		endDateFormatMonthUTC := utils.FormatMonthUTC(*subscriptionModel.EndDate)
		getOutputDTO.EndDate = &endDateFormatMonthUTC
	}

	return getOutputDTO, nil
}
