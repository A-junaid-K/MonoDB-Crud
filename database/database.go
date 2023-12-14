package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

	var Collections *mongo.Collection

func ConnectToDb() {
	const uri = "mongodb//:localhost:27017"
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	Collections = client.Database("mongo-crud").Collection("movies")
}
