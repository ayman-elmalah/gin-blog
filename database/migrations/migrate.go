package migrations

import (
	"gin-blog/app/models"
	"gin-blog/pkg/database"
)

func MysqlMigrate() {
	database.Connection().AutoMigrate(
		&models.User{},
		&models.Article{},
	)
}
