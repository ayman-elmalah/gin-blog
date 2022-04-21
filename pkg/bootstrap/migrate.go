package bootstrap

import (
	"gin-blog/database/migrations"
	"gin-blog/pkg/database"
	"gin-blog/pkg/env"
)

func Migrate() {
	env.Loader()

	new(database.Mysql).Connect()

	migrations.MysqlMigrate()
}
