// Package v1 handlers
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// ArticleCateHandler category
type ArticleCateHandler struct {
}

// RegisterRouter 路由注册
func (h *ArticleCateHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/articleCate")
	group.GET("/getAllList", h.GetAllList)
}

// GetAllList 获取全部
func (h *ArticleCateHandler) GetAllList(c *gin.Context) {
	fmt.Println(c.Query(""))
}
