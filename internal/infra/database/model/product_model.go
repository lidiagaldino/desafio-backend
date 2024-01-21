package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string  `bson:"title"`
	Price       float64 `bson:"price"`
	Description string  `bson:"description"`
	OwnerID     string  `bson:"owner_id"`
	CategoryID  string  `bson:"category_id"`
}
