package user

import (
	DB "../../DBConnections"
	. "../../schema"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)


//[WIP] This function checks for the user credentials
//[TODO] return token
func Login(c *gin.Context) {
	var usr User
	var userResponse map[string]string
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tmp, err := DB.DbRead("email", usr.Email, "User")
	userResponse = tmp.(map[string]string)
	if comparePasswords(userResponse["password"], usr.Password)!= true {
		c.JSON(400, gin.H{
			"message": "tare2 l salama enta",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Etfdal Ya Bashaaa!!",
	})
}

func CreateGroup(firstName string, lastName string) bool {
	return true
}

//[WIP] This function checks for conflict between time periods
//[TODO] return which times are causing conflict
func CheckTimeConflicts(list []FreeTime) bool {
	for i, time1 := range list {
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

//[WIP] This function adds new free times if possible
//[TODO] return which times are causing conflict
func AddFreeTime(c *gin.Context)  {
	var newFreeTimeArr FreeTimesArr
	err := c.ShouldBind(&newFreeTimeArr)
	userId := c.Param("id")
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{})
	}

	prmUserID, _ := primitive.ObjectIDFromHex(userId)

	// this code dynamically reads form a collection by key
	tmpUser, _ := DB.DbReadByID("freetimearr", prmUserID, "User")
	oldFreeTimeArr := tmpUser.(User).FreeTimeArr
	var allFreeTimeArr FreeTimesArr
	allFreeTimeArr.FreeTime = append(newFreeTimeArr.FreeTime, oldFreeTimeArr.FreeTime...)
	if CheckTimeConflicts(allFreeTimeArr.FreeTime) != true {
		c.JSON(400, gin.H{
			"message": "Get your times right bitch",
		})
		return
	}

	// inserting the freetimearr object into an already found user
	DB.DbUpdate(prmUserID, "User", "freetimearr", allFreeTimeArr)
	c.JSON(200, gin.H{
		"message": "Shokrn ya sadeq, ele mara el gya send nudes",
	})
	return
}

func VoteForPlace(firstName string, lastName string) bool {
	return true
}

func AddPlaceInGroup(firstName string, lastName string) bool {
	return true
}

//[WIP] Registers new account
//[TODO] Check that all the mandatory fields are filled with appropriate data
func RegisterAccount(c *gin.Context) {
	var usr User
	err := c.ShouldBindJSON(&usr)
	//fmt.Println(usr)
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

//[Done] compares hashed password with the input pasword
func comparePasswords(hashedPwd string, plainPwd string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//[Done] returns the password after hashing and spicing it
func getPwd(pwd string) string {
	return hashAndSalt([]byte(pwd))
}

//[Done] hashes the password and adds salt to make it tasty
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
