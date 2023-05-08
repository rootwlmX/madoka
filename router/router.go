// Package router 路由
package router

import (
	"github.com/gin-gonic/gin"
	"madoka/handler/oauth"
	v1 "madoka/handler/v1"
)

// InitRouter 初始化路由handler
func InitRouter(engine *gin.Engine) {
	new(v1.ArticleHandler).RegisterRouter(engine)
	new(v1.CommentHandler).RegisterRouter(engine)
	new(v1.ArticleCateHandler).RegisterRouter(engine)
	new(v1.LoveHandler).RegisterRouter(engine)

	// OAuth
	new(oauth.GithubHandler).RegisterRouter(engine)
}
