package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProjectsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Projects",
	})
}