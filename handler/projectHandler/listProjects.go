package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulval/gofreela/schemas"
)

func ListProjectsHandler(ctx *gin.Context) {
	projects := []schemas.Project{}

	if err := db.Find(&projects).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing projects")
		return
	}

	sendSuccess(ctx, projects)
}
