package user

import (
	DB "../../DBConnections"
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
	var usr User
	var userResponse map[string]string
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userResponse, err = DB.DbRead("email", usr.Email, "User")
	fmt.Println("compare pwd", comparePasswords(userResponse["password"], usr.Password))
	c.JSON(200, gin.H{
		"message": "Etfdal Ya Bashaaa!!",
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
	var usr User
	err := c.ShouldBindJSON(&usr)
	fmt.Println(usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr.Password = getPwd(usr.Password)
	DB.DbInsert(usr, "User")
	c.JSON(200, gin.H{
		"message": usr,
	})
}
func comparePasswords(hashedPwd string, plainPwd string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func getPwd(pwd string) string {
	return hashAndSalt([]byte(pwd))
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
