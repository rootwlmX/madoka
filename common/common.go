package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": message,
	})
	c.Abort()
}

func Failed(c *gin.Context, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusBadRequest,
		"message": message,
	})
}
