package database

import (
	"fmt"
	"gin-blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Mysql struct{}

var DB *gorm.DB

func (db *Mysql) Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.Database()["username"],
		config.Database()["password"],
		config.Database()["host"],
		config.Database()["port"],
		config.Database()["name"],
		config.Database()["charset"],
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting database")
		return
	}

	DB = database
}

func Connection() *gorm.DB {
	return DB
}
