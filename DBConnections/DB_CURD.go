package DBConnections

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

type Student struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Aykofta string `json:"aykofta,omitempty" bson:"aykofta,omitempty"`
}

func DB_Insert_Student(res http.ResponseWriter, req *http.Request){
	client, err, conContext := CreateDBconnection("mongodb+srv://meetup:9tS3qY8BKwXb1n8d@test-cluster-dvxai.mongodb.net/meetup")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	res.Header().Add("content-type", "application/json")
	var stud Student
	_ = json.NewDecoder(req.Body).Decode(&stud)
	collection := client.Database("meetup").Collection("user51")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx,stud)
	if err!=nil{
		log.Fatal(err)
	}
	json.NewEncoder(res).Encode(result)
}

func DB_Delete_Student(res http.ResponseWriter, req *http.Request){
	client, err, conContext := CreateDBconnection("mongodb+srv://meetup:9tS3qY8BKwXb1n8d@test-cluster-dvxai.mongodb.net/meetup")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	var stud Student
	_ = json.NewDecoder(req.Body).Decode(&stud)
	collection := client.Database("meetup").Collection("user51")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.DeleteOne(ctx,stud)
	if err!=nil{
		log.Fatal(err)
	}
	json.NewEncoder(res).Encode(result)
}