package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raulval/gofreela/schemas"
)

// @BasePath /api/v1

// @Summary Create project
// @Description Create a new project
// @Tags Project
// @Accept json
// @Produce json
// @Param request body CreateProjectRequest true "Request body"
// @Success 200 {object} CreateProjectResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /project [post]
func CreateProjectHandler(ctx *gin.Context) {
	request := &CreateProjectRequest{}

	ctx.BindJSON(request)

	if err := request.ValidateCreateProjectRequest(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	project := schemas.Project{
		ID:          uuid.New(),
		Title:       request.Title,
		Client:      request.Client,
		Description: request.Description,
		Deadline:    request.Deadline,
		Status:      request.Status,
		Value:       request.Value,
		IsPaid:      *request.IsPaid,
	}

	if err := db.Create(&project).Error; err != nil {
		logger.Errorf("error creating project: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, project)
}
