package rules

import (
	"fmt"
	"major/internal/handler/constHandler"
	"major/internal/models"
	"major/pkg/customErrors/handlerErrors"
	"major/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Удаление правила
// @Description Удаление правила из системы по его идентификатору
// @Tags rules
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param ruleId path string true "Идентификатор правила"
// @Success 200 {object} models.ResponseAPI{result=models.RuleResponse} "Успешное удаление"
// @Failure 404 {object} models.ResponseAPI
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/rules/{ruleId} [delete]
func (uh *RulesHandler) DeleteRule(ctx *gin.Context) {
	idRule := ctx.Param("ruleId")

	deleteRuleError := uh.service.RulesService.DeleteRule(idRule)
	if deleteRuleError != nil {
		messageError := fmt.Sprintf("Error delete rule: %s", deleteRuleError.Error())
		httpStatus := handlerErrors.MapError(deleteRuleError)
		logger.NewErrorResponse(ctx, uh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   fmt.Sprintf("Delete rule: %s", idRule),
		Result: models.RuleResponse{
			Id: idRule,
		},
	}

	uh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}
