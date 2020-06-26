package main

import (
	. "MeetUp/User/model"
	. "MeetUp/config"
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
	router.Run(Config.PORT)
}
