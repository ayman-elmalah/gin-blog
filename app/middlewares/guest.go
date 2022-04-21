package middlewares

import (
	"gin-blog/app/models"
	"gin-blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IsGuest() gin.HandlerFunc {
	var usersModel = new(models.User)

	return func(c *gin.Context) {
		var isExistingUser int64

		auth := sessions.Get(c, "auth")

		userId, _ := strconv.Atoi(auth)

		usersModel.FindUserById(userId).Count(&isExistingUser)

		if isExistingUser != 0 {
			c.Redirect(http.StatusFound, "/")
		}

		c.Next()
	}
}
