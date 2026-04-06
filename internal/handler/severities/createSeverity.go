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

// @Summary Добавление уровня опасности
// @Description Создание уровня опасности
// @Tags severities
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param request body models.CreateSeverityRequest true "Данные нового уровня опасности"
// @Success 201 {object} models.ResponseAPI{result=models.SeverityResponse}
// @Failure 400 {object} models.ResponseAPI
// @Failure 409 {object} models.ResponseAPI
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/severities [post]
func (sh *SeveritiesHandler) CreateSeverity(ctx *gin.Context) {
	var request models.CreateSeverityRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		messageError := fmt.Sprintf("Error parsing severity data: %s", err.Error())
		logger.NewErrorResponse(ctx, sh.log, true, http.StatusBadRequest, messageError)
		return
	}

	createSeverityId, createSeverityError := sh.service.SeverityService.CreateSeverity(&request)
	if createSeverityError != nil {
		messageError := fmt.Sprintf("Error creating severity: %s", createSeverityError.Error())
		httpStatus := handlerErrors.MapError(createSeverityError)
		logger.NewErrorResponse(ctx, sh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)
	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Create severity",
		Result: models.SeverityResponse{
			Id: createSeverityId,
		},
	}

	sh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusCreated, responseApi)
}
