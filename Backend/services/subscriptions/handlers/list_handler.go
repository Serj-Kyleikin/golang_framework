package handlers

import (
	"github.com/gin-gonic/gin"

	"subscriptions/Backend/kernel/utils"

	subscriptionsDB "subscriptions/Backend/db/repositories/subscriptions"
	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
)

type ListHandler struct {
	subscriptionsRepository subscriptionsDB.SubscriptionsRepositoryInterface
}

func NewListHandler() *ListHandler {
	return &ListHandler{
		subscriptionsRepository: subscriptionsDB.Construct(),
	}
}

func (listHandler *ListHandler) Handle(context *gin.Context, listInputDTO subscriptionsContracts.ListInputDTO) (subscriptionsContracts.ListOutputDTO, error) {

	subscriptionModelList, err := listHandler.subscriptionsRepository.List(context, listInputDTO.Limit, listInputDTO.Offset)
	if err != nil {
		return subscriptionsContracts.ListOutputDTO{}, err
	}

	getOutputDTOItems := make([]subscriptionsContracts.GetOutputDTO, 0, len(subscriptionModelList))

	for _, subscriptionModel := range subscriptionModelList {

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

		getOutputDTOItems = append(getOutputDTOItems, getOutputDTO)
	}

	return subscriptionsContracts.ListOutputDTO{
		Items:  getOutputDTOItems,
		Limit:  listInputDTO.Limit,
		Offset: listInputDTO.Offset,
	}, nil
}
