package db

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
	"os"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func ConnectMongoDB() {
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_DB")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("MongoDB ping error: %v", err)
	}

	MongoClient = client
	MongoDB = client.Database(dbName)

	log.Println("Connected to MongoDB")
}
