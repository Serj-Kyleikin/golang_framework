package handlers

import (
	"github.com/gin-gonic/gin"

	"subscriptions/Backend/kernel/utils"

	subscriptionsDB "subscriptions/Backend/db/repositories/subscriptions"
	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
)

type UpdateHandler struct {
	subscriptionsRepository subscriptionsDB.SubscriptionsRepositoryInterface
}

func NewUpdateHandler() *UpdateHandler {
	return &UpdateHandler{
		subscriptionsRepository: subscriptionsDB.Construct(),
	}
}

func (updateHandler *UpdateHandler) Handle(context *gin.Context, updateInputDTO subscriptionsContracts.UpdateInputDTO) (subscriptionsContracts.UpdateOutputDTO, error) {

	newSubscriptionModel := subscriptionsDB.SubscriptionModel{
		ServiceName: updateInputDTO.ServiceName,
		Price:       updateInputDTO.Price,
		UserID:      updateInputDTO.UserID,
		StartDate:   updateInputDTO.StartDate,
		EndDate:     updateInputDTO.EndDate,
	}

	subscriptionModel, err := updateHandler.subscriptionsRepository.UpdateByID(context, updateInputDTO.ID, newSubscriptionModel)
	if err != nil {
		return subscriptionsContracts.UpdateOutputDTO{}, err
	}

	updateOutputDTO := subscriptionsContracts.UpdateOutputDTO{
		ID:          subscriptionModel.ID.String(),
		ServiceName: subscriptionModel.ServiceName,
		Price:       subscriptionModel.Price,
		UserID:      subscriptionModel.UserID.String(),
		StartDate:   utils.FormatMonthUTC(subscriptionModel.StartDate),
	}

	if subscriptionModel.EndDate != nil {
		endDateFormatMonthUTC := utils.FormatMonthUTC(*subscriptionModel.EndDate)
		updateOutputDTO.EndDate = &endDateFormatMonthUTC
	}

	return updateOutputDTO, nil
}
