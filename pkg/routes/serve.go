package routes

import (
	"fmt"
	"gin-blog/config"
	"github.com/gin-gonic/gin"
	"log"
)

func ServeRoutes(router *gin.Engine) {
	err := router.Run(fmt.Sprintf("%v", config.App()["url"]))

	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}
