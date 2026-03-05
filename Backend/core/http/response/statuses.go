package response

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, data interface{}) {

	Status(http.StatusOK).
		Data(data).
		Success(c)
}

func Created(c *gin.Context, data interface{}, message string) {

	Status(http.StatusCreated).
		Data(data).
		Message(message).
		Success(c)
}

func UnprocessableEntity(c *gin.Context, message string) {

	Status(http.StatusUnprocessableEntity).
		Message(message).
		Fail(c)
}

func BadRequest(c *gin.Context, message string, logMessage string) {

	slog.Error(message, "error", logMessage)

	Status(http.StatusBadRequest).
		Message(message).
		Fail(c)
}

func NotFound(c *gin.Context, message string) {

	Status(http.StatusNotFound).
		Message(message).
		Fail(c)
}

func InternalServer(c *gin.Context, message string, logMessage string) {

	slog.Error(message, "error", logMessage)

	Status(http.StatusInternalServerError).
		Message(message).
		Fail(c)
}
