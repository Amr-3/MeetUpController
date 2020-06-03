package main

import (
	"./DBConnections"
	"./Place"
	user "./User/model"
	. "./config"
	"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	LoadConfig()
	fmt.Println(Config.Port)
	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.RegisterAccount)
	}
	p := Place.Place{
		Name:        "Primos",
		Votes:       5,
		Rating:      4.8,
		Description: "the best pizza restaurant",
		UserID:      []string{"1","2","3","4"},
	}

	DBConnections.DbInsert(p,"AyEsm")
	//idPrimitive, _ := primitive.ObjectIDFromHex("5ed43386749d03263305f7ca")
	//DBConnections.DbDelete(idPrimitive, "AyEsm")
	router.Run(":9000")
}

