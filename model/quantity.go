package model

type Quantity struct {
	Variant string `bson:"variant" json:"variant"`
	Value   int    `bson:"value" json:"value"`
}
