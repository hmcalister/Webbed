package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PongResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
