package database

import (
	"go.mongodb.org/mongo-driver/mongo"

)

// MongoRepo - maintain mongo connection
type MongoRepo struct {
	mongoConn *mongo.Client 
}

