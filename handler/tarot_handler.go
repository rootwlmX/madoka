package handler

import (
	"github.com/gin-gonic/gin"
	"madoka/common"
	"madoka/service"
)

type TarotHandler struct {
}

func (h *TarotHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/api/tarot")
	group.GET("gacha", h.Gacha)
}

func (h *TarotHandler) Gacha(c *gin.Context) {
	tarotService := service.TarotService{}
	err := tarotService.GetRandomTarot()
	if err != nil {
		return
	}
	common.Success(c, nil, "测试")
}
