package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	Director string             `json:"director,omitempty" bson:"director,omitempty"`
	Actor    string             `json:"actor,omitempty" bson:"actor,omitempty"`
}
