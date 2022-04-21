package view

import (
	"gin-blog/app/helpers/auth"
	"gin-blog/pkg/converters"
	"gin-blog/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["errors"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["old"] = converters.StringToUrlValues(sessions.Flash(c, "old"))
	data["auth"] = auth.Auth(c)

	return data
}
