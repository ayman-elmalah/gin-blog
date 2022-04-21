package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func StartSession(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("sessions", store))
}

func Set(c *gin.Context, key string, value string) {
	session := sessions.Default(c)

	session.Set(key, value)
	_ = session.Save()
}

func Get(c *gin.Context, key string) string {
	session := sessions.Default(c)

	response := session.Get(key)

	if response != nil {
		return response.(string)
	}

	return ""
}

func Flash(c *gin.Context, key string) string {
	session := sessions.Default(c)

	response := session.Get(key)
	_ = session.Save()

	session.Delete(key)
	_ = session.Save()

	if response != nil {
		return response.(string)
	}

	return ""
}

func Remove(c *gin.Context, key string) {
	session := sessions.Default(c)

	session.Delete(key)
	_ = session.Save()
}
