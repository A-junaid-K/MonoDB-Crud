package helper

import (
	"context"

	"github.com/Crud-Mongodb/database"
	models "github.com/Crud-Mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertMovie(movie models.Movie) (interface{}, error) {
	insert, err := database.Collections.InsertOne(context.Background(), movie)
	if err != nil {
		return nil, err
	}
	return insert, nil
}

// update a record
func ListAllMovies() ([]bson.M, error) {
	cur, err := database.Collections.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var movies []bson.M

	for cur.Next(context.Background()) {
		var movie bson.M
		if err = cur.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func UpdateMovie(movie models.Movie, id string) (int, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"title": movie.Title, "director": movie.Director, "actor": movie.Actor}}
	res, err := database.Collections.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	return int(res.MatchedCount), nil
}

func GetOneMovie(id string) (*models.Movie, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var movie models.Movie
	filter := bson.M{"_id": oid}
	err = database.Collections.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func DeleteOneMovie(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	_, err = database.Collections.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllMovie() error {
	_, err := database.Collections.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	return nil
}
