package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type User struct {
	id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty"`
	Lastname  string             `json:"lastName,omitempty"`
	Email     string             `json:"email,omitempty" binding:"required"`
	Password  string             `json:"password,omitempty" binding:"required"`
	FreeTimes []string           `json:"freetimes,omitempty" `
	Groups    []string           `json:"groups,omitempty"`
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
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input.Firstname)
	fmt.Println(input.Lastname)

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
