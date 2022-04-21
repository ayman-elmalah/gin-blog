package models

import (
	"gin-blog/pkg/database"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar(191)"`
	Content string `gorm:"text"`
	UserID  uint
	User    User
}

func (article *Article) GetArticles(limit int) *gorm.DB {
	var response []Article

	return database.Connection().Limit(limit).Joins("User").Order("rand()").Find(&response)
}

func (article *Article) GetArticle(id int) *gorm.DB {
	var response Article

	return database.Connection().Joins("User").First(&response, id)
}

func (article *Article) Create(data Article) *gorm.DB {
	var response Article

	return database.Connection().Create(&data).Scan(&response)
}
