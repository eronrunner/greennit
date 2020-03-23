package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// DBName - Name of the database.
	DBName = "greennit"
	// UserColl - USER collection
	UserColl = "users"
	// URI - mongodb URI
	URI = "mongodb://greennit:123456@192.168.116.128:27017/greennit"
)


// GetConnection - Get db connection
func GetConnection() *mongo.Client {
	// Base context.
	ctx := context.Background()
	// Options to the database.
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
			fmt.Println(err)
	}
	return client
}