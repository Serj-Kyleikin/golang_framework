package subscriptions

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"subscriptions/Backend/kernel/utils"
)

type CreateSubscriptionRequest struct {
	ServiceName string  `json:"service_name"`
	Price       int     `json:"price"`
	UserID      string  `json:"user_id"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date,omitempty"`

	StartTime time.Time  `json:"-"`
	EndTime   *time.Time `json:"-"`
}

var monthYearRegExp = regexp.MustCompile(`^(0[1-9]|1[0-2])-[12]\d{3}$`)

func ValidateCreate(context *gin.Context) (request CreateSubscriptionRequest, userID uuid.UUID, err error) {

	var uid uuid.UUID
	var parseErr error
	var errorMessage string

	if err := context.ShouldBindJSON(&request); err != nil {
		errorMessage = "invalid json"
	} else {

		if request.ServiceName == "" {
			errorMessage = "service_name is empty string"
		}

		if request.Price < 0 {
			errorMessage = "price must be >= 0"
		}

		uid, parseErr = uuid.Parse(request.UserID)
		if parseErr != nil {
			errorMessage = "invalid user_id"
		} else {

			if !monthYearRegExp.MatchString(request.StartDate) {
				errorMessage = "start_date must be in MM-YYYY format"
			} else {

				start, err := utils.ParseMonthUTC(request.StartDate)
				if err != nil {
					errorMessage = err.Error()
				} else {

					request.StartTime = start

					if request.EndDate != nil {
						if !monthYearRegExp.MatchString(*request.EndDate) {
							errorMessage = "end_date must be in MM-YYYY format"
						} else {

							end, err := parseAndValidateEndDate(request.EndDate, start)
							if err != nil {
								errorMessage = err.Error()
							} else if end != nil {
								request.EndTime = end
							}
						}
					}
				}
			}
		}
	}

	if errorMessage != "" {
		return CreateSubscriptionRequest{}, uid, fmt.Errorf("%s", errorMessage)
	}

	return request, uid, nil
}

func parseAndValidateEndDate(endDate *string, start time.Time) (*time.Time, error) {
	if endDate == nil {
		return nil, nil
	}

	endDateParseMonthUTC, err := utils.ParseMonthUTC(*endDate)
	if err != nil {
		return nil, fmt.Errorf("%s", "end_date must be in MM-YYYY format")
	}

	if endDateParseMonthUTC.Before(start) {
		return nil, fmt.Errorf("%s", "end_date must be >= start_date")
	}

	return &endDateParseMonthUTC, nil
}
