package v1

import (
	"github.com/gin-gonic/gin"

	subscriptionsController "subscriptions/Backend/http/controllers/v1/subscriptions"
)

type Routes struct{}

func Bind() *Routes {
	subscriptionsController.Construct()
	return &Routes{}
}

func (r *Routes) Register(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		sub := v1.Group("/subscriptions")
		{
			sub.POST("", subscriptionsController.Create)
			sub.GET("/total", subscriptionsController.Total)
			sub.GET("/:id", subscriptionsController.Get)
			sub.GET("", subscriptionsController.List)
			sub.PUT("/:id", subscriptionsController.Update)
			sub.DELETE("/:id", subscriptionsController.Delete)
		}
	}
}
