package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netfilx struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"movie,omitempty"`
	Wathced bool               `json:"watched,omitempty"`
}
