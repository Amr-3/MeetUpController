package schema


type FreeTime struct {
	StartTime int  `json:"starttime,omitempty" bson:"starttime,omitempty"`
	EndTime int `json:"endtime,omitempty" bson:"endtime,omitempty"`
	Span int
}
