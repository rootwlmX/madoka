package v1

import (
	"github.com/gin-gonic/gin"
	"madoka/common"
	"madoka/service"
	"strconv"
)

type ArticleHandler struct {
}

func (h *ArticleHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/article")
	group.GET("/getList", h.GetList)
}

func (h *ArticleHandler) GetList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))

	articleService := service.ArticleService{}
	articles, err := articleService.ListArticles(pageSize, currentPage)
	if err != nil {
		common.Failed(c, err.Error())
	}
	common.Success(c, articles, "success")
}
