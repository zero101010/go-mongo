package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func (db *DB) Connect() *mongo.Database {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://root:12345678@localhost:27017")
	db.client, _ = mongo.Connect(context.Background(), clientOptions)
	database := db.client.Database("staging-db")
	return database
}
