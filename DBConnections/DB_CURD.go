package DBConnections

import (
	. "../config"
	. "../schema"
	"context"
	. "fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//[Done] This function inserts a new field in a collection
func DbInsert(data interface{}, collection string) bool {
	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
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

//[WIP] This function reads one or more value in a field using any other value
//[TODO] fix the find by ID problem
func DbRead(findByKey string, findByValue string, collection string, readKey ...string) (interface{}, error) {
	Println(findByKey+" "+findByValue+" / "+readKey[0])
	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)

	var projection bson.D
	for _, projKey := range readKey {
		projection = append(projection, bson.E{projKey, 1})
	}
	Println("3aaaaa")
	Println(projection)

	result, err := mongoClient.Find(context.Background(), bson.D{{findByKey, findByValue}}, options.Find().SetProjection(projection))
	Println("3aaaaa")
	Println(result)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			return nil, err
		}
		log.Fatal(err)
	}

	for result.Next(context.Background()) {
		Println("da5al gamed")
		var usr User
		// Decode the document
		if err := result.Decode(&usr); err != nil {
			log.Fatal("cursor.Decode ERROR:", err)
		}
		return usr, nil
	}
	return nil, nil
}

//[Done] This function reads a certain object inside a field in a collection using ID
func DbReadByID(key string, id primitive.ObjectID, collection string) (interface{}, error) {
	tmp := bson.D{{"_id", id}}
	Println("3aaaaaa")
	Println(tmp)

	/*tmpID := id.String()
	result, err := DbRead("_id",tmpID , collection, key)
	if err!=nil{
		return nil,err
	}
	return result,nil*/

	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)

	projection := bson.D{
		{key, 1},
	}

	result,err := mongoClient.Find(context.Background(), bson.D{{"_id", id}},options.Find().SetProjection(projection))

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			return  nil,err
		}
		log.Fatal(err)
	}

	for result.Next(context.Background()) {
		var usr User
		// Decode the document
		if err := result.Decode(&usr); err != nil {
			log.Fatal("cursor.Decode ERROR:", err)
		}
		return usr,nil
	}
	return nil,nil

}

//[Done?] This function deletes a field in a collection by it's ID
//to be tested
func DbDelete(ID primitive.ObjectID, collection string) bool {
	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
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

//[Done] This function updates a certain object inside a field in a collection using ID
func DbUpdate(prmUserID primitive.ObjectID, collection string, key string, data interface{}) bool {

	client, err, conContext := CreateDBConnection(Config.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(conContext)
	mongoClient := client.Database("meetup").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"_id", prmUserID}}
	update := bson.D{{"$set", bson.D{{key, data}}}}
	result, err := mongoClient.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	Println(result)
	return true
}
