package subscriptions

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListSubscriptionsRequest struct {
	UserID *uuid.UUID
	Limit  int
	Offset int
}

func ValidateList(context *gin.Context) (ListSubscriptionsRequest, error) {

	var out ListSubscriptionsRequest

	limit := context.Query("limit")
	offset := context.Query("offset")

	out.Limit = 50
	out.Offset = 0

	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil || limitInt < 1 || limitInt > 500 {
			return ListSubscriptionsRequest{}, fmt.Errorf("%s", "limit must be 1..500")
		}
		out.Limit = limitInt
	}

	if offset != "" {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil || offsetInt < 0 {
			return ListSubscriptionsRequest{}, fmt.Errorf("%s", "offset must be >= 0")
		}
		out.Offset = offsetInt
	}

	return out, nil
}
