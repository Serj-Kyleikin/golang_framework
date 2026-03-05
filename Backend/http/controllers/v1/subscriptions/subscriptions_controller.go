package controllers

import (
	"github.com/gin-gonic/gin"

	"subscriptions/Backend/core/http/response"
	"subscriptions/Backend/services/subscriptions"

	subscriptionsRequests "subscriptions/Backend/http/requests/subscriptions"
	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
)

var subscriptionsService subscriptionsContracts.SubscriptionsServiceContract

func Construct() {
	subscriptionsService = subscriptions.NewSubscriptionsService()
}

func Create(context *gin.Context) {

	request, uid, err := subscriptionsRequests.ValidateCreate(context)
	if err != nil {
		response.UnprocessableEntity(context, err.Error())
		return
	}

	createOutputDTO, err := subscriptionsService.Create(context, subscriptionsContracts.CreateInputDTO{
		ServiceName: request.ServiceName,
		Price:       request.Price,
		UserID:      uid,
		StartDate:   request.StartTime,
		EndDate:     request.EndTime,
	})

	if err != nil {
		response.InternalServer(context, "Create subscription failed", err.Error())
		return
	}

	response.Created(context, createOutputDTO, "Subscription created")
}

func Get(context *gin.Context) {
	id, err := subscriptionsRequests.ValidateGet(context)
	if err != nil {
		response.UnprocessableEntity(context, err.Error())
		return
	}

	createOutputDTO, err := subscriptionsService.Get(context, subscriptionsContracts.GetInputDTO{ID: id})
	if err != nil {
		response.InternalServer(context, "Get subscription failed", err.Error())
		return
	}

	response.Ok(context, createOutputDTO)
}

func List(context *gin.Context) {
	req, err := subscriptionsRequests.ValidateList(context)
	if err != nil {
		response.UnprocessableEntity(context, err.Error())
		return
	}

	createOutputDTO, err := subscriptionsService.List(context, subscriptionsContracts.ListInputDTO{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		response.InternalServer(context, "List subscriptions failed", err.Error())
		return
	}

	response.Ok(context, createOutputDTO)
}

func Update(context *gin.Context) {
	id, req, uid, err := subscriptionsRequests.ValidateUpdate(context)
	if err != nil {
		response.UnprocessableEntity(context, err.Error())
		return
	}

	createOutputDTO, err := subscriptionsService.Update(context, subscriptionsContracts.UpdateInputDTO{
		ID:          id,
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      uid,
		StartDate:   req.StartTime,
		EndDate:     req.EndTime,
	})
	if err != nil {
		response.InternalServer(context, "Update subscription failed", err.Error())
		return
	}

	response.Created(context, createOutputDTO, "Subscription updated")
}

func Delete(context *gin.Context) {
	id, err := subscriptionsRequests.ValidateDelete(context)
	if err != nil {
		response.UnprocessableEntity(context, err.Error())
		return
	}

	err = subscriptionsService.Delete(context, subscriptionsContracts.DeleteInputDTO{ID: id})
	if err != nil {
		response.InternalServer(context, "Delete subscription failed", err.Error())
		return
	}

	response.Ok(context, "Subscription deleted")
}

func Total(context *gin.Context) {
	req, err := subscriptionsRequests.ValidateTotal(context)
	if err != nil {
		response.UnprocessableEntity(context, err.Error())
		return
	}

	createOutputDTO, err := subscriptionsService.Total(context, subscriptionsContracts.TotalInputDTO{
		StartDate:   req.StartTime,
		EndDate:     req.EndTime,
		UserID:      req.UserID,
		ServiceName: req.ServiceName,
	})
	if err != nil {
		response.InternalServer(context, "Total subscriptions cost failed", err.Error())
		return
	}

	response.Ok(context, createOutputDTO)
}
