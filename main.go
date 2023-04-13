package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"madoka/engine"
	"madoka/router"
	"madoka/util"
)

func main() {

	config, err := util.ParseConfig("./config/db.json")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	_, err = engine.NewOrmEngine(config)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	r := gin.Default()
	router.InitRouter(r)
	r.Static("/data", "./data")
	err = r.Run(":8800")
	if err != nil {
		return
	}
}
