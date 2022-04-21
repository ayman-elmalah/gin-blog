package auth

import (
	"gin-blog/app/models"
	"gin-blog/app/requests/auth"
	"gin-blog/pkg/converters"
	"gin-blog/pkg/errors"
	"gin-blog/pkg/old"
	"gin-blog/pkg/sessions"
	"gin-blog/pkg/templates"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type AuthController struct{}

var usersModel = new(models.User)

func (controller *AuthController) Register(c *gin.Context) {
	title := "Register"

	templates.Render(c, http.StatusOK, "web/auth/register", gin.H{
		"title": title,
	})
}

func (controller *AuthController) HandleRegister(c *gin.Context) {
	var user models.User
	var request auth.RegisterRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.AddFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	var isExistingUser int64
	usersModel.FindUserByEmail(request.Email).Count(&isExistingUser)

	if isExistingUser > 0 {
		errors.Init()
		errors.Add("Email", "Email address already registered before")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	usersModel.Create(user).Scan(&user)

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	c.Redirect(http.StatusFound, "/")
	return
}

func (controller *AuthController) Login(c *gin.Context) {
	title := "Login"

	templates.Render(c, http.StatusOK, "web/auth/login", gin.H{
		"title": title,
	})
}

func (controller *AuthController) HandleLogin(c *gin.Context) {
	var user models.User
	var isExistingUser int64

	usersModel.FindUserByEmail(c.PostForm("email")).Scan(&user).Count(&isExistingUser)

	if isExistingUser == 0 {
		errors.Init()
		errors.Add("Email", "Invalid credentials")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
	if err != nil {
		errors.Init()
		errors.Add("Email", "Invalid credentials")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	c.Redirect(http.StatusFound, "/")
	return
}

func (controller *AuthController) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")

	c.Redirect(http.StatusFound, "/")
	return
}
