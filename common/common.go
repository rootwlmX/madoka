// Package common restful return
package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success response of success
func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": message,
	})
	c.Abort()
}

// Failed response of failed
func Failed(c *gin.Context, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusBadRequest,
		"message": message,
	})
}
