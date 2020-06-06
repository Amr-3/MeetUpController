package DBConnections

import (
	. "../config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func DbInsert(data interface{}, collection string) bool {
	client, err, conContext := CreateDBconnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := mongoClient.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(result)
	return true
}
func DbRead(key string, value string, collection string) (map[string]string, error) {
	var result map[string]string
	client, err, conContext := CreateDBconnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)
	err = mongoClient.FindOne(context.TODO(), bson.D{{key, value}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			return nil, err
		}
		log.Fatal(err)
	}
	return result, nil
}
func DbDelete(ID primitive.ObjectID, collection string) bool {
	client, err, conContext := CreateDBconnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)

	mongoClient := client.Database("meetup").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err = mongoClient.DeleteOne(ctx, bson.M{"_id": ID})

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func DbUpdate(ID primitive.ObjectID, collection string) bool{
	// TODO
	return true;
}
