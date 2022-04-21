package env

import (
	"github.com/joho/godotenv"
	"log"
)

func Loader() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}
