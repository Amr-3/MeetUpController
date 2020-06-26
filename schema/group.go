package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Duration    int                `json:"duration,omitempty"`
	AdminUser   primitive.ObjectID `json:"admin_user,omitempty"`
	LifeSpan    int                `json:"life_span,omitempty"`
	Users       GroupUsers         `json:"users,omitempty"`
	Places      GroupPlaces        `json:"places,omitempty"`
}

type GroupUsers struct {
	Users []primitive.ObjectID `json:"users,omitempty"`
}
type GroupPlaces struct {
	Places []string `json:"places,omitempty"`
}
