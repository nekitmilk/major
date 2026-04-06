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

// @Summary Удаление угрозы
// @Description Удаление угрозы из системы по ее идентификатору
// @Tags issues
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param issueId path string true "Идентификатор угрозы"
// @Success 200 {object} models.ResponseAPI{result=models.IssueResponse} "Успешное удаление"
// @Failure 404 {object} models.ResponseAPI
// @Failure 500 {object} models.ResponseAPI
// @Router /v1/issues/{issueId} [delete]
func (uh *IssuesHandler) DeleteIssue(ctx *gin.Context) {
	idIssue := ctx.Param("issueId")

	deleteIssueError := uh.service.IssueService.DeleteIssue(idIssue)
	if deleteIssueError != nil {
		messageError := fmt.Sprintf("Error delete issue: %s", deleteIssueError.Error())
		httpStatus := handlerErrors.MapError(deleteIssueError)
		logger.NewErrorResponse(ctx, uh.log, true, httpStatus, messageError)
		return
	}

	currentRequestId, _ := ctx.Get(constHandler.REQUEST_ID)

	responseApi := &models.ResponseAPI{
		Success:   true,
		RequestId: fmt.Sprint(currentRequestId),
		Message:   fmt.Sprintf("Delete issue: %s", idIssue),
		Result: models.IssueResponse{
			Id: idIssue,
		},
	}

	uh.log.Debugf("response %s: %+v", ctx.Request.URL.Path, responseApi)

	ctx.JSON(http.StatusOK, responseApi)
}
