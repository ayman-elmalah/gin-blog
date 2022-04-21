package auth

import (
	"gin-blog/app/models"
	"gin-blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Auth(c *gin.Context) models.User {
	var usersModel = new(models.User)
	var authUser models.User

	authID := sessions.Get(c, "auth")

	userId, _ := strconv.Atoi(authID)

	usersModel.FindUserById(userId).Scan(&authUser)

	return authUser
}
