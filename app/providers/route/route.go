package route

import (
	"gin-blog/routes"
	"github.com/gin-gonic/gin"
)

func RouteRegister(router *gin.Engine) {
	new(routes.HomeRoutes).HomeRoutes(router)
	new(routes.AuthRoutes).AuthRoutes(router)
	new(routes.ArticleRoutes).ArticleRoutes(router)
}
