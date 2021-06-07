package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Storage struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"string" json:"name"`
}
