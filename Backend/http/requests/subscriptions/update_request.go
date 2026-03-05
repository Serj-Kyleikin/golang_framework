package subscriptions

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateUpdate(context *gin.Context) (id uuid.UUID, req CreateSubscriptionRequest, userID uuid.UUID, err error) {

	raw := context.Param("id")

	parsedID, parseErr := uuid.Parse(raw)
	if parseErr != nil {
		return uuid.Nil, CreateSubscriptionRequest{}, uuid.Nil, fmt.Errorf("%s", "Invalid id")
	}

	validateCreate, uid, vErr := ValidateCreate(context)
	if vErr != nil {
		return uuid.Nil, CreateSubscriptionRequest{}, uuid.Nil, vErr
	}

	return parsedID, validateCreate, uid, nil
}
