package home

import (
	"gin-blog/app/models"
	"gin-blog/pkg/templates"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct{}

var articlesModel = new(models.Article)

func (controller *HomeController) Index(c *gin.Context) {
	var featured []models.Article
	articlesModel.GetArticles(4).Scan(&featured)

	var stories []models.Article
	articlesModel.GetArticles(6).Scan(&stories)

	title := "Home"

	templates.Render(c, http.StatusOK, "web/home/index", gin.H{
		"featured": featured,
		"stories":  stories,
		"title":    title,
	})
}
