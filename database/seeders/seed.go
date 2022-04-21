package seeders

import (
	"fmt"
	"gin-blog/app/models"
	"gin-blog/pkg/database"
)

func MysqlSeed() {
	var user = models.User{Name: "Ayman Elmalah", Email: "info@ayman-elmalah.com", Password: "12345678"}
	database.Connection().Create(&user)

	for i := 1; i < 10; i++ {
		var article = models.Article{Title: fmt.Sprintf("Title %d", i), Content: fmt.Sprintf("Content %d", i), UserID: user.ID}
		database.Connection().Create(&article)
	}

	var user2 = models.User{Name: "The Ayman", Email: "the@ayman.com", Password: "12345678"}
	database.Connection().Create(&user2)

	for i := 1; i < 10; i++ {
		var article = models.Article{Title: fmt.Sprintf("Title %d", i), Content: fmt.Sprintf("Content %d", i), UserID: user2.ID}
		database.Connection().Create(&article)
	}
}
