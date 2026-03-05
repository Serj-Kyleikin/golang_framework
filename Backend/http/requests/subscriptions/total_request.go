package subscriptions

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"subscriptions/Backend/kernel/utils"
)

type TotalSubscriptionsRequest struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`

	UserID      *uuid.UUID `json:"-"`
	ServiceName *string    `json:"-"`

	StartTime time.Time `json:"-"`
	EndTime   time.Time `json:"-"`
}

func ValidateTotal(context *gin.Context) (TotalSubscriptionsRequest, error) {

	var req TotalSubscriptionsRequest
	var errorMessage string

	startStr := context.Query("start_date")
	endStr := context.Query("end_date")

	if startStr == "" {
		errorMessage = "start_date is required"
	} else if !monthYearRegExp.MatchString(startStr) {
		errorMessage = "start_date must be in MM-YYYY format"
	}

	if errorMessage == "" {
		if endStr == "" {
			errorMessage = "end_date is required"
		} else if !monthYearRegExp.MatchString(endStr) {
			errorMessage = "end_date must be in MM-YYYY format"
		}
	}

	if errorMessage == "" {
		start, err := utils.ParseMonthUTC(startStr)
		if err != nil {
			errorMessage = err.Error()
		} else {
			end, err := utils.ParseMonthUTC(endStr)
			if err != nil {
				errorMessage = err.Error()
			} else if end.Before(start) {
				errorMessage = "end_date must be >= start_date"
			} else {
				req.StartDate = startStr
				req.EndDate = endStr
				req.StartTime = start
				req.EndTime = end
			}
		}
	}

	if errorMessage == "" {
		if raw := context.Query("user_id"); raw != "" {
			uid, err := uuid.Parse(raw)
			if err != nil {
				errorMessage = "invalid user_id"
			} else {
				req.UserID = &uid
			}
		}
	}

	if errorMessage == "" {
		if sn := context.Query("service_name"); sn != "" {
			req.ServiceName = &sn
		}
	}

	if errorMessage != "" {
		return TotalSubscriptionsRequest{}, fmt.Errorf("%s", errorMessage)
	}

	return req, nil
}
