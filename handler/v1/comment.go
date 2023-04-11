package v1

import "github.com/gin-gonic/gin"

type CommentHandler struct {
}

func (h *CommentHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/comment")
	group.GET("/getTopComment", h.GetTopComment)
}

func (h *CommentHandler) GetTopComment(c *gin.Context) {

}
