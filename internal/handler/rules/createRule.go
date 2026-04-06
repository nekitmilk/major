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

// @Summary Добавление правила
// @Description Создание правила
// @Tags rules
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param request body models.CreateRuleRequest true "Данные нового правила"
// @Success 201 {object} models.ResponseAPI{result=models.RuleResponse}
// @Failure 400 {object} models.ResponseAPI
// @Failure 409 {object} models.ResponseAPI
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/rules [post]
func (rh *RulesHandler) CreateRule(ctx *gin.Context) {
	var request models.CreateRuleRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		messageError := fmt.Sprintf("Error parsing rule data: %s", err.Error())
		logger.NewErrorResponse(ctx, rh.log, true, http.StatusBadRequest, messageError)
		return
	}

	createRuleId, createRuleError := rh.service.RulesService.CreateRule(&request)
	if createRuleError != nil {
		messageError := fmt.Sprintf("Error creating rule: %s", createRuleError.Error())
		httpStatus := handlerErrors.MapError(createRuleError)
		logger.NewErrorResponse(ctx, rh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)
	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Create rule",
		Result: models.RuleResponse{
			Id: createRuleId,
		},
	}

	rh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusCreated, responseApi)
}
