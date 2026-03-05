package http

import (
	"net/http"

	"subscriptions/Backend/openapi"

	"github.com/gin-gonic/gin"
)

func ServeOpenAPIYAML() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Data(http.StatusOK, "application/yaml; charset=utf-8", openapi.Spec)
	}
}
