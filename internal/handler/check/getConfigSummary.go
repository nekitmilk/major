package check

import (
	"fmt"
	"major/internal/handler/constHandler"
	"major/internal/models"
	"major/pkg/customErrors/handlerErrors"
	"major/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Получение сводки для конфига
// @Description Получение сводки для конфига
// @Tags check
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param request body models.CheckRequest true "Данные конфига"
// @Success 200 {object} models.ResponseAPI{result=models.CheckResponseFullData}
// @Failure 400 {object} models.ResponseAPI
// @Failure 409 {object} models.ResponseAPI
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/check [post]
func (nh *CheckHandler) GetConfigSummary(ctx *gin.Context) {
	var request models.CheckRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		messageError := fmt.Sprintf("Error parsing config data: %s", err.Error())
		logger.NewErrorResponse(ctx, nh.log, true, http.StatusBadRequest, messageError)
		return
	}

	getConfigSummary, getConfigSummaryError := nh.service.CheckService.GetConfigSummary(&request)
	if getConfigSummaryError != nil {
		messageError := fmt.Sprintf("Error get config summary: %s", getConfigSummaryError.Error())
		httpStatus := handlerErrors.MapError(getConfigSummaryError)
		logger.NewErrorResponse(ctx, nh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Get config summary",
		Result:    getConfigSummary,
	}

	nh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}
