package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"encoding/json"
)

type User struct {
	first_name string             `json:"first_name" binding:"required"`
}

func New(firstName string, lastName string, password string, email string) bool {
	return true
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func CreateGroup(firstName string, lastName string) bool {
	return true
}
func AddFreeTime(firstName string, lastName string) bool {
	return true
}
func VoteForPlace(firstName string, lastName string) bool {
	return true
}

func AddPlaceInGroup(firstName string, lastName string) bool {
	return true
}
func RegisterAccount(c *gin.Context) {

	var input User
	err:= c.BindJSON(&input)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)
	c.JSON(200, gin.H{
		"message": input,
	})
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
func getPwd(pwd string) []byte {
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
