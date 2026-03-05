package http

import (
	"github.com/gin-gonic/gin"

	"subscriptions/Backend/providers"
)

func ConstructRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(RecoverWithLog())
	engine.Use(RequestLogger())

	engine.GET("/openapi.yaml", ServeOpenAPIYAML())

	providers.BindRoutes().RegisterRoutes(engine)

	return engine
}
