package handlers

import (
	"github.com/gin-gonic/gin"

	"subscriptions/Backend/kernel/utils"

	subscriptionsDB "subscriptions/Backend/db/repositories/subscriptions"
	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
)

type CreateHandler struct {
	subscriptionsRepository subscriptionsDB.SubscriptionsRepositoryInterface
}

func NewCreateHandler() *CreateHandler {
	return &CreateHandler{
		subscriptionsRepository: subscriptionsDB.Construct(),
	}
}

func (createHandler *CreateHandler) Handle(context *gin.Context, createInputDTO subscriptionsContracts.CreateInputDTO) (subscriptionsContracts.CreateOutputDTO, error) {

	newSubscriptionModel := subscriptionsDB.SubscriptionModel{
		ServiceName: createInputDTO.ServiceName,
		Price:       createInputDTO.Price,
		UserID:      createInputDTO.UserID,
		StartDate:   createInputDTO.StartDate,
		EndDate:     createInputDTO.EndDate,
	}

	subscriptionModel, err := createHandler.subscriptionsRepository.Create(context, newSubscriptionModel)
	if err != nil {
		return subscriptionsContracts.CreateOutputDTO{}, err
	}

	createOutputDTO := subscriptionsContracts.CreateOutputDTO{
		ID:          subscriptionModel.ID.String(),
		ServiceName: subscriptionModel.ServiceName,
		Price:       subscriptionModel.Price,
		UserID:      subscriptionModel.UserID.String(),
		StartDate:   utils.FormatMonthUTC(subscriptionModel.StartDate),
	}

	if subscriptionModel.EndDate != nil {
		endDateFormatMonthUTC := utils.FormatMonthUTC(*subscriptionModel.EndDate)
		createOutputDTO.EndDate = &endDateFormatMonthUTC
	}

	return createOutputDTO, nil
}
