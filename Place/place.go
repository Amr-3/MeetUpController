package Place

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type Place struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Votes       int                `json:"votes,omitempty" bson:"votes,omitempty"`
	Rating      float64            `json:"rating,omitempty" bson:"rating,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	UserID      []string           `json:"userid,omitempty" bson:"userid,omitempty"`
}

func (p Place) Test() {
	tmp := strings.Join(p.UserID, " ")
	tmpp := strings.Fields(tmp)
	fmt.Println(tmp)
	fmt.Println(tmpp)
	fmt.Println(p.ID, p.Name, p.Votes, p.Rating, p.Description, p.UserID)
}

func (p Place) AddFreeTime() {

}

func (p Place) RemoveFreeTime() {

}

func (p Place) GetFreeTime() {

}
