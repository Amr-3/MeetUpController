package DBConnections

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDBconnection(connectionString string) (*mongo.Client, error, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	return client, err, ctx
}
