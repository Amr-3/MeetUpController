package DBConnections

import (
	"../User/model"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func DB_Insert_Student(res http.ResponseWriter, req *http.Request) {
	client, err, conContext := CreateDBconnection("mongodb+srv://meetup:9tS3qY8BKwXb1n8d@test-cluster-dvxai.mongodb.net/meetup")
	defer client.Disconnect(conContext)
	res.Header().Add("content-type", "application/json")
	var stud user.User
	_ = json.NewDecoder(req.Body).Decode(&stud)
	collection := client.Database("meetup").Collection("user51")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, stud)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(res).Encode(result)
}
