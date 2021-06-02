package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	IncrementalIndex int                `bson:"incremental_index" json:"incremental_index"`
	CustomerID       primitive.ObjectID `bson:"customer_id" json:"-"`
	Customer         *Customer          `json:"customer"`
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`
	Items            []OrderItem        `bson:"items" json:"items"`
}
