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

// @Summary Удаление угрозы
// @Description Удаление угрозы из системы по ее идентификатору
// @Tags severities
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param severityId path string true "Идентификатор угрозы"
// @Success 200 {object} models.ResponseAPI{result=models.SeverityResponse} "Успешное удаление"
// @Failure 404 {object} models.ResponseAPI
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/severities/{severityId} [delete]
func (uh *SeveritiesHandler) DeleteSeverity(ctx *gin.Context) {
	idSeverity := ctx.Param("severityId")

	deleteSeverityError := uh.service.SeverityService.DeleteSeverity(idSeverity)
	if deleteSeverityError != nil {
		messageError := fmt.Sprintf("Error delete severity: %s", deleteSeverityError.Error())
		httpStatus := handlerErrors.MapError(deleteSeverityError)
		logger.NewErrorResponse(ctx, uh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   fmt.Sprintf("Delete severity: %s", idSeverity),
		Result: models.SeverityResponse{
			Id: idSeverity,
		},
	}

	uh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}
