package main

import (
	. "./User/model"
	. "./config"
	."fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	LoadConfig()
	x := time.Now().Unix()
	y := time.Unix(x,0)
	Println(y)
	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", RegisterAccount)
		userGroup.POST("/login", Login)
		userGroup.POST("/id/:id/add-free-time", AddFreeTime)
		userGroup.POST("/id/:id/add-friend",AddFriend)
	}
	router.Run(Config.PORT)
}
