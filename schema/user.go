package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty"`
	Lastname  string             `json:"lastName,omitempty"`
	Email     string             `json:"email,omitempty" `
	Password  string             `json:"password,omitempty" `
	FreeTimes FreeTimes          `json:"freetimes,omitempty"`
	Groups    Groups             `json:"groups,omitempty"`
	Friends   Friends            `json:"friends,omitempty"`
}

type Friends struct {
	DisplayName []string
	Id          []primitive.ObjectID
}

type FreeTimes struct {
	FreeTime []FreeTime `json:"freetime,omitempty"`
}

type Groups struct {
	Groups []primitive.ObjectID
	GroupName []string
}

type FreeTime struct {
	StartTime int `json:"starttime,omitempty"`
	EndTime   int `json:"endtime,omitempty"`
	Span      int
}
