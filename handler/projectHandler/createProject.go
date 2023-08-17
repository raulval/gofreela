package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulval/gofreela/schemas"
)

func CreateProjectHandler(ctx *gin.Context) {
	request := &CreateProjectRequest{}

	ctx.BindJSON(request)

	if err := request.ValidateCreateProjectRequest(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	project := schemas.Project{
		Title:           request.Title,
		Client:          request.Client,
		Description:     request.Description,
		Deadline:        request.Deadline,
		Status:          request.Status,
		Value:           request.Value,
		IsPaid: *request.IsPaid,
	}

	if err := db.Create(&project).Error; err != nil {
		logger.Errorf("error creating project: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, project)
}