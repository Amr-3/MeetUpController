package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `json:"_id,omitempty"`
	Firstname   string             `json:"firstname,omitempty"`
	Lastname    string             `json:"lastName,omitempty"`
	Email       string             `json:"email,omitempty" `
	Password    string             `json:"password,omitempty" `
	FreeTimeArr FreeTimesArr       `json:"freetimearr,omitempty"`
	Groups      Groups             `json:"groups,omitempty"`
}

type FreeTimesArr struct {
	FreeTime []FreeTime `json:"freetime,omitempty"`
}

type Groups struct {
	Groups []string `json:"groups,omitempty"`
}

type FreeTime struct {
	StartTime int `json:"starttime,omitempty"`
	EndTime   int `json:"endtime,omitempty"`
	Span      int
}
