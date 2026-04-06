package issues

import (
	"fmt"
	"major/internal/handler/constHandler"
	"major/internal/models"
	"major/pkg/customErrors/handlerErrors"
	"major/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Вывод списка угроз
// @Description Получение списка всех угроз
// @Tags issues
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} models.ResponseAPI
// @Success 200 {object} models.ResponseAPI{result=[]models.Issue}
// @Failure 400 {object} models.ResponseAPI
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/issues/ [get]
func (nh *IssuesHandler) GetAllIssues(ctx *gin.Context) {
	getAllIssues, getAllIssuesError := nh.service.IssueService.GetAllIssues()
	if getAllIssuesError != nil {
		messageError := fmt.Sprintf("Error get list issues: %s", getAllIssuesError.Error())
		httpStatus := handlerErrors.MapError(getAllIssuesError)
		logger.NewErrorResponse(ctx, nh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Get list issues",
		Result:    getAllIssues,
	}

	nh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}
