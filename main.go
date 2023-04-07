package main

import (
	"github.com/gin-gonic/gin"
	"madoka/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)

	err := r.Run()
	if err != nil {
		return
	}
}
