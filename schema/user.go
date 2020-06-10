package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty"`
	Lastname  string             `json:"lastName,omitempty"`
	Email     string             `json:"email,omitempty" binding:"required"`
	Password  string             `json:"password,omitempty" binding:"required"`
	FreeTimes []FreeTime         `json:"freetime,omitempty" `
	Groups    []string           `json:"groups,omitempty"`
}