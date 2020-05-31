package main

import (
	"./DBConnections"
	user "./User/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	http.HandleFunc("/insertintodb", DBConnections.DB_Insert_Student)
	http.HandleFunc("/removefromdb", DBConnections.DB_Delete_Student)
	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.RegisterAccount)

	}
	router.Run(":9000")
}
