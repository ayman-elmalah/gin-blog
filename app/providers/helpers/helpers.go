package helpers

import (
	"gin-blog/app/helpers/app"
	"gin-blog/app/helpers/auth"
	"gin-blog/app/helpers/date"
	"github.com/gin-gonic/gin"
	"html/template"
)

func HelperLoad(router *gin.Engine) {
	router.SetFuncMap(template.FuncMap{
		"formatAsDate":     date.FormatAsDate,
		"formatAsDateTime": date.FormatAsDateTime,
		"asTitle":          app.AsTitle,
		"auth":             auth.Auth,
	})
}
