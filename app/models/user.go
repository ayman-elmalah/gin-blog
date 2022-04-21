package models

import (
	"gin-blog/pkg/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"varchar(191)"`
	Email    string `gorm:"varchar(191)"`
	Password string `gorm:"varchar(191)"`
}

func (user *User) FindUserByEmail(email string) *gorm.DB {
	var response User

	return database.Connection().First(&response, "email = ?", email)
}

func (user *User) FindUserById(id int) *gorm.DB {
	var response User

	return database.Connection().First(&response, "id = ?", id)
}

func (user *User) Create(data User) *gorm.DB {
	var response User

	return database.Connection().Create(&data).Scan(&response)
}
