package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectinString string = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.10.1"

type DB struct {
	client *mongo.Client
}

func connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectinString))
}
