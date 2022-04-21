package articles

import (
	"fmt"
	"gin-blog/app/helpers/auth"
	"gin-blog/app/models"
	"gin-blog/app/requests/articles"
	"gin-blog/pkg/converters"
	"gin-blog/pkg/errors"
	"gin-blog/pkg/old"
	"gin-blog/pkg/sessions"
	"gin-blog/pkg/templates"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticlesController struct{}

var articlesModel = new(models.Article)

func (controller *ArticlesController) Create(c *gin.Context) {
	title := "Create article"

	templates.Render(c, http.StatusOK, "web/articles/create", gin.H{
		"title": title,
	})
}

func (controller *ArticlesController) Store(c *gin.Context) {
	var request articles.StoreRequest
	var article models.Article

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.AddFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = auth.Auth(c).ID

	articlesModel.Create(article).Scan(&article)

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
	return
}

func (controller *ArticlesController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var article models.Article
	var count int64

	articlesModel.GetArticle(id).Scan(&article).Count(&count)

	if count == 0 {
		templates.Render(c, http.StatusNotFound, "web/errors/404", gin.H{
			"title": "Page not found",
		})
		return
	}

	title := article.Title

	templates.Render(c, http.StatusOK, "web/articles/show", gin.H{
		"article": article,
		"title":   title,
	})
}
