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
func updatemovie(Id string){
	movieId,_ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id":movieId}
	update := bson.M{"$set":bson.M{"watched":true}}
	result , err := database.Collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal("failed to update movie : ",err)
	}
	fmt.Println("Successfully updated movie \n modified count : ",result.ModifiedCount)
}