package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulval/gofreela/schemas"
)

// @BasePath /api/v1

// @Summary List projects
// @Description List all projects
// @Tags Project
// @Accept json
// @Produce json
// @Success 200 {object} ListProjectsResponse
// @Failure 500 {object} ErrorResponse
// @Router /projects [get]
func ListProjectsHandler(ctx *gin.Context) {
	projects := []schemas.Project{}

	if err := db.Find(&projects).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing projects")
		return
	}

	sendSuccess(ctx, projects)
}
