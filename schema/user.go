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
	DisplayName []string `json:"displayname,omitempty"`
	Id          []primitive.ObjectID `json:"id,omitempty"`
}

type FreeTimes struct {
	FreeTime []FreeTime `json:"freetime,omitempty"`
}

type Groups struct {
	Groups []primitive.ObjectID `json:"groups,omitempty"`
	GroupName []string `json:"groupsname,omitempty"`
}

type FreeTime struct {
	StartTime int `json:"starttime,omitempty"`
	EndTime   int `json:"endtime,omitempty"`
	Span      int
}
