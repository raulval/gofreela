package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raulval/gofreela/schemas"
)

// @BasePath /api/v1

// @Summary Get project
// @Description Get a project
// @Tags Project
// @Accept json
// @Produce json
// @Param id query string true "Project ID"
// @Success 200 {object} GetProjectResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /project [get]
func GetProjectHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, project)
}
