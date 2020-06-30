package main

import (
	. "./User/model"
	. "./config"
	"github.com/gin-gonic/gin"
)

func main() {
	LoadConfig()

	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", RegisterAccount)
		userGroup.POST("/login", Login)
		userGroup.POST("/id/:id/add-free-time", AddFreeTime)
		userGroup.POST("/id/:id/add-friend", AddFriend)
		userGroup.POST("/id/:id/create-group", CreateGroup)
	}
	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"message": "welcome to bate5 ",
		})
	})
	router.Run(Config.PORT)
}
