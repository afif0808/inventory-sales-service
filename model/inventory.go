package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	ProductID primitive.ObjectID `bson:"product_id" json:"-"`
	Product   *Product           `json:"product"`
	Quantity  *Quantity          `bson:"quantity" json:"quantity"`
}
