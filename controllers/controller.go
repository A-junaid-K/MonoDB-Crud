package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/Crud-Mongodb/database"
	"github.com/Crud-Mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insertMovie(movie model.Netfilx) {
	insert, err := database.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal("Failed to insert movie : ", err)
	}
	fmt.Println("Inserted one movei in database with id : ", insert.InsertedID)
}

// update a record
func updatemovie(Id string) {
	movieId, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": movieId}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := database.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal("failed to update movie : ", err)
	}
	fmt.Println("Successfully updated movie \n modified count : ", result.ModifiedCount)
}

// delete one record from DB
func deletemovie(Id string) {
	movieId, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": movieId}
	deletecount, _ := database.Collection.DeleteOne(context.Background(), filter)
	fmt.Println("Movie successfully deleted with delete count : ", deletecount)
}

// delte multiple record from the database
func deleteMany(Id string) int64{
	result, err := database.Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal("Failed to dalate all records from the Db : ", err)
	}
	fmt.Println("Succeffully deleted all records with delet count : ", result.DeletedCount)
	return result.DeletedCount
}
