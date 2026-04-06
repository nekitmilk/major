package health

import (
	"fmt"
	"major/internal/handler/constHandler"
	"major/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (healthHandler *HealthHandler) HealthCheck(ctx *gin.Context) {
	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Health check app",
	}

	healthHandler.log.Debugf("🚀 response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}
