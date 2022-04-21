package routes

import (
	"gin-blog/app/controllers/web/auth"
	"gin-blog/app/middlewares"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct{}

func (routes *AuthRoutes) AuthRoutes(router *gin.Engine) {
	var authController = new(auth.AuthController)

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.GET("/register", authController.Register)
		guestGroup.POST("/register", authController.HandleRegister)

		guestGroup.GET("/login", authController.Login)
		guestGroup.POST("/login", authController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/logout", authController.HandleLogout)
	}
}
