package contracts

import "github.com/gin-gonic/gin"

type SubscriptionsServiceContract interface {
	Create(context *gin.Context, in CreateInputDTO) (CreateOutputDTO, error)
	Get(context *gin.Context, in GetInputDTO) (GetOutputDTO, error)
	List(context *gin.Context, in ListInputDTO) (ListOutputDTO, error)
	Update(context *gin.Context, in UpdateInputDTO) (UpdateOutputDTO, error)
	Delete(context *gin.Context, in DeleteInputDTO) error
	Total(context *gin.Context, in TotalInputDTO) (TotalOutputDTO, error)
}

type CreateHandlerContract interface {
	Handle(context *gin.Context, in CreateInputDTO) (CreateOutputDTO, error)
}

type GetHandlerContract interface {
	Handle(c *gin.Context, in GetInputDTO) (GetOutputDTO, error)
}

type ListHandlerContract interface {
	Handle(c *gin.Context, in ListInputDTO) (ListOutputDTO, error)
}

type UpdateHandlerContract interface {
	Handle(c *gin.Context, in UpdateInputDTO) (UpdateOutputDTO, error)
}

type DeleteHandlerContract interface {
	Handle(c *gin.Context, in DeleteInputDTO) error
}

type TotalHandlerContract interface {
	Handle(c *gin.Context, in TotalInputDTO) (TotalOutputDTO, error)
}
