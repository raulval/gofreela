package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raulval/gofreela/handler/projectHandler"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/project", projectHandler.GetProjectHandler)
		v1.POST("/project", projectHandler.CreateProjectHandler)
		v1.PUT("/project", projectHandler.UpdateProjectHandler)
		v1.DELETE("/project", projectHandler.DeleteProjectHandler)
		v1.GET("/projects", projectHandler.ListProjectsHandler)
	}
}