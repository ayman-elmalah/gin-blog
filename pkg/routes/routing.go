package routes

import (
	"gin-blog/app/providers/route"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	return gin.Default()
}

func RegisterRoutes(router *gin.Engine) {
	route.RouteRegister(router)
}
