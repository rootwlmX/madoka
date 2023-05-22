// Package v1 handlers
package v1

import (
	"github.com/gin-gonic/gin"
	"madoka/common"
	"madoka/midware"
	"madoka/service"
)

// ArticleCateHandler category
type ArticleCateHandler struct {
}

// RegisterRouter 路由注册
func (h *ArticleCateHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/articleCate")
	group.GET("/getAllList", h.GetAllList, midware.BrowseMiddleWare())
}

// GetAllList 获取全部
func (h *ArticleCateHandler) GetAllList(c *gin.Context) {
	articleCateService := service.ArticleCateService{}
	classes, err := articleCateService.GetAllArticleClass()
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	common.Success(c, classes, "ok")
}
