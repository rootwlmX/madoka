// Package midware 中间件
package midware

import (
	"github.com/gin-gonic/gin"
)

// BrowseMiddleWare 浏览量
func BrowseMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
