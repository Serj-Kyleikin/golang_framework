package handlers

import (
	"github.com/gin-gonic/gin"

	subscriptionsDB "subscriptions/Backend/db/repositories/subscriptions"
	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
)

type DeleteHandler struct {
	subscriptionsRepository subscriptionsDB.SubscriptionsRepositoryInterface
}

func NewDeleteHandler() *DeleteHandler {
	return &DeleteHandler{
		subscriptionsRepository: subscriptionsDB.Construct(),
	}
}

func (deleteHandler *DeleteHandler) Handle(context *gin.Context, deleteInputDTO subscriptionsContracts.DeleteInputDTO) error {
	return deleteHandler.subscriptionsRepository.DeleteByID(context, deleteInputDTO.ID)
}
