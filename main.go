package main

import (
	"./DBConnections"
	user "./User/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.RegisterAccount)

	}

	/*p := Place.Place{
		Name:        "Primos",
		Votes:       5,
		Rating:      4.8,
		Description: "the best pizza restaurant",
		UserID:      []string{"1","2","3","4"},
	}*/

	//DBConnections.DbInsert(p,"AyEsm")
	idPrimitive, _ := primitive.ObjectIDFromHex("5ed43386749d03263305f7ca")
	DBConnections.DbDelete(idPrimitive, "AyEsm")
	router.Run(":9000")
}

