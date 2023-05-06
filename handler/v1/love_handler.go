package v1

import (
	"github.com/gin-gonic/gin"
	"madoka/common"
	"madoka/service"
)

// LoveHandler 点赞
type LoveHandler struct {
}

// RegisterRouter 路由注册
func (h *LoveHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/love")
	group.POST("/add", h.AddLove)
}

// AddLove 点赞增加
func (h *LoveHandler) AddLove(c *gin.Context) {
	loveService := service.LoveService{}
	res, err := loveService.AddLove(c.ClientIP())
	if err != nil {
		common.Failed(c, "error")
		return
	}
	common.Success(c, struct{ status int }{res}, "ok")
}
