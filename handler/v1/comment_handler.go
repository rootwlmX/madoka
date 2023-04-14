package v1

import (
	"github.com/gin-gonic/gin"
	"madoka/common"
	"madoka/service"
)

// CommentHandler 评论handler
type CommentHandler struct {
}

// RegisterRouter 路由注册
func (h *CommentHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/comment")
	group.GET("/getTopComment", h.GetTopComment)
}

// GetTopComment 获取最新评论
func (h *CommentHandler) GetTopComment(c *gin.Context) {
	commentService := service.CommentService{}
	res, err := commentService.GetTopCommentAndTopBrowse()
	if err != nil {
		common.Failed(c, "error")
		return
	}
	common.Success(c, res, "success")
}
