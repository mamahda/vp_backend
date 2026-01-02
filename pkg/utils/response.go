package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}
