package main

import (
	. "./User/model"
	. "./config"
	. "fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	LoadConfig()
	Println(Config.PORT)
	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", RegisterAccount)
		userGroup.POST("/login", Login)
	}

	router.Run(Config.PORT)
}
