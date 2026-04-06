package logger

import (
	"fmt"
	"major/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewErrorResponse(c *gin.Context, log *logrus.Logger, writeLog bool, code int, message string) {
	currRequestId, _ := c.Get("requestId")

	if writeLog {
		log.Errorf("%d %s %s", code, currRequestId, message)
	}

	errorResponse := models.ErrorResponse{
		Code:    code,
		Message: message,
	}

	c.AbortWithStatusJSON(code, models.ResponseAPI{
		Errors:    errorResponse,
		Success:   false,
		RequestId: fmt.Sprint(currRequestId),
	})
}
