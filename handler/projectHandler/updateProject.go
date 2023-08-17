package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulval/gofreela/schemas"
)

func UpdateProjectHandler(ctx *gin.Context) {
	request := &UpdateProjectRequest{}

	ctx.BindJSON(request)

	if err := request.ValidateUpdateProjectRequest(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	if request.Title != "" {
		project.Title = request.Title
	}
	if request.Client != "" {
		project.Client = request.Client
	}
	if request.Description != "" {
		project.Description = request.Description
	}
	if !request.Deadline.IsZero() {
		project.Deadline = request.Deadline
	}
	if request.Status != "" {
		project.Status = request.Status
	}
	if request.Value > 0 {
		project.Value = request.Value
	}
	if request.IsPaid != nil {
		project.IsPaid = *request.IsPaid
	}

	if err := db.Save(&project).Error; err != nil {
		logger.Errorf("error updating project: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating project")
		return
	}

	sendSuccess(ctx, project)
}
