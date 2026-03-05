package subscriptions

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
)

type SubscriptionsRepositoryInterface interface {
	Create(c *gin.Context, entity SubscriptionModel) (SubscriptionModel, error)

	GetByID(c *gin.Context, id uuid.UUID) (SubscriptionModel, error)
	UpdateByID(c *gin.Context, id uuid.UUID, entity SubscriptionModel) (SubscriptionModel, error)
	DeleteByID(c *gin.Context, id uuid.UUID) error
	List(c *gin.Context, limit int, offset int) ([]SubscriptionModel, error)

	SumTotalCost(c *gin.Context, totalInputDTO subscriptionsContracts.TotalInputDTO) (int64, error)
}
