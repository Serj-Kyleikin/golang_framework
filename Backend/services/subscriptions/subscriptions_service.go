package subscriptions

import (
	"github.com/gin-gonic/gin"

	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
	subscriptionsHandlers "subscriptions/Backend/services/subscriptions/handlers"
)

type SubscriptionsService struct {
	createHandler subscriptionsContracts.CreateHandlerContract
	getHandler    subscriptionsContracts.GetHandlerContract
	listHandler   subscriptionsContracts.ListHandlerContract
	updateHandler subscriptionsContracts.UpdateHandlerContract
	deleteHandler subscriptionsContracts.DeleteHandlerContract
	totalHandler  subscriptionsContracts.TotalHandlerContract
}

func NewSubscriptionsService() subscriptionsContracts.SubscriptionsServiceContract {
	return &SubscriptionsService{
		createHandler: subscriptionsHandlers.NewCreateHandler(),
		getHandler:    subscriptionsHandlers.NewGetHandler(),
		listHandler:   subscriptionsHandlers.NewListHandler(),
		updateHandler: subscriptionsHandlers.NewUpdateHandler(),
		deleteHandler: subscriptionsHandlers.NewDeleteHandler(),
		totalHandler:  subscriptionsHandlers.NewTotalHandler(),
	}
}

func (service *SubscriptionsService) Create(context *gin.Context, createInputDTO subscriptionsContracts.CreateInputDTO) (subscriptionsContracts.CreateOutputDTO, error) {
	return service.createHandler.Handle(context, createInputDTO)
}

func (service *SubscriptionsService) Get(context *gin.Context, getInputDTO subscriptionsContracts.GetInputDTO) (subscriptionsContracts.GetOutputDTO, error) {
	return service.getHandler.Handle(context, getInputDTO)
}

func (service *SubscriptionsService) List(context *gin.Context, listInputDTO subscriptionsContracts.ListInputDTO) (subscriptionsContracts.ListOutputDTO, error) {
	return service.listHandler.Handle(context, listInputDTO)
}

func (service *SubscriptionsService) Update(context *gin.Context, updateInputDTO subscriptionsContracts.UpdateInputDTO) (subscriptionsContracts.UpdateOutputDTO, error) {
	return service.updateHandler.Handle(context, updateInputDTO)
}

func (service *SubscriptionsService) Delete(context *gin.Context, deleteInputDTO subscriptionsContracts.DeleteInputDTO) error {
	return service.deleteHandler.Handle(context, deleteInputDTO)
}

func (service *SubscriptionsService) Total(context *gin.Context, totalInputDTO subscriptionsContracts.TotalInputDTO) (subscriptionsContracts.TotalOutputDTO, error) {
	return service.totalHandler.Handle(context, totalInputDTO)
}
