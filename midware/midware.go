package midware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func BrowseMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("success")
	}
}
