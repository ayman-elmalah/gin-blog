package bootstrap

import (
	"gin-blog/pkg/database"
	"gin-blog/pkg/env"
	"gin-blog/pkg/helpers"
	"gin-blog/pkg/routes"
	"gin-blog/pkg/sessions"
	"gin-blog/pkg/static"
	"gin-blog/pkg/templates"
)

func Serve() {
	env.Loader()

	new(database.Mysql).Connect()

	router := routes.GetRouter()

	sessions.StartSession(router)

	routes.RegisterRoutes(router)

	helpers.Load(router)

	templates.LoadHtml(router)

	static.LoadStatic(router)

	routes.ServeRoutes(router)
}
