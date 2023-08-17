package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulval/gofreela/schemas"
)

func DeleteProjectHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	project := schemas.Project{}

	if err := db.First(&project, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "project not found")
		return
	}

	if err := db.Delete(&project).Error; err != nil {
		logger.Errorf("error deleting project: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error deleting project")
		return
	}

	sendSuccess(ctx, project)
}
