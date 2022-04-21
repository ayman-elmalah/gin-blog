package bootstrap

import (
	"gin-blog/database/seeders"
	"gin-blog/pkg/database"
	"gin-blog/pkg/env"
)

func Seed() {
	env.Loader()

	new(database.Mysql).Connect()

	seeders.MysqlSeed()
}
