package http

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader("X-Request-Id")
		if reqID == "" {
			reqID = uuid.NewString()
		}
		c.Writer.Header().Set("X-Request-Id", reqID)
		c.Set("request_id", reqID)

		start := time.Now()
		c.Next()
		lat := time.Since(start)

		slog.Info("request",
			"request_id", reqID,
			"method", c.Request.Method,
			"path", c.FullPath(),
			"status", c.Writer.Status(),
			"latency_ms", lat.Milliseconds(),
			"ip", c.ClientIP(),
		)
	}
}

func RecoverWithLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("panic", "error", r)
				c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
			}
		}()
		c.Next()
	}
}
