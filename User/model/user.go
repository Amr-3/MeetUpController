package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	firstName string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	lastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	email     string             `json:"email,omitempty" bson:"email,omitempty"`
	password  string             `json:"password,omitempty" bson:"password,omitempty"`
	freeTimes []string           `json:"freeTimes,omitempty" bson:"freeTimes,omitempty"`
	groups    []string           `json:"groups,omitempty" bson:"groups,omitempty"`
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
	c.JSON(200, gin.H{
		"message": "pong",
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
