package main

import (
	"github.com/gin-gonic/gin"
	"madoka/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Static("/data", "./data")
	err := r.Run(":8800")
	if err != nil {
		return
	}
}
