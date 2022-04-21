package routes

import (
	"gin-blog/app/controllers/web/home"
	"github.com/gin-gonic/gin"
)

type HomeRoutes struct{}

func (routes *HomeRoutes) HomeRoutes(router *gin.Engine) {
	var homeController = new(home.HomeController)

	router.GET("/", homeController.Index)
}
