package routes

import (
	"gin-blog/app/controllers/web/articles"
	"gin-blog/app/middlewares"
	"github.com/gin-gonic/gin"
)

type ArticleRoutes struct{}

func (routes *ArticleRoutes) ArticleRoutes(router *gin.Engine) {
	var articlesController = new(articles.ArticlesController)

	router.GET("/articles/:id", articlesController.Show)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/articles/create", articlesController.Create)
		authGroup.POST("/articles/store", articlesController.Store)
	}
}
