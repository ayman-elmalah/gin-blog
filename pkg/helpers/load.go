package helpers

import (
	"gin-blog/app/providers/helpers"
	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) {
	helpers.HelperLoad(router)
}
