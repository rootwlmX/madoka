package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	fmt.Println(c.Query(""))
}
