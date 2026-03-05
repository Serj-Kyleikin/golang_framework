package subscriptions

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateDelete(context *gin.Context) (uuid.UUID, error) {

	raw := context.Param("id")

	id, err := uuid.Parse(raw)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s", "Invalid id")
	}

	return id, nil
}
