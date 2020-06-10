package user

import (
	DB "../../DBConnections"
	. "../../schema"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)



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

func CheckTimeConflicts(list []FreeTime) bool {
	for i, time1 := range list{
		for j, time2 := range list {
			if j <= i {
				continue
			}
			if time2.StartTime > time1.StartTime && time2.StartTime < time1.EndTime {
				return false
			}
			if time2.EndTime > time1.StartTime && time2.EndTime < time1.EndTime {
				return false
			}
		}
	}
	return true
}

func AddFreeTime(c *gin.Context) {
	var ft []FreeTime
	err := c.ShouldBind(&ft)
	userId := c.Param("id")
	fmt.Println("user Id:: "+userId)
	if err != nil{
		log.Fatal(err)
		c.JSON(500, gin.H{})
	}
	/*User, _ := DB.DbRead("id",userId,"User")
	fmt.Println(User["freetime"])
	UserTime := strings.Split(User["freetime"],",")
	var userTimesV3 []int
	json.Unmarshal([]byte(User["freetime"]), &userTimesV3)

	if CheckTimeConflicts (userTimesV3) != true {
		c.JSON(69, gin.H{
			"message": "Bara ya ibn el wes5a mn hna",
		})
	}*/
	var usr User
	usr.Id, _ = primitive.ObjectIDFromHex(userId)
	usr.FreeTimes = ft
	DB.DbUpdate(usr,"User")
	c.JSON(200, gin.H{
		"message": "Shokrn ya setaq, ele mara el gya send nudes",
	})
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

