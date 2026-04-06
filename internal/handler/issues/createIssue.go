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

// @Summary Добавление угрозы
// @Description Создание угрозы
// @Tags issues
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param request body models.CreateIssueRequest true "Данные новой угрозы"
// @Success 201 {object} models.ResponseAPI{result=models.IssueResponse}
// @Failure 400 {object} models.ResponseAPI
// @Failure 409 {object} models.ResponseAPI
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/issues [post]
func (uh *IssuesHandler) CreateIssue(ctx *gin.Context) {
	var request models.CreateIssueRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		messageError := fmt.Sprintf("Error parsing issue data: %s", err.Error())
		logger.NewErrorResponse(ctx, uh.log, true, http.StatusBadRequest, messageError)
		return
	}

	createIssueId, createIssueError := uh.service.IssueService.CreateIssue(&request)
	if createIssueError != nil {
		messageError := fmt.Sprintf("Error creating issue: %s", createIssueError.Error())
		httpStatus := handlerErrors.MapError(createIssueError)
		logger.NewErrorResponse(ctx, uh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)
	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   "Create issue",
		Result: models.IssueResponse{
			Id: createIssueId,
		},
	}

	uh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusCreated, responseApi)
}
