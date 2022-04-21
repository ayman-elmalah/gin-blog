package static

import (
	"gin-blog/app/providers/static"
	"github.com/gin-gonic/gin"
)

func LoadStatic(router *gin.Engine) {
	for key, value := range static.GetStaticPaths() {
		router.Static(key, value)
	}
}
