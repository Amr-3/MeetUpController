package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	Id 			primitive.ObjectID
	Name        string
	Description string
	Duration    int
	AdminUser   primitive.ObjectID
	LifeSpan    int
	Users       []primitive.ObjectID
	Places      []string
}
