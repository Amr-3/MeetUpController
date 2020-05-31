package DBConnections

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Aykofta   string             `json:"aykofta,omitempty" bson:"aykofta,omitempty"`
}

func DbInsert (data, collection string) bool {
	client, err, conContext := CreateDBconnection("mongodb+srv://meetup:9tS3qY8BKwXb1n8d@test-cluster-dvxai.mongodb.net/meetup")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = mongoClient.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func DbDelete(ID, collection string) bool {
	client, err, conContext := CreateDBconnection("mongodb+srv://meetup:9tS3qY8BKwXb1n8d@test-cluster-dvxai.mongodb.net/meetup")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)

	mongoClient := client.Database("meetup").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err = mongoClient.DeleteOne(ctx, ID)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
