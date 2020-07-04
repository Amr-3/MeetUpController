package user

import (
	. "../../schema"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	url2 "net/url"

	//"github.com/parnurzeal/gorequest"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	//"os"
)

//[WIP] This function checks for the user credentials
//[TODO] return token
/*func Login(c *gin.Context) {
	var usr User
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//tmp, err := DB.DbRead("email", usr.Email, "User", "password")
	userResponse := tmp.(User)
	if comparePasswords(userResponse.Password, usr.Password) != true {
		c.JSON(400, gin.H{
			"message": "tare2 l salama enta",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Etfdal Ya Bashaaa!!",
	})
}*/

/*func CreateGroup(c *gin.Context) {
	var newGroup Group
	err := c.ShouldBind(&newGroup)
	userId := c.Param("id")
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{})
	}
	newGroup.AdminUser, err = primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{})
	}

	newGroup.Users.Users = append(newGroup.Users.Users, newGroup.AdminUser)
	//doneFlag := DB.DbInsert(newGroup, "Group")

	/*if !doneFlag {
		c.JSON(500, gin.H{"message": "Wrong Hole"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Group Added",
	})
	return

}*/

//[WIP] This function checks for conflict between time periods
//[TODO] return which times are causing conflict
/*func CheckTimeConflicts(list []FreeTime) bool {
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
*/
//[WIP] This function adds new free times if possible
//[TODO] return which times are causing conflict
/*func AddFreeTime(c *gin.Context) {
	var newFreeTimeArr FreeTimes
	err := c.ShouldBind(&newFreeTimeArr)
	userId := c.Param("id")
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{})
	}

	prmUserID, _ := primitive.ObjectIDFromHex(userId)

	// this code dynamically reads form a collection by key
	//tmpUser, _ := DB.DbRead("_id", prmUserID, "User", "freetimes")
	oldFreeTimeArr := tmpUser.(User).FreeTimes
	var allFreeTimeArr FreeTimes
	allFreeTimeArr.FreeTime = append(newFreeTimeArr.FreeTime, oldFreeTimeArr.FreeTime...)
	if CheckTimeConflicts(allFreeTimeArr.FreeTime) != true {
		c.JSON(400, gin.H{
			"message": "Get your times right bitch",
		})
		return
	}

	// inserting the freetimes object into an already found user
	//DB.DbUpdate(prmUserID, "User", "freetimes", allFreeTimeArr)
	c.JSON(200, gin.H{
		"message": "Shokrn ya sadeq, ele mara el gya send nudes",
	})
	return
}
*/
func VoteForPlace(firstName string, lastName string) bool {
	return true
}

func AddPlaceInGroup(firstName string, lastName string) bool {
	return true
}

func UserToString(user User) (string){
	tmpmsg, err := json.Marshal(user)
	if err!=nil{
		log.Fatal(err)
		return ""
	}
	return string(tmpmsg)
}

//[WIP] Registers new account with unique email
//[TODO] Check that all the mandatory fields are filled with appropriate data
func RegisterAccount(c *gin.Context) {

	var usr User
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if the email is unique
	//found, err := DB.DbRead("email", usr.Email, "User", "email")
	tmp,_ := url2.Parse("http://localhost:9000/CRUD/read")
	url:=tmp.String()
	request := gorequest.New()

	msg := `{"findByKey":"email", "findByValue":"`+usr.Email+`","collection":"User","readKey":["email"]}`
	fmt.Println(msg)
	resp, _, errs := request.Post(url).
		Send(msg).
		End()

	if errs != nil {
		log.Fatal(errs)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Post Error"})
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", resp.Body)

	if resp.StatusCode == http.StatusAccepted {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This email is already taken ro7 shoflak kalba"})
		return
	}

	usr.Password = getPwd(usr.Password)
	//DB.DbInsert(usr, "User")

	tmp,_ = url2.Parse("http://localhost:9000/CRUD/insert")
	url = tmp.String()
	msg=UserToString(usr)
	msg=`{"user":`+msg+`,"collection":"User"}`
	fmt.Println(msg)
	resp, _, errs = request.Post(url).
		Send(msg).
		End()

	if errs != nil {
		log.Fatal(errs)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Post Error"})
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", resp.Body)

	if resp.StatusCode == http.StatusAccepted {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This email is already taken ro7 shoflak kalba"})
		return
	}

	c.JSON(200, gin.H{
		"message": usr,
	})
}

//[Done] Adding a friend in the friends list
/*func AddFriend(c *gin.Context) {

	var emailUser User
	err := c.ShouldBind(&emailUser)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{})
	}

	userId := c.Param("id")

	//get the friend user data
	interFriendUser, err := DB.DbRead("email", emailUser.Email, "User", "firstname", "lastname", "_id")
	if interFriendUser == nil {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}
	friendUser := interFriendUser.(User)

	//get the old list of friends
	id, err := primitive.ObjectIDFromHex(userId)
	currentFriends, err := DB.DbRead("_id", id, "User", "friends")
	newFriends := currentFriends.(User)

	//check if friend already exists
	for _, friendId := range newFriends.Friends.Id {
		if friendId == friendUser.ID {
			c.JSON(400, gin.H{
				"message": "Mawgod ya sakallyyy",
			})
			return
		}
	}
	//add the new friend data
	newFriends.Friends.DisplayName = append(newFriends.Friends.DisplayName, friendUser.Firstname+" "+friendUser.Lastname)
	newFriends.Friends.Id = append(newFriends.Friends.Id, friendUser.ID)

	//insert the new friend list in the user object
	done := DB.DbUpdate(id, "User", "friends", newFriends.Friends)
	if done == false {
		c.JSON(404, gin.H{
			"message": "bos fe 7aga 3'alat 7asalet bs ana msh 3rf eh lessa",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "mabrook 3alek l friend l gdeed 3o2bal kol mara",
	})
}*/
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
