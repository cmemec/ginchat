package routers

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/index", service.Index)
	return r
}
