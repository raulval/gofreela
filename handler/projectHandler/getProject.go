package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulval/gofreela/schemas"
)

func GetProjectHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, project)
}
