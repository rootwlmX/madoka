package v1

import (
	"github.com/gin-gonic/gin"
	"madoka/common"
	"madoka/service"
	"strconv"
)

// CommentHandler 评论handler
type CommentHandler struct {
}

// RegisterRouter 路由注册
func (h *CommentHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/comment")
	group.GET("/getTopComment", h.GetTopComment)
	group.GET("/getList", h.GetList)
}

// GetList 获取评论列表
func (h *CommentHandler) GetList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))
	articleID, _ := strconv.Atoi(c.Query("articleId"))
	commentService := service.CommentService{}
	res, err := commentService.GetCommentList(pageSize, currentPage, articleID)
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	common.Success(c, res, "ok")
}

// GetTopComment 获取最新评论
func (h *CommentHandler) GetTopComment(c *gin.Context) {
	commentService := service.CommentService{}
	res, err := commentService.GetTopCommentAndTopBrowse()
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	common.Success(c, res, "success")
}
