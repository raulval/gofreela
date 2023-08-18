package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raulval/gofreela/schemas"
)

func DeleteProjectHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	// Verify if id is a valid UUID
	parsedID, err := uuid.Parse(id)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid project id")
		return
	}

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	project := schemas.Project{}

	if err := db.First(&project, parsedID).Error; err != nil {
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
