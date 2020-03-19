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
	// URI - mongodb URI
	URI = "mongodb://greennit:123456@localhost/greennit"
)


// InitDB - init database
func InitDB() {
	// Base context.
	ctx := context.Background()
	// Options to the database.
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
			fmt.Println(err)
			return
	}
	db := client.Database(DBName)
	fmt.Println(db.Name())
}