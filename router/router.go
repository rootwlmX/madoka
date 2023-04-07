package router

import (
	"github.com/gin-gonic/gin"
	"madoka/handler"
)

func InitRouter(engine *gin.Engine) {
	new(handler.TarotHandler).RegisterRouter(engine)
}
