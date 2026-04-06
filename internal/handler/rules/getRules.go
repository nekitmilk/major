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

// @Summary Вывод списка правил выявлений
// @Description Получение списка всех правил выявления
// @Tags rules
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} models.ResponseAPI{result=[]models.Rule}
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/rules/ [get]
func (rh *RulesHandler) GetAllRules(ctx *gin.Context) {
	getAllRules, getAllRulesError := rh.service.RulesService.GetAllRules()
	if getAllRulesError != nil {
		messageError := fmt.Sprintf("Error get list rules: %s", getAllRulesError.Error())
		httpStatus := handlerErrors.MapError(getAllRulesError)
		logger.NewErrorResponse(ctx, rh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Get list rules",
		Result:    getAllRules,
	}

	rh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}

// @Summary Вывод списка правил выявлений
// @Description Получение списка всех правил выявления
// @Tags rules
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} models.ResponseAPI{result=[]models.RuleFullData}
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/rules/full [get]
func (rh *RulesHandler) GetAllRulesFullData(ctx *gin.Context) {
	getAllRules, getAllRulesError := rh.service.RulesService.GetAllRulesFullData()
	if getAllRulesError != nil {
		messageError := fmt.Sprintf("Error get list rules: %s", getAllRulesError.Error())
		httpStatus := handlerErrors.MapError(getAllRulesError)
		logger.NewErrorResponse(ctx, rh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Get list rules",
		Result:    getAllRules,
	}

	rh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}

// @Summary Вывод правила выявления по id
// @Description Получение правила выявления по id
// @Tags rules
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ruleId path string true "Идентификатор связи"
// @Success 200 {object} models.ResponseAPI
// @Success 200 {object} models.ResponseAPI{result=models.Rule}
// @Failure 400 {object} models.ResponseAPI
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/rules/full/{ruleId} [get]
func (lh *RulesHandler) GetRuleFullDataByID(ctx *gin.Context) {
	idRule := ctx.Param("ruleId")

	getRule, getRuleError := lh.service.RulesService.GetRuleFullDataByID(idRule)
	if getRuleError != nil {
		messageError := fmt.Sprintf("Error data rule: %s", getRuleError.Error())
		httpStatus := handlerErrors.MapError(getRuleError)
		logger.NewErrorResponse(ctx, lh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   fmt.Sprintf("Get data rule: %s", idRule),
		Result:    getRule,
	}

	lh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}
