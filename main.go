package main

import (
	"github.com/gin-gonic/gin"
	"madoka/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Static("/data", "./data")
	err := r.Run()
	if err != nil {
		return
	}
}
