package main

import (
	"ginchat/resource"
	"ginchat/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	if err := resource.InitConfig(); err != nil {
		panic(err.Error())
	}
	// db
	if err := resource.InitDB(); err != nil {
		panic(err.Error())
	}
	r := gin.Default()
	r = routers.Router(r)
	r.Run(":8081")
}
