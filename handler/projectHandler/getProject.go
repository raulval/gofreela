package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjectHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Project",
	})
}