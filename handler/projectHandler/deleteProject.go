package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteProjectHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE Project",
	})
}