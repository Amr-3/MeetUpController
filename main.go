package main

import (
	"./DBConnections"
	"./Place"
	user "./User/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	http.HandleFunc("/insertintodb", DBConnections.DbInsert)
	http.HandleFunc("/removefromdb", DBConnections.DbDelete)
	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.RegisterAccount)

	}
	router.Run(":9000")

	p := Place.Place{
		ID:          1,
		Name:        "Primos",
		Votes:       5,
		Rating:      4.8,
		Description: "the best pizza restaurant",
		UserID:      []string{"1","2","3","4"},
	}

	DBConnections.DbInsert(p,"AyEsm")
}

