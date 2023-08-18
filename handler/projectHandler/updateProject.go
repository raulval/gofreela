package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raulval/gofreela/schemas"
)

// @BasePath /api/v1

// @Summary Update project
// @Description Update a project
// @Tags Project
// @Accept json
// @Produce json
// @Param id query string true "Project ID"
// @Param request body UpdateProjectRequest true "Request body"
// @Success 200 {object} UpdateProjectResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /project [put]
func UpdateProjectHandler(ctx *gin.Context) {
	request := &UpdateProjectRequest{}

	ctx.BindJSON(request)

	if err := request.ValidateUpdateProjectRequest(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	if request.Title != "" {
		project.Title = request.Title
	}
	if request.Client != "" {
		project.Client = request.Client
	}
	if request.Description != "" {
		project.Description = request.Description
	}
	if request.Deadline != "" {
		deadlineTime, err := ParseDeadline(request.Deadline)
		if err != nil {
			sendError(ctx, http.StatusBadRequest, err.Error())
			return
		}
		project.Deadline = deadlineTime
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
