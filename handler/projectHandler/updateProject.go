package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProjectHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT Project",
	})
}