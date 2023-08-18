package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raulval/gofreela/docs"
	"github.com/raulval/gofreela/handler/projectHandler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	projectHandler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/project", projectHandler.GetProjectHandler)
		v1.POST("/project", projectHandler.CreateProjectHandler)
		v1.PUT("/project", projectHandler.UpdateProjectHandler)
		v1.DELETE("/project", projectHandler.DeleteProjectHandler)
		v1.GET("/projects", projectHandler.ListProjectsHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
