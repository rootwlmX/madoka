package v1

import (
	"github.com/gin-gonic/gin"
)

type ArticleCateHandler struct {
}

func (h *ArticleCateHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/articleCate")
	group.GET("/getAllList", h.GetAllList)
}

func (h *ArticleCateHandler) GetAllList(c *gin.Context) {

}
