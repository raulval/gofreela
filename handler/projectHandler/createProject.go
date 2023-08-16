package projectHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProjectHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST Project",
	})
}