package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulval/gofreela/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type CreateProjectResponse struct {
	Data schemas.ProjectResponse `json:"data"`
}

type GetProjectResponse struct {
	Data schemas.ProjectResponse `json:"data"`
}

type UpdateProjectResponse struct {
	Data schemas.ProjectResponse `json:"data"`
}

type DeleteProjectResponse struct {
	Data schemas.ProjectResponse `json:"data"`
}

type ListProjectsResponse struct {
	Data []schemas.ProjectResponse `json:"data"`
}
