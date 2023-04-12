package router

import (
	"github.com/gin-gonic/gin"
	"madoka/handler"
	v1 "madoka/handler/v1"
)

func InitRouter(engine *gin.Engine) {
	new(handler.TarotHandler).RegisterRouter(engine)
	new(v1.ArticleHandler).RegisterRouter(engine)
	new(v1.CommentHandler).RegisterRouter(engine)
	new(v1.ArticleCateHandler).RegisterRouter(engine)
}
