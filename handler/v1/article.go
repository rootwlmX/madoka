package v1

import (
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
}

func (h *ArticleHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/article")
	group.GET("/getList", h.GetList)
}

func (h *ArticleHandler) GetList(c *gin.Context) {

}
