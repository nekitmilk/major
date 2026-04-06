package severities

import (
	"fmt"
	"major/internal/handler/constHandler"
	"major/internal/models"
	"major/pkg/customErrors/handlerErrors"
	"major/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Вывод списка уровней опасности
// @Description Получение списка всех уровней опасности
// @Tags severities
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} models.ResponseAPI{result=[]models.Severity}
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/severities/ [get]
func (sh *SeveritiesHandler) GetAllSeverities(ctx *gin.Context) {
	getAllSeverities, getAllSeveritiesError := sh.service.SeverityService.GetAllSeverities()
	if getAllSeveritiesError != nil {
		messageError := fmt.Sprintf("Error get list severities: %s", getAllSeveritiesError.Error())
		httpStatus := handlerErrors.MapError(getAllSeveritiesError)
		logger.NewErrorResponse(ctx, sh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Get list severities",
		Result:    getAllSeverities,
	}

	sh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}
