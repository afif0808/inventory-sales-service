package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaleItem struct {
	Product    *Product           `json:"product"`
	ProductID  primitive.ObjectID `bson:"product_id" json:"-"`
	Quantities []Quantity         `bson:"qtys" json:"qtys"`
	Price      int                `bson:"price" json:"price"`
	Discount   int                `bson:"discount" json:"discount"`
}
