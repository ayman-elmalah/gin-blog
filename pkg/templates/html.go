package templates

import (
	"fmt"
	"gin-blog/app/providers/view"
	"gin-blog/config"
	"github.com/gin-gonic/gin"
)

func LoadHtml(router *gin.Engine) {
	router.LoadHTMLGlob(fmt.Sprintf("%s.%s", config.Templates()["pattern"], config.Templates()["extension"]))
}

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = view.WithGlobalData(c, data)

	c.HTML(code, name, data)
}
