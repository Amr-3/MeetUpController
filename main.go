package main

import (
	user "./User/model"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.RegisterAccount)

	}
	router.Run(":9000")
}
