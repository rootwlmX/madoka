package v1

import (
	"github.com/gin-gonic/gin"
	"madoka/common"
)

// LoveHandler 点赞
type LoveHandler struct {
}

// RegisterRouter 路由注册
func (h *LoveHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/love")
	group.GET("/getAllList", h.AddLove)
}

// AddLove 点赞增加
func (h *LoveHandler) AddLove(c *gin.Context) {
	common.Success(c, nil, "")
}
